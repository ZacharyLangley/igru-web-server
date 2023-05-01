package garden

import (
	gocontext "context"
	"errors"
	"fmt"
	"time"

	"github.com/ZacharyLangley/igru-web-server/pkg/context"
	models "github.com/ZacharyLangley/igru-web-server/pkg/models/garden"
	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1"
	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) CreateRecipe(baseCtx gocontext.Context, req *connect_go.Request[v1.CreateRecipeRequest]) (*connect_go.Response[v1.CreateRecipeResponse], error) {
	var err error
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.CreateRecipeResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
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
			GroupID:             groupID.UUID,
			Comment:             req.Msg.Comment,
			Ingredients:         req.Msg.Ingredients,
			Instructions:        req.Msg.Instructions,
			Ph:                  req.Msg.Ph,
			MixTimeMilliseconds: req.Msg.MixTimeMilliseconds,
			Tags:                req.Msg.Tags,
			CreatedAt:           time.Now(),
		}
		recipe, err = queries.CreateRecipe(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Recipe = &v1.Recipe{
		Id:                  recipe.ID.String(),
		GroupId:             recipe.GroupID.String(),
		Name:                recipe.Name,
		Comment:             recipe.Comment,
		Ingredients:         recipe.Ingredients,
		Instructions:        recipe.Instructions,
		Ph:                  recipe.Ph,
		MixTimeMilliseconds: recipe.MixTimeMilliseconds,
		Tags:                recipe.Tags,
		CreatedAt:           timestamppb.New(recipe.CreatedAt),
	}
	return res, nil
}

func (s *Service) DeleteRecipe(baseCtx gocontext.Context, req *connect_go.Request[v1.DeleteRecipeRequest]) (*connect_go.Response[v1.DeleteRecipeResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.DeleteRecipeResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
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
	recipeID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid recipe id format: %w", err))
	}
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		queries := models.New(tx)
		params := models.DeleteRecipeParams{
			ID:      recipeID,
			GroupID: groupID.UUID,
		}
		return queries.DeleteRecipe(ctx, params)
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateRecipe(baseCtx gocontext.Context, req *connect_go.Request[v1.UpdateRecipeRequest]) (*connect_go.Response[v1.UpdateRecipeResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.UpdateRecipeResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
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
	recipeID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid user id format: %w", err))
	}
	var recipe models.Recipe
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		var err error
		query := models.New(tx)
		params := models.UpdateRecipeParams{
			ID:                  recipeID,
			GroupID:             groupID.UUID,
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
		Id:                  recipe.ID.String(),
		GroupId:             recipe.GroupID.String(),
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

func (s *Service) GetRecipes(baseCtx gocontext.Context, req *connect_go.Request[v1.GetRecipesRequest]) (*connect_go.Response[v1.GetRecipesResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetRecipesResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
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
		recipes, err = queries.GetRecipes(ctx, groupID.UUID)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Recipes = make([]*v1.Recipe, 0, len(recipes))
	for _, recipe := range recipes {
		newRecipe := v1.Recipe{
			Id:                  recipe.ID.String(),
			GroupId:             recipe.GroupID.String(),
			Name:                recipe.Name,
			Comment:             recipe.Comment,
			Ingredients:         recipe.Ingredients,
			Instructions:        recipe.Instructions,
			Ph:                  recipe.Ph,
			MixTimeMilliseconds: recipe.MixTimeMilliseconds,
			Tags:                recipe.Tags,
			CreatedAt:           timestamppb.New(recipe.CreatedAt),
		}
		if recipe.UpdatedAt.Valid {
			newRecipe.UpdatedAt = timestamppb.New(recipe.UpdatedAt.Time)
		}
		res.Msg.Recipes = append(res.Msg.Recipes, &newRecipe)
	}
	return res, nil
}

func (s *Service) GetRecipe(baseCtx gocontext.Context, req *connect_go.Request[v1.GetRecipeRequest]) (*connect_go.Response[v1.GetRecipeResponse], error) {
	ctx := context.New(baseCtx)
	res := connect.NewResponse(&v1.GetRecipeResponse{})
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("missing request body"))
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
	recipeID, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid recipe id format: %w", err))
	}

	var recipe models.Recipe
	if err := s.pool.RunTransaction(ctx, func(ctx context.Context, tx pgx.Tx) error {
		query := models.New(tx)
		params := models.GetRecipeParams{
			ID:      recipeID,
			GroupID: groupID.UUID,
		}
		recipe, err = query.GetRecipe(ctx, params)
		return err
	}); err != nil {
		return nil, err
	}

	res.Msg.Recipe = &v1.Recipe{
		Id:                  recipe.ID.String(),
		Name:                recipe.Name,
		Comment:             recipe.Comment,
		Ingredients:         recipe.Ingredients,
		Instructions:        recipe.Instructions,
		Ph:                  recipe.Ph,
		MixTimeMilliseconds: recipe.MixTimeMilliseconds,
		Tags:                recipe.Tags,
		CreatedAt:           timestamppb.New(recipe.CreatedAt),
	}
	if recipe.UpdatedAt.Valid {
		res.Msg.Recipe.UpdatedAt = timestamppb.New(recipe.UpdatedAt.Time)
	}
	return res, nil
}
