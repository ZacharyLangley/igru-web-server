package authentication

import (
	gocontext "context"
	"errors"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/auth"
	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	commonv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/common/v1"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var errInactiveUser = errors.New("inactive user")
var errInvalidPassword = errors.New("invalid password")
var errInvalidLogin = errors.New("invalid email/password")

func (s *Service) CreateSession(baseCtx gocontext.Context, req *connect.Request[v1.CreateSessionRequest]) (*connect.Response[v1.CreateSessionResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateSessionResponse{})
	if err := validateCreateSessionRequest(req.Msg); err != nil {
		return nil, err
	}
	ctx.AddStringAttribute("email", req.Msg.Email)
	var sess models.Session
	now := time.Now()
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		user, err := queries.GetUserByEmail(ctx, req.Msg.Email)
		if err != nil {
			return err
		}
		if user.Active.Valid && !user.Active.Bool {
			return errInactiveUser
		}
		if !checkHash(user, req.Msg.Password) {
			return errInvalidPassword
		}
		sess, err = queries.CreateSession(ctx, models.CreateSessionParams{
			UserID: user.ID,
			CreatedAt: pgtype.Timestamp{
				Time:  now,
				Valid: true,
			},
			ExpiredAt: pgtype.Timestamp{
				Time:  now.Add(s.SessionDuration),
				Valid: true,
			},
		})
		// Need to do a conditional here where if fullName is not submitted, then we use the email
		res.Msg.User = &v1.User{
			Id:         uuid.UUID(user.ID.Bytes).String(),
			Email:      user.Email,
			FullName:   user.FullName.String,
			GlobalRole: v1.GroupRole(user.GlobalRole.Int32),
		}
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errInvalidLogin)
	}
	token, err := auth.EncodeToken(sess)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	ctx.L().Info("Started new user session", zap.String("user_id", uuid.UUID(sess.UserID.Bytes).String()), zap.String("session_id", uuid.UUID(sess.ID.Bytes).String()))
	AddSessionToken(res.Header(), token)
	return res, nil
}

func (s *Service) GetSessionUser(baseCtx gocontext.Context, req *connect.Request[v1.GetSessionUserRequest]) (*connect.Response[v1.GetSessionUserResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetSessionUserResponse{})
	token, err := ExtractSessionToken(req.Header())
	if err != nil {
		return nil, err
	}
	sess, err := auth.DecodeToken(token)
	if err != nil {
		ctx.L().Error("Failed to verify session", zap.Error(err))
		return nil, connect.NewError(connect.CodeUnauthenticated, auth.ErrMissingToken)
	}
	now := time.Now()
	if !sess.ExpiredAt.Valid || now.After(sess.ExpiredAt.Time) {
		ctx.L().Error("Token is Expired", zap.Error(err))
		return nil, connect.NewError(connect.CodeUnauthenticated, auth.ErrPermissionDenied)
	}

	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		user, err := queries.GetUser(ctx, sess.UserID)
		if err != nil {
			return err
		}
		res.Msg.User = &v1.User{
			Id:         uuid.UUID(user.ID.Bytes).String(),
			Email:      user.Email,
			FullName:   user.FullName.String,
			GlobalRole: v1.GroupRole(user.GlobalRole.Int32),
		}
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, auth.ErrPermissionDenied)
	}
	ctx.L().Info("Retrieved session information", zap.String("user_id", uuid.UUID(sess.UserID.Bytes).String()), zap.String("session_id", uuid.UUID(sess.ID.Bytes).String()))

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
		return nil, connect.NewError(connect.CodeUnauthenticated, auth.ErrMissingToken)
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
		newSession.Id = uuid.UUID(sess.ID.Bytes).String()
		newSession.UserId = uuid.UUID(sess.UserID.Bytes).String()
		newSession.CreatedAt = timestamppb.New(sess.CreatedAt.Time)
		newSession.ExpiredAt = timestamppb.New(sess.ExpiredAt.Time)
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
		return nil, connect.NewError(connect.CodeUnauthenticated, auth.ErrMissingToken)
	}
	// Only the target user ID can delete their own sessions
	if uuid.UUID(sess.UserID.Bytes) != uuid.MustParse(req.Msg.UserId) {
		return nil, connect.NewError(connect.CodeUnauthenticated, auth.ErrPermissionDenied)
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		err = queries.DeleteSession(ctx, models.DeleteSessionParams{
			ID:     database.NewFromUUID(uuid.MustParse(req.Msg.Id)),
			UserID: database.NewFromUUID(uuid.MustParse(req.Msg.UserId)),
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

func (s *Service) CheckSessionPermissions(baseCtx gocontext.Context, req *connect.Request[v1.CheckSessionPermissionsRequest]) (*connect.Response[v1.CheckSessionPermissionsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CheckSessionPermissionsResponse{})
	token, err := ExtractSessionToken(req.Header())
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}
	sess, err := auth.DecodeToken(token)
	if err != nil {
		ctx.L().Error("Failed to verify session", zap.Error(err))
		return nil, connect.NewError(connect.CodeUnauthenticated, auth.ErrMissingToken)
	}
	res.Msg.Responses = make([]*v1.PermissionResponse, len(req.Msg.Requests))
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		user, err := queries.GetUser(ctx, sess.UserID)
		if err != nil {
			return err
		}
		rawRoles, err := queries.GetUserGroupRoles(ctx, sess.UserID)
		if err != nil {
			return err
		}
		roles := make(map[string]v1.GroupRole)
		for _, role := range rawRoles {
			if role.GroupID.Valid {
				roles[uuid.UUID(role.GroupID.Bytes).String()] = v1.GroupRole(role.Role)
			}
		}
		for i, permissionRequest := range req.Msg.Requests {
			if permissionRequest.GroupId != nil {
				// Check group role
				grantedRole, ok := roles[*permissionRequest.GroupId]
				res.Msg.Responses[i] = &v1.PermissionResponse{
					Request: permissionRequest,
					Allowed: ok && grantedRole == permissionRequest.Role,
				}
			} else {
				// Check global role
				ok := user.GlobalRole.Valid
				res.Msg.Responses[i] = &v1.PermissionResponse{
					Request: permissionRequest,
					Allowed: ok && v1.GroupRole(user.GlobalRole.Int32) == permissionRequest.Role,
				}
			}
		}
		return err
	}); err != nil {
		ctx.L().Error("Failed to run transaction", zap.Error(err))
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return res, nil
}
