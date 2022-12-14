// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gardens/v1/recipes.proto

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

// RecipesServiceClient is the client API for RecipesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecipesServiceClient interface {
	CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error)
	UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error)
	GetRecipes(ctx context.Context, in *GetRecipesRequest, opts ...grpc.CallOption) (*GetRecipesResponse, error)
	GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error)
}

type recipesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipesServiceClient(cc grpc.ClientConnInterface) RecipesServiceClient {
	return &recipesServiceClient{cc}
}

func (c *recipesServiceClient) CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.RecipesService/CreateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesServiceClient) DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error) {
	out := new(DeleteRecipeResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.RecipesService/DeleteRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesServiceClient) UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error) {
	out := new(UpdateRecipeResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.RecipesService/UpdateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesServiceClient) GetRecipes(ctx context.Context, in *GetRecipesRequest, opts ...grpc.CallOption) (*GetRecipesResponse, error) {
	out := new(GetRecipesResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.RecipesService/GetRecipes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesServiceClient) GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error) {
	out := new(GetRecipeResponse)
	err := c.cc.Invoke(ctx, "/gardens.v1.RecipesService/GetRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipesServiceServer is the server API for RecipesService service.
// All implementations must embed UnimplementedRecipesServiceServer
// for forward compatibility
type RecipesServiceServer interface {
	CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error)
	DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error)
	UpdateRecipe(context.Context, *UpdateRecipeRequest) (*UpdateRecipeResponse, error)
	GetRecipes(context.Context, *GetRecipesRequest) (*GetRecipesResponse, error)
	GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error)
	mustEmbedUnimplementedRecipesServiceServer()
}

// UnimplementedRecipesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecipesServiceServer struct {
}

func (UnimplementedRecipesServiceServer) CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecipe not implemented")
}
func (UnimplementedRecipesServiceServer) DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
func (UnimplementedRecipesServiceServer) UpdateRecipe(context.Context, *UpdateRecipeRequest) (*UpdateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipe not implemented")
}
func (UnimplementedRecipesServiceServer) GetRecipes(context.Context, *GetRecipesRequest) (*GetRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipes not implemented")
}
func (UnimplementedRecipesServiceServer) GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (UnimplementedRecipesServiceServer) mustEmbedUnimplementedRecipesServiceServer() {}

// UnsafeRecipesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecipesServiceServer will
// result in compilation errors.
type UnsafeRecipesServiceServer interface {
	mustEmbedUnimplementedRecipesServiceServer()
}

func RegisterRecipesServiceServer(s grpc.ServiceRegistrar, srv RecipesServiceServer) {
	s.RegisterService(&RecipesService_ServiceDesc, srv)
}

func _RecipesService_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServiceServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.RecipesService/CreateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServiceServer).CreateRecipe(ctx, req.(*CreateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipesService_DeleteRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServiceServer).DeleteRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.RecipesService/DeleteRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServiceServer).DeleteRecipe(ctx, req.(*DeleteRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipesService_UpdateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServiceServer).UpdateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.RecipesService/UpdateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServiceServer).UpdateRecipe(ctx, req.(*UpdateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipesService_GetRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServiceServer).GetRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.RecipesService/GetRecipes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServiceServer).GetRecipes(ctx, req.(*GetRecipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipesService_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServiceServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gardens.v1.RecipesService/GetRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServiceServer).GetRecipe(ctx, req.(*GetRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecipesService_ServiceDesc is the grpc.ServiceDesc for RecipesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecipesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gardens.v1.RecipesService",
	HandlerType: (*RecipesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecipe",
			Handler:    _RecipesService_CreateRecipe_Handler,
		},
		{
			MethodName: "DeleteRecipe",
			Handler:    _RecipesService_DeleteRecipe_Handler,
		},
		{
			MethodName: "UpdateRecipe",
			Handler:    _RecipesService_UpdateRecipe_Handler,
		},
		{
			MethodName: "GetRecipes",
			Handler:    _RecipesService_GetRecipes_Handler,
		},
		{
			MethodName: "GetRecipe",
			Handler:    _RecipesService_GetRecipe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gardens/v1/recipes.proto",
}
