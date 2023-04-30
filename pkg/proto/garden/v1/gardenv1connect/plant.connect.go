// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: garden/v1/plant.proto

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
	// PlantServiceName is the fully-qualified name of the PlantService service.
	PlantServiceName = "garden.v1.PlantService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PlantServiceCreatePlantProcedure is the fully-qualified name of the PlantService's CreatePlant
	// RPC.
	PlantServiceCreatePlantProcedure = "/garden.v1.PlantService/CreatePlant"
	// PlantServiceDeletePlantProcedure is the fully-qualified name of the PlantService's DeletePlant
	// RPC.
	PlantServiceDeletePlantProcedure = "/garden.v1.PlantService/DeletePlant"
	// PlantServiceUpdatePlantProcedure is the fully-qualified name of the PlantService's UpdatePlant
	// RPC.
	PlantServiceUpdatePlantProcedure = "/garden.v1.PlantService/UpdatePlant"
	// PlantServiceGetPlantsProcedure is the fully-qualified name of the PlantService's GetPlants RPC.
	PlantServiceGetPlantsProcedure = "/garden.v1.PlantService/GetPlants"
	// PlantServiceGetPlantProcedure is the fully-qualified name of the PlantService's GetPlant RPC.
	PlantServiceGetPlantProcedure = "/garden.v1.PlantService/GetPlant"
)

// PlantServiceClient is a client for the garden.v1.PlantService service.
type PlantServiceClient interface {
	CreatePlant(context.Context, *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error)
	DeletePlant(context.Context, *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error)
	UpdatePlant(context.Context, *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error)
	GetPlants(context.Context, *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error)
	GetPlant(context.Context, *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error)
}

// NewPlantServiceClient constructs a client for the garden.v1.PlantService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPlantServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PlantServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &plantServiceClient{
		createPlant: connect_go.NewClient[v1.CreatePlantRequest, v1.CreatePlantResponse](
			httpClient,
			baseURL+PlantServiceCreatePlantProcedure,
			opts...,
		),
		deletePlant: connect_go.NewClient[v1.DeletePlantRequest, v1.DeletePlantResponse](
			httpClient,
			baseURL+PlantServiceDeletePlantProcedure,
			opts...,
		),
		updatePlant: connect_go.NewClient[v1.UpdatePlantRequest, v1.UpdatePlantResponse](
			httpClient,
			baseURL+PlantServiceUpdatePlantProcedure,
			opts...,
		),
		getPlants: connect_go.NewClient[v1.GetPlantsRequest, v1.GetPlantsResponse](
			httpClient,
			baseURL+PlantServiceGetPlantsProcedure,
			opts...,
		),
		getPlant: connect_go.NewClient[v1.GetPlantRequest, v1.GetPlantResponse](
			httpClient,
			baseURL+PlantServiceGetPlantProcedure,
			opts...,
		),
	}
}

// plantServiceClient implements PlantServiceClient.
type plantServiceClient struct {
	createPlant *connect_go.Client[v1.CreatePlantRequest, v1.CreatePlantResponse]
	deletePlant *connect_go.Client[v1.DeletePlantRequest, v1.DeletePlantResponse]
	updatePlant *connect_go.Client[v1.UpdatePlantRequest, v1.UpdatePlantResponse]
	getPlants   *connect_go.Client[v1.GetPlantsRequest, v1.GetPlantsResponse]
	getPlant    *connect_go.Client[v1.GetPlantRequest, v1.GetPlantResponse]
}

// CreatePlant calls garden.v1.PlantService.CreatePlant.
func (c *plantServiceClient) CreatePlant(ctx context.Context, req *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error) {
	return c.createPlant.CallUnary(ctx, req)
}

// DeletePlant calls garden.v1.PlantService.DeletePlant.
func (c *plantServiceClient) DeletePlant(ctx context.Context, req *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error) {
	return c.deletePlant.CallUnary(ctx, req)
}

// UpdatePlant calls garden.v1.PlantService.UpdatePlant.
func (c *plantServiceClient) UpdatePlant(ctx context.Context, req *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error) {
	return c.updatePlant.CallUnary(ctx, req)
}

// GetPlants calls garden.v1.PlantService.GetPlants.
func (c *plantServiceClient) GetPlants(ctx context.Context, req *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error) {
	return c.getPlants.CallUnary(ctx, req)
}

// GetPlant calls garden.v1.PlantService.GetPlant.
func (c *plantServiceClient) GetPlant(ctx context.Context, req *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error) {
	return c.getPlant.CallUnary(ctx, req)
}

// PlantServiceHandler is an implementation of the garden.v1.PlantService service.
type PlantServiceHandler interface {
	CreatePlant(context.Context, *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error)
	DeletePlant(context.Context, *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error)
	UpdatePlant(context.Context, *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error)
	GetPlants(context.Context, *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error)
	GetPlant(context.Context, *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error)
}

// NewPlantServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPlantServiceHandler(svc PlantServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(PlantServiceCreatePlantProcedure, connect_go.NewUnaryHandler(
		PlantServiceCreatePlantProcedure,
		svc.CreatePlant,
		opts...,
	))
	mux.Handle(PlantServiceDeletePlantProcedure, connect_go.NewUnaryHandler(
		PlantServiceDeletePlantProcedure,
		svc.DeletePlant,
		opts...,
	))
	mux.Handle(PlantServiceUpdatePlantProcedure, connect_go.NewUnaryHandler(
		PlantServiceUpdatePlantProcedure,
		svc.UpdatePlant,
		opts...,
	))
	mux.Handle(PlantServiceGetPlantsProcedure, connect_go.NewUnaryHandler(
		PlantServiceGetPlantsProcedure,
		svc.GetPlants,
		opts...,
	))
	mux.Handle(PlantServiceGetPlantProcedure, connect_go.NewUnaryHandler(
		PlantServiceGetPlantProcedure,
		svc.GetPlant,
		opts...,
	))
	return "/garden.v1.PlantService/", mux
}

// UnimplementedPlantServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPlantServiceHandler struct{}

func (UnimplementedPlantServiceHandler) CreatePlant(context.Context, *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.PlantService.CreatePlant is not implemented"))
}

func (UnimplementedPlantServiceHandler) DeletePlant(context.Context, *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.PlantService.DeletePlant is not implemented"))
}

func (UnimplementedPlantServiceHandler) UpdatePlant(context.Context, *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.PlantService.UpdatePlant is not implemented"))
}

func (UnimplementedPlantServiceHandler) GetPlants(context.Context, *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.PlantService.GetPlants is not implemented"))
}

func (UnimplementedPlantServiceHandler) GetPlant(context.Context, *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("garden.v1.PlantService.GetPlant is not implemented"))
}
