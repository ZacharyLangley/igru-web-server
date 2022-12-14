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

func (s *Service) CreateStrain(baseCtx gocontext.Context, req *connect_go.Request[v1.CreateStrainRequest]) (*connect_go.Response[v1.CreateStrainResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateStrainResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	parentageID, err := uuid.Parse(req.Msg.Parentage)

	var strain models.Strain
	if err = s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.CreateStrainParams{
			Name:       req.Msg.Name,
			Comment:    req.Msg.Comment,
			Notes:      req.Msg.Notes,
			Type:       req.Msg.Type,
			Price:      req.Msg.Price,
			ThcPercent: req.Msg.ThcPercent,
			CbdPercent: req.Msg.CbdPercent,
			Parentage:  parentageID,
			Aroma:      req.Msg.Aroma,
			Taste:      req.Msg.Taste,
			Tags:       req.Msg.Tags,
			CreatedAt:  time.Now(),
		}
		strain, err = queries.CreateStrain(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Strain = &v1.Strain{
		Id:         strain.ID.String(),
		Name:       strain.Name,
		Comment:    strain.Comment,
		Notes:      strain.Notes,
		Type:       strain.Type,
		Price:      strain.Price,
		ThcPercent: strain.ThcPercent,
		CbdPercent: strain.CbdPercent,
		Parentage:  strain.Parentage.String(),
		Aroma:      strain.Aroma,
		Taste:      strain.Taste,
		Tags:       strain.Tags,
		CreatedAt:  timestamppb.New(strain.CreatedAt),
	}
	return res, nil
}

func (s *Service) DeleteStrain(baseCtx gocontext.Context, req *connect_go.Request[v1.DeleteStrainRequest]) (*connect_go.Response[v1.DeleteStrainResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteStrainResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	strainID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid strain id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		return queries.DeleteStrain(ctx, strainID)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateStrain(baseCtx gocontext.Context, req *connect_go.Request[v1.UpdateStrainRequest]) (*connect_go.Response[v1.UpdateStrainResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateStrainResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	strainID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	parentageID, err := uuid.Parse(req.Msg.Parentage)
	var strain models.Strain
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		params := models.UpdateStrainParams{
			ID:         strainID,
			Name:       req.Msg.Name,
			Comment:    req.Msg.Comment,
			Notes:      req.Msg.Notes,
			Type:       req.Msg.Type,
			Price:      req.Msg.Price,
			ThcPercent: req.Msg.ThcPercent,
			CbdPercent: req.Msg.CbdPercent,
			Parentage:  parentageID,
			Aroma:      req.Msg.Aroma,
			Taste:      req.Msg.Taste,
			Tags:       req.Msg.Tags,
		}
		strain, err = query.UpdateStrain(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Strain = &v1.Strain{
		Id:         strain.ID.String(),
		Name:       strain.Name,
		Comment:    strain.Comment,
		Notes:      strain.Notes,
		Type:       strain.Type,
		Price:      strain.Price,
		ThcPercent: strain.ThcPercent,
		CbdPercent: strain.CbdPercent,
		Parentage:  strain.Parentage.String(),
		Aroma:      strain.Aroma,
		Taste:      strain.Taste,
		Tags:       strain.Tags,
	}
	if strain.UpdatedAt.Valid {
		res.Msg.Strain.UpdatedAt = timestamppb.New(strain.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) GetStrains(baseCtx gocontext.Context, req *connect_go.Request[v1.GetStrainsRequest]) (*connect_go.Response[v1.GetStrainsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetStrainsResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	// TODO: ZL | Pagination Functionality
	var strains []models.Strain
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		strains, err = queries.GetStrains(ctx)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Strains = make([]*v1.Strain, 0, len(strains))
	for _, strain := range strains {
		newStrain := v1.Strain{
			Id:         strain.ID.String(),
			Name:       strain.Name,
			Comment:    strain.Comment,
			Notes:      strain.Notes,
			Type:       strain.Type,
			Price:      strain.Price,
			ThcPercent: strain.ThcPercent,
			CbdPercent: strain.CbdPercent,
			Parentage:  strain.Parentage.String(),
			Aroma:      strain.Aroma,
			Taste:      strain.Taste,
			Tags:       strain.Tags,
			CreatedAt:  timestamppb.New(strain.CreatedAt),
		}
		if strain.UpdatedAt.Valid {
			newStrain.UpdatedAt = timestamppb.New(strain.UpdatedAt.Time)
		}
		res.Msg.Strains = append(res.Msg.Strains, &newStrain)
	}
	return res, nil
}

func (s *Service) GetStrain(baseCtx gocontext.Context, req *connect_go.Request[v1.GetStrainRequest]) (*connect_go.Response[v1.GetStrainResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetStrainResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
	}
	strainID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid strain id format: %w", err))
	}

	var strain models.Strain
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		strain, err = query.GetStrain(ctx, strainID)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Strain = &v1.Strain{
		Id:         strain.ID.String(),
		Name:       strain.Name,
		Comment:    strain.Comment,
		Notes:      strain.Notes,
		Type:       strain.Type,
		Price:      strain.Price,
		ThcPercent: strain.ThcPercent,
		CbdPercent: strain.CbdPercent,
		Parentage:  strain.Parentage.String(),
		Aroma:      strain.Aroma,
		Taste:      strain.Taste,
		Tags:       strain.Tags,
		CreatedAt:  timestamppb.New(strain.CreatedAt),
	}
	if strain.UpdatedAt.Valid {
		res.Msg.Strain.UpdatedAt = timestamppb.New(strain.UpdatedAt.Time)
	}
	return res, nil
}
