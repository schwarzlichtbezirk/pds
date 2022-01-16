// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.19.3
// source: pb/pds.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Ping message content.
type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Port description.
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string    `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,9,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{1}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// Summary result of ports streaming.
type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of ports received.
	PortCount int32 `protobuf:"varint,1,opt,name=port_count,json=portCount,proto3" json:"port_count,omitempty"`
	// The duration of the traversal in milliseconds.
	ElapsedTime int32 `protobuf:"varint,2,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{2}
}

func (x *Summary) GetPortCount() int32 {
	if x != nil {
		return x.PortCount
	}
	return 0
}

func (x *Summary) GetElapsedTime() int32 {
	if x != nil {
		return x.ElapsedTime
	}
	return 0
}

// Port key.
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{3}
}

func (x *Key) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Port name.
type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{4}
}

func (x *Name) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Quest with text to find in object fields.
type Quest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Sensitive bool   `protobuf:"varint,2,opt,name=sensitive,proto3" json:"sensitive,omitempty"`
	Whole     bool   `protobuf:"varint,3,opt,name=whole,proto3" json:"whole,omitempty"`
}

func (x *Quest) Reset() {
	*x = Quest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quest) ProtoMessage() {}

func (x *Quest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quest.ProtoReflect.Descriptor instead.
func (*Quest) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{5}
}

func (x *Quest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Quest) GetSensitive() bool {
	if x != nil {
		return x.Sensitive
	}
	return false
}

func (x *Quest) GetWhole() bool {
	if x != nil {
		return x.Whole
	}
	return false
}

// Point with geo coordinates as latitude-longitude pair.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{6}
}

func (x *Point) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// Circle with center at given Point, and radius in meters.
type Circle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Center *Point  `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Radius float32 `protobuf:"fixed32,2,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *Circle) Reset() {
	*x = Circle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Circle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Circle) ProtoMessage() {}

func (x *Circle) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Circle.ProtoReflect.Descriptor instead.
func (*Circle) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{7}
}

func (x *Circle) GetCenter() *Point {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *Circle) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

// List on founded ports for given condition.
type Ports struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Port `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Ports) Reset() {
	*x = Ports{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pds_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ports) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ports) ProtoMessage() {}

func (x *Ports) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pds_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ports.ProtoReflect.Descriptor instead.
func (*Ports) Descriptor() ([]byte, []int) {
	return file_pb_pds_proto_rawDescGZIP(), []int{8}
}

func (x *Ports) GetList() []*Port {
	if x != nil {
		return x.List
	}
	return nil
}

var File_pb_pds_proto protoreflect.FileDescriptor

var file_pb_pds_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x70, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x70, 0x64, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1f, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xfe, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x4b, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65,
	0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x1b,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x68, 0x6f, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x68, 0x6f, 0x6c, 0x65, 0x22, 0x41, 0x0a, 0x05,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22,
	0x44, 0x0a, 0x06, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x64, 0x73, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70,
	0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xb8, 0x01,
	0x0a, 0x09, 0x54, 0x6f, 0x6f, 0x6c, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22,
	0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x57, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x22, 0x0e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x5a, 0x18,
	0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x65, 0x63, 0x68, 0x6f,
	0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x32, 0xb0, 0x03, 0x0a, 0x09, 0x50, 0x6f, 0x72,
	0x74, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x1a,
	0x0c, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x39, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x09, 0x2e,
	0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x08, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x4b,
	0x65, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x39, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x08, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x4b,
	0x65, 0x79, 0x1a, 0x09, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0x09, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x3f, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x65, 0x61,
	0x72, 0x65, 0x73, 0x74, 0x12, 0x0a, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x1a, 0x09, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6e,
	0x65, 0x61, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x44, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x6e,
	0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x43, 0x69, 0x72,
	0x63, 0x6c, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x3d, 0x0a, 0x08,
	0x46, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x0a, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x26, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x68, 0x77, 0x61, 0x72,
	0x7a, 0x6c, 0x69, 0x63, 0x68, 0x74, 0x62, 0x65, 0x7a, 0x69, 0x72, 0x6b, 0x2f, 0x70, 0x64, 0x73,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_pds_proto_rawDescOnce sync.Once
	file_pb_pds_proto_rawDescData = file_pb_pds_proto_rawDesc
)

func file_pb_pds_proto_rawDescGZIP() []byte {
	file_pb_pds_proto_rawDescOnce.Do(func() {
		file_pb_pds_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_pds_proto_rawDescData)
	})
	return file_pb_pds_proto_rawDescData
}

var file_pb_pds_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_pds_proto_goTypes = []interface{}{
	(*Content)(nil),               // 0: pds.Content
	(*Port)(nil),                  // 1: pds.Port
	(*Summary)(nil),               // 2: pds.Summary
	(*Key)(nil),                   // 3: pds.Key
	(*Name)(nil),                  // 4: pds.Name
	(*Quest)(nil),                 // 5: pds.Quest
	(*Point)(nil),                 // 6: pds.Point
	(*Circle)(nil),                // 7: pds.Circle
	(*Ports)(nil),                 // 8: pds.Ports
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_pb_pds_proto_depIdxs = []int32{
	6,  // 0: pds.Circle.center:type_name -> pds.Point
	1,  // 1: pds.Ports.list:type_name -> pds.Port
	9,  // 2: pds.ToolGuide.Ping:input_type -> google.protobuf.Empty
	0,  // 3: pds.ToolGuide.Echo:input_type -> pds.Content
	1,  // 4: pds.PortGuide.RecordList:input_type -> pds.Port
	1,  // 5: pds.PortGuide.SetByKey:input_type -> pds.Port
	3,  // 6: pds.PortGuide.GetByKey:input_type -> pds.Key
	4,  // 7: pds.PortGuide.GetByName:input_type -> pds.Name
	6,  // 8: pds.PortGuide.FindNearest:input_type -> pds.Point
	7,  // 9: pds.PortGuide.FindInCircle:input_type -> pds.Circle
	5,  // 10: pds.PortGuide.FindText:input_type -> pds.Quest
	10, // 11: pds.ToolGuide.Ping:output_type -> google.protobuf.Timestamp
	0,  // 12: pds.ToolGuide.Echo:output_type -> pds.Content
	2,  // 13: pds.PortGuide.RecordList:output_type -> pds.Summary
	3,  // 14: pds.PortGuide.SetByKey:output_type -> pds.Key
	1,  // 15: pds.PortGuide.GetByKey:output_type -> pds.Port
	1,  // 16: pds.PortGuide.GetByName:output_type -> pds.Port
	1,  // 17: pds.PortGuide.FindNearest:output_type -> pds.Port
	8,  // 18: pds.PortGuide.FindInCircle:output_type -> pds.Ports
	8,  // 19: pds.PortGuide.FindText:output_type -> pds.Ports
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_pb_pds_proto_init() }
func file_pb_pds_proto_init() {
	if File_pb_pds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_pds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_pb_pds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_pb_pds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
		file_pb_pds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_pb_pds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_pb_pds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quest); i {
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
		file_pb_pds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_pb_pds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Circle); i {
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
		file_pb_pds_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ports); i {
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
			RawDescriptor: file_pb_pds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pb_pds_proto_goTypes,
		DependencyIndexes: file_pb_pds_proto_depIdxs,
		MessageInfos:      file_pb_pds_proto_msgTypes,
	}.Build()
	File_pb_pds_proto = out.File
	file_pb_pds_proto_rawDesc = nil
	file_pb_pds_proto_goTypes = nil
	file_pb_pds_proto_depIdxs = nil
}
