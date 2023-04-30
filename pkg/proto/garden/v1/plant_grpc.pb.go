// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: garden/v1/plant.proto

package gardenv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PlantService_CreatePlant_FullMethodName = "/garden.v1.PlantService/CreatePlant"
	PlantService_DeletePlant_FullMethodName = "/garden.v1.PlantService/DeletePlant"
	PlantService_UpdatePlant_FullMethodName = "/garden.v1.PlantService/UpdatePlant"
	PlantService_GetPlants_FullMethodName   = "/garden.v1.PlantService/GetPlants"
	PlantService_GetPlant_FullMethodName    = "/garden.v1.PlantService/GetPlant"
)

// PlantServiceClient is the client API for PlantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlantServiceClient interface {
	CreatePlant(ctx context.Context, in *CreatePlantRequest, opts ...grpc.CallOption) (*CreatePlantResponse, error)
	DeletePlant(ctx context.Context, in *DeletePlantRequest, opts ...grpc.CallOption) (*DeletePlantResponse, error)
	UpdatePlant(ctx context.Context, in *UpdatePlantRequest, opts ...grpc.CallOption) (*UpdatePlantResponse, error)
	GetPlants(ctx context.Context, in *GetPlantsRequest, opts ...grpc.CallOption) (*GetPlantsResponse, error)
	GetPlant(ctx context.Context, in *GetPlantRequest, opts ...grpc.CallOption) (*GetPlantResponse, error)
}

type plantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlantServiceClient(cc grpc.ClientConnInterface) PlantServiceClient {
	return &plantServiceClient{cc}
}

func (c *plantServiceClient) CreatePlant(ctx context.Context, in *CreatePlantRequest, opts ...grpc.CallOption) (*CreatePlantResponse, error) {
	out := new(CreatePlantResponse)
	err := c.cc.Invoke(ctx, PlantService_CreatePlant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plantServiceClient) DeletePlant(ctx context.Context, in *DeletePlantRequest, opts ...grpc.CallOption) (*DeletePlantResponse, error) {
	out := new(DeletePlantResponse)
	err := c.cc.Invoke(ctx, PlantService_DeletePlant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plantServiceClient) UpdatePlant(ctx context.Context, in *UpdatePlantRequest, opts ...grpc.CallOption) (*UpdatePlantResponse, error) {
	out := new(UpdatePlantResponse)
	err := c.cc.Invoke(ctx, PlantService_UpdatePlant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plantServiceClient) GetPlants(ctx context.Context, in *GetPlantsRequest, opts ...grpc.CallOption) (*GetPlantsResponse, error) {
	out := new(GetPlantsResponse)
	err := c.cc.Invoke(ctx, PlantService_GetPlants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plantServiceClient) GetPlant(ctx context.Context, in *GetPlantRequest, opts ...grpc.CallOption) (*GetPlantResponse, error) {
	out := new(GetPlantResponse)
	err := c.cc.Invoke(ctx, PlantService_GetPlant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlantServiceServer is the server API for PlantService service.
// All implementations must embed UnimplementedPlantServiceServer
// for forward compatibility
type PlantServiceServer interface {
	CreatePlant(context.Context, *CreatePlantRequest) (*CreatePlantResponse, error)
	DeletePlant(context.Context, *DeletePlantRequest) (*DeletePlantResponse, error)
	UpdatePlant(context.Context, *UpdatePlantRequest) (*UpdatePlantResponse, error)
	GetPlants(context.Context, *GetPlantsRequest) (*GetPlantsResponse, error)
	GetPlant(context.Context, *GetPlantRequest) (*GetPlantResponse, error)
	mustEmbedUnimplementedPlantServiceServer()
}

// UnimplementedPlantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlantServiceServer struct {
}

func (UnimplementedPlantServiceServer) CreatePlant(context.Context, *CreatePlantRequest) (*CreatePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlant not implemented")
}
func (UnimplementedPlantServiceServer) DeletePlant(context.Context, *DeletePlantRequest) (*DeletePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlant not implemented")
}
func (UnimplementedPlantServiceServer) UpdatePlant(context.Context, *UpdatePlantRequest) (*UpdatePlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlant not implemented")
}
func (UnimplementedPlantServiceServer) GetPlants(context.Context, *GetPlantsRequest) (*GetPlantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlants not implemented")
}
func (UnimplementedPlantServiceServer) GetPlant(context.Context, *GetPlantRequest) (*GetPlantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlant not implemented")
}
func (UnimplementedPlantServiceServer) mustEmbedUnimplementedPlantServiceServer() {}

// UnsafePlantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlantServiceServer will
// result in compilation errors.
type UnsafePlantServiceServer interface {
	mustEmbedUnimplementedPlantServiceServer()
}

func RegisterPlantServiceServer(s grpc.ServiceRegistrar, srv PlantServiceServer) {
	s.RegisterService(&PlantService_ServiceDesc, srv)
}

func _PlantService_CreatePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantServiceServer).CreatePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlantService_CreatePlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantServiceServer).CreatePlant(ctx, req.(*CreatePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlantService_DeletePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantServiceServer).DeletePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlantService_DeletePlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantServiceServer).DeletePlant(ctx, req.(*DeletePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlantService_UpdatePlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantServiceServer).UpdatePlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlantService_UpdatePlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantServiceServer).UpdatePlant(ctx, req.(*UpdatePlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlantService_GetPlants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantServiceServer).GetPlants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlantService_GetPlants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantServiceServer).GetPlants(ctx, req.(*GetPlantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlantService_GetPlant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlantServiceServer).GetPlant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlantService_GetPlant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlantServiceServer).GetPlant(ctx, req.(*GetPlantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlantService_ServiceDesc is the grpc.ServiceDesc for PlantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "garden.v1.PlantService",
	HandlerType: (*PlantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlant",
			Handler:    _PlantService_CreatePlant_Handler,
		},
		{
			MethodName: "DeletePlant",
			Handler:    _PlantService_DeletePlant_Handler,
		},
		{
			MethodName: "UpdatePlant",
			Handler:    _PlantService_UpdatePlant_Handler,
		},
		{
			MethodName: "GetPlants",
			Handler:    _PlantService_GetPlants_Handler,
		},
		{
			MethodName: "GetPlant",
			Handler:    _PlantService_GetPlant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "garden/v1/plant.proto",
}
