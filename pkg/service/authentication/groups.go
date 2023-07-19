package authentication

import (
	gocontext "context"
	"fmt"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/authentication"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	commonv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/common/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/common"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateGroup(baseCtx gocontext.Context, req *connect.Request[v1.CreateGroupRequest]) (*connect.Response[v1.CreateGroupResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateGroupResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	if err := validateCreateGroupRequest(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	var group models.Group
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		group, err = query.CreateGroup(ctx, req.Msg.Name)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Group = &v1.Group{
		Id:        uuid.UUID(group.ID.Bytes).String(),
		Name:      group.Name,
		CreatedAt: timestamppb.New(group.CreatedAt.Time),
	}
	return res, nil
}

func (s *Service) UpdateGroup(baseCtx gocontext.Context, req *connect.Request[v1.UpdateGroupRequest]) (*connect.Response[v1.UpdateGroupResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateGroupResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	if err := validateUpdateGroupRequest(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	groupID, err := uuid.Parse(req.Msg.Group.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	var group models.Group
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		params := models.UpdateGroupParams{
			ID:   database.NewFromUUID(groupID),
			Name: req.Msg.Group.Name,
		}
		group, err = query.UpdateGroup(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Group = &v1.Group{
		Id:        uuid.UUID(group.ID.Bytes).String(),
		Name:      group.Name,
		CreatedAt: timestamppb.New(group.CreatedAt.Time),
	}
	if group.UpdatedAt.Valid {
		res.Msg.Group.UpdatedAt = timestamppb.New(group.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) DeleteGroup(baseCtx gocontext.Context, req *connect.Request[v1.DeleteGroupRequest]) (*connect.Response[v1.DeleteGroupResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteGroupResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	if err := validateDeleteGroupRequest(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	groupID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		return query.DeleteGroup(ctx, database.NewFromUUID(groupID))
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetGroup(baseCtx gocontext.Context, req *connect.Request[v1.GetGroupRequest]) (*connect.Response[v1.GetGroupResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetGroupResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	groupID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	var group models.Group
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		group, err = query.GetGroup(ctx, database.NewFromUUID(groupID))
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Group = &v1.Group{
		Id:        uuid.UUID(group.ID.Bytes).String(),
		Name:      group.Name,
		CreatedAt: timestamppb.New(group.CreatedAt.Time),
	}
	if group.UpdatedAt.Valid {
		res.Msg.Group.UpdatedAt = timestamppb.New(group.UpdatedAt.Time)
	}
	return res, nil
}

var defaultPagination = &commonv1.PaginationRequest{
	Length: defaultPageLength,
}

func (s *Service) GetGroups(baseCtx gocontext.Context, req *connect.Request[v1.GetGroupsRequest]) (*connect.Response[v1.GetGroupsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetGroupsResponse{})
	if err := validatePaginationRequest(req.Msg); err != nil {
		return nil, err
	}
	if req.Msg.Pagination == nil {
		req.Msg.Pagination = defaultPagination
	}

	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		var groups []models.Group
		var err error
		if req.Msg.IncludeUserGroups {
			groups, err = query.GetAllGroups(ctx, models.GetAllGroupsParams{
				Limit:  req.Msg.Pagination.Length,
				Offset: req.Msg.Pagination.Cursor,
			})
		} else {
			groups, err = query.GetGroups(ctx, models.GetGroupsParams{
				Limit:  req.Msg.Pagination.Length,
				Offset: req.Msg.Pagination.Cursor,
			})
		}
		if err != nil {
			return fmt.Errorf("failed to get groups: %w", err)
		}
		res.Msg.Groups = make([]*v1.Group, 0, len(groups))
		for _, group := range groups {
			newGroup := v1.Group{
				Id:        uuid.UUID(group.ID.Bytes).String(),
				Name:      group.Name,
				CreatedAt: timestamppb.New(group.CreatedAt.Time),
			}
			if group.UpdatedAt.Valid {
				newGroup.UpdatedAt = timestamppb.New(group.UpdatedAt.Time)
			}
			newGroup.NumMembers, err = query.CountGroupMembers(ctx, group.ID)
			if err != nil {
				return fmt.Errorf("failed to count group members: %w", err)
			}
			res.Msg.Groups = append(res.Msg.Groups, &newGroup)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) AddGroupMember(baseCtx gocontext.Context, req *connect.Request[v1.AddGroupMemberRequest]) (*connect.Response[v1.AddGroupMemberResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.AddGroupMemberResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	groupID, err := uuid.Parse(req.Msg.GroupId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		return query.AddGroupMember(ctx, models.AddGroupMemberParams{
			UserID:  database.NewFromUUID(userID),
			GroupID: database.NewFromUUID(groupID),
			Role:    int32(v1.GroupRole_GROUP_ROLE_UNSPECIFIED),
		})
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateGroupMember(baseCtx gocontext.Context, req *connect.Request[v1.UpdateGroupMemberRequest]) (*connect.Response[v1.UpdateGroupMemberResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateGroupMemberResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	groupID, err := uuid.Parse(req.Msg.GroupId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		return query.UpdateGroupMember(ctx, models.UpdateGroupMemberParams{
			UserID:  database.NewFromUUID(userID),
			GroupID: database.NewFromUUID(groupID),
			Role:    int32(req.Msg.Role),
		})
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) RemoveGroupMember(baseCtx gocontext.Context, req *connect.Request[v1.RemoveGroupMemberRequest]) (*connect.Response[v1.RemoveGroupMemberResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.RemoveGroupMemberResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	groupID, err := uuid.Parse(req.Msg.GroupId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		return query.DeleteGroupMember(ctx, models.DeleteGroupMemberParams{
			UserID:  database.NewFromUUID(userID),
			GroupID: database.NewFromUUID(groupID),
		})
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetUserGroups(baseCtx gocontext.Context, req *connect.Request[v1.GetUserGroupsRequest]) (*connect.Response[v1.GetUserGroupsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetUserGroupsResponse{})
	if err := validatePaginationRequest(req.Msg); err != nil {
		return nil, err
	}
	if req.Msg.Pagination == nil {
		req.Msg.Pagination = defaultPagination
	}
	userID, err := uuid.Parse(req.Msg.UserId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		groups, err := query.GetUserGroups(ctx, models.GetUserGroupsParams{
			Limit:  req.Msg.Pagination.Length,
			Offset: req.Msg.Pagination.Cursor,
			UserID: database.NewFromUUID(userID),
		})
		if err != nil {
			return fmt.Errorf("failed to get user groups: %w", err)
		}
		res.Msg.Groups = make([]*v1.Group, 0, len(groups))
		for _, group := range groups {
			newGroup := v1.Group{
				Id:        uuid.UUID(group.ID.Bytes).String(),
				Name:      group.Name,
				CreatedAt: timestamppb.New(group.CreatedAt.Time),
			}
			if group.UpdatedAt.Valid {
				newGroup.UpdatedAt = timestamppb.New(group.UpdatedAt.Time)
			}
			newGroup.NumMembers, err = query.CountGroupMembers(ctx, group.ID)
			if err != nil {
				return fmt.Errorf("failed to count group members: %w", err)
			}
			res.Msg.Groups = append(res.Msg.Groups, &newGroup)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}
