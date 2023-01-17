package garden

import (
	gocontext "context"
	"errors"
	"fmt"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/garden"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1"
	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateGarden(baseCtx gocontext.Context, req *connect_go.Request[v1.CreateGardenRequest]) (*connect_go.Response[v1.CreateGardenResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateGardenResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	allowed, err := s.checker.AssertWrite(ctx, req, req.Msg.GroupId)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
	}
	var garden models.Garden
	if err = s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.CreateGardenParams{
			Name:          req.Msg.Name,
			GroupID:       uuid.MustParse(req.Msg.GroupId),
			Comment:       req.Msg.Comment,
			Location:      req.Msg.Location,
			GrowType:      req.Msg.GrowType,
			GrowSize:      req.Msg.GrowSize,
			GrowStyle:     req.Msg.GrowStyle,
			ContainerSize: req.Msg.ContainerSize,
			Tags:          req.Msg.Tags,
			CreatedAt:     time.Now(),
		}
		garden, err = queries.CreateGarden(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Garden = &v1.Garden{
		Id:            garden.ID.String(),
		GroupId:       garden.GroupID.String(),
		Name:          garden.Name,
		Comment:       garden.Comment,
		Location:      garden.Location,
		GrowType:      garden.GrowType,
		GrowSize:      garden.GrowSize,
		GrowStyle:     garden.GrowStyle,
		ContainerSize: garden.ContainerSize,
		Tags:          garden.Tags,
		CreatedAt:     timestamppb.New(garden.CreatedAt),
	}
	return res, nil
}

func (s *Service) DeleteGarden(baseCtx gocontext.Context, req *connect_go.Request[v1.DeleteGardenRequest]) (*connect_go.Response[v1.DeleteGardenResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteGardenResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	gardenID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid garden id format: %w", err))
	}
	// Check write access
	groupID, err := s.resolveGardenGroupID(ctx, gardenID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id: %w", err))
	}
	allowed, err := s.checker.AssertWrite(ctx, req, groupID.String())
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		return queries.DeleteGarden(ctx, gardenID)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateGarden(baseCtx gocontext.Context, req *connect_go.Request[v1.UpdateGardenRequest]) (*connect_go.Response[v1.UpdateGardenResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateGardenResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	gardenID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id format: %w", err))
	}
	// Check write access
	groupID, err := s.resolveGardenGroupID(ctx, gardenID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid group id: %w", err))
	}
	allowed, err := s.checker.AssertWrite(ctx, req, groupID.String())
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
	}
	var garden models.Garden
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		params := models.UpdateGardenParams{
			ID:            gardenID,
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
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Garden = &v1.Garden{
		Id:            garden.ID.String(),
		Name:          garden.Name,
		Comment:       garden.Comment,
		Location:      garden.Location,
		GrowType:      garden.GrowType,
		GrowSize:      garden.GrowSize,
		GrowStyle:     garden.GrowStyle,
		ContainerSize: garden.ContainerSize,
		Tags:          garden.Tags,
		CreatedAt:     timestamppb.New(garden.CreatedAt),
	}
	if garden.UpdatedAt.Valid {
		res.Msg.Garden.UpdatedAt = timestamppb.New(garden.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) GetGardens(baseCtx gocontext.Context, req *connect_go.Request[v1.GetGardensRequest]) (*connect_go.Response[v1.GetGardensResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetGardensResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	allowed, err := s.checker.AssertRead(ctx, req, req.Msg.GroupId)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
	}
	// TODO: ZL | Pagination Functionality
	var gardens []models.Garden
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		gardens, err = queries.GetGardens(ctx)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Gardens = make([]*v1.Garden, 0, len(gardens))
	for _, garden := range gardens {
		newGarden := v1.Garden{
			Id:            garden.ID.String(),
			Name:          garden.Name,
			Comment:       garden.Comment,
			Location:      garden.Location,
			GrowType:      garden.GrowType,
			GrowSize:      garden.GrowSize,
			GrowStyle:     garden.GrowStyle,
			ContainerSize: garden.ContainerSize,
			Tags:          garden.Tags,
			CreatedAt:     timestamppb.New(garden.CreatedAt),
		}
		if garden.UpdatedAt.Valid {
			newGarden.UpdatedAt = timestamppb.New(garden.UpdatedAt.Time)
		}
		res.Msg.Gardens = append(res.Msg.Gardens, &newGarden)
	}
	return res, nil
}

func (s *Service) GetGarden(baseCtx gocontext.Context, req *connect_go.Request[v1.GetGardenRequest]) (*connect_go.Response[v1.GetGardenResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetGardenResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	gardenID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid garden id format: %w", err))
	}

	var garden models.Garden
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		garden, err = query.GetGarden(ctx, gardenID)
		return err
	}); err != nil {
		return nil, err
	}
	allowed, err := s.checker.AssertRead(ctx, req, garden.GroupID.String())
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("access denied"))
	}
	res.Msg.Garden = &v1.Garden{
		Id:            garden.ID.String(),
		Name:          garden.Name,
		Comment:       garden.Comment,
		Location:      garden.Location,
		GrowType:      garden.GrowType,
		GrowSize:      garden.GrowSize,
		GrowStyle:     garden.GrowStyle,
		ContainerSize: garden.ContainerSize,
		Tags:          garden.Tags,
		CreatedAt:     timestamppb.New(garden.CreatedAt),
	}
	if garden.UpdatedAt.Valid {
		res.Msg.Garden.UpdatedAt = timestamppb.New(garden.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) resolveGardenGroupID(ctx context.Context, gardenID uuid.UUID) (uuid.UUID, error) {
	var groupID uuid.UUID
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		garden, err := queries.GetGarden(ctx, gardenID)
		if err != nil {
			return err
		}
		groupID = garden.GroupID
		return nil
	}); err != nil {
		return uuid.UUID{}, err
	}
	return groupID, nil
}
