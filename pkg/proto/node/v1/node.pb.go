// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: node/v1/node.proto

package nodev1

import (
	v1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MacAddress   string            `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Name         string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CustomLabels map[string]string `protobuf:"bytes,4,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateNodeRequest) Reset() {
	*x = UpdateNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeRequest) ProtoMessage() {}

func (x *UpdateNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateNodeRequest) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *UpdateNodeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNodeRequest) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

type UpdateNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *UpdateNodeResponse) Reset() {
	*x = UpdateNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeResponse) ProtoMessage() {}

func (x *UpdateNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeResponse.ProtoReflect.Descriptor instead.
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateNodeResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type DeleteNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    *string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3,oneof" json:"group_id,omitempty"`
	MacAddress string  `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
}

func (x *DeleteNodeRequest) Reset() {
	*x = DeleteNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeRequest) ProtoMessage() {}

func (x *DeleteNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteNodeRequest) GetGroupId() string {
	if x != nil && x.GroupId != nil {
		return *x.GroupId
	}
	return ""
}

func (x *DeleteNodeRequest) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

type DeleteNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNodeResponse) Reset() {
	*x = DeleteNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeResponse) ProtoMessage() {}

func (x *DeleteNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeResponse.ProtoReflect.Descriptor instead.
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{3}
}

type GetNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *v1.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	GroupId    *string               `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3,oneof" json:"group_id,omitempty"`
}

func (x *GetNodesRequest) Reset() {
	*x = GetNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesRequest) ProtoMessage() {}

func (x *GetNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesRequest.ProtoReflect.Descriptor instead.
func (*GetNodesRequest) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{4}
}

func (x *GetNodesRequest) GetPagination() *v1.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetNodesRequest) GetGroupId() string {
	if x != nil && x.GroupId != nil {
		return *x.GroupId
	}
	return ""
}

type GetNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GetNodesResponse) Reset() {
	*x = GetNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesResponse) ProtoMessage() {}

func (x *GetNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesResponse.ProtoReflect.Descriptor instead.
func (*GetNodesResponse) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{5}
}

func (x *GetNodesResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    *string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3,oneof" json:"group_id,omitempty"`
	MacAddress string  `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{6}
}

func (x *GetNodeRequest) GetGroupId() string {
	if x != nil && x.GroupId != nil {
		return *x.GroupId
	}
	return ""
}

func (x *GetNodeRequest) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

type GetNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GetNodeResponse) Reset() {
	*x = GetNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeResponse) ProtoMessage() {}

func (x *GetNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeResponse.ProtoReflect.Descriptor instead.
func (*GetNodeResponse) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{7}
}

func (x *GetNodeResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

var File_node_v1_node_proto protoreflect.FileDescriptor

var file_node_v1_node_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdc, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x3f, 0x0a,
	0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x61, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x7c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x37,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x32, 0xa2, 0x02,
	0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x99, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x61, 0x63, 0x68,
	0x61, 0x72, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x79, 0x2f, 0x69, 0x67, 0x72, 0x75, 0x2d,
	0x77, 0x65, 0x62, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x64,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4e, 0x6f, 0x64, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x4e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_v1_node_proto_rawDescOnce sync.Once
	file_node_v1_node_proto_rawDescData = file_node_v1_node_proto_rawDesc
)

func file_node_v1_node_proto_rawDescGZIP() []byte {
	file_node_v1_node_proto_rawDescOnce.Do(func() {
		file_node_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_v1_node_proto_rawDescData)
	})
	return file_node_v1_node_proto_rawDescData
}

var file_node_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_node_v1_node_proto_goTypes = []interface{}{
	(*UpdateNodeRequest)(nil),    // 0: node.v1.UpdateNodeRequest
	(*UpdateNodeResponse)(nil),   // 1: node.v1.UpdateNodeResponse
	(*DeleteNodeRequest)(nil),    // 2: node.v1.DeleteNodeRequest
	(*DeleteNodeResponse)(nil),   // 3: node.v1.DeleteNodeResponse
	(*GetNodesRequest)(nil),      // 4: node.v1.GetNodesRequest
	(*GetNodesResponse)(nil),     // 5: node.v1.GetNodesResponse
	(*GetNodeRequest)(nil),       // 6: node.v1.GetNodeRequest
	(*GetNodeResponse)(nil),      // 7: node.v1.GetNodeResponse
	nil,                          // 8: node.v1.UpdateNodeRequest.CustomLabelsEntry
	(*Node)(nil),                 // 9: node.v1.Node
	(*v1.PaginationRequest)(nil), // 10: common.v1.PaginationRequest
}
var file_node_v1_node_proto_depIdxs = []int32{
	8,  // 0: node.v1.UpdateNodeRequest.custom_labels:type_name -> node.v1.UpdateNodeRequest.CustomLabelsEntry
	9,  // 1: node.v1.UpdateNodeResponse.node:type_name -> node.v1.Node
	10, // 2: node.v1.GetNodesRequest.pagination:type_name -> common.v1.PaginationRequest
	9,  // 3: node.v1.GetNodesResponse.nodes:type_name -> node.v1.Node
	9,  // 4: node.v1.GetNodeResponse.node:type_name -> node.v1.Node
	0,  // 5: node.v1.NodeService.UpdateNode:input_type -> node.v1.UpdateNodeRequest
	2,  // 6: node.v1.NodeService.DeleteNode:input_type -> node.v1.DeleteNodeRequest
	4,  // 7: node.v1.NodeService.GetNodes:input_type -> node.v1.GetNodesRequest
	6,  // 8: node.v1.NodeService.GetNode:input_type -> node.v1.GetNodeRequest
	1,  // 9: node.v1.NodeService.UpdateNode:output_type -> node.v1.UpdateNodeResponse
	3,  // 10: node.v1.NodeService.DeleteNode:output_type -> node.v1.DeleteNodeResponse
	5,  // 11: node.v1.NodeService.GetNodes:output_type -> node.v1.GetNodesResponse
	7,  // 12: node.v1.NodeService.GetNode:output_type -> node.v1.GetNodeResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_node_v1_node_proto_init() }
func file_node_v1_node_proto_init() {
	if File_node_v1_node_proto != nil {
		return
	}
	file_node_v1_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_node_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_v1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_v1_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_v1_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_v1_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_v1_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_v1_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_v1_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_node_v1_node_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_node_v1_node_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_node_v1_node_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_v1_node_proto_goTypes,
		DependencyIndexes: file_node_v1_node_proto_depIdxs,
		MessageInfos:      file_node_v1_node_proto_msgTypes,
	}.Build()
	File_node_v1_node_proto = out.File
	file_node_v1_node_proto_rawDesc = nil
	file_node_v1_node_proto_goTypes = nil
	file_node_v1_node_proto_depIdxs = nil
}
