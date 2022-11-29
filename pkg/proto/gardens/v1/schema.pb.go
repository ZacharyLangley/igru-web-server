// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gardens/v1/schema.proto

package gardensv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Garden struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Comment       string                 `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Location      string                 `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	GrowType      string                 `protobuf:"bytes,5,opt,name=grow_type,json=growType,proto3" json:"grow_type,omitempty"`
	GrowSize      string                 `protobuf:"bytes,6,opt,name=grow_size,json=growSize,proto3" json:"grow_size,omitempty"`
	GrowStyle     string                 `protobuf:"bytes,7,opt,name=grow_style,json=growStyle,proto3" json:"grow_style,omitempty"`
	ContainerSize string                 `protobuf:"bytes,8,opt,name=container_size,json=containerSize,proto3" json:"container_size,omitempty"`
	Tags          string                 `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Garden) Reset() {
	*x = Garden{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Garden) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Garden) ProtoMessage() {}

func (x *Garden) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Garden.ProtoReflect.Descriptor instead.
func (*Garden) Descriptor() ([]byte, []int) {
	return file_gardens_v1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *Garden) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Garden) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Garden) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Garden) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Garden) GetGrowType() string {
	if x != nil {
		return x.GrowType
	}
	return ""
}

func (x *Garden) GetGrowSize() string {
	if x != nil {
		return x.GrowSize
	}
	return ""
}

func (x *Garden) GetGrowStyle() string {
	if x != nil {
		return x.GrowStyle
	}
	return ""
}

func (x *Garden) GetContainerSize() string {
	if x != nil {
		return x.ContainerSize
	}
	return ""
}

func (x *Garden) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Garden) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Garden) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Plant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Comment         string                 `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Notes           string                 `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
	GrowCycleLength string                 `protobuf:"bytes,5,opt,name=grow_cycle_length,json=growCycleLength,proto3" json:"grow_cycle_length,omitempty"`
	Parentage       string                 `protobuf:"bytes,6,opt,name=parentage,proto3" json:"parentage,omitempty"`
	Origin          string                 `protobuf:"bytes,7,opt,name=origin,proto3" json:"origin,omitempty"`
	Gender          string                 `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
	DaysFlowering   float64                `protobuf:"fixed64,9,opt,name=days_flowering,json=daysFlowering,proto3" json:"days_flowering,omitempty"`
	DaysCured       float64                `protobuf:"fixed64,10,opt,name=days_cured,json=daysCured,proto3" json:"days_cured,omitempty"`
	HarvestedWeight string                 `protobuf:"bytes,11,opt,name=harvested_weight,json=harvestedWeight,proto3" json:"harvested_weight,omitempty"`
	BudDensity      float64                `protobuf:"fixed64,12,opt,name=bud_density,json=budDensity,proto3" json:"bud_density,omitempty"`
	BudPistils      bool                   `protobuf:"varint,13,opt,name=bud_pistils,json=budPistils,proto3" json:"bud_pistils,omitempty"`
	Tags            string                 `protobuf:"bytes,14,opt,name=tags,proto3" json:"tags,omitempty"`
	AcquiredAt      *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=acquired_at,json=acquiredAt,proto3" json:"acquired_at,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Plant) Reset() {
	*x = Plant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plant) ProtoMessage() {}

func (x *Plant) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plant.ProtoReflect.Descriptor instead.
func (*Plant) Descriptor() ([]byte, []int) {
	return file_gardens_v1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Plant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plant) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Plant) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Plant) GetGrowCycleLength() string {
	if x != nil {
		return x.GrowCycleLength
	}
	return ""
}

func (x *Plant) GetParentage() string {
	if x != nil {
		return x.Parentage
	}
	return ""
}

func (x *Plant) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Plant) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Plant) GetDaysFlowering() float64 {
	if x != nil {
		return x.DaysFlowering
	}
	return 0
}

func (x *Plant) GetDaysCured() float64 {
	if x != nil {
		return x.DaysCured
	}
	return 0
}

func (x *Plant) GetHarvestedWeight() string {
	if x != nil {
		return x.HarvestedWeight
	}
	return ""
}

func (x *Plant) GetBudDensity() float64 {
	if x != nil {
		return x.BudDensity
	}
	return 0
}

func (x *Plant) GetBudPistils() bool {
	if x != nil {
		return x.BudPistils
	}
	return false
}

func (x *Plant) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Plant) GetAcquiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AcquiredAt
	}
	return nil
}

func (x *Plant) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Plant) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Strain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Comment    string                 `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Notes      string                 `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
	Type       string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Price      float64                `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	ThcPercent float64                `protobuf:"fixed64,7,opt,name=thc_percent,json=thcPercent,proto3" json:"thc_percent,omitempty"`
	CbdPercent float64                `protobuf:"fixed64,8,opt,name=cbd_percent,json=cbdPercent,proto3" json:"cbd_percent,omitempty"`
	Parentage  string                 `protobuf:"bytes,9,opt,name=parentage,proto3" json:"parentage,omitempty"`
	Aroma      string                 `protobuf:"bytes,10,opt,name=aroma,proto3" json:"aroma,omitempty"`
	Taste      string                 `protobuf:"bytes,11,opt,name=taste,proto3" json:"taste,omitempty"`
	Tags       string                 `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Strain) Reset() {
	*x = Strain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gardens_v1_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strain) ProtoMessage() {}

func (x *Strain) ProtoReflect() protoreflect.Message {
	mi := &file_gardens_v1_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strain.ProtoReflect.Descriptor instead.
func (*Strain) Descriptor() ([]byte, []int) {
	return file_gardens_v1_schema_proto_rawDescGZIP(), []int{2}
}

func (x *Strain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Strain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Strain) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Strain) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Strain) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Strain) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Strain) GetThcPercent() float64 {
	if x != nil {
		return x.ThcPercent
	}
	return 0
}

func (x *Strain) GetCbdPercent() float64 {
	if x != nil {
		return x.CbdPercent
	}
	return 0
}

func (x *Strain) GetParentage() string {
	if x != nil {
		return x.Parentage
	}
	return ""
}

func (x *Strain) GetAroma() string {
	if x != nil {
		return x.Aroma
	}
	return ""
}

func (x *Strain) GetTaste() string {
	if x != nil {
		return x.Taste
	}
	return ""
}

func (x *Strain) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Strain) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Strain) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_gardens_v1_schema_proto protoreflect.FileDescriptor

var file_gardens_v1_schema_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x06, 0x47, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x72, 0x6f, 0x77, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x72, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x77,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f,
	0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x77, 0x53,
	0x74, 0x79, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xcf, 0x04, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x67, 0x72, 0x6f, 0x77, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x46, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x63, 0x75, 0x72,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x64, 0x61, 0x79, 0x73, 0x43, 0x75,
	0x72, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x68, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x68,
	0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x75, 0x64, 0x5f, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x62, 0x75, 0x64, 0x44, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x64, 0x5f, 0x70, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x75, 0x64, 0x50, 0x69, 0x73, 0x74, 0x69, 0x6c, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9c, 0x03, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x68, 0x63, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x62, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x62, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x72, 0x6f, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x73, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xb0, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x61, 0x63, 0x68, 0x61, 0x72, 0x79, 0x4c, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x79, 0x2f, 0x69, 0x67, 0x72, 0x75, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x47, 0x61, 0x72, 0x64, 0x65,
	0x6e, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x16, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x47, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gardens_v1_schema_proto_rawDescOnce sync.Once
	file_gardens_v1_schema_proto_rawDescData = file_gardens_v1_schema_proto_rawDesc
)

func file_gardens_v1_schema_proto_rawDescGZIP() []byte {
	file_gardens_v1_schema_proto_rawDescOnce.Do(func() {
		file_gardens_v1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_gardens_v1_schema_proto_rawDescData)
	})
	return file_gardens_v1_schema_proto_rawDescData
}

var file_gardens_v1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gardens_v1_schema_proto_goTypes = []interface{}{
	(*Garden)(nil),                // 0: gardens.v1.Garden
	(*Plant)(nil),                 // 1: gardens.v1.Plant
	(*Strain)(nil),                // 2: gardens.v1.Strain
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_gardens_v1_schema_proto_depIdxs = []int32{
	3, // 0: gardens.v1.Garden.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: gardens.v1.Garden.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: gardens.v1.Plant.acquired_at:type_name -> google.protobuf.Timestamp
	3, // 3: gardens.v1.Plant.created_at:type_name -> google.protobuf.Timestamp
	3, // 4: gardens.v1.Plant.updated_at:type_name -> google.protobuf.Timestamp
	3, // 5: gardens.v1.Strain.created_at:type_name -> google.protobuf.Timestamp
	3, // 6: gardens.v1.Strain.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_gardens_v1_schema_proto_init() }
func file_gardens_v1_schema_proto_init() {
	if File_gardens_v1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gardens_v1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Garden); i {
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
		file_gardens_v1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plant); i {
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
		file_gardens_v1_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strain); i {
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
			RawDescriptor: file_gardens_v1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gardens_v1_schema_proto_goTypes,
		DependencyIndexes: file_gardens_v1_schema_proto_depIdxs,
		MessageInfos:      file_gardens_v1_schema_proto_msgTypes,
	}.Build()
	File_gardens_v1_schema_proto = out.File
	file_gardens_v1_schema_proto_rawDesc = nil
	file_gardens_v1_schema_proto_goTypes = nil
	file_gardens_v1_schema_proto_depIdxs = nil
}
