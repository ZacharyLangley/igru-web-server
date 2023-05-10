package garden

import (
	gocontext "context"
	"fmt"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	"github.com/ZacharyLangley/igru-web-server/pkg/database"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/garden"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/service/common"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateGarden(baseCtx gocontext.Context, req *connect.Request[v1.CreateGardenRequest]) (*connect.Response[v1.CreateGardenResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	res := connect.NewResponse(&v1.CreateGardenResponse{})
	// Check write access
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
	)
	if err != nil {
		return nil, err
	}
	var garden models.Garden
	if err = s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.CreateGardenParams{
			Name:          req.Msg.Name,
			GroupID:       database.NewFromNullableUUID(groupID),
			Comment:       req.Msg.Comment,
			Location:      req.Msg.Location,
			GrowType:      req.Msg.GrowType,
			GrowSize:      req.Msg.GrowSize,
			GrowStyle:     req.Msg.GrowStyle,
			ContainerSize: req.Msg.ContainerSize,
			Tags:          req.Msg.Tags,
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
		}
		garden, err = queries.CreateGarden(ctx, params)
		return fmt.Errorf("failed to create garden: %w", err)
	}); err != nil {
		return nil, fmt.Errorf("transaction failure: %w", err)
	}
	res.Msg.Garden = &v1.Garden{
		Id:            uuid.UUID(garden.ID.Bytes).String(),
		GroupId:       uuid.UUID(garden.GroupID.Bytes).String(),
		Name:          garden.Name,
		Comment:       garden.Comment,
		Location:      garden.Location,
		GrowType:      garden.GrowType,
		GrowSize:      garden.GrowSize,
		GrowStyle:     garden.GrowStyle,
		ContainerSize: garden.ContainerSize,
		Tags:          garden.Tags,
		CreatedAt:     timestamppb.New(garden.CreatedAt.Time),
	}
	return res, nil
}

func (s *Service) DeleteGarden(baseCtx gocontext.Context, req *connect.Request[v1.DeleteGardenRequest]) (*connect.Response[v1.DeleteGardenResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteGardenResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	gardenID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid garden id format: %w", err))
	}
	// Check write access
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
	)
	if err != nil {
		return nil, err
	}
	// Delete garden
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.DeleteGardenParams{
			ID:      database.NewFromUUID(gardenID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		return queries.DeleteGarden(ctx, params)
	}); err != nil {
		return nil, fmt.Errorf("transaction failure: %w", err)
	}
	return res, nil
}

func (s *Service) UpdateGarden(baseCtx gocontext.Context, req *connect.Request[v1.UpdateGardenRequest]) (*connect.Response[v1.UpdateGardenResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateGardenResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	gardenID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id format: %w", err))
	}
	// Check write access
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
	)
	if err != nil {
		return nil, err
	}
	var garden models.Garden
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		params := models.UpdateGardenParams{
			ID:            database.NewFromUUID(gardenID),
			GroupID:       database.NewFromNullableUUID(groupID),
			Name:          req.Msg.Name,
			Comment:       req.Msg.Comment,
			Location:      req.Msg.Location,
			GrowType:      req.Msg.GrowType,
			GrowSize:      req.Msg.GrowSize,
			GrowStyle:     req.Msg.GrowStyle,
			ContainerSize: req.Msg.ContainerSize,
			Tags:          req.Msg.Tags,
		}
		garden, err = query.UpdateGarden(ctx, params)
		return fmt.Errorf("failed to update garden: %w", err)
	}); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("failed transaction: %w", err))
	}
	res.Msg.Garden = &v1.Garden{
		Id:            uuid.UUID(garden.ID.Bytes).String(),
		GroupId:       uuid.UUID(garden.GroupID.Bytes).String(),
		Name:          garden.Name,
		Comment:       garden.Comment,
		Location:      garden.Location,
		GrowType:      garden.GrowType,
		GrowSize:      garden.GrowSize,
		GrowStyle:     garden.GrowStyle,
		ContainerSize: garden.ContainerSize,
		Tags:          garden.Tags,
		CreatedAt:     timestamppb.New(garden.CreatedAt.Time),
	}
	if garden.UpdatedAt.Valid {
		res.Msg.Garden.UpdatedAt = timestamppb.New(garden.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) GetGardens(baseCtx gocontext.Context, req *connect.Request[v1.GetGardensRequest]) (*connect.Response[v1.GetGardensResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetGardensResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	// Check read access
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
		authenticationv1.GroupRole_GROUP_ROLE_READ_ONLY,
	)
	if err != nil {
		return nil, err
	}
	// TODO: ZL | Pagination Functionality
	var gardens []models.Garden
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		gardens, err = queries.GetGardens(ctx, database.NewFromNullableUUID(groupID))
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("database error: %w", err))
	}
	res.Msg.Gardens = make([]*v1.Garden, 0, len(gardens))
	for _, garden := range gardens {
		newGarden := v1.Garden{
			Id:            uuid.UUID(garden.ID.Bytes).String(),
			GroupId:       groupID.UUID.String(),
			Name:          garden.Name,
			Comment:       garden.Comment,
			Location:      garden.Location,
			GrowType:      garden.GrowType,
			GrowSize:      garden.GrowSize,
			GrowStyle:     garden.GrowStyle,
			ContainerSize: garden.ContainerSize,
			Tags:          garden.Tags,
			CreatedAt:     timestamppb.New(garden.CreatedAt.Time),
		}
		if garden.UpdatedAt.Valid {
			newGarden.UpdatedAt = timestamppb.New(garden.UpdatedAt.Time)
		}
		res.Msg.Gardens = append(res.Msg.Gardens, &newGarden)
	}
	return res, nil
}

func (s *Service) GetGarden(baseCtx gocontext.Context, req *connect.Request[v1.GetGardenRequest]) (*connect.Response[v1.GetGardenResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetGardenResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	gardenID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid garden id format: %w", err))
	}
	// Check read access
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
		authenticationv1.GroupRole_GROUP_ROLE_READ_ONLY,
	)
	if err != nil {
		return nil, err
	}
	var garden models.Garden
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		params := models.GetGardenParams{
			ID:      database.NewFromUUID(gardenID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		garden, err = query.GetGarden(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Garden = &v1.Garden{
		Id:            uuid.UUID(garden.ID.Bytes).String(),
		Name:          garden.Name,
		Comment:       garden.Comment,
		Location:      garden.Location,
		GrowType:      garden.GrowType,
		GrowSize:      garden.GrowSize,
		GrowStyle:     garden.GrowStyle,
		ContainerSize: garden.ContainerSize,
		Tags:          garden.Tags,
		CreatedAt:     timestamppb.New(garden.CreatedAt.Time),
	}
	if garden.UpdatedAt.Valid {
		res.Msg.Garden.UpdatedAt = timestamppb.New(garden.UpdatedAt.Time)
	}
	return res, nil
}
