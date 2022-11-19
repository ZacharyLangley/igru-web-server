package authentication

import (
	gocontext "context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateUser(baseCtx gocontext.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateUserResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	hash, salt, err := generateCred(req.Msg.Password)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	var user models.User
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.CreateUserParams{
			Email: req.Msg.Email,
			Salt:  salt,
			Hash:  hash,
		}
		if req.Msg.FullName != "" {
			params.FullName = sql.NullString{Valid: true, String: req.Msg.FullName}
		}
		user, err = queries.CreateUser(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.User = &v1.User{
		Id:        user.ID.String(),
		Email:     user.Email,
		FullName:  user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
	return res, nil
}

func (s *Service) DeleteUser(baseCtx gocontext.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteUserResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if req.Msg != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		return queries.DeleteUser(ctx, userID)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) ResetUserPassword(baseCtx gocontext.Context, req *connect.Request[v1.ResetUserPasswordRequest]) (*connect.Response[v1.ResetUserPasswordResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.ResetUserPasswordResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if req.Msg != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	hash, salt, err := generateCred(req.Msg.Password)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.UpdateUserPasswordParams{
			ID:   userID,
			Salt: salt,
			Hash: hash,
		}
		return queries.UpdateUserPassword(ctx, params)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateUser(baseCtx gocontext.Context, req *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateUserResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	userID, err := uuid.Parse(req.Msg.User.Id)
	if req.Msg != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	var user models.User
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.UpdateUserParams{
			ID: userID,
		}
		if req.Msg.User.FullName != "" {
			params.FullName = sql.NullString{Valid: true, String: req.Msg.User.FullName}
		}
		user, err = queries.UpdateUser(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.User = &v1.User{
		Id:        user.ID.String(),
		Email:     user.Email,
		FullName:  user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
	return res, nil
}

func (s *Service) GetUsers(baseCtx gocontext.Context, req *connect.Request[v1.GetUsersRequest]) (*connect.Response[v1.GetUsersResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetUsersResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	if req.Msg.Pagination.Length <= 0 {
		req.Msg.Pagination.Length = 100
	}
	var users []models.User
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		params := models.GetUsersParams{
			Limit:  req.Msg.Pagination.Length,
			Offset: req.Msg.Pagination.Cursor,
		}
		users, err = queries.GetUsers(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Users = make([]*v1.User, 0, len(users))
	for _, user := range users {
		newUser := v1.User{
			Id:        user.ID.String(),
			Email:     user.Email,
			FullName:  user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
		}
		if user.UpdatedAt.Valid {
			newUser.UpdatedAt = timestamppb.New(user.UpdatedAt.Time)
		}
		res.Msg.Users = append(res.Msg.Users, &newUser)
	}
	return res, nil
}
