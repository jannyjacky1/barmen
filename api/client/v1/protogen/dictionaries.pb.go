// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        (unknown)
// source: dictionaries.proto

package protogen

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NameItem_Type int32

const (
	NameItem_COCKTAIL   NameItem_Type = 0
	NameItem_INGREDIENT NameItem_Type = 1
	NameItem_INSTRUMENT NameItem_Type = 2
)

// Enum value maps for NameItem_Type.
var (
	NameItem_Type_name = map[int32]string{
		0: "COCKTAIL",
		1: "INGREDIENT",
		2: "INSTRUMENT",
	}
	NameItem_Type_value = map[string]int32{
		"COCKTAIL":   0,
		"INGREDIENT": 1,
		"INSTRUMENT": 2,
	}
)

func (x NameItem_Type) Enum() *NameItem_Type {
	p := new(NameItem_Type)
	*p = x
	return p
}

func (x NameItem_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NameItem_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_dictionaries_proto_enumTypes[0].Descriptor()
}

func (NameItem_Type) Type() protoreflect.EnumType {
	return &file_dictionaries_proto_enumTypes[0]
}

func (x NameItem_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NameItem_Type.Descriptor instead.
func (NameItem_Type) EnumDescriptor() ([]byte, []int) {
	return file_dictionaries_proto_rawDescGZIP(), []int{4, 0}
}

type DictionariesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DictionariesRequest) Reset() {
	*x = DictionariesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionaries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictionariesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictionariesRequest) ProtoMessage() {}

func (x *DictionariesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dictionaries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictionariesRequest.ProtoReflect.Descriptor instead.
func (*DictionariesRequest) Descriptor() ([]byte, []int) {
	return file_dictionaries_proto_rawDescGZIP(), []int{0}
}

type Dictionary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Dictionary) Reset() {
	*x = Dictionary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionaries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dictionary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dictionary) ProtoMessage() {}

func (x *Dictionary) ProtoReflect() protoreflect.Message {
	mi := &file_dictionaries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dictionary.ProtoReflect.Descriptor instead.
func (*Dictionary) Descriptor() ([]byte, []int) {
	return file_dictionaries_proto_rawDescGZIP(), []int{1}
}

func (x *Dictionary) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dictionary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DictionariesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComplicationLevels []*Dictionary `protobuf:"bytes,1,rep,name=complication_levels,json=complicationLevels,proto3" json:"complication_levels,omitempty"`
	FortressLevels     []*Dictionary `protobuf:"bytes,2,rep,name=fortress_levels,json=fortressLevels,proto3" json:"fortress_levels,omitempty"`
	Volumes            []*Dictionary `protobuf:"bytes,3,rep,name=volumes,proto3" json:"volumes,omitempty"`
	Ingredients        []*Dictionary `protobuf:"bytes,4,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
	Other              []*Dictionary `protobuf:"bytes,5,rep,name=other,proto3" json:"other,omitempty"`
}

func (x *DictionariesResponse) Reset() {
	*x = DictionariesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionaries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictionariesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictionariesResponse) ProtoMessage() {}

func (x *DictionariesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dictionaries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictionariesResponse.ProtoReflect.Descriptor instead.
func (*DictionariesResponse) Descriptor() ([]byte, []int) {
	return file_dictionaries_proto_rawDescGZIP(), []int{2}
}

func (x *DictionariesResponse) GetComplicationLevels() []*Dictionary {
	if x != nil {
		return x.ComplicationLevels
	}
	return nil
}

func (x *DictionariesResponse) GetFortressLevels() []*Dictionary {
	if x != nil {
		return x.FortressLevels
	}
	return nil
}

func (x *DictionariesResponse) GetVolumes() []*Dictionary {
	if x != nil {
		return x.Volumes
	}
	return nil
}

func (x *DictionariesResponse) GetIngredients() []*Dictionary {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

func (x *DictionariesResponse) GetOther() []*Dictionary {
	if x != nil {
		return x.Other
	}
	return nil
}

type NameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Page int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *NameRequest) Reset() {
	*x = NameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionaries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameRequest) ProtoMessage() {}

func (x *NameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dictionaries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameRequest.ProtoReflect.Descriptor instead.
func (*NameRequest) Descriptor() ([]byte, []int) {
	return file_dictionaries_proto_rawDescGZIP(), []int{3}
}

func (x *NameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NameRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type NameItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type NameItem_Type `protobuf:"varint,3,opt,name=type,proto3,enum=proto.NameItem_Type" json:"type,omitempty"`
}

func (x *NameItem) Reset() {
	*x = NameItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionaries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameItem) ProtoMessage() {}

func (x *NameItem) ProtoReflect() protoreflect.Message {
	mi := &file_dictionaries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameItem.ProtoReflect.Descriptor instead.
func (*NameItem) Descriptor() ([]byte, []int) {
	return file_dictionaries_proto_rawDescGZIP(), []int{4}
}

func (x *NameItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NameItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NameItem) GetType() NameItem_Type {
	if x != nil {
		return x.Type
	}
	return NameItem_COCKTAIL
}

type NameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NameItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *NameResponse) Reset() {
	*x = NameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictionaries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameResponse) ProtoMessage() {}

func (x *NameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dictionaries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameResponse.ProtoReflect.Descriptor instead.
func (*NameResponse) Descriptor() ([]byte, []int) {
	return file_dictionaries_proto_rawDescGZIP(), []int{5}
}

func (x *NameResponse) GetItems() []*NameItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_dictionaries_proto protoreflect.FileDescriptor

var file_dictionaries_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x30, 0x0a, 0x0a, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x14, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x12, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x73, 0x12, 0x3a, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x74, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x0e, 0x66,
	0x6f, 0x72, 0x74, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x2b, 0x0a,
	0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x27, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x8e, 0x01, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x43, 0x4b, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x44, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02,
	0x22, 0x35, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x94, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dictionaries_proto_rawDescOnce sync.Once
	file_dictionaries_proto_rawDescData = file_dictionaries_proto_rawDesc
)

func file_dictionaries_proto_rawDescGZIP() []byte {
	file_dictionaries_proto_rawDescOnce.Do(func() {
		file_dictionaries_proto_rawDescData = protoimpl.X.CompressGZIP(file_dictionaries_proto_rawDescData)
	})
	return file_dictionaries_proto_rawDescData
}

var file_dictionaries_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dictionaries_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dictionaries_proto_goTypes = []interface{}{
	(NameItem_Type)(0),           // 0: proto.NameItem.Type
	(*DictionariesRequest)(nil),  // 1: proto.DictionariesRequest
	(*Dictionary)(nil),           // 2: proto.Dictionary
	(*DictionariesResponse)(nil), // 3: proto.DictionariesResponse
	(*NameRequest)(nil),          // 4: proto.NameRequest
	(*NameItem)(nil),             // 5: proto.NameItem
	(*NameResponse)(nil),         // 6: proto.NameResponse
}
var file_dictionaries_proto_depIdxs = []int32{
	2, // 0: proto.DictionariesResponse.complication_levels:type_name -> proto.Dictionary
	2, // 1: proto.DictionariesResponse.fortress_levels:type_name -> proto.Dictionary
	2, // 2: proto.DictionariesResponse.volumes:type_name -> proto.Dictionary
	2, // 3: proto.DictionariesResponse.ingredients:type_name -> proto.Dictionary
	2, // 4: proto.DictionariesResponse.other:type_name -> proto.Dictionary
	0, // 5: proto.NameItem.type:type_name -> proto.NameItem.Type
	5, // 6: proto.NameResponse.items:type_name -> proto.NameItem
	1, // 7: proto.Dictionaries.GetDictionaries:input_type -> proto.DictionariesRequest
	4, // 8: proto.Dictionaries.GetByName:input_type -> proto.NameRequest
	3, // 9: proto.Dictionaries.GetDictionaries:output_type -> proto.DictionariesResponse
	6, // 10: proto.Dictionaries.GetByName:output_type -> proto.NameResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_dictionaries_proto_init() }
func file_dictionaries_proto_init() {
	if File_dictionaries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dictionaries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictionariesRequest); i {
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
		file_dictionaries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dictionary); i {
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
		file_dictionaries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictionariesResponse); i {
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
		file_dictionaries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameRequest); i {
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
		file_dictionaries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameItem); i {
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
		file_dictionaries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameResponse); i {
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
			RawDescriptor: file_dictionaries_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dictionaries_proto_goTypes,
		DependencyIndexes: file_dictionaries_proto_depIdxs,
		EnumInfos:         file_dictionaries_proto_enumTypes,
		MessageInfos:      file_dictionaries_proto_msgTypes,
	}.Build()
	File_dictionaries_proto = out.File
	file_dictionaries_proto_rawDesc = nil
	file_dictionaries_proto_goTypes = nil
	file_dictionaries_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DictionariesClient is the client API for Dictionaries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DictionariesClient interface {
	GetDictionaries(ctx context.Context, in *DictionariesRequest, opts ...grpc.CallOption) (*DictionariesResponse, error)
	GetByName(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error)
}

type dictionariesClient struct {
	cc grpc.ClientConnInterface
}

func NewDictionariesClient(cc grpc.ClientConnInterface) DictionariesClient {
	return &dictionariesClient{cc}
}

func (c *dictionariesClient) GetDictionaries(ctx context.Context, in *DictionariesRequest, opts ...grpc.CallOption) (*DictionariesResponse, error) {
	out := new(DictionariesResponse)
	err := c.cc.Invoke(ctx, "/proto.Dictionaries/GetDictionaries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionariesClient) GetByName(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/proto.Dictionaries/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictionariesServer is the server API for Dictionaries service.
type DictionariesServer interface {
	GetDictionaries(context.Context, *DictionariesRequest) (*DictionariesResponse, error)
	GetByName(context.Context, *NameRequest) (*NameResponse, error)
}

// UnimplementedDictionariesServer can be embedded to have forward compatible implementations.
type UnimplementedDictionariesServer struct {
}

func (*UnimplementedDictionariesServer) GetDictionaries(context.Context, *DictionariesRequest) (*DictionariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDictionaries not implemented")
}
func (*UnimplementedDictionariesServer) GetByName(context.Context, *NameRequest) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}

func RegisterDictionariesServer(s *grpc.Server, srv DictionariesServer) {
	s.RegisterService(&_Dictionaries_serviceDesc, srv)
}

func _Dictionaries_GetDictionaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictionariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionariesServer).GetDictionaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dictionaries/GetDictionaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionariesServer).GetDictionaries(ctx, req.(*DictionariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dictionaries_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionariesServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dictionaries/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionariesServer).GetByName(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dictionaries_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dictionaries",
	HandlerType: (*DictionariesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDictionaries",
			Handler:    _Dictionaries_GetDictionaries_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _Dictionaries_GetByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dictionaries.proto",
}
