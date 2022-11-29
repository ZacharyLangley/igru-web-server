// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gardens/v1/strains.proto

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

type CreateStrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Comment    string  `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Notes      string  `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
	Type       string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Price      float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	ThcPercent float64 `protobuf:"fixed64,6,opt,name=thc_percent,json=thcPercent,proto3" json:"thc_percent,omitempty"`
	CbdPercent float64 `protobuf:"fixed64,7,opt,name=cbd_percent,json=cbdPercent,proto3" json:"cbd_percent,omitempty"`
	Parentage  string  `protobuf:"bytes,8,opt,name=parentage,proto3" json:"parentage,omitempty"`
	Aroma      string  `protobuf:"bytes,9,opt,name=aroma,proto3" json:"aroma,omitempty"`
	Taste      string  `protobuf:"bytes,10,opt,name=taste,proto3" json:"taste,omitempty"`
	Tags       string  `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateStrainRequest) Reset() {
	*x = CreateStrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStrainRequest) ProtoMessage() {}

func (x *CreateStrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStrainRequest.ProtoReflect.Descriptor instead.
func (*CreateStrainRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStrainRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStrainRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *CreateStrainRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *CreateStrainRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateStrainRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateStrainRequest) GetThcPercent() float64 {
	if x != nil {
		return x.ThcPercent
	}
	return 0
}

func (x *CreateStrainRequest) GetCbdPercent() float64 {
	if x != nil {
		return x.CbdPercent
	}
	return 0
}

func (x *CreateStrainRequest) GetParentage() string {
	if x != nil {
		return x.Parentage
	}
	return ""
}

func (x *CreateStrainRequest) GetAroma() string {
	if x != nil {
		return x.Aroma
	}
	return ""
}

func (x *CreateStrainRequest) GetTaste() string {
	if x != nil {
		return x.Taste
	}
	return ""
}

func (x *CreateStrainRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type CreateStrainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strain *Strain `protobuf:"bytes,1,opt,name=strain,proto3" json:"strain,omitempty"`
}

func (x *CreateStrainResponse) Reset() {
	*x = CreateStrainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStrainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStrainResponse) ProtoMessage() {}

func (x *CreateStrainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStrainResponse.ProtoReflect.Descriptor instead.
func (*CreateStrainResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStrainResponse) GetStrain() *Strain {
	if x != nil {
		return x.Strain
	}
	return nil
}

type DeleteStrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStrainRequest) Reset() {
	*x = DeleteStrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStrainRequest) ProtoMessage() {}

func (x *DeleteStrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStrainRequest.ProtoReflect.Descriptor instead.
func (*DeleteStrainRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteStrainRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteStrainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strain *Strain `protobuf:"bytes,1,opt,name=strain,proto3" json:"strain,omitempty"`
}

func (x *DeleteStrainResponse) Reset() {
	*x = DeleteStrainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStrainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStrainResponse) ProtoMessage() {}

func (x *DeleteStrainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStrainResponse.ProtoReflect.Descriptor instead.
func (*DeleteStrainResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteStrainResponse) GetStrain() *Strain {
	if x != nil {
		return x.Strain
	}
	return nil
}

type UpdateStrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Comment    string  `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Notes      string  `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
	Type       string  `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Price      float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	ThcPercent float64 `protobuf:"fixed64,7,opt,name=thc_percent,json=thcPercent,proto3" json:"thc_percent,omitempty"`
	CbdPercent float64 `protobuf:"fixed64,8,opt,name=cbd_percent,json=cbdPercent,proto3" json:"cbd_percent,omitempty"`
	Parentage  string  `protobuf:"bytes,9,opt,name=parentage,proto3" json:"parentage,omitempty"`
	Aroma      string  `protobuf:"bytes,10,opt,name=aroma,proto3" json:"aroma,omitempty"`
	Taste      string  `protobuf:"bytes,11,opt,name=taste,proto3" json:"taste,omitempty"`
	Tags       string  `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateStrainRequest) Reset() {
	*x = UpdateStrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStrainRequest) ProtoMessage() {}

func (x *UpdateStrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStrainRequest.ProtoReflect.Descriptor instead.
func (*UpdateStrainRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStrainRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStrainRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStrainRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *UpdateStrainRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *UpdateStrainRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateStrainRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateStrainRequest) GetThcPercent() float64 {
	if x != nil {
		return x.ThcPercent
	}
	return 0
}

func (x *UpdateStrainRequest) GetCbdPercent() float64 {
	if x != nil {
		return x.CbdPercent
	}
	return 0
}

func (x *UpdateStrainRequest) GetParentage() string {
	if x != nil {
		return x.Parentage
	}
	return ""
}

func (x *UpdateStrainRequest) GetAroma() string {
	if x != nil {
		return x.Aroma
	}
	return ""
}

func (x *UpdateStrainRequest) GetTaste() string {
	if x != nil {
		return x.Taste
	}
	return ""
}

func (x *UpdateStrainRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type UpdateStrainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strain *Strain `protobuf:"bytes,1,opt,name=strain,proto3" json:"strain,omitempty"`
}

func (x *UpdateStrainResponse) Reset() {
	*x = UpdateStrainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStrainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStrainResponse) ProtoMessage() {}

func (x *UpdateStrainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStrainResponse.ProtoReflect.Descriptor instead.
func (*UpdateStrainResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateStrainResponse) GetStrain() *Strain {
	if x != nil {
		return x.Strain
	}
	return nil
}

type GetStrainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStrainsRequest) Reset() {
	*x = GetStrainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStrainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStrainsRequest) ProtoMessage() {}

func (x *GetStrainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStrainsRequest.ProtoReflect.Descriptor instead.
func (*GetStrainsRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{6}
}

type GetStrainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strains []*Strain `protobuf:"bytes,1,rep,name=strains,proto3" json:"strains,omitempty"`
}

func (x *GetStrainsResponse) Reset() {
	*x = GetStrainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStrainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStrainsResponse) ProtoMessage() {}

func (x *GetStrainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStrainsResponse.ProtoReflect.Descriptor instead.
func (*GetStrainsResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{7}
}

func (x *GetStrainsResponse) GetStrains() []*Strain {
	if x != nil {
		return x.Strains
	}
	return nil
}

type GetStrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStrainRequest) Reset() {
	*x = GetStrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStrainRequest) ProtoMessage() {}

func (x *GetStrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStrainRequest.ProtoReflect.Descriptor instead.
func (*GetStrainRequest) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{8}
}

func (x *GetStrainRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetStrainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strain *Strain `protobuf:"bytes,1,opt,name=strain,proto3" json:"strain,omitempty"`
}

func (x *GetStrainResponse) Reset() {
	*x = GetStrainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_strains_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStrainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStrainResponse) ProtoMessage() {}

func (x *GetStrainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_strains_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStrainResponse.ProtoReflect.Descriptor instead.
func (*GetStrainResponse) Descriptor() ([]byte, []int) {
	return file_gardens_v1_strains_proto_rawDescGZIP(), []int{9}
}

func (x *GetStrainResponse) GetStrain() *Strain {
	if x != nil {
		return x.Strain
	}
	return nil
}

var File_gardens_v1_strains_proto protoreflect.FileDescriptor

var file_gardens_v1_strains_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa3, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68, 0x63, 0x5f, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x68, 0x63, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x62, 0x64, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x62, 0x64,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x73, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x52, 0x06, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x42, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x22, 0xb3, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68,
	0x63, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x74, 0x68, 0x63, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x62, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x63, 0x62, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72,
	0x6f, 0x6d, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x72, 0x6f, 0x6d, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x61, 0x73, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x07,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x32, 0xaa, 0x03, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x53, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x12,
	0x1f, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x61, 0x72,
	0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb1, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69,
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
	file_gardens_v1_strains_proto_rawDescOnce sync.Once
	file_gardens_v1_strains_proto_rawDescData = file_gardens_v1_strains_proto_rawDesc
)

func file_gardens_v1_strains_proto_rawDescGZIP() []byte {
	file_gardens_v1_strains_proto_rawDescOnce.Do(func() {
		file_gardens_v1_strains_proto_rawDescData = protoimpl.X.CompressGZIP(file_gardens_v1_strains_proto_rawDescData)
	})
	return file_gardens_v1_strains_proto_rawDescData
}

var file_gardens_v1_strains_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gardens_v1_strains_proto_goTypes = []interface{}{
	(*CreateStrainRequest)(nil),  // 0: gardens.v1.CreateStrainRequest
	(*CreateStrainResponse)(nil), // 1: gardens.v1.CreateStrainResponse
	(*DeleteStrainRequest)(nil),  // 2: gardens.v1.DeleteStrainRequest
	(*DeleteStrainResponse)(nil), // 3: gardens.v1.DeleteStrainResponse
	(*UpdateStrainRequest)(nil),  // 4: gardens.v1.UpdateStrainRequest
	(*UpdateStrainResponse)(nil), // 5: gardens.v1.UpdateStrainResponse
	(*GetStrainsRequest)(nil),    // 6: gardens.v1.GetStrainsRequest
	(*GetStrainsResponse)(nil),   // 7: gardens.v1.GetStrainsResponse
	(*GetStrainRequest)(nil),     // 8: gardens.v1.GetStrainRequest
	(*GetStrainResponse)(nil),    // 9: gardens.v1.GetStrainResponse
	(*Strain)(nil),               // 10: gardens.v1.Strain
}
var file_gardens_v1_strains_proto_depIdxs = []int32{
	10, // 0: gardens.v1.CreateStrainResponse.strain:type_name -> gardens.v1.Strain
	10, // 1: gardens.v1.DeleteStrainResponse.strain:type_name -> gardens.v1.Strain
	10, // 2: gardens.v1.UpdateStrainResponse.strain:type_name -> gardens.v1.Strain
	10, // 3: gardens.v1.GetStrainsResponse.strains:type_name -> gardens.v1.Strain
	10, // 4: gardens.v1.GetStrainResponse.strain:type_name -> gardens.v1.Strain
	0,  // 5: gardens.v1.StrainsService.CreateStrain:input_type -> gardens.v1.CreateStrainRequest
	2,  // 6: gardens.v1.StrainsService.DeleteStrain:input_type -> gardens.v1.DeleteStrainRequest
	4,  // 7: gardens.v1.StrainsService.UpdateStrain:input_type -> gardens.v1.UpdateStrainRequest
	6,  // 8: gardens.v1.StrainsService.GetStrains:input_type -> gardens.v1.GetStrainsRequest
	8,  // 9: gardens.v1.StrainsService.GetStrain:input_type -> gardens.v1.GetStrainRequest
	1,  // 10: gardens.v1.StrainsService.CreateStrain:output_type -> gardens.v1.CreateStrainResponse
	3,  // 11: gardens.v1.StrainsService.DeleteStrain:output_type -> gardens.v1.DeleteStrainResponse
	5,  // 12: gardens.v1.StrainsService.UpdateStrain:output_type -> gardens.v1.UpdateStrainResponse
	7,  // 13: gardens.v1.StrainsService.GetStrains:output_type -> gardens.v1.GetStrainsResponse
	9,  // 14: gardens.v1.StrainsService.GetStrain:output_type -> gardens.v1.GetStrainResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_gardens_v1_strains_proto_init() }
func file_gardens_v1_strains_proto_init() {
	if File_gardens_v1_strains_proto != nil {
		return
	}
	file_gardens_v1_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gardens_v1_strains_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStrainRequest); i {
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
		file_gardens_v1_strains_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStrainResponse); i {
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
		file_gardens_v1_strains_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStrainRequest); i {
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
		file_gardens_v1_strains_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStrainResponse); i {
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
		file_gardens_v1_strains_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStrainRequest); i {
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
		file_gardens_v1_strains_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStrainResponse); i {
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
		file_gardens_v1_strains_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStrainsRequest); i {
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
		file_gardens_v1_strains_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStrainsResponse); i {
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
		file_gardens_v1_strains_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStrainRequest); i {
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
		file_gardens_v1_strains_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStrainResponse); i {
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
			RawDescriptor: file_gardens_v1_strains_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gardens_v1_strains_proto_goTypes,
		DependencyIndexes: file_gardens_v1_strains_proto_depIdxs,
		MessageInfos:      file_gardens_v1_strains_proto_msgTypes,
	}.Build()
	File_gardens_v1_strains_proto = out.File
	file_gardens_v1_strains_proto_rawDesc = nil
	file_gardens_v1_strains_proto_goTypes = nil
	file_gardens_v1_strains_proto_depIdxs = nil
}
