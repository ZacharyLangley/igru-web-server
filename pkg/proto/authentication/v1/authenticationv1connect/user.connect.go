// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: authentication/v1/user.proto

package authenticationv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
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
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "authentication.v1.UserService"
)

// UserServiceClient is a client for the authentication.v1.UserService service.
type UserServiceClient interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
	ResetUserPassword(context.Context, *connect_go.Request[v1.ResetUserPasswordRequest]) (*connect_go.Response[v1.ResetUserPasswordResponse], error)
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
	GetUsers(context.Context, *connect_go.Request[v1.GetUsersRequest]) (*connect_go.Response[v1.GetUsersResponse], error)
}

// NewUserServiceClient constructs a client for the authentication.v1.UserService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+"/authentication.v1.UserService/CreateUser",
			opts...,
		),
		deleteUser: connect_go.NewClient[v1.DeleteUserRequest, v1.DeleteUserResponse](
			httpClient,
			baseURL+"/authentication.v1.UserService/DeleteUser",
			opts...,
		),
		resetUserPassword: connect_go.NewClient[v1.ResetUserPasswordRequest, v1.ResetUserPasswordResponse](
			httpClient,
			baseURL+"/authentication.v1.UserService/ResetUserPassword",
			opts...,
		),
		updateUser: connect_go.NewClient[v1.UpdateUserRequest, v1.UpdateUserResponse](
			httpClient,
			baseURL+"/authentication.v1.UserService/UpdateUser",
			opts...,
		),
		getUsers: connect_go.NewClient[v1.GetUsersRequest, v1.GetUsersResponse](
			httpClient,
			baseURL+"/authentication.v1.UserService/GetUsers",
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	createUser        *connect_go.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	deleteUser        *connect_go.Client[v1.DeleteUserRequest, v1.DeleteUserResponse]
	resetUserPassword *connect_go.Client[v1.ResetUserPasswordRequest, v1.ResetUserPasswordResponse]
	updateUser        *connect_go.Client[v1.UpdateUserRequest, v1.UpdateUserResponse]
	getUsers          *connect_go.Client[v1.GetUsersRequest, v1.GetUsersResponse]
}

// CreateUser calls authentication.v1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// DeleteUser calls authentication.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// ResetUserPassword calls authentication.v1.UserService.ResetUserPassword.
func (c *userServiceClient) ResetUserPassword(ctx context.Context, req *connect_go.Request[v1.ResetUserPasswordRequest]) (*connect_go.Response[v1.ResetUserPasswordResponse], error) {
	return c.resetUserPassword.CallUnary(ctx, req)
}

// UpdateUser calls authentication.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// GetUsers calls authentication.v1.UserService.GetUsers.
func (c *userServiceClient) GetUsers(ctx context.Context, req *connect_go.Request[v1.GetUsersRequest]) (*connect_go.Response[v1.GetUsersResponse], error) {
	return c.getUsers.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the authentication.v1.UserService service.
type UserServiceHandler interface {
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
	ResetUserPassword(context.Context, *connect_go.Request[v1.ResetUserPasswordRequest]) (*connect_go.Response[v1.ResetUserPasswordResponse], error)
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
	GetUsers(context.Context, *connect_go.Request[v1.GetUsersRequest]) (*connect_go.Response[v1.GetUsersResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/authentication.v1.UserService/CreateUser", connect_go.NewUnaryHandler(
		"/authentication.v1.UserService/CreateUser",
		svc.CreateUser,
		opts...,
	))
	mux.Handle("/authentication.v1.UserService/DeleteUser", connect_go.NewUnaryHandler(
		"/authentication.v1.UserService/DeleteUser",
		svc.DeleteUser,
		opts...,
	))
	mux.Handle("/authentication.v1.UserService/ResetUserPassword", connect_go.NewUnaryHandler(
		"/authentication.v1.UserService/ResetUserPassword",
		svc.ResetUserPassword,
		opts...,
	))
	mux.Handle("/authentication.v1.UserService/UpdateUser", connect_go.NewUnaryHandler(
		"/authentication.v1.UserService/UpdateUser",
		svc.UpdateUser,
		opts...,
	))
	mux.Handle("/authentication.v1.UserService/GetUsers", connect_go.NewUnaryHandler(
		"/authentication.v1.UserService/GetUsers",
		svc.GetUsers,
		opts...,
	))
	return "/authentication.v1.UserService/", mux
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.UserService.DeleteUser is not implemented"))
}

func (UnimplementedUserServiceHandler) ResetUserPassword(context.Context, *connect_go.Request[v1.ResetUserPasswordRequest]) (*connect_go.Response[v1.ResetUserPasswordResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.UserService.ResetUserPassword is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUsers(context.Context, *connect_go.Request[v1.GetUsersRequest]) (*connect_go.Response[v1.GetUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.UserService.GetUsers is not implemented"))
}
