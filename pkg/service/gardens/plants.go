package gardens

import (
	gocontext "context"
	"errors"
	"fmt"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/gardens"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreatePlant(baseCtx gocontext.Context, req *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreatePlantResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	var plant models.Plant
	if err = s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
		params := models.CreatePlantParams{
			Name:            req.Msg.Name,
			Comment:         req.Msg.Comment,
			Notes:           req.Msg.Notes,
			GrowCycleLength: req.Msg.GrowCycleLength,
			Parentage:       req.Msg.Parentage,
			Origin:          req.Msg.Origin,
			Gender:          req.Msg.Gender,
			DaysFlowering:   req.Msg.DaysFlowering,
			DaysCured:       req.Msg.DaysCured,
			HarvestedWeight: req.Msg.HarvestedWeight,
			BudDensity:      req.Msg.BudDensity,
			BudPistils:      req.Msg.BudPistils,
			Tags:            req.Msg.Tags,
			AcquiredAt:      time.Now(),
			CreatedAt:       time.Now(),
		}
		plant, err = queries.CreatePlant(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
	res.Msg.Plant = &v1.Plant{
		Id:              plant.ID.String(),
		Name:            plant.Name,
		Comment:         plant.Comment,
		Notes:           plant.Notes,
		GrowCycleLength: plant.GrowCycleLength,
		Parentage:       plant.Parentage,
		Origin:          plant.Origin,
		Gender:          plant.Gender,
		DaysFlowering:   plant.DaysFlowering,
		DaysCured:       plant.DaysCured,
		HarvestedWeight: plant.HarvestedWeight,
		BudDensity:      plant.BudDensity,
		BudPistils:      plant.BudPistils,
		Tags:            plant.Tags,
		AcquiredAt:      timestamppb.New(plant.AcquiredAt),
		CreatedAt:       timestamppb.New(plant.CreatedAt),
	}
	return res, nil
}

func (s *Service) DeletePlant(baseCtx gocontext.Context, req *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeletePlantResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	plantID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid plant id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		return queries.DeletePlant(ctx, plantID)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdatePlant(baseCtx gocontext.Context, req *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdatePlantResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	plantID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	var plant models.Plant
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
		params := models.UpdatePlantParams{
			ID:              plantID,
			Name:            req.Msg.Name,
			Comment:         req.Msg.Comment,
			Notes:           req.Msg.Notes,
			GrowCycleLength: req.Msg.GrowCycleLength,
			Parentage:       req.Msg.Parentage,
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
		Id:              plant.ID.String(),
		Comment:         plant.Comment,
		Notes:           plant.Notes,
		GrowCycleLength: plant.GrowCycleLength,
		Parentage:       plant.Parentage,
		Origin:          plant.Origin,
		Gender:          plant.Gender,
		DaysFlowering:   plant.DaysFlowering,
		DaysCured:       plant.DaysCured,
		HarvestedWeight: plant.HarvestedWeight,
		BudDensity:      plant.BudDensity,
		BudPistils:      plant.BudPistils,
		Tags:            plant.Tags,
		CreatedAt:       timestamppb.New(plant.CreatedAt),
	}
	if plant.UpdatedAt.Valid {
		res.Msg.Plant.UpdatedAt = timestamppb.New(plant.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) GetPlants(baseCtx gocontext.Context, req *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetPlantsResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	// TODO: ZL | Pagination Functionality
	var plants []models.Plant
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		plants, err = queries.GetPlants(ctx)
		return err
	}); err != nil {
		return nil, err
	}
	// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
	res.Msg.Plants = make([]*v1.Plant, 0, len(plants))
	for _, plant := range plants {
		newPlant := v1.Plant{
			Id:              plant.ID.String(),
			Name:            plant.Name,
			Comment:         plant.Comment,
			Notes:           plant.Notes,
			GrowCycleLength: plant.GrowCycleLength,
			Parentage:       plant.Parentage,
			Origin:          plant.Origin,
			Gender:          plant.Gender,
			DaysFlowering:   plant.DaysFlowering,
			DaysCured:       plant.DaysCured,
			HarvestedWeight: plant.HarvestedWeight,
			BudDensity:      plant.BudDensity,
			BudPistils:      plant.BudPistils,
			Tags:            plant.Tags,
			CreatedAt:       timestamppb.New(plant.CreatedAt),
		}
		if plant.UpdatedAt.Valid {
			newPlant.UpdatedAt = timestamppb.New(plant.UpdatedAt.Time)
		}
		res.Msg.Plants = append(res.Msg.Plants, &newPlant)
	}
	return res, nil
}

func (s *Service) GetPlant(baseCtx gocontext.Context, req *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetPlantResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	plantID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid plant id format: %w", err))
	}

	var plant models.Plant
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		plant, err = query.GetPlant(ctx, plantID)
		return err
	}); err != nil {
		return nil, err
	}
	// TODO: ZL | ADD AcquiredAt Data Point for when it was obtained vs when the data is made.
	res.Msg.Plant = &v1.Plant{
		Id:              plant.ID.String(),
		Name:            plant.Name,
		Comment:         plant.Comment,
		Notes:           plant.Notes,
		GrowCycleLength: plant.GrowCycleLength,
		Parentage:       plant.Parentage,
		Origin:          plant.Origin,
		Gender:          plant.Gender,
		DaysFlowering:   plant.DaysFlowering,
		DaysCured:       plant.DaysCured,
		HarvestedWeight: plant.HarvestedWeight,
		BudDensity:      plant.BudDensity,
		BudPistils:      plant.BudPistils,
		Tags:            plant.Tags,
		AcquiredAt:      timestamppb.New(plant.AcquiredAt),
		CreatedAt:       timestamppb.New(plant.CreatedAt),
	}
	if plant.UpdatedAt.Valid {
		res.Msg.Plant.UpdatedAt = timestamppb.New(plant.UpdatedAt.Time)
	}
	return res, nil
}
