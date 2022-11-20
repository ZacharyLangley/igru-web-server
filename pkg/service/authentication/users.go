package authentication

import (
	gocontext "context"
	"database/sql"
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
	if err := validateCreateUserRequest(req.Msg); err != nil {
		return nil, err
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
	if err := validateDeleteUserRequest(req.Msg); err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to re-parse UUID: %w", err))
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
	if err := validateResetUserPasswordRequest(req.Msg); err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to re-parse UUID: %w", err))
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
	if err := validateUpdateUserRequest(req.Msg); err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.Msg.User.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to re-parse UUID: %w", err))
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
	if err := validatePaginationRequest(req.Msg); err != nil {
		return nil, err
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
