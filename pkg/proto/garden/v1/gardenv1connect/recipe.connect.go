// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: garden/v1/recipe.proto

package gardenv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/garden/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// RecipeServiceName is the fully-qualified name of the RecipeService service.
	RecipeServiceName = "garden.v1.RecipeService"
)

// RecipeServiceClient is a client for the garden.v1.RecipeService service.
type RecipeServiceClient interface {
	CreateRecipe(context.Context, *connect_go.Request[v1.CreateRecipeRequest]) (*connect_go.Response[v1.CreateRecipeResponse], error)
	DeleteRecipe(context.Context, *connect_go.Request[v1.DeleteRecipeRequest]) (*connect_go.Response[v1.DeleteRecipeResponse], error)
	UpdateRecipe(context.Context, *connect_go.Request[v1.UpdateRecipeRequest]) (*connect_go.Response[v1.UpdateRecipeResponse], error)
	GetRecipes(context.Context, *connect_go.Request[v1.GetRecipesRequest]) (*connect_go.Response[v1.GetRecipesResponse], error)
	GetRecipe(context.Context, *connect_go.Request[v1.GetRecipeRequest]) (*connect_go.Response[v1.GetRecipeResponse], error)
}

// NewRecipeServiceClient constructs a client for the garden.v1.RecipeService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRecipeServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RecipeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &recipeServiceClient{
		createRecipe: connect_go.NewClient[v1.CreateRecipeRequest, v1.CreateRecipeResponse](
			httpClient,
			baseURL+"/garden.v1.RecipeService/CreateRecipe",
			opts...,
		),
		deleteRecipe: connect_go.NewClient[v1.DeleteRecipeRequest, v1.DeleteRecipeResponse](
			httpClient,
			baseURL+"/garden.v1.RecipeService/DeleteRecipe",
			opts...,
		),
		updateRecipe: connect_go.NewClient[v1.UpdateRecipeRequest, v1.UpdateRecipeResponse](
			httpClient,
			baseURL+"/garden.v1.RecipeService/UpdateRecipe",
			opts...,
		),
		getRecipes: connect_go.NewClient[v1.GetRecipesRequest, v1.GetRecipesResponse](
			httpClient,
			baseURL+"/garden.v1.RecipeService/GetRecipes",
			opts...,
		),
		getRecipe: connect_go.NewClient[v1.GetRecipeRequest, v1.GetRecipeResponse](
			httpClient,
			baseURL+"/garden.v1.RecipeService/GetRecipe",
			opts...,
		),
	}
}

// recipeServiceClient implements RecipeServiceClient.
type recipeServiceClient struct {
	createRecipe *connect_go.Client[v1.CreateRecipeRequest, v1.CreateRecipeResponse]
	deleteRecipe *connect_go.Client[v1.DeleteRecipeRequest, v1.DeleteRecipeResponse]
	updateRecipe *connect_go.Client[v1.UpdateRecipeRequest, v1.UpdateRecipeResponse]
	getRecipes   *connect_go.Client[v1.GetRecipesRequest, v1.GetRecipesResponse]
	getRecipe    *connect_go.Client[v1.GetRecipeRequest, v1.GetRecipeResponse]
}

// CreateRecipe calls garden.v1.RecipeService.CreateRecipe.
func (c *recipeServiceClient) CreateRecipe(ctx context.Context, req *connect_go.Request[v1.CreateRecipeRequest]) (*connect_go.Response[v1.CreateRecipeResponse], error) {
	return c.createRecipe.CallUnary(ctx, req)
}

// DeleteRecipe calls garden.v1.RecipeService.DeleteRecipe.
func (c *recipeServiceClient) DeleteRecipe(ctx context.Context, req *connect_go.Request[v1.DeleteRecipeRequest]) (*connect_go.Response[v1.DeleteRecipeResponse], error) {
	return c.deleteRecipe.CallUnary(ctx, req)
}

// UpdateRecipe calls garden.v1.RecipeService.UpdateRecipe.
func (c *recipeServiceClient) UpdateRecipe(ctx context.Context, req *connect_go.Request[v1.UpdateRecipeRequest]) (*connect_go.Response[v1.UpdateRecipeResponse], error) {
	return c.updateRecipe.CallUnary(ctx, req)
}

// GetRecipes calls garden.v1.RecipeService.GetRecipes.
func (c *recipeServiceClient) GetRecipes(ctx context.Context, req *connect_go.Request[v1.GetRecipesRequest]) (*connect_go.Response[v1.GetRecipesResponse], error) {
	return c.getRecipes.CallUnary(ctx, req)
}

// GetRecipe calls garden.v1.RecipeService.GetRecipe.
func (c *recipeServiceClient) GetRecipe(ctx context.Context, req *connect_go.Request[v1.GetRecipeRequest]) (*connect_go.Response[v1.GetRecipeResponse], error) {
	return c.getRecipe.CallUnary(ctx, req)
}

// RecipeServiceHandler is an implementation of the garden.v1.RecipeService service.
type RecipeServiceHandler interface {
	CreateRecipe(context.Context, *connect_go.Request[v1.CreateRecipeRequest]) (*connect_go.Response[v1.CreateRecipeResponse], error)
	DeleteRecipe(context.Context, *connect_go.Request[v1.DeleteRecipeRequest]) (*connect_go.Response[v1.DeleteRecipeResponse], error)
	UpdateRecipe(context.Context, *connect_go.Request[v1.UpdateRecipeRequest]) (*connect_go.Response[v1.UpdateRecipeResponse], error)
	GetRecipes(context.Context, *connect_go.Request[v1.GetRecipesRequest]) (*connect_go.Response[v1.GetRecipesResponse], error)
	GetRecipe(context.Context, *connect_go.Request[v1.GetRecipeRequest]) (*connect_go.Response[v1.GetRecipeResponse], error)
}

// NewRecipeServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRecipeServiceHandler(svc RecipeServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/garden.v1.RecipeService/CreateRecipe", connect_go.NewUnaryHandler(
		"/garden.v1.RecipeService/CreateRecipe",
		svc.CreateRecipe,
		opts...,
	))
	mux.Handle("/garden.v1.RecipeService/DeleteRecipe", connect_go.NewUnaryHandler(
		"/garden.v1.RecipeService/DeleteRecipe",
		svc.DeleteRecipe,
		opts...,
	))
	mux.Handle("/garden.v1.RecipeService/UpdateRecipe", connect_go.NewUnaryHandler(
		"/garden.v1.RecipeService/UpdateRecipe",
		svc.UpdateRecipe,
		opts...,
	))
	mux.Handle("/garden.v1.RecipeService/GetRecipes", connect_go.NewUnaryHandler(
		"/garden.v1.RecipeService/GetRecipes",
		svc.GetRecipes,
		opts...,
	))
	mux.Handle("/garden.v1.RecipeService/GetRecipe", connect_go.NewUnaryHandler(
		"/garden.v1.RecipeService/GetRecipe",
		svc.GetRecipe,
		opts...,
	))
	return "/garden.v1.RecipeService/", mux
}

// UnimplementedRecipeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRecipeServiceHandler struct{}

func (UnimplementedRecipeServiceHandler) CreateRecipe(context.Context, *connect_go.Request[v1.CreateRecipeRequest]) (*connect_go.Response[v1.CreateRecipeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.RecipeService.CreateRecipe is not implemented"))
}

func (UnimplementedRecipeServiceHandler) DeleteRecipe(context.Context, *connect_go.Request[v1.DeleteRecipeRequest]) (*connect_go.Response[v1.DeleteRecipeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.RecipeService.DeleteRecipe is not implemented"))
}

func (UnimplementedRecipeServiceHandler) UpdateRecipe(context.Context, *connect_go.Request[v1.UpdateRecipeRequest]) (*connect_go.Response[v1.UpdateRecipeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.RecipeService.UpdateRecipe is not implemented"))
}

func (UnimplementedRecipeServiceHandler) GetRecipes(context.Context, *connect_go.Request[v1.GetRecipesRequest]) (*connect_go.Response[v1.GetRecipesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.RecipeService.GetRecipes is not implemented"))
}

func (UnimplementedRecipeServiceHandler) GetRecipe(context.Context, *connect_go.Request[v1.GetRecipeRequest]) (*connect_go.Response[v1.GetRecipeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.RecipeService.GetRecipe is not implemented"))
}
