// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: authentication/v1/session.proto

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
	// SessionServiceName is the fully-qualified name of the SessionService service.
	SessionServiceName = "authentication.v1.SessionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SessionServiceCreateSessionProcedure is the fully-qualified name of the SessionService's
	// CreateSession RPC.
	SessionServiceCreateSessionProcedure = "/authentication.v1.SessionService/CreateSession"
	// SessionServiceGetSessionsProcedure is the fully-qualified name of the SessionService's
	// GetSessions RPC.
	SessionServiceGetSessionsProcedure = "/authentication.v1.SessionService/GetSessions"
	// SessionServiceGetSessionUserProcedure is the fully-qualified name of the SessionService's
	// GetSessionUser RPC.
	SessionServiceGetSessionUserProcedure = "/authentication.v1.SessionService/GetSessionUser"
	// SessionServiceDeleteSessionProcedure is the fully-qualified name of the SessionService's
	// DeleteSession RPC.
	SessionServiceDeleteSessionProcedure = "/authentication.v1.SessionService/DeleteSession"
	// SessionServiceCheckSessionPermissionsProcedure is the fully-qualified name of the
	// SessionService's CheckSessionPermissions RPC.
	SessionServiceCheckSessionPermissionsProcedure = "/authentication.v1.SessionService/CheckSessionPermissions"
)

// SessionServiceClient is a client for the authentication.v1.SessionService service.
type SessionServiceClient interface {
	CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error)
	GetSessions(context.Context, *connect_go.Request[v1.GetSessionsRequest]) (*connect_go.Response[v1.GetSessionsResponse], error)
	GetSessionUser(context.Context, *connect_go.Request[v1.GetSessionUserRequest]) (*connect_go.Response[v1.GetSessionUserResponse], error)
	DeleteSession(context.Context, *connect_go.Request[v1.DeleteSessionRequest]) (*connect_go.Response[v1.DeleteSessionResponse], error)
	CheckSessionPermissions(context.Context, *connect_go.Request[v1.CheckSessionPermissionsRequest]) (*connect_go.Response[v1.CheckSessionPermissionsResponse], error)
}

// NewSessionServiceClient constructs a client for the authentication.v1.SessionService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSessionServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SessionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sessionServiceClient{
		createSession: connect_go.NewClient[v1.CreateSessionRequest, v1.CreateSessionResponse](
			httpClient,
			baseURL+SessionServiceCreateSessionProcedure,
			opts...,
		),
		getSessions: connect_go.NewClient[v1.GetSessionsRequest, v1.GetSessionsResponse](
			httpClient,
			baseURL+SessionServiceGetSessionsProcedure,
			opts...,
		),
		getSessionUser: connect_go.NewClient[v1.GetSessionUserRequest, v1.GetSessionUserResponse](
			httpClient,
			baseURL+SessionServiceGetSessionUserProcedure,
			opts...,
		),
		deleteSession: connect_go.NewClient[v1.DeleteSessionRequest, v1.DeleteSessionResponse](
			httpClient,
			baseURL+SessionServiceDeleteSessionProcedure,
			opts...,
		),
		checkSessionPermissions: connect_go.NewClient[v1.CheckSessionPermissionsRequest, v1.CheckSessionPermissionsResponse](
			httpClient,
			baseURL+SessionServiceCheckSessionPermissionsProcedure,
			opts...,
		),
	}
}

// sessionServiceClient implements SessionServiceClient.
type sessionServiceClient struct {
	createSession           *connect_go.Client[v1.CreateSessionRequest, v1.CreateSessionResponse]
	getSessions             *connect_go.Client[v1.GetSessionsRequest, v1.GetSessionsResponse]
	getSessionUser          *connect_go.Client[v1.GetSessionUserRequest, v1.GetSessionUserResponse]
	deleteSession           *connect_go.Client[v1.DeleteSessionRequest, v1.DeleteSessionResponse]
	checkSessionPermissions *connect_go.Client[v1.CheckSessionPermissionsRequest, v1.CheckSessionPermissionsResponse]
}

// CreateSession calls authentication.v1.SessionService.CreateSession.
func (c *sessionServiceClient) CreateSession(ctx context.Context, req *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error) {
	return c.createSession.CallUnary(ctx, req)
}

// GetSessions calls authentication.v1.SessionService.GetSessions.
func (c *sessionServiceClient) GetSessions(ctx context.Context, req *connect_go.Request[v1.GetSessionsRequest]) (*connect_go.Response[v1.GetSessionsResponse], error) {
	return c.getSessions.CallUnary(ctx, req)
}

// GetSessionUser calls authentication.v1.SessionService.GetSessionUser.
func (c *sessionServiceClient) GetSessionUser(ctx context.Context, req *connect_go.Request[v1.GetSessionUserRequest]) (*connect_go.Response[v1.GetSessionUserResponse], error) {
	return c.getSessionUser.CallUnary(ctx, req)
}

// DeleteSession calls authentication.v1.SessionService.DeleteSession.
func (c *sessionServiceClient) DeleteSession(ctx context.Context, req *connect_go.Request[v1.DeleteSessionRequest]) (*connect_go.Response[v1.DeleteSessionResponse], error) {
	return c.deleteSession.CallUnary(ctx, req)
}

// CheckSessionPermissions calls authentication.v1.SessionService.CheckSessionPermissions.
func (c *sessionServiceClient) CheckSessionPermissions(ctx context.Context, req *connect_go.Request[v1.CheckSessionPermissionsRequest]) (*connect_go.Response[v1.CheckSessionPermissionsResponse], error) {
	return c.checkSessionPermissions.CallUnary(ctx, req)
}

// SessionServiceHandler is an implementation of the authentication.v1.SessionService service.
type SessionServiceHandler interface {
	CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error)
	GetSessions(context.Context, *connect_go.Request[v1.GetSessionsRequest]) (*connect_go.Response[v1.GetSessionsResponse], error)
	GetSessionUser(context.Context, *connect_go.Request[v1.GetSessionUserRequest]) (*connect_go.Response[v1.GetSessionUserResponse], error)
	DeleteSession(context.Context, *connect_go.Request[v1.DeleteSessionRequest]) (*connect_go.Response[v1.DeleteSessionResponse], error)
	CheckSessionPermissions(context.Context, *connect_go.Request[v1.CheckSessionPermissionsRequest]) (*connect_go.Response[v1.CheckSessionPermissionsResponse], error)
}

// NewSessionServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSessionServiceHandler(svc SessionServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(SessionServiceCreateSessionProcedure, connect_go.NewUnaryHandler(
		SessionServiceCreateSessionProcedure,
		svc.CreateSession,
		opts...,
	))
	mux.Handle(SessionServiceGetSessionsProcedure, connect_go.NewUnaryHandler(
		SessionServiceGetSessionsProcedure,
		svc.GetSessions,
		opts...,
	))
	mux.Handle(SessionServiceGetSessionUserProcedure, connect_go.NewUnaryHandler(
		SessionServiceGetSessionUserProcedure,
		svc.GetSessionUser,
		opts...,
	))
	mux.Handle(SessionServiceDeleteSessionProcedure, connect_go.NewUnaryHandler(
		SessionServiceDeleteSessionProcedure,
		svc.DeleteSession,
		opts...,
	))
	mux.Handle(SessionServiceCheckSessionPermissionsProcedure, connect_go.NewUnaryHandler(
		SessionServiceCheckSessionPermissionsProcedure,
		svc.CheckSessionPermissions,
		opts...,
	))
	return "/authentication.v1.SessionService/", mux
}

// UnimplementedSessionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSessionServiceHandler struct{}

func (UnimplementedSessionServiceHandler) CreateSession(context.Context, *connect_go.Request[v1.CreateSessionRequest]) (*connect_go.Response[v1.CreateSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.SessionService.CreateSession is not implemented"))
}

func (UnimplementedSessionServiceHandler) GetSessions(context.Context, *connect_go.Request[v1.GetSessionsRequest]) (*connect_go.Response[v1.GetSessionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.SessionService.GetSessions is not implemented"))
}

func (UnimplementedSessionServiceHandler) GetSessionUser(context.Context, *connect_go.Request[v1.GetSessionUserRequest]) (*connect_go.Response[v1.GetSessionUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.SessionService.GetSessionUser is not implemented"))
}

func (UnimplementedSessionServiceHandler) DeleteSession(context.Context, *connect_go.Request[v1.DeleteSessionRequest]) (*connect_go.Response[v1.DeleteSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.SessionService.DeleteSession is not implemented"))
}

func (UnimplementedSessionServiceHandler) CheckSessionPermissions(context.Context, *connect_go.Request[v1.CheckSessionPermissionsRequest]) (*connect_go.Response[v1.CheckSessionPermissionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("authentication.v1.SessionService.CheckSessionPermissions is not implemented"))
}
