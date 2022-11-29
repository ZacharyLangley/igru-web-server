// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gardens/v1/gardens.proto

package gardensv1

import (
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

type CreateGardenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Comment       string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Location      string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	GrowType      string `protobuf:"bytes,4,opt,name=grow_type,json=growType,proto3" json:"grow_type,omitempty"`
	GrowSize      string `protobuf:"bytes,5,opt,name=grow_size,json=growSize,proto3" json:"grow_size,omitempty"`
	GrowStyle     string `protobuf:"bytes,6,opt,name=grow_style,json=growStyle,proto3" json:"grow_style,omitempty"`
	ContainerSize string `protobuf:"bytes,7,opt,name=container_size,json=containerSize,proto3" json:"container_size,omitempty"`
	Tags          string `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateGardenRequest) Reset() {
	*x = CreateGardenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGardenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGardenRequest) ProtoMessage() {}

func (x *CreateGardenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGardenRequest.ProtoReflect.Descriptor instead.
func (*CreateGardenRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGardenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGardenRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CreateGardenRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateGardenRequest) GetGrowType() string {
	if x != nil {
		return x.GrowType
	}
	return ""
}

func (x *CreateGardenRequest) GetGrowSize() string {
	if x != nil {
		return x.GrowSize
	}
	return ""
}

func (x *CreateGardenRequest) GetGrowStyle() string {
	if x != nil {
		return x.GrowStyle
	}
	return ""
}

func (x *CreateGardenRequest) GetContainerSize() string {
	if x != nil {
		return x.ContainerSize
	}
	return ""
}

func (x *CreateGardenRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type CreateGardenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Garden *Garden `protobuf:"bytes,1,opt,name=garden,proto3" json:"garden,omitempty"`
}

func (x *CreateGardenResponse) Reset() {
	*x = CreateGardenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGardenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGardenResponse) ProtoMessage() {}

func (x *CreateGardenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGardenResponse.ProtoReflect.Descriptor instead.
func (*CreateGardenResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGardenResponse) GetGarden() *Garden {
	if x != nil {
		return x.Garden
	}
	return nil
}

type DeleteGardenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGardenRequest) Reset() {
	*x = DeleteGardenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGardenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGardenRequest) ProtoMessage() {}

func (x *DeleteGardenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGardenRequest.ProtoReflect.Descriptor instead.
func (*DeleteGardenRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGardenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGardenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Garden *Garden `protobuf:"bytes,1,opt,name=garden,proto3" json:"garden,omitempty"`
}

func (x *DeleteGardenResponse) Reset() {
	*x = DeleteGardenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGardenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGardenResponse) ProtoMessage() {}

func (x *DeleteGardenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGardenResponse.ProtoReflect.Descriptor instead.
func (*DeleteGardenResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteGardenResponse) GetGarden() *Garden {
	if x != nil {
		return x.Garden
	}
	return nil
}

type UpdateGardenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Comment       string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Location      string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	GrowType      string `protobuf:"bytes,5,opt,name=grow_type,json=growType,proto3" json:"grow_type,omitempty"`
	GrowSize      string `protobuf:"bytes,6,opt,name=grow_size,json=growSize,proto3" json:"grow_size,omitempty"`
	GrowStyle     string `protobuf:"bytes,7,opt,name=grow_style,json=growStyle,proto3" json:"grow_style,omitempty"`
	ContainerSize string `protobuf:"bytes,8,opt,name=container_size,json=containerSize,proto3" json:"container_size,omitempty"`
	Tags          string `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateGardenRequest) Reset() {
	*x = UpdateGardenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGardenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGardenRequest) ProtoMessage() {}

func (x *UpdateGardenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGardenRequest.ProtoReflect.Descriptor instead.
func (*UpdateGardenRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGardenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateGardenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGardenRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *UpdateGardenRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UpdateGardenRequest) GetGrowType() string {
	if x != nil {
		return x.GrowType
	}
	return ""
}

func (x *UpdateGardenRequest) GetGrowSize() string {
	if x != nil {
		return x.GrowSize
	}
	return ""
}

func (x *UpdateGardenRequest) GetGrowStyle() string {
	if x != nil {
		return x.GrowStyle
	}
	return ""
}

func (x *UpdateGardenRequest) GetContainerSize() string {
	if x != nil {
		return x.ContainerSize
	}
	return ""
}

func (x *UpdateGardenRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type UpdateGardenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Garden *Garden `protobuf:"bytes,1,opt,name=garden,proto3" json:"garden,omitempty"`
}

func (x *UpdateGardenResponse) Reset() {
	*x = UpdateGardenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGardenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGardenResponse) ProtoMessage() {}

func (x *UpdateGardenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGardenResponse.ProtoReflect.Descriptor instead.
func (*UpdateGardenResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateGardenResponse) GetGarden() *Garden {
	if x != nil {
		return x.Garden
	}
	return nil
}

type GetGardensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGardensRequest) Reset() {
	*x = GetGardensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGardensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGardensRequest) ProtoMessage() {}

func (x *GetGardensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGardensRequest.ProtoReflect.Descriptor instead.
func (*GetGardensRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{6}
}

type GetGardensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gardens []*Garden `protobuf:"bytes,1,rep,name=gardens,proto3" json:"gardens,omitempty"`
}

func (x *GetGardensResponse) Reset() {
	*x = GetGardensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGardensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGardensResponse) ProtoMessage() {}

func (x *GetGardensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGardensResponse.ProtoReflect.Descriptor instead.
func (*GetGardensResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{7}
}

func (x *GetGardensResponse) GetGardens() []*Garden {
	if x != nil {
		return x.Gardens
	}
	return nil
}

type GetGardenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGardenRequest) Reset() {
	*x = GetGardenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGardenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGardenRequest) ProtoMessage() {}

func (x *GetGardenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGardenRequest.ProtoReflect.Descriptor instead.
func (*GetGardenRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{8}
}

func (x *GetGardenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGardenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Garden *Garden `protobuf:"bytes,1,opt,name=garden,proto3" json:"garden,omitempty"`
}

func (x *GetGardenResponse) Reset() {
	*x = GetGardenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_gardens_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGardenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGardenResponse) ProtoMessage() {}

func (x *GetGardenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_gardens_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGardenResponse.ProtoReflect.Descriptor instead.
func (*GetGardenResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_gardens_proto_rawDescGZIP(), []int{9}
}

func (x *GetGardenResponse) GetGarden() *Garden {
	if x != nil {
		return x.Garden
	}
	return nil
}

var File_gardens_v1_gardens_proto protoreflect.FileDescriptor

var file_gardens_v1_gardens_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf3, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67,
	0x72, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x77, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x06, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x52, 0x06, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x42, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x67, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x06, 0x67, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x22, 0x83, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x77, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x77, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x77, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x06, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x07,
	0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x52, 0x06, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x32, 0xaa, 0x03, 0x0a,
	0x0e, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x53, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x12,
	0x1f, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb1, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x47, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x61, 0x63, 0x68, 0x61, 0x72, 0x79,
	0x4c, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x79, 0x2f, 0x69, 0x67, 0x72, 0x75, 0x2d, 0x77, 0x65, 0x62,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x47,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x47, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0b, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gardens_v1_gardens_proto_rawDescOnce sync.Once
	file_gardens_v1_gardens_proto_rawDescData = file_gardens_v1_gardens_proto_rawDesc
)

func file_gardens_v1_gardens_proto_rawDescGZIP() []byte {
	file_gardens_v1_gardens_proto_rawDescOnce.Do(func() {
		file_gardens_v1_gardens_proto_rawDescData = protoimpl.X.CompressGZIP(file_gardens_v1_gardens_proto_rawDescData)
	})
	return file_gardens_v1_gardens_proto_rawDescData
}

var file_gardens_v1_gardens_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gardens_v1_gardens_proto_goTypes = []interface{}{
	(*CreateGardenRequest)(nil),  // 0: gardens.v1.CreateGardenRequest
	(*CreateGardenResponse)(nil), // 1: gardens.v1.CreateGardenResponse
	(*DeleteGardenRequest)(nil),  // 2: gardens.v1.DeleteGardenRequest
	(*DeleteGardenResponse)(nil), // 3: gardens.v1.DeleteGardenResponse
	(*UpdateGardenRequest)(nil),  // 4: gardens.v1.UpdateGardenRequest
	(*UpdateGardenResponse)(nil), // 5: gardens.v1.UpdateGardenResponse
	(*GetGardensRequest)(nil),    // 6: gardens.v1.GetGardensRequest
	(*GetGardensResponse)(nil),   // 7: gardens.v1.GetGardensResponse
	(*GetGardenRequest)(nil),     // 8: gardens.v1.GetGardenRequest
	(*GetGardenResponse)(nil),    // 9: gardens.v1.GetGardenResponse
	(*Garden)(nil),               // 10: gardens.v1.Garden
}
var file_gardens_v1_gardens_proto_depIdxs = []int32{
	10, // 0: gardens.v1.CreateGardenResponse.garden:type_name -> gardens.v1.Garden
	10, // 1: gardens.v1.DeleteGardenResponse.garden:type_name -> gardens.v1.Garden
	10, // 2: gardens.v1.UpdateGardenResponse.garden:type_name -> gardens.v1.Garden
	10, // 3: gardens.v1.GetGardensResponse.gardens:type_name -> gardens.v1.Garden
	10, // 4: gardens.v1.GetGardenResponse.garden:type_name -> gardens.v1.Garden
	0,  // 5: gardens.v1.GardensService.CreateGarden:input_type -> gardens.v1.CreateGardenRequest
	2,  // 6: gardens.v1.GardensService.DeleteGarden:input_type -> gardens.v1.DeleteGardenRequest
	4,  // 7: gardens.v1.GardensService.UpdateGarden:input_type -> gardens.v1.UpdateGardenRequest
	6,  // 8: gardens.v1.GardensService.GetGardens:input_type -> gardens.v1.GetGardensRequest
	8,  // 9: gardens.v1.GardensService.GetGarden:input_type -> gardens.v1.GetGardenRequest
	1,  // 10: gardens.v1.GardensService.CreateGarden:output_type -> gardens.v1.CreateGardenResponse
	3,  // 11: gardens.v1.GardensService.DeleteGarden:output_type -> gardens.v1.DeleteGardenResponse
	5,  // 12: gardens.v1.GardensService.UpdateGarden:output_type -> gardens.v1.UpdateGardenResponse
	7,  // 13: gardens.v1.GardensService.GetGardens:output_type -> gardens.v1.GetGardensResponse
	9,  // 14: gardens.v1.GardensService.GetGarden:output_type -> gardens.v1.GetGardenResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_gardens_v1_gardens_proto_init() }
func file_gardens_v1_gardens_proto_init() {
	if File_gardens_v1_gardens_proto != nil {
		return
	}
	file_gardens_v1_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gardens_v1_gardens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGardenRequest); i {
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
		file_gardens_v1_gardens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGardenResponse); i {
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
		file_gardens_v1_gardens_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGardenRequest); i {
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
		file_gardens_v1_gardens_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGardenResponse); i {
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
		file_gardens_v1_gardens_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGardenRequest); i {
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
		file_gardens_v1_gardens_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGardenResponse); i {
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
		file_gardens_v1_gardens_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGardensRequest); i {
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
		file_gardens_v1_gardens_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGardensResponse); i {
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
		file_gardens_v1_gardens_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGardenRequest); i {
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
		file_gardens_v1_gardens_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGardenResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gardens_v1_gardens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gardens_v1_gardens_proto_goTypes,
		DependencyIndexes: file_gardens_v1_gardens_proto_depIdxs,
		MessageInfos:      file_gardens_v1_gardens_proto_msgTypes,
	}.Build()
	File_gardens_v1_gardens_proto = out.File
	file_gardens_v1_gardens_proto_rawDesc = nil
	file_gardens_v1_gardens_proto_goTypes = nil
	file_gardens_v1_gardens_proto_depIdxs = nil
}
