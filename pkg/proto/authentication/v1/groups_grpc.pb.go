// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: authentication/v1/groups.proto

package authenticationv1

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

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error)
	AddGroupMember(ctx context.Context, in *AddGroupMemberRequest, opts ...grpc.CallOption) (*AddGroupMemberResponse, error)
	UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*UpdateGroupMemberResponse, error)
	RemoveGroupMember(ctx context.Context, in *RemoveGroupMemberRequest, opts ...grpc.CallOption) (*RemoveGroupMemberResponse, error)
	GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error)
	GetUserGroups(ctx context.Context, in *GetUserGroupsRequest, opts ...grpc.CallOption) (*GetUserGroupsResponse, error)
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	out := new(UpdateGroupResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	out := new(GetGroupsResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) AddGroupMember(ctx context.Context, in *AddGroupMemberRequest, opts ...grpc.CallOption) (*AddGroupMemberResponse, error) {
	out := new(AddGroupMemberResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/AddGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*UpdateGroupMemberResponse, error) {
	out := new(UpdateGroupMemberResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/UpdateGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) RemoveGroupMember(ctx context.Context, in *RemoveGroupMemberRequest, opts ...grpc.CallOption) (*RemoveGroupMemberResponse, error) {
	out := new(RemoveGroupMemberResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/RemoveGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroupMembers(ctx context.Context, in *GetGroupMembersRequest, opts ...grpc.CallOption) (*GetGroupMembersResponse, error) {
	out := new(GetGroupMembersResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/GetGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetUserGroups(ctx context.Context, in *GetUserGroupsRequest, opts ...grpc.CallOption) (*GetUserGroupsResponse, error) {
	out := new(GetUserGroupsResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.GroupService/GetUserGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServiceServer is the server API for GroupService service.
// All implementations must embed UnimplementedGroupServiceServer
// for forward compatibility
type GroupServiceServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error)
	AddGroupMember(context.Context, *AddGroupMemberRequest) (*AddGroupMemberResponse, error)
	UpdateGroupMember(context.Context, *UpdateGroupMemberRequest) (*UpdateGroupMemberResponse, error)
	RemoveGroupMember(context.Context, *RemoveGroupMemberRequest) (*RemoveGroupMemberResponse, error)
	GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error)
	GetUserGroups(context.Context, *GetUserGroupsRequest) (*GetUserGroupsResponse, error)
	mustEmbedUnimplementedGroupServiceServer()
}

// UnimplementedGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func (UnimplementedGroupServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupServiceServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGroupServiceServer) GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedGroupServiceServer) AddGroupMember(context.Context, *AddGroupMemberRequest) (*AddGroupMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupMember not implemented")
}
func (UnimplementedGroupServiceServer) UpdateGroupMember(context.Context, *UpdateGroupMemberRequest) (*UpdateGroupMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMember not implemented")
}
func (UnimplementedGroupServiceServer) RemoveGroupMember(context.Context, *RemoveGroupMemberRequest) (*RemoveGroupMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupMember not implemented")
}
func (UnimplementedGroupServiceServer) GetGroupMembers(context.Context, *GetGroupMembersRequest) (*GetGroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMembers not implemented")
}
func (UnimplementedGroupServiceServer) GetUserGroups(context.Context, *GetUserGroupsRequest) (*GetUserGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroups not implemented")
}
func (UnimplementedGroupServiceServer) mustEmbedUnimplementedGroupServiceServer() {}

// UnsafeGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServiceServer will
// result in compilation errors.
type UnsafeGroupServiceServer interface {
	mustEmbedUnimplementedGroupServiceServer()
}

func RegisterGroupServiceServer(s grpc.ServiceRegistrar, srv GroupServiceServer) {
	s.RegisterService(&GroupService_ServiceDesc, srv)
}

func _GroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroups(ctx, req.(*GetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_AddGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).AddGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/AddGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).AddGroupMember(ctx, req.(*AddGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_UpdateGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).UpdateGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/UpdateGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).UpdateGroupMember(ctx, req.(*UpdateGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_RemoveGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).RemoveGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/RemoveGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).RemoveGroupMember(ctx, req.(*RemoveGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/GetGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetGroupMembers(ctx, req.(*GetGroupMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_GetUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).GetUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.GroupService/GetUserGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).GetUserGroups(ctx, req.(*GetUserGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupService_ServiceDesc is the grpc.ServiceDesc for GroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.v1.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _GroupService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupService_DeleteGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupService_GetGroup_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _GroupService_GetGroups_Handler,
		},
		{
			MethodName: "AddGroupMember",
			Handler:    _GroupService_AddGroupMember_Handler,
		},
		{
			MethodName: "UpdateGroupMember",
			Handler:    _GroupService_UpdateGroupMember_Handler,
		},
		{
			MethodName: "RemoveGroupMember",
			Handler:    _GroupService_RemoveGroupMember_Handler,
		},
		{
			MethodName: "GetGroupMembers",
			Handler:    _GroupService_GetGroupMembers_Handler,
		},
		{
			MethodName: "GetUserGroups",
			Handler:    _GroupService_GetUserGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication/v1/groups.proto",
}
