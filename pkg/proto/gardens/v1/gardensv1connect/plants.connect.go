// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gardens/v1/plants.proto

package gardensv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/gardens/v1"
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
	// PlantsServiceName is the fully-qualified name of the PlantsService service.
	PlantsServiceName = "gardens.v1.PlantsService"
)

// PlantsServiceClient is a client for the gardens.v1.PlantsService service.
type PlantsServiceClient interface {
	CreatePlant(context.Context, *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error)
	DeletePlant(context.Context, *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error)
	UpdatePlant(context.Context, *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error)
	GetPlants(context.Context, *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error)
	GetPlant(context.Context, *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error)
}

// NewPlantsServiceClient constructs a client for the gardens.v1.PlantsService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPlantsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PlantsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &plantsServiceClient{
		createPlant: connect_go.NewClient[v1.CreatePlantRequest, v1.CreatePlantResponse](
			httpClient,
			baseURL+"/gardens.v1.PlantsService/CreatePlant",
			opts...,
		),
		deletePlant: connect_go.NewClient[v1.DeletePlantRequest, v1.DeletePlantResponse](
			httpClient,
			baseURL+"/gardens.v1.PlantsService/DeletePlant",
			opts...,
		),
		updatePlant: connect_go.NewClient[v1.UpdatePlantRequest, v1.UpdatePlantResponse](
			httpClient,
			baseURL+"/gardens.v1.PlantsService/UpdatePlant",
			opts...,
		),
		getPlants: connect_go.NewClient[v1.GetPlantsRequest, v1.GetPlantsResponse](
			httpClient,
			baseURL+"/gardens.v1.PlantsService/GetPlants",
			opts...,
		),
		getPlant: connect_go.NewClient[v1.GetPlantRequest, v1.GetPlantResponse](
			httpClient,
			baseURL+"/gardens.v1.PlantsService/GetPlant",
			opts...,
		),
	}
}

// plantsServiceClient implements PlantsServiceClient.
type plantsServiceClient struct {
	createPlant *connect_go.Client[v1.CreatePlantRequest, v1.CreatePlantResponse]
	deletePlant *connect_go.Client[v1.DeletePlantRequest, v1.DeletePlantResponse]
	updatePlant *connect_go.Client[v1.UpdatePlantRequest, v1.UpdatePlantResponse]
	getPlants   *connect_go.Client[v1.GetPlantsRequest, v1.GetPlantsResponse]
	getPlant    *connect_go.Client[v1.GetPlantRequest, v1.GetPlantResponse]
}

// CreatePlant calls gardens.v1.PlantsService.CreatePlant.
func (c *plantsServiceClient) CreatePlant(ctx context.Context, req *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error) {
	return c.createPlant.CallUnary(ctx, req)
}

// DeletePlant calls gardens.v1.PlantsService.DeletePlant.
func (c *plantsServiceClient) DeletePlant(ctx context.Context, req *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error) {
	return c.deletePlant.CallUnary(ctx, req)
}

// UpdatePlant calls gardens.v1.PlantsService.UpdatePlant.
func (c *plantsServiceClient) UpdatePlant(ctx context.Context, req *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error) {
	return c.updatePlant.CallUnary(ctx, req)
}

// GetPlants calls gardens.v1.PlantsService.GetPlants.
func (c *plantsServiceClient) GetPlants(ctx context.Context, req *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error) {
	return c.getPlants.CallUnary(ctx, req)
}

// GetPlant calls gardens.v1.PlantsService.GetPlant.
func (c *plantsServiceClient) GetPlant(ctx context.Context, req *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error) {
	return c.getPlant.CallUnary(ctx, req)
}

// PlantsServiceHandler is an implementation of the gardens.v1.PlantsService service.
type PlantsServiceHandler interface {
	CreatePlant(context.Context, *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error)
	DeletePlant(context.Context, *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error)
	UpdatePlant(context.Context, *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error)
	GetPlants(context.Context, *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error)
	GetPlant(context.Context, *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error)
}

// NewPlantsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPlantsServiceHandler(svc PlantsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gardens.v1.PlantsService/CreatePlant", connect_go.NewUnaryHandler(
		"/gardens.v1.PlantsService/CreatePlant",
		svc.CreatePlant,
		opts...,
	))
	mux.Handle("/gardens.v1.PlantsService/DeletePlant", connect_go.NewUnaryHandler(
		"/gardens.v1.PlantsService/DeletePlant",
		svc.DeletePlant,
		opts...,
	))
	mux.Handle("/gardens.v1.PlantsService/UpdatePlant", connect_go.NewUnaryHandler(
		"/gardens.v1.PlantsService/UpdatePlant",
		svc.UpdatePlant,
		opts...,
	))
	mux.Handle("/gardens.v1.PlantsService/GetPlants", connect_go.NewUnaryHandler(
		"/gardens.v1.PlantsService/GetPlants",
		svc.GetPlants,
		opts...,
	))
	mux.Handle("/gardens.v1.PlantsService/GetPlant", connect_go.NewUnaryHandler(
		"/gardens.v1.PlantsService/GetPlant",
		svc.GetPlant,
		opts...,
	))
	return "/gardens.v1.PlantsService/", mux
}

// UnimplementedPlantsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPlantsServiceHandler struct{}

func (UnimplementedPlantsServiceHandler) CreatePlant(context.Context, *connect_go.Request[v1.CreatePlantRequest]) (*connect_go.Response[v1.CreatePlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gardens.v1.PlantsService.CreatePlant is not implemented"))
}

func (UnimplementedPlantsServiceHandler) DeletePlant(context.Context, *connect_go.Request[v1.DeletePlantRequest]) (*connect_go.Response[v1.DeletePlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gardens.v1.PlantsService.DeletePlant is not implemented"))
}

func (UnimplementedPlantsServiceHandler) UpdatePlant(context.Context, *connect_go.Request[v1.UpdatePlantRequest]) (*connect_go.Response[v1.UpdatePlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gardens.v1.PlantsService.UpdatePlant is not implemented"))
}

func (UnimplementedPlantsServiceHandler) GetPlants(context.Context, *connect_go.Request[v1.GetPlantsRequest]) (*connect_go.Response[v1.GetPlantsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gardens.v1.PlantsService.GetPlants is not implemented"))
}

func (UnimplementedPlantsServiceHandler) GetPlant(context.Context, *connect_go.Request[v1.GetPlantRequest]) (*connect_go.Response[v1.GetPlantResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gardens.v1.PlantsService.GetPlant is not implemented"))
}
