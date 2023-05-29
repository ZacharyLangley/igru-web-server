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

func (s *Service) CreateStrain(baseCtx gocontext.Context, req *connect.Request[v1.CreateStrainRequest]) (*connect.Response[v1.CreateStrainResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateStrainResponse{})
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
	// TODO: Add validation function
	parentageID, err := uuid.Parse(req.Msg.Parentage)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid parentage id format: %w", err))
	}
	var strain models.Strain
	if err = s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.CreateStrainParams{
			Name:       req.Msg.Name,
			GroupID:    database.NewFromNullableUUID(groupID),
			Comment:    req.Msg.Comment,
			Notes:      req.Msg.Notes,
			Type:       req.Msg.Type,
			Price:      req.Msg.Price,
			ThcPercent: req.Msg.ThcPercent,
			CbdPercent: req.Msg.CbdPercent,
			Parentage:  database.NewFromUUID(parentageID),
			Aroma:      req.Msg.Aroma,
			Taste:      req.Msg.Taste,
			Tags:       req.Msg.Tags,
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
		}
		strain, err = queries.CreateStrain(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Strain = &v1.Strain{
		Id:         uuid.UUID(strain.ID.Bytes).String(),
		GroupId:    uuid.UUID(strain.GroupID.Bytes).String(),
		Name:       strain.Name,
		Comment:    strain.Comment,
		Notes:      strain.Notes,
		Type:       strain.Type,
		Price:      strain.Price,
		ThcPercent: strain.ThcPercent,
		CbdPercent: strain.CbdPercent,
		Parentage:  uuid.UUID(strain.Parentage.Bytes).String(),
		Aroma:      strain.Aroma,
		Taste:      strain.Taste,
		Tags:       strain.Tags,
		CreatedAt:  timestamppb.New(strain.CreatedAt.Time),
	}
	return res, nil
}

func (s *Service) DeleteStrain(baseCtx gocontext.Context, req *connect.Request[v1.DeleteStrainRequest]) (*connect.Response[v1.DeleteStrainResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteStrainResponse{})
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
	strainID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid strain id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.DeleteStrainParams{
			ID:      database.NewFromUUID(strainID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		return queries.DeleteStrain(ctx, params)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateStrain(baseCtx gocontext.Context, req *connect.Request[v1.UpdateStrainRequest]) (*connect.Response[v1.UpdateStrainResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateStrainResponse{})
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
	// TODO: Add validation functions
	strainID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	parentageID, err := uuid.Parse(req.Msg.Parentage)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid parentage id format: %w", err))
	}
	var strain models.Strain
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		params := models.UpdateStrainParams{
			ID:         database.NewFromUUID(strainID),
			GroupID:    database.NewFromNullableUUID(groupID),
			Name:       req.Msg.Name,
			Comment:    req.Msg.Comment,
			Notes:      req.Msg.Notes,
			Type:       req.Msg.Type,
			Price:      req.Msg.Price,
			ThcPercent: req.Msg.ThcPercent,
			CbdPercent: req.Msg.CbdPercent,
			Parentage:  database.NewFromUUID(parentageID),
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
		Id:         uuid.UUID(strain.ID.Bytes).String(),
		GroupId:    uuid.UUID(strain.GroupID.Bytes).String(),
		Name:       strain.Name,
		Comment:    strain.Comment,
		Notes:      strain.Notes,
		Type:       strain.Type,
		Price:      strain.Price,
		ThcPercent: strain.ThcPercent,
		CbdPercent: strain.CbdPercent,
		Parentage:  uuid.UUID(strain.Parentage.Bytes).String(),
		Aroma:      strain.Aroma,
		Taste:      strain.Taste,
		Tags:       strain.Tags,
	}
	if strain.UpdatedAt.Valid {
		res.Msg.Strain.UpdatedAt = timestamppb.New(strain.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) GetStrains(baseCtx gocontext.Context, req *connect.Request[v1.GetStrainsRequest]) (*connect.Response[v1.GetStrainsResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetStrainsResponse{})
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
	var strains []models.Strain
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		strains, err = queries.GetStrains(ctx, database.NewFromNullableUUID(groupID))
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Strains = make([]*v1.Strain, 0, len(strains))
	for _, strain := range strains {
		newStrain := v1.Strain{
			Id:         uuid.UUID(strain.ID.Bytes).String(),
			GroupId:    uuid.UUID(strain.GroupID.Bytes).String(),
			Name:       strain.Name,
			Comment:    strain.Comment,
			Notes:      strain.Notes,
			Type:       strain.Type,
			Price:      strain.Price,
			ThcPercent: strain.ThcPercent,
			CbdPercent: strain.CbdPercent,
			Parentage:  uuid.UUID(strain.Parentage.Bytes).String(),
			Aroma:      strain.Aroma,
			Taste:      strain.Taste,
			Tags:       strain.Tags,
			CreatedAt:  timestamppb.New(strain.CreatedAt.Time),
		}
		if strain.UpdatedAt.Valid {
			newStrain.UpdatedAt = timestamppb.New(strain.UpdatedAt.Time)
		}
		res.Msg.Strains = append(res.Msg.Strains, &newStrain)
	}
	return res, nil
}

func (s *Service) GetStrain(baseCtx gocontext.Context, req *connect.Request[v1.GetStrainRequest]) (*connect.Response[v1.GetStrainResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetStrainResponse{})
	if err := common.CheckMessage(req); err != nil {
		return nil, err
	}
	// Check read access
	var err error
	groupID, err := s.checker.AssertAny(ctx,
		req,
		&req.Msg.GroupId,
		authenticationv1.GroupRole_GROUP_ROLE_ADMIN,
		authenticationv1.GroupRole_GROUP_ROLE_READ_ONLY,
	)
	if err != nil {
		return nil, err
	}
	strainID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid strain id format: %w", err))
	}

	var strain models.Strain
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		params := models.GetStrainParams{
			ID:      database.NewFromUUID(strainID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		strain, err = query.GetStrain(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Strain = &v1.Strain{
		Id:         uuid.UUID(strain.ID.Bytes).String(),
		GroupId:    uuid.UUID(strain.GroupID.Bytes).String(),
		Name:       strain.Name,
		Comment:    strain.Comment,
		Notes:      strain.Notes,
		Type:       strain.Type,
		Price:      strain.Price,
		ThcPercent: strain.ThcPercent,
		CbdPercent: strain.CbdPercent,
		Parentage:  uuid.UUID(strain.Parentage.Bytes).String(),
		Aroma:      strain.Aroma,
		Taste:      strain.Taste,
		Tags:       strain.Tags,
		CreatedAt:  timestamppb.New(strain.CreatedAt.Time),
	}
	if strain.UpdatedAt.Valid {
		res.Msg.Strain.UpdatedAt = timestamppb.New(strain.UpdatedAt.Time)
	}
	return res, nil
}
