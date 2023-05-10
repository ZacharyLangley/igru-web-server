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

func (s *Service) CreatePlant(baseCtx gocontext.Context, req *connect.Request[v1.CreatePlantRequest]) (*connect.Response[v1.CreatePlantResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreatePlantResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
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
	parentageID, err := uuid.Parse(req.Msg.Parentage)
	var plant models.Plant
	if err = s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
		params := models.CreatePlantParams{
			Name:            req.Msg.Name,
			GroupID:         database.NewFromNullableUUID(groupID),
			Comment:         req.Msg.Comment,
			Notes:           req.Msg.Notes,
			GrowCycleLength: req.Msg.GrowCycleLength,
			Parentage:       database.NewFromUUID(parentageID),
			Origin:          req.Msg.Origin,
			Gender:          req.Msg.Gender,
			DaysFlowering:   req.Msg.DaysFlowering,
			DaysCured:       req.Msg.DaysCured,
			HarvestedWeight: req.Msg.HarvestedWeight,
			BudDensity:      req.Msg.BudDensity,
			BudPistils:      req.Msg.BudPistils,
			Tags:            req.Msg.Tags,
			AcquiredAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
		}
		plant, err = queries.CreatePlant(ctx, params)
		return fmt.Errorf("failed to create plant: %w", err)
	}); err != nil {
		return nil, fmt.Errorf("transaction failure: %w", err)
	}
	// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
	res.Msg.Plant = &v1.Plant{
		Id:              uuid.UUID(plant.ID.Bytes).String(),
		GroupId:         uuid.UUID(plant.GroupID.Bytes).String(),
		Name:            plant.Name,
		Comment:         plant.Comment,
		Notes:           plant.Notes,
		GrowCycleLength: plant.GrowCycleLength,
		Parentage:       uuid.UUID(plant.Parentage.Bytes).String(),
		Origin:          plant.Origin,
		Gender:          plant.Gender,
		DaysFlowering:   plant.DaysFlowering,
		DaysCured:       plant.DaysCured,
		HarvestedWeight: plant.HarvestedWeight,
		BudDensity:      plant.BudDensity,
		BudPistils:      plant.BudPistils,
		Tags:            plant.Tags,
		AcquiredAt:      timestamppb.New(plant.AcquiredAt.Time),
		CreatedAt:       timestamppb.New(plant.CreatedAt.Time),
	}
	return res, nil
}

func (s *Service) DeletePlant(baseCtx gocontext.Context, req *connect.Request[v1.DeletePlantRequest]) (*connect.Response[v1.DeletePlantResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeletePlantResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	// Check write access
	var err error
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
	)
	if err != nil {
		return nil, err
	}
	plantID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid plant id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.DeletePlantParams{
			ID:      database.NewFromUUID(plantID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		return queries.DeletePlant(ctx, params)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdatePlant(baseCtx gocontext.Context, req *connect.Request[v1.UpdatePlantRequest]) (*connect.Response[v1.UpdatePlantResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdatePlantResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	// Check write access
	var err error
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
	)
	if err != nil {
		return nil, err
	}
	// TODO: Create validation function
	plantID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	parentageID, err := uuid.Parse(req.Msg.Parentage)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid parentage id format: %w", err))
	}
	var plant models.Plant
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
		params := models.UpdatePlantParams{
			ID:              database.NewFromUUID(plantID),
			GroupID:         database.NewFromNullableUUID(groupID),
			Name:            req.Msg.Name,
			Comment:         req.Msg.Comment,
			Notes:           req.Msg.Notes,
			GrowCycleLength: req.Msg.GrowCycleLength,
			Parentage:       database.NewFromUUID(parentageID),
			Origin:          req.Msg.Origin,
			Gender:          req.Msg.Gender,
			DaysFlowering:   req.Msg.DaysFlowering,
			DaysCured:       req.Msg.DaysCured,
			HarvestedWeight: req.Msg.HarvestedWeight,
			BudDensity:      req.Msg.BudDensity,
			BudPistils:      req.Msg.BudPistils,
			Tags:            req.Msg.Tags,
		}
		plant, err = query.UpdatePlant(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
	res.Msg.Plant = &v1.Plant{
		Id:              uuid.UUID(plant.ID.Bytes).String(),
		GroupId:         uuid.UUID(plant.GroupID.Bytes).String(),
		Name:            plant.Name,
		Comment:         plant.Comment,
		Notes:           plant.Notes,
		GrowCycleLength: plant.GrowCycleLength,
		Parentage:       uuid.UUID(plant.Parentage.Bytes).String(),
		Origin:          plant.Origin,
		Gender:          plant.Gender,
		DaysFlowering:   plant.DaysFlowering,
		DaysCured:       plant.DaysCured,
		HarvestedWeight: plant.HarvestedWeight,
		BudDensity:      plant.BudDensity,
		BudPistils:      plant.BudPistils,
		Tags:            plant.Tags,
		CreatedAt:       timestamppb.New(plant.CreatedAt.Time),
	}
	if plant.UpdatedAt.Valid {
		res.Msg.Plant.UpdatedAt = timestamppb.New(plant.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) GetPlants(baseCtx gocontext.Context, req *connect.Request[v1.GetPlantsRequest]) (*connect.Response[v1.GetPlantsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetPlantsResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	// Check write access
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
	var plants []models.Plant
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		plants, err = queries.GetPlants(ctx, database.NewFromNullableUUID(groupID))
		return err
	}); err != nil {
		return nil, err
	}
	// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
	res.Msg.Plants = make([]*v1.Plant, 0, len(plants))
	for _, plant := range plants {
		newPlant := v1.Plant{
			Id:              uuid.UUID(plant.ID.Bytes).String(),
			GroupId:         uuid.UUID(plant.GroupID.Bytes).String(),
			Name:            plant.Name,
			Comment:         plant.Comment,
			Notes:           plant.Notes,
			GrowCycleLength: plant.GrowCycleLength,
			Parentage:       uuid.UUID(plant.Parentage.Bytes).String(),
			Origin:          plant.Origin,
			Gender:          plant.Gender,
			DaysFlowering:   plant.DaysFlowering,
			DaysCured:       plant.DaysCured,
			HarvestedWeight: plant.HarvestedWeight,
			BudDensity:      plant.BudDensity,
			BudPistils:      plant.BudPistils,
			Tags:            plant.Tags,
			CreatedAt:       timestamppb.New(plant.CreatedAt.Time),
		}
		if plant.UpdatedAt.Valid {
			newPlant.UpdatedAt = timestamppb.New(plant.UpdatedAt.Time)
		}
		res.Msg.Plants = append(res.Msg.Plants, &newPlant)
	}
	return res, nil
}

func (s *Service) GetPlant(baseCtx gocontext.Context, req *connect.Request[v1.GetPlantRequest]) (*connect.Response[v1.GetPlantResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetPlantResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	plantID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid plant id format: %w", err))
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
	var plant models.Plant
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		params := models.GetPlantParams{
			ID:      database.NewFromUUID(plantID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		plant, err = query.GetPlant(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
	res.Msg.Plant = &v1.Plant{
		Id:              uuid.UUID(plant.ID.Bytes).String(),
		GroupId:         uuid.UUID(plant.GroupID.Bytes).String(),
		Name:            plant.Name,
		Comment:         plant.Comment,
		Notes:           plant.Notes,
		GrowCycleLength: plant.GrowCycleLength,
		Parentage:       uuid.UUID(plant.Parentage.Bytes).String(),
		Origin:          plant.Origin,
		Gender:          plant.Gender,
		DaysFlowering:   plant.DaysFlowering,
		DaysCured:       plant.DaysCured,
		HarvestedWeight: plant.HarvestedWeight,
		BudDensity:      plant.BudDensity,
		BudPistils:      plant.BudPistils,
		Tags:            plant.Tags,
		AcquiredAt:      timestamppb.New(plant.AcquiredAt.Time),
		CreatedAt:       timestamppb.New(plant.CreatedAt.Time),
	}
	if plant.UpdatedAt.Valid {
		res.Msg.Plant.UpdatedAt = timestamppb.New(plant.UpdatedAt.Time)
	}
	return res, nil
}
