package authentication

import (
	gocontext "context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const sessionDuration = time.Minute * 5

var errUnauthorizedUser = errors.New("could not authenticate user")

func (s *Service) Authenticate(baseCtx gocontext.Context, req *connect_go.Request[v1.AuthenticateRequest]) (*connect_go.Response[v1.AuthenticateResponse], error) {
	ctx := context.New(baseCtx).Named("authenticate")
	logger := ctx.L()
	res := connect.NewResponse(&v1.AuthenticateResponse{})
	if req.Msg == nil {
		logger.Error("failed beginning transaction", zap.Error(errors.New("missing request body")))
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("missing request body"))
	}
	ctx.AddEvent("rolled back")
	logger.Debug("starting transaction")
	tx, err := s.conn.Begin()
	if err != nil {
		logger.Error("failed beginning transaction", zap.Error(err))
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	var success bool
	defer func() {
		logger.Debug("ending transaction")
		if !success {
			logger.Warn("rolling back")
			ctx.AddEvent("rolled back")
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	queries := models.New(tx)
	user, err := queries.GetUser(ctx, req.Msg.Email)
	if err != nil {
		logger.Error("failed getting user", zap.Error(err), zap.String("email", req.Msg.Email))
		ctx.AddEvent("rolled back")
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, errUnauthorizedUser)
	}
	if !checkHash(user, req.Msg.Password) {
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, errUnauthorizedUser)
	}
	sessionID, err := queries.CreateSession(ctx, models.CreateSessionParams{
		UserID:    user.ID.UUID,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(sessionDuration),
	})
	if err != nil {
		logger.Error("failed creating session", zap.Error(err), zap.String("userID", user.ID.UUID.String()))
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, errUnauthorizedUser)
	}
	success = true
	AddSessionToken(res.Header(), sessionID.String())
	return res, nil
}

func (s *Service) Create(baseCtx gocontext.Context, req *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	ctx := context.New(baseCtx).Named("create")
	logger := ctx.L()
	res := connect.NewResponse(&v1.CreateResponse{})
	if req.Msg == nil {
		logger.Error("failed beginning transaction", zap.Error(errors.New("missing request body")))
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("missing request body"))
	}
	hash, salt, err := generateCred(req.Msg.Password)
	if err != nil {
		logger.Error("failed generating credentials", zap.Error(err))
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	queries := models.New(s.conn)
	user, err := queries.CreateUser(ctx, models.CreateUserParams{
		Email:     req.Msg.Email,
		FirstName: req.Msg.FirstName,
		LastName:  req.Msg.LastName,
		Hash:      hash,
		Salt:      salt,
	})
	if err != nil {
		logger.Error("failed creating user", zap.Error(err))
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	ctx.AddEvent("Created new user")
	res.Msg.Id = user.ID.UUID.String()
	res.Msg.Email = user.Email
	res.Msg.FirstName = user.FirstName
	res.Msg.LastName = user.LastName
	return res, nil
}

func (s *Service) Whoami(baseCtx gocontext.Context, req *connect_go.Request[v1.WhoamiRequest]) (*connect_go.Response[v1.WhoamiResponse], error) {
	ctx := context.New(baseCtx).Named("whoami")
	logger := ctx.L()
	res := connect.NewResponse(&v1.WhoamiResponse{})
	sessionID, err := ExtractSessionToken(req.Header())
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, err)
	}
	logger = logger.With(zap.String("sessionID", sessionID.String()))
	logger.Debug("starting transaction")
	tx, err := s.conn.Begin()
	if err != nil {
		logger.Error("failed beginning transaction", zap.Error(err))
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	var success bool
	defer func() {
		logger.Debug("ending transaction")
		if !success {
			logger.Warn("rolling back")
			ctx.AddEvent("rolled back")
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	queries := models.New(tx)
	row, err := queries.GetSessionUserID(ctx, sessionID)
	if err != nil {
		logger.Error("failed getting session")
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	if row.ExpiredAt.Before(time.Now()) {
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, err)
	}
	user, err := queries.GetUserByID(ctx, uuid.NullUUID{
		UUID:  row.UserID,
		Valid: true,
	})
	if err != nil {
		logger.Error("failed getting user id", zap.String("userID", row.UserID.String()))
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	res.Msg.Id = user.ID.UUID.String()
	res.Msg.Email = user.Email
	res.Msg.FirstName = user.FirstName
	res.Msg.LastName = user.LastName
	success = true
	return res, nil
}

func generateCred(password string) (string, string, error) {
	rawSalt := make([]byte, 128)
	if _, err := rand.Read(rawSalt); err != nil {
		return "", "", err
	}
	salt := base64.StdEncoding.EncodeToString(rawSalt)
	digest := append(rawSalt, []byte(password)...)
	rawHash, err := bcrypt.GenerateFromPassword(digest, 0)
	if err != nil {
		return "", "", err
	}
	hash := base64.StdEncoding.EncodeToString(rawHash)
	return hash, salt, nil
}

func checkHash(user models.User, password string) bool {
	salt, err := base64.StdEncoding.DecodeString(user.Salt)
	if err != nil {
		return false
	}
	userDigest := append(salt, []byte(password)...)
	hashDigest, err := base64.StdEncoding.DecodeString(user.Hash)
	if err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword(hashDigest, userDigest); err != nil {
		// TODO check for non-mismatched errors
		return false
	}
	return true
}

func AddSessionToken(h http.Header, token string) {
	h.Add("Authentication", "Bearer: "+token)
}

func ExtractSessionToken(h http.Header) (uuid.UUID, error) {
	authHeader := h.Get("Authentication")
	if authHeader == "" {
		return uuid.Nil, errUnauthorizedUser
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 {
		return uuid.Nil, errUnauthorizedUser
	}
	return uuid.Parse(authParts[1])
}
