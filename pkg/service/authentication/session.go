package authentication

import (
	gocontext "context"
	"errors"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/bufbuild/connect-go"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetToken(baseCtx gocontext.Context, req *connect.Request[v1.GetTokenRequest]) (*connect.Response[v1.GetTokenResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetTokenResponse{})
	if err := validateGetTokenRequest(req.Msg); err != nil {
		return nil, err
	}
	var sess models.Session
	now := time.Now()
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		user, err := queries.GetUserByEmail(ctx, req.Msg.Email)
		if err != nil {
			return err
		}
		if user.Active.Valid && !user.Active.Bool {
			return errors.New("inactive user")
		}
		if !checkHash(user, req.Msg.Password) {
			return errors.New("invalid password")
		}
		sess, err = queries.CreateSession(ctx, models.CreateSessionParams{
			UserID:    user.ID,
			CreatedAt: now,
			ExpiredAt: now.Add(s.SessionDuration),
		})
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid email/password"))
	}
	token, err := s.authServer.CreateSessionJWT(sess)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	AddSessionToken(res.Header(), token)
	return res, nil
}

func (s *Service) GetSessions(baseCtx gocontext.Context, req *connect.Request[v1.GetSessionsRequest]) (*connect.Response[v1.GetSessionsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetSessionsResponse{})
	token, err := ExtractSessionToken(req.Header())
	if err != nil {
		return nil, err
	}
	sess, err := s.authServer.VerifySessionJWT(token)
	if err != nil {
		ctx.L().Error("Failed to verify session JTW", zap.Error(err))
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing token"))
	}
	var sessions []models.Session
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		sessions, err = queries.GetSessions(ctx, models.GetSessionsParams{
			Limit:  req.Msg.Pagination.Length,
			Offset: req.Msg.Pagination.Cursor,
			UserID: sess.UserID,
		})
		if err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res.Msg.Sessions = make([]*v1.Session, len(sessions))
	for i, sess := range sessions {
		res.Msg.Sessions[i].Id = sess.ID.String()
		res.Msg.Sessions[i].CreatedAt = timestamppb.New(sess.CreatedAt)
		res.Msg.Sessions[i].ExpiredAt = timestamppb.New(sess.ExpiredAt)
	}
	return res, nil
}

func (s *Service) DeleteSession(gocontext.Context, *connect.Request[v1.DeleteSessionRequest]) (*connect.Response[v1.DeleteSessionResponse], error) {
	return nil, nil
}
