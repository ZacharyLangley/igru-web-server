// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: garden/v1/recipe.proto

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

// RecipeServiceClient is the client API for RecipeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecipeServiceClient interface {
	CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error)
	UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error)
	GetRecipes(ctx context.Context, in *GetRecipesRequest, opts ...grpc.CallOption) (*GetRecipesResponse, error)
	GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error)
}

type recipeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipeServiceClient(cc grpc.ClientConnInterface) RecipeServiceClient {
	return &recipeServiceClient{cc}
}

func (c *recipeServiceClient) CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.RecipeService/CreateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) DeleteRecipe(ctx context.Context, in *DeleteRecipeRequest, opts ...grpc.CallOption) (*DeleteRecipeResponse, error) {
	out := new(DeleteRecipeResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.RecipeService/DeleteRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) UpdateRecipe(ctx context.Context, in *UpdateRecipeRequest, opts ...grpc.CallOption) (*UpdateRecipeResponse, error) {
	out := new(UpdateRecipeResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.RecipeService/UpdateRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GetRecipes(ctx context.Context, in *GetRecipesRequest, opts ...grpc.CallOption) (*GetRecipesResponse, error) {
	out := new(GetRecipesResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.RecipeService/GetRecipes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipeServiceClient) GetRecipe(ctx context.Context, in *GetRecipeRequest, opts ...grpc.CallOption) (*GetRecipeResponse, error) {
	out := new(GetRecipeResponse)
	err := c.cc.Invoke(ctx, "/garden.v1.RecipeService/GetRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecipeServiceServer is the server API for RecipeService service.
// All implementations must embed UnimplementedRecipeServiceServer
// for forward compatibility
type RecipeServiceServer interface {
	CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error)
	DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error)
	UpdateRecipe(context.Context, *UpdateRecipeRequest) (*UpdateRecipeResponse, error)
	GetRecipes(context.Context, *GetRecipesRequest) (*GetRecipesResponse, error)
	GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error)
	mustEmbedUnimplementedRecipeServiceServer()
}

// UnimplementedRecipeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecipeServiceServer struct {
}

func (UnimplementedRecipeServiceServer) CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) DeleteRecipe(context.Context, *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) UpdateRecipe(context.Context, *UpdateRecipeRequest) (*UpdateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) GetRecipes(context.Context, *GetRecipesRequest) (*GetRecipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipes not implemented")
}
func (UnimplementedRecipeServiceServer) GetRecipe(context.Context, *GetRecipeRequest) (*GetRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipe not implemented")
}
func (UnimplementedRecipeServiceServer) mustEmbedUnimplementedRecipeServiceServer() {}

// UnsafeRecipeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecipeServiceServer will
// result in compilation errors.
type UnsafeRecipeServiceServer interface {
	mustEmbedUnimplementedRecipeServiceServer()
}

func RegisterRecipeServiceServer(s grpc.ServiceRegistrar, srv RecipeServiceServer) {
	s.RegisterService(&RecipeService_ServiceDesc, srv)
}

func _RecipeService_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.RecipeService/CreateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).CreateRecipe(ctx, req.(*CreateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_DeleteRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).DeleteRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.RecipeService/DeleteRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).DeleteRecipe(ctx, req.(*DeleteRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_UpdateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).UpdateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.RecipeService/UpdateRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).UpdateRecipe(ctx, req.(*UpdateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GetRecipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.RecipeService/GetRecipes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipes(ctx, req.(*GetRecipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipeService_GetRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipeServiceServer).GetRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/garden.v1.RecipeService/GetRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipeServiceServer).GetRecipe(ctx, req.(*GetRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecipeService_ServiceDesc is the grpc.ServiceDesc for RecipeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecipeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "garden.v1.RecipeService",
	HandlerType: (*RecipeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecipe",
			Handler:    _RecipeService_CreateRecipe_Handler,
		},
		{
			MethodName: "DeleteRecipe",
			Handler:    _RecipeService_DeleteRecipe_Handler,
		},
		{
			MethodName: "UpdateRecipe",
			Handler:    _RecipeService_UpdateRecipe_Handler,
		},
		{
			MethodName: "GetRecipes",
			Handler:    _RecipeService_GetRecipes_Handler,
		},
		{
			MethodName: "GetRecipe",
			Handler:    _RecipeService_GetRecipe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "garden/v1/recipe.proto",
}