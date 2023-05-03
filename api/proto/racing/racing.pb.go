// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: racing/racing.proto

package racing

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListRacesRequestSort_Order int32

const (
	ListRacesRequestSort_ASC  ListRacesRequestSort_Order = 0
	ListRacesRequestSort_DESC ListRacesRequestSort_Order = 1
)

// Enum value maps for ListRacesRequestSort_Order.
var (
	ListRacesRequestSort_Order_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	ListRacesRequestSort_Order_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x ListRacesRequestSort_Order) Enum() *ListRacesRequestSort_Order {
	p := new(ListRacesRequestSort_Order)
	*p = x
	return p
}

func (x ListRacesRequestSort_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListRacesRequestSort_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_racing_racing_proto_enumTypes[0].Descriptor()
}

func (ListRacesRequestSort_Order) Type() protoreflect.EnumType {
	return &file_racing_racing_proto_enumTypes[0]
}

func (x ListRacesRequestSort_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListRacesRequestSort_Order.Descriptor instead.
func (ListRacesRequestSort_Order) EnumDescriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{4, 0}
}

type ListRacesRequestSort_Field int32

const (
	ListRacesRequestSort_ID                    ListRacesRequestSort_Field = 0
	ListRacesRequestSort_NAME                  ListRacesRequestSort_Field = 1
	ListRacesRequestSort_NUMBER                ListRacesRequestSort_Field = 2
	ListRacesRequestSort_ADVERTISED_START_TIME ListRacesRequestSort_Field = 3
)

// Enum value maps for ListRacesRequestSort_Field.
var (
	ListRacesRequestSort_Field_name = map[int32]string{
		0: "ID",
		1: "NAME",
		2: "NUMBER",
		3: "ADVERTISED_START_TIME",
	}
	ListRacesRequestSort_Field_value = map[string]int32{
		"ID":                    0,
		"NAME":                  1,
		"NUMBER":                2,
		"ADVERTISED_START_TIME": 3,
	}
)

func (x ListRacesRequestSort_Field) Enum() *ListRacesRequestSort_Field {
	p := new(ListRacesRequestSort_Field)
	*p = x
	return p
}

func (x ListRacesRequestSort_Field) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListRacesRequestSort_Field) Descriptor() protoreflect.EnumDescriptor {
	return file_racing_racing_proto_enumTypes[1].Descriptor()
}

func (ListRacesRequestSort_Field) Type() protoreflect.EnumType {
	return &file_racing_racing_proto_enumTypes[1]
}

func (x ListRacesRequestSort_Field) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListRacesRequestSort_Field.Descriptor instead.
func (ListRacesRequestSort_Field) EnumDescriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{4, 1}
}

type Race_Status int32

const (
	Race_CLOSED Race_Status = 0
	Race_OPEN   Race_Status = 1
)

// Enum value maps for Race_Status.
var (
	Race_Status_name = map[int32]string{
		0: "CLOSED",
		1: "OPEN",
	}
	Race_Status_value = map[string]int32{
		"CLOSED": 0,
		"OPEN":   1,
	}
)

func (x Race_Status) Enum() *Race_Status {
	p := new(Race_Status)
	*p = x
	return p
}

func (x Race_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Race_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_racing_racing_proto_enumTypes[2].Descriptor()
}

func (Race_Status) Type() protoreflect.EnumType {
	return &file_racing_racing_proto_enumTypes[2]
}

func (x Race_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Race_Status.Descriptor instead.
func (Race_Status) EnumDescriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{5, 0}
}

type GetRaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRaceRequest) Reset() {
	*x = GetRaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaceRequest) ProtoMessage() {}

func (x *GetRaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaceRequest.ProtoReflect.Descriptor instead.
func (*GetRaceRequest) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{0}
}

func (x *GetRaceRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Request for ListRaces call.
type ListRacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *ListRacesRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Sort   *ListRacesRequestSort   `protobuf:"bytes,2,opt,name=sort,proto3,oneof" json:"sort,omitempty"`
}

func (x *ListRacesRequest) Reset() {
	*x = ListRacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequest) ProtoMessage() {}

func (x *ListRacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequest.ProtoReflect.Descriptor instead.
func (*ListRacesRequest) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{1}
}

func (x *ListRacesRequest) GetFilter() *ListRacesRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListRacesRequest) GetSort() *ListRacesRequestSort {
	if x != nil {
		return x.Sort
	}
	return nil
}

// Response to ListRaces call.
type ListRacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Races []*Race `protobuf:"bytes,1,rep,name=races,proto3" json:"races,omitempty"`
}

func (x *ListRacesResponse) Reset() {
	*x = ListRacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesResponse) ProtoMessage() {}

func (x *ListRacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesResponse.ProtoReflect.Descriptor instead.
func (*ListRacesResponse) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{2}
}

func (x *ListRacesResponse) GetRaces() []*Race {
	if x != nil {
		return x.Races
	}
	return nil
}

// Filter for listing races.
type ListRacesRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetingIds []int64 `protobuf:"varint,1,rep,packed,name=meeting_ids,json=meetingIds,proto3" json:"meeting_ids,omitempty"`
	Visible    *bool   `protobuf:"varint,2,opt,name=visible,proto3,oneof" json:"visible,omitempty"`
}

func (x *ListRacesRequestFilter) Reset() {
	*x = ListRacesRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequestFilter) ProtoMessage() {}

func (x *ListRacesRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequestFilter.ProtoReflect.Descriptor instead.
func (*ListRacesRequestFilter) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{3}
}

func (x *ListRacesRequestFilter) GetMeetingIds() []int64 {
	if x != nil {
		return x.MeetingIds
	}
	return nil
}

func (x *ListRacesRequestFilter) GetVisible() bool {
	if x != nil && x.Visible != nil {
		return *x.Visible
	}
	return false
}

type ListRacesRequestSort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *ListRacesRequestSort_Order `protobuf:"varint,1,opt,name=order,proto3,enum=racing.ListRacesRequestSort_Order,oneof" json:"order,omitempty"`
	Field *ListRacesRequestSort_Field `protobuf:"varint,2,opt,name=field,proto3,enum=racing.ListRacesRequestSort_Field,oneof" json:"field,omitempty"`
}

func (x *ListRacesRequestSort) Reset() {
	*x = ListRacesRequestSort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequestSort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequestSort) ProtoMessage() {}

func (x *ListRacesRequestSort) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequestSort.ProtoReflect.Descriptor instead.
func (*ListRacesRequestSort) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{4}
}

func (x *ListRacesRequestSort) GetOrder() ListRacesRequestSort_Order {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return ListRacesRequestSort_ASC
}

func (x *ListRacesRequestSort) GetField() ListRacesRequestSort_Field {
	if x != nil && x.Field != nil {
		return *x.Field
	}
	return ListRacesRequestSort_ID
}

// A race resource.
type Race struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the race.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// MeetingID represents a unique identifier for the races meeting.
	MeetingId int64 `protobuf:"varint,2,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// Name is the official name given to the race.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Number represents the number of the race.
	Number int64 `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	// Visible represents whether or not the race is visible.
	Visible bool `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	// AdvertisedStartTime is the time the race is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
	// Status reflects whether the Race is open or closed, based off the
	// advertised start time being in the past.
	Status Race_Status `protobuf:"varint,7,opt,name=status,proto3,enum=racing.Race_Status" json:"status,omitempty"`
}

func (x *Race) Reset() {
	*x = Race{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Race) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Race) ProtoMessage() {}

func (x *Race) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Race.ProtoReflect.Descriptor instead.
func (*Race) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{5}
}

func (x *Race) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Race) GetMeetingId() int64 {
	if x != nil {
		return x.MeetingId
	}
	return 0
}

func (x *Race) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Race) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Race) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Race) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

func (x *Race) GetStatus() Race_Status {
	if x != nil {
		return x.Status
	}
	return Race_CLOSED
}

var File_racing_racing_proto protoreflect.FileDescriptor

var file_racing_racing_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8a,
	0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x05, 0x72,
	0x61, 0x63, 0x65, 0x73, 0x22, 0x64, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x12,
	0x1d, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x6f, 0x72, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x72, 0x74,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x3d, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x22, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x01, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01,
	0x01, 0x22, 0x1a, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53,
	0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x22, 0x40, 0x0a,
	0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53,
	0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x22, 0x98, 0x02, 0x0a, 0x04, 0x52, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1e,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x32, 0xad,
	0x01, 0x0a, 0x06, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x5b, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x2d, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x63,
	0x65, 0x12, 0x16, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x09,
	0x5a, 0x07, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_racing_racing_proto_rawDescOnce sync.Once
	file_racing_racing_proto_rawDescData = file_racing_racing_proto_rawDesc
)

func file_racing_racing_proto_rawDescGZIP() []byte {
	file_racing_racing_proto_rawDescOnce.Do(func() {
		file_racing_racing_proto_rawDescData = protoimpl.X.CompressGZIP(file_racing_racing_proto_rawDescData)
	})
	return file_racing_racing_proto_rawDescData
}

var file_racing_racing_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_racing_racing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_racing_racing_proto_goTypes = []interface{}{
	(ListRacesRequestSort_Order)(0), // 0: racing.ListRacesRequestSort.Order
	(ListRacesRequestSort_Field)(0), // 1: racing.ListRacesRequestSort.Field
	(Race_Status)(0),                // 2: racing.Race.Status
	(*GetRaceRequest)(nil),          // 3: racing.GetRaceRequest
	(*ListRacesRequest)(nil),        // 4: racing.ListRacesRequest
	(*ListRacesResponse)(nil),       // 5: racing.ListRacesResponse
	(*ListRacesRequestFilter)(nil),  // 6: racing.ListRacesRequestFilter
	(*ListRacesRequestSort)(nil),    // 7: racing.ListRacesRequestSort
	(*Race)(nil),                    // 8: racing.Race
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_racing_racing_proto_depIdxs = []int32{
	6, // 0: racing.ListRacesRequest.filter:type_name -> racing.ListRacesRequestFilter
	7, // 1: racing.ListRacesRequest.sort:type_name -> racing.ListRacesRequestSort
	8, // 2: racing.ListRacesResponse.races:type_name -> racing.Race
	0, // 3: racing.ListRacesRequestSort.order:type_name -> racing.ListRacesRequestSort.Order
	1, // 4: racing.ListRacesRequestSort.field:type_name -> racing.ListRacesRequestSort.Field
	9, // 5: racing.Race.advertised_start_time:type_name -> google.protobuf.Timestamp
	2, // 6: racing.Race.status:type_name -> racing.Race.Status
	4, // 7: racing.Racing.ListRaces:input_type -> racing.ListRacesRequest
	3, // 8: racing.Racing.GetRace:input_type -> racing.GetRaceRequest
	5, // 9: racing.Racing.ListRaces:output_type -> racing.ListRacesResponse
	8, // 10: racing.Racing.GetRace:output_type -> racing.Race
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_racing_racing_proto_init() }
func file_racing_racing_proto_init() {
	if File_racing_racing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_racing_racing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaceRequest); i {
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
		file_racing_racing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequest); i {
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
		file_racing_racing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesResponse); i {
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
		file_racing_racing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequestFilter); i {
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
		file_racing_racing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequestSort); i {
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
		file_racing_racing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Race); i {
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
	file_racing_racing_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_racing_racing_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_racing_racing_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_racing_racing_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_racing_racing_proto_goTypes,
		DependencyIndexes: file_racing_racing_proto_depIdxs,
		EnumInfos:         file_racing_racing_proto_enumTypes,
		MessageInfos:      file_racing_racing_proto_msgTypes,
	}.Build()
	File_racing_racing_proto = out.File
	file_racing_racing_proto_rawDesc = nil
	file_racing_racing_proto_goTypes = nil
	file_racing_racing_proto_depIdxs = nil
}
