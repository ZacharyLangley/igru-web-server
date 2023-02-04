// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: garden/v1/strains.proto

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

// StrainServiceClient is the client API for StrainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrainServiceClient interface {
	CreateStrain(ctx context.Context, in *CreateStrainRequest, opts ...grpc.CallOption) (*CreateStrainResponse, error)
	DeleteStrain(ctx context.Context, in *DeleteStrainRequest, opts ...grpc.CallOption) (*DeleteStrainResponse, error)
	UpdateStrain(ctx context.Context, in *UpdateStrainRequest, opts ...grpc.CallOption) (*UpdateStrainResponse, error)
	GetStrains(ctx context.Context, in *GetStrainsRequest, opts ...grpc.CallOption) (*GetStrainsResponse, error)
	GetStrain(ctx context.Context, in *GetStrainRequest, opts ...grpc.CallOption) (*GetStrainResponse, error)
}

type strainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStrainServiceClient(cc grpc.ClientConnInterface) StrainServiceClient {
	return &strainServiceClient{cc}
}

func (c *strainServiceClient) CreateStrain(ctx context.Context, in *CreateStrainRequest, opts ...grpc.CallOption) (*CreateStrainResponse, error) {
	out := new(CreateStrainResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.StrainService/CreateStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainServiceClient) DeleteStrain(ctx context.Context, in *DeleteStrainRequest, opts ...grpc.CallOption) (*DeleteStrainResponse, error) {
	out := new(DeleteStrainResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.StrainService/DeleteStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainServiceClient) UpdateStrain(ctx context.Context, in *UpdateStrainRequest, opts ...grpc.CallOption) (*UpdateStrainResponse, error) {
	out := new(UpdateStrainResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.StrainService/UpdateStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainServiceClient) GetStrains(ctx context.Context, in *GetStrainsRequest, opts ...grpc.CallOption) (*GetStrainsResponse, error) {
	out := new(GetStrainsResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.StrainService/GetStrains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainServiceClient) GetStrain(ctx context.Context, in *GetStrainRequest, opts ...grpc.CallOption) (*GetStrainResponse, error) {
	out := new(GetStrainResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.StrainService/GetStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrainServiceServer is the server API for StrainService service.
// All implementations must embed UnimplementedStrainServiceServer
// for forward compatibility
type StrainServiceServer interface {
	CreateStrain(context.Context, *CreateStrainRequest) (*CreateStrainResponse, error)
	DeleteStrain(context.Context, *DeleteStrainRequest) (*DeleteStrainResponse, error)
	UpdateStrain(context.Context, *UpdateStrainRequest) (*UpdateStrainResponse, error)
	GetStrains(context.Context, *GetStrainsRequest) (*GetStrainsResponse, error)
	GetStrain(context.Context, *GetStrainRequest) (*GetStrainResponse, error)
	mustEmbedUnimplementedStrainServiceServer()
}

// UnimplementedStrainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStrainServiceServer struct {
}

func (UnimplementedStrainServiceServer) CreateStrain(context.Context, *CreateStrainRequest) (*CreateStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStrain not implemented")
}
func (UnimplementedStrainServiceServer) DeleteStrain(context.Context, *DeleteStrainRequest) (*DeleteStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStrain not implemented")
}
func (UnimplementedStrainServiceServer) UpdateStrain(context.Context, *UpdateStrainRequest) (*UpdateStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStrain not implemented")
}
func (UnimplementedStrainServiceServer) GetStrains(context.Context, *GetStrainsRequest) (*GetStrainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStrains not implemented")
}
func (UnimplementedStrainServiceServer) GetStrain(context.Context, *GetStrainRequest) (*GetStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStrain not implemented")
}
func (UnimplementedStrainServiceServer) mustEmbedUnimplementedStrainServiceServer() {}

// UnsafeStrainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrainServiceServer will
// result in compilation errors.
type UnsafeStrainServiceServer interface {
	mustEmbedUnimplementedStrainServiceServer()
}

func RegisterStrainServiceServer(s grpc.ServiceRegistrar, srv StrainServiceServer) {
	s.RegisterService(&StrainService_ServiceDesc, srv)
}

func _StrainService_CreateStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainServiceServer).CreateStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.StrainService/CreateStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainServiceServer).CreateStrain(ctx, req.(*CreateStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainService_DeleteStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainServiceServer).DeleteStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.StrainService/DeleteStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainServiceServer).DeleteStrain(ctx, req.(*DeleteStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainService_UpdateStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainServiceServer).UpdateStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.StrainService/UpdateStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainServiceServer).UpdateStrain(ctx, req.(*UpdateStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainService_GetStrains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStrainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainServiceServer).GetStrains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.StrainService/GetStrains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainServiceServer).GetStrains(ctx, req.(*GetStrainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainService_GetStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainServiceServer).GetStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.StrainService/GetStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainServiceServer).GetStrain(ctx, req.(*GetStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StrainService_ServiceDesc is the grpc.ServiceDesc for StrainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StrainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "garden.v1.StrainService",
	HandlerType: (*StrainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStrain",
			Handler:    _StrainService_CreateStrain_Handler,
		},
		{
			MethodName: "DeleteStrain",
			Handler:    _StrainService_DeleteStrain_Handler,
		},
		{
			MethodName: "UpdateStrain",
			Handler:    _StrainService_UpdateStrain_Handler,
		},
		{
			MethodName: "GetStrains",
			Handler:    _StrainService_GetStrains_Handler,
		},
		{
			MethodName: "GetStrain",
			Handler:    _StrainService_GetStrain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "garden/v1/strains.proto",
}