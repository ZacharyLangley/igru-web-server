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

func (s *Service) CreateRecipe(baseCtx gocontext.Context, req *connect.Request[v1.CreateRecipeRequest]) (*connect.Response[v1.CreateRecipeResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateRecipeResponse{})
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

	var recipe models.Recipe
	if err = s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.CreateRecipeParams{
			Name:                req.Msg.Name,
			GroupID:             database.NewFromNullableUUID(groupID),
			Comment:             req.Msg.Comment,
			Ingredients:         req.Msg.Ingredients,
			Instructions:        req.Msg.Instructions,
			Ph:                  req.Msg.Ph,
			MixTimeMilliseconds: req.Msg.MixTimeMilliseconds,
			Tags:                req.Msg.Tags,
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
		}
		recipe, err = queries.CreateRecipe(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Recipe = &v1.Recipe{
		Id:                  uuid.UUID(recipe.ID.Bytes).String(),
		GroupId:             uuid.UUID(recipe.GroupID.Bytes).String(),
		Name:                recipe.Name,
		Comment:             recipe.Comment,
		Ingredients:         recipe.Ingredients,
		Instructions:        recipe.Instructions,
		Ph:                  recipe.Ph,
		MixTimeMilliseconds: recipe.MixTimeMilliseconds,
		Tags:                recipe.Tags,
		CreatedAt:           timestamppb.New(recipe.CreatedAt.Time),
	}
	return res, nil
}

func (s *Service) DeleteRecipe(baseCtx gocontext.Context, req *connect.Request[v1.DeleteRecipeRequest]) (*connect.Response[v1.DeleteRecipeResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteRecipeResponse{})
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
	recipeID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid recipe id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.DeleteRecipeParams{
			ID:      database.NewFromUUID(recipeID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		return queries.DeleteRecipe(ctx, params)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateRecipe(baseCtx gocontext.Context, req *connect.Request[v1.UpdateRecipeRequest]) (*connect.Response[v1.UpdateRecipeResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateRecipeResponse{})
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
	recipeID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	var recipe models.Recipe
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		params := models.UpdateRecipeParams{
			ID:                  database.NewFromUUID(recipeID),
			GroupID:             database.NewFromNullableUUID(groupID),
			Name:                req.Msg.Name,
			Comment:             req.Msg.Comment,
			Ingredients:         req.Msg.Ingredients,
			Instructions:        req.Msg.Instructions,
			Ph:                  req.Msg.Ph,
			MixTimeMilliseconds: req.Msg.MixTimeMilliseconds,
			Tags:                req.Msg.Tags,
		}
		recipe, err = query.UpdateRecipe(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}
	res.Msg.Recipe = &v1.Recipe{
		Id:                  uuid.UUID(recipe.ID.Bytes).String(),
		GroupId:             uuid.UUID(recipe.GroupID.Bytes).String(),
		Name:                recipe.Name,
		Comment:             recipe.Comment,
		Ingredients:         recipe.Ingredients,
		Instructions:        recipe.Instructions,
		Ph:                  recipe.Ph,
		MixTimeMilliseconds: recipe.MixTimeMilliseconds,
		Tags:                recipe.Tags,
	}
	if recipe.UpdatedAt.Valid {
		res.Msg.Recipe.UpdatedAt = timestamppb.New(recipe.UpdatedAt.Time)
	}
	return res, nil
}

func (s *Service) GetRecipes(baseCtx gocontext.Context, req *connect.Request[v1.GetRecipesRequest]) (*connect.Response[v1.GetRecipesResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetRecipesResponse{})
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
	var recipes []models.Recipe
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		queries := models.New(tx)
		recipes, err = queries.GetRecipes(ctx, database.NewFromNullableUUID(groupID))
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Recipes = make([]*v1.Recipe, 0, len(recipes))
	for _, recipe := range recipes {
		newRecipe := v1.Recipe{
			Id:                  uuid.UUID(recipe.ID.Bytes).String(),
			GroupId:             uuid.UUID(recipe.GroupID.Bytes).String(),
			Name:                recipe.Name,
			Comment:             recipe.Comment,
			Ingredients:         recipe.Ingredients,
			Instructions:        recipe.Instructions,
			Ph:                  recipe.Ph,
			MixTimeMilliseconds: recipe.MixTimeMilliseconds,
			Tags:                recipe.Tags,
			CreatedAt:           timestamppb.New(recipe.CreatedAt.Time),
		}
		if recipe.UpdatedAt.Valid {
			newRecipe.UpdatedAt = timestamppb.New(recipe.UpdatedAt.Time)
		}
		res.Msg.Recipes = append(res.Msg.Recipes, &newRecipe)
	}
	return res, nil
}

func (s *Service) GetRecipe(baseCtx gocontext.Context, req *connect.Request[v1.GetRecipeRequest]) (*connect.Response[v1.GetRecipeResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetRecipeResponse{})
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
	recipeID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid recipe id format: %w", err))
	}

	var recipe models.Recipe
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		params := models.GetRecipeParams{
			ID:      database.NewFromUUID(recipeID),
			GroupID: database.NewFromNullableUUID(groupID),
		}
		recipe, err = query.GetRecipe(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Recipe = &v1.Recipe{
		Id:                  uuid.UUID(recipe.ID.Bytes).String(),
		GroupId:             uuid.UUID(recipe.GroupID.Bytes).String(),
		Name:                recipe.Name,
		Comment:             recipe.Comment,
		Ingredients:         recipe.Ingredients,
		Instructions:        recipe.Instructions,
		Ph:                  recipe.Ph,
		MixTimeMilliseconds: recipe.MixTimeMilliseconds,
		Tags:                recipe.Tags,
		CreatedAt:           timestamppb.New(recipe.CreatedAt.Time),
	}
	if recipe.UpdatedAt.Valid {
		res.Msg.Recipe.UpdatedAt = timestamppb.New(recipe.UpdatedAt.Time)
	}
	return res, nil
}
