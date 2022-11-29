// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gardens/v1/strains.proto

package gardensv1

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

// StrainsServiceClient is the client API for StrainsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrainsServiceClient interface {
	CreateStrain(ctx context.Context, in *CreateStrainRequest, opts ...grpc.CallOption) (*CreateStrainResponse, error)
	DeleteStrain(ctx context.Context, in *DeleteStrainRequest, opts ...grpc.CallOption) (*DeleteStrainResponse, error)
	UpdateStrain(ctx context.Context, in *UpdateStrainRequest, opts ...grpc.CallOption) (*UpdateStrainResponse, error)
	GetStrains(ctx context.Context, in *GetStrainsRequest, opts ...grpc.CallOption) (*GetStrainsResponse, error)
	GetStrain(ctx context.Context, in *GetStrainRequest, opts ...grpc.CallOption) (*GetStrainResponse, error)
}

type strainsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStrainsServiceClient(cc grpc.ClientConnInterface) StrainsServiceClient {
	return &strainsServiceClient{cc}
}

func (c *strainsServiceClient) CreateStrain(ctx context.Context, in *CreateStrainRequest, opts ...grpc.CallOption) (*CreateStrainResponse, error) {
	out := new(CreateStrainResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.StrainsService/CreateStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainsServiceClient) DeleteStrain(ctx context.Context, in *DeleteStrainRequest, opts ...grpc.CallOption) (*DeleteStrainResponse, error) {
	out := new(DeleteStrainResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.StrainsService/DeleteStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainsServiceClient) UpdateStrain(ctx context.Context, in *UpdateStrainRequest, opts ...grpc.CallOption) (*UpdateStrainResponse, error) {
	out := new(UpdateStrainResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.StrainsService/UpdateStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainsServiceClient) GetStrains(ctx context.Context, in *GetStrainsRequest, opts ...grpc.CallOption) (*GetStrainsResponse, error) {
	out := new(GetStrainsResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.StrainsService/GetStrains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strainsServiceClient) GetStrain(ctx context.Context, in *GetStrainRequest, opts ...grpc.CallOption) (*GetStrainResponse, error) {
	out := new(GetStrainResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.StrainsService/GetStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrainsServiceServer is the server API for StrainsService service.
// All implementations must embed UnimplementedStrainsServiceServer
// for forward compatibility
type StrainsServiceServer interface {
	CreateStrain(context.Context, *CreateStrainRequest) (*CreateStrainResponse, error)
	DeleteStrain(context.Context, *DeleteStrainRequest) (*DeleteStrainResponse, error)
	UpdateStrain(context.Context, *UpdateStrainRequest) (*UpdateStrainResponse, error)
	GetStrains(context.Context, *GetStrainsRequest) (*GetStrainsResponse, error)
	GetStrain(context.Context, *GetStrainRequest) (*GetStrainResponse, error)
	mustEmbedUnimplementedStrainsServiceServer()
}

// UnimplementedStrainsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStrainsServiceServer struct {
}

func (UnimplementedStrainsServiceServer) CreateStrain(context.Context, *CreateStrainRequest) (*CreateStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStrain not implemented")
}
func (UnimplementedStrainsServiceServer) DeleteStrain(context.Context, *DeleteStrainRequest) (*DeleteStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStrain not implemented")
}
func (UnimplementedStrainsServiceServer) UpdateStrain(context.Context, *UpdateStrainRequest) (*UpdateStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStrain not implemented")
}
func (UnimplementedStrainsServiceServer) GetStrains(context.Context, *GetStrainsRequest) (*GetStrainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStrains not implemented")
}
func (UnimplementedStrainsServiceServer) GetStrain(context.Context, *GetStrainRequest) (*GetStrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStrain not implemented")
}
func (UnimplementedStrainsServiceServer) mustEmbedUnimplementedStrainsServiceServer() {}

// UnsafeStrainsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrainsServiceServer will
// result in compilation errors.
type UnsafeStrainsServiceServer interface {
	mustEmbedUnimplementedStrainsServiceServer()
}

func RegisterStrainsServiceServer(s grpc.ServiceRegistrar, srv StrainsServiceServer) {
	s.RegisterService(&StrainsService_ServiceDesc, srv)
}

func _StrainsService_CreateStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainsServiceServer).CreateStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.StrainsService/CreateStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainsServiceServer).CreateStrain(ctx, req.(*CreateStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainsService_DeleteStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainsServiceServer).DeleteStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.StrainsService/DeleteStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainsServiceServer).DeleteStrain(ctx, req.(*DeleteStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainsService_UpdateStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainsServiceServer).UpdateStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.StrainsService/UpdateStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainsServiceServer).UpdateStrain(ctx, req.(*UpdateStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainsService_GetStrains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStrainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainsServiceServer).GetStrains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.StrainsService/GetStrains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainsServiceServer).GetStrains(ctx, req.(*GetStrainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StrainsService_GetStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrainsServiceServer).GetStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.StrainsService/GetStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrainsServiceServer).GetStrain(ctx, req.(*GetStrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StrainsService_ServiceDesc is the grpc.ServiceDesc for StrainsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StrainsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gardens.v1.StrainsService",
	HandlerType: (*StrainsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStrain",
			Handler:    _StrainsService_CreateStrain_Handler,
		},
		{
			MethodName: "DeleteStrain",
			Handler:    _StrainsService_DeleteStrain_Handler,
		},
		{
			MethodName: "UpdateStrain",
			Handler:    _StrainsService_UpdateStrain_Handler,
		},
		{
			MethodName: "GetStrains",
			Handler:    _StrainsService_GetStrains_Handler,
		},
		{
			MethodName: "GetStrain",
			Handler:    _StrainsService_GetStrain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gardens/v1/strains.proto",
}
