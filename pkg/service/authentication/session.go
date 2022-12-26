package authentication

import (
	gocontext "context"
	"errors"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	commonv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/common/v1"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateSession(baseCtx gocontext.Context, req *connect.Request[v1.CreateSessionRequest]) (*connect.Response[v1.CreateSessionResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateSessionResponse{})
	if err := validateCreateSessionRequest(req.Msg); err != nil {
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
	token, err := auth.EncodeToken(sess)
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
	pageInfo := GetPage(req.Msg)
	sess, err := auth.DecodeToken(token)
	if err != nil {
		ctx.L().Error("Failed to verify session", zap.Error(err))
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing token"))
	}
	var sessions []models.Session
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		sessions, err = queries.GetSessions(ctx, models.GetSessionsParams{
			Limit:  pageInfo.Length,
			Offset: pageInfo.Cursor,
			UserID: sess.UserID,
		})
		req.Msg.GetPagination()
		if err != nil {
			return err
		}
		return err
	}); err != nil {
		ctx.L().Error("Failed to run transaction", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res.Msg.Sessions = make([]*v1.Session, 0, len(sessions))
	for _, sess := range sessions {
		var newSession v1.Session
		newSession.Id = sess.ID.String()
		newSession.CreatedAt = timestamppb.New(sess.CreatedAt)
		newSession.ExpiredAt = timestamppb.New(sess.ExpiredAt)
		res.Msg.Sessions = append(res.Msg.Sessions, &newSession)
	}
	return res, nil
}

func GetPage(req interface {
	GetPagination() *commonv1.PaginationRequest
}) (output commonv1.PaginationRequest) {
	output.Cursor = 0
	output.Length = 10
	if req != nil {
		if page := req.GetPagination(); page != nil {
			output.Cursor = page.Cursor
			output.Length = page.Length
		}
	}
	return
}

func (s *Service) DeleteSession(baseCtx gocontext.Context, req *connect.Request[v1.DeleteSessionRequest]) (*connect.Response[v1.DeleteSessionResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteSessionResponse{})
	token, err := ExtractSessionToken(req.Header())
	if err != nil {
		return nil, err
	}
	sess, err := auth.DecodeToken(token)
	if err != nil {
		ctx.L().Error("Failed to verify session", zap.Error(err))
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing token"))
	}
	if sess.UserID != uuid.MustParse(req.Msg.UserId) {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("permission denied"))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		err = queries.DeleteSession(ctx, models.DeleteSessionParams{
			ID:     uuid.MustParse(req.Msg.Id),
			UserID: uuid.MustParse(req.Msg.UserId),
		})
		if err != nil {
			return err
		}
		return err
	}); err != nil {
		ctx.L().Error("Failed to run transaction", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return res, nil
}
