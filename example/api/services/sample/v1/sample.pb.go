// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: listify/services/sample/v1/sample.proto

package sspb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/pentops/protoc-gen-listify/listify/v1"
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

type WidgetStatus int32

const (
	WidgetStatus_WIDGET_STATUS_UNSPECIFIED WidgetStatus = 0
	WidgetStatus_WIDGET_STATUS_REQUESTED   WidgetStatus = 1
	WidgetStatus_WIDGET_STATUS_CREATED     WidgetStatus = 2
)

// Enum value maps for WidgetStatus.
var (
	WidgetStatus_name = map[int32]string{
		0: "WIDGET_STATUS_UNSPECIFIED",
		1: "WIDGET_STATUS_REQUESTED",
		2: "WIDGET_STATUS_CREATED",
	}
	WidgetStatus_value = map[string]int32{
		"WIDGET_STATUS_UNSPECIFIED": 0,
		"WIDGET_STATUS_REQUESTED":   1,
		"WIDGET_STATUS_CREATED":     2,
	}
)

func (x WidgetStatus) Enum() *WidgetStatus {
	p := new(WidgetStatus)
	*p = x
	return p
}

func (x WidgetStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WidgetStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_listify_services_sample_v1_sample_proto_enumTypes[0].Descriptor()
}

func (WidgetStatus) Type() protoreflect.EnumType {
	return &file_listify_services_sample_v1_sample_proto_enumTypes[0]
}

func (x WidgetStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WidgetStatus.Descriptor instead.
func (WidgetStatus) EnumDescriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{0}
}

type GetWidgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWidgetRequest) Reset() {
	*x = GetWidgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWidgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWidgetRequest) ProtoMessage() {}

func (x *GetWidgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWidgetRequest.ProtoReflect.Descriptor instead.
func (*GetWidgetRequest) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{0}
}

func (x *GetWidgetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWidgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Widget *Widget `protobuf:"bytes,1,opt,name=widget,proto3" json:"widget,omitempty"`
}

func (x *GetWidgetResponse) Reset() {
	*x = GetWidgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWidgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWidgetResponse) ProtoMessage() {}

func (x *GetWidgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWidgetResponse.ProtoReflect.Descriptor instead.
func (*GetWidgetResponse) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{1}
}

func (x *GetWidgetResponse) GetWidget() *Widget {
	if x != nil {
		return x.Widget
	}
	return nil
}

type Widget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string                 `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Details    *WidgetDetails         `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Status     WidgetStatus           `protobuf:"varint,4,opt,name=status,proto3,enum=listify.services.sample.v1.WidgetStatus" json:"status,omitempty"`
	Created    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Widget) Reset() {
	*x = Widget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget) ProtoMessage() {}

func (x *Widget) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget.ProtoReflect.Descriptor instead.
func (*Widget) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{2}
}

func (x *Widget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Widget) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Widget) GetDetails() *WidgetDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Widget) GetStatus() WidgetStatus {
	if x != nil {
		return x.Status
	}
	return WidgetStatus_WIDGET_STATUS_UNSPECIFIED
}

func (x *Widget) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type WidgetDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight  int64  `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Special bool   `protobuf:"varint,3,opt,name=special,proto3" json:"special,omitempty"`
	// Types that are assignable to Type:
	//
	//	*WidgetDetails_Square
	//	*WidgetDetails_Round
	Type isWidgetDetails_Type `protobuf_oneof:"type"`
}

func (x *WidgetDetails) Reset() {
	*x = WidgetDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetDetails) ProtoMessage() {}

func (x *WidgetDetails) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetDetails.ProtoReflect.Descriptor instead.
func (*WidgetDetails) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{3}
}

func (x *WidgetDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WidgetDetails) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *WidgetDetails) GetSpecial() bool {
	if x != nil {
		return x.Special
	}
	return false
}

func (m *WidgetDetails) GetType() isWidgetDetails_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *WidgetDetails) GetSquare() *SquareWidget {
	if x, ok := x.GetType().(*WidgetDetails_Square); ok {
		return x.Square
	}
	return nil
}

func (x *WidgetDetails) GetRound() *RoundWidget {
	if x, ok := x.GetType().(*WidgetDetails_Round); ok {
		return x.Round
	}
	return nil
}

type isWidgetDetails_Type interface {
	isWidgetDetails_Type()
}

type WidgetDetails_Square struct {
	Square *SquareWidget `protobuf:"bytes,4,opt,name=square,proto3,oneof"`
}

type WidgetDetails_Round struct {
	Round *RoundWidget `protobuf:"bytes,5,opt,name=round,proto3,oneof"`
}

func (*WidgetDetails_Square) isWidgetDetails_Type() {}

func (*WidgetDetails_Round) isWidgetDetails_Type() {}

type SquareWidget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int64 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *SquareWidget) Reset() {
	*x = SquareWidget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareWidget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareWidget) ProtoMessage() {}

func (x *SquareWidget) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareWidget.ProtoReflect.Descriptor instead.
func (*SquareWidget) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{4}
}

func (x *SquareWidget) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *SquareWidget) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type RoundWidget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diameter int64 `protobuf:"varint,1,opt,name=diameter,proto3" json:"diameter,omitempty"`
}

func (x *RoundWidget) Reset() {
	*x = RoundWidget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoundWidget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoundWidget) ProtoMessage() {}

func (x *RoundWidget) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoundWidget.ProtoReflect.Descriptor instead.
func (*RoundWidget) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{5}
}

func (x *RoundWidget) GetDiameter() int64 {
	if x != nil {
		return x.Diameter
	}
	return 0
}

type ListWidgetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    *v1.PageRequest `protobuf:"bytes,40,opt,name=page,proto3" json:"page,omitempty"`
	Search  *v1.Search      `protobuf:"bytes,41,opt,name=search,proto3" json:"search,omitempty"`
	Sorts   []*v1.Sort      `protobuf:"bytes,42,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Filters []*v1.Filter    `protobuf:"bytes,43,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ListWidgetsRequest) Reset() {
	*x = ListWidgetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWidgetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWidgetsRequest) ProtoMessage() {}

func (x *ListWidgetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWidgetsRequest.ProtoReflect.Descriptor instead.
func (*ListWidgetsRequest) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{6}
}

func (x *ListWidgetsRequest) GetPage() *v1.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListWidgetsRequest) GetSearch() *v1.Search {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *ListWidgetsRequest) GetSorts() []*v1.Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *ListWidgetsRequest) GetFilters() []*v1.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ListWidgetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Widgets []*Widget        `protobuf:"bytes,1,rep,name=widgets,proto3" json:"widgets,omitempty"`
	Page    *v1.PageResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListWidgetsResponse) Reset() {
	*x = ListWidgetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listify_services_sample_v1_sample_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWidgetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWidgetsResponse) ProtoMessage() {}

func (x *ListWidgetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listify_services_sample_v1_sample_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWidgetsResponse.ProtoReflect.Descriptor instead.
func (*ListWidgetsResponse) Descriptor() ([]byte, []int) {
	return file_listify_services_sample_v1_sample_proto_rawDescGZIP(), []int{7}
}

func (x *ListWidgetsResponse) GetWidgets() []*Widget {
	if x != nil {
		return x.Widgets
	}
	return nil
}

func (x *ListWidgetsResponse) GetPage() *v1.PageResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_listify_services_sample_v1_sample_proto protoreflect.FileDescriptor

var file_listify_services_sample_v1_sample_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6c, 0x69, 0x73, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x59, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x77, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22, 0xd0, 0x02,
	0x0a, 0x06, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x8a, 0x9b, 0x81, 0xca, 0x02, 0x08, 0x72, 0x06, 0x1a, 0x04, 0x12, 0x02, 0x08, 0x01, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x57, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x15, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00,
	0x8a, 0x9b, 0x81, 0xca, 0x02, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x15, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x8a, 0x9b, 0x81, 0xca, 0x02, 0x07,
	0xf2, 0x01, 0x04, 0x08, 0x01, 0x10, 0x01, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x22, 0xa5, 0x02, 0x0a, 0x0d, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x13, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x8a, 0x9b, 0x81, 0xca, 0x02, 0x06, 0x72,
	0x04, 0x0a, 0x02, 0x08, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x13, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x8a, 0x9b, 0x81, 0xca, 0x02, 0x06, 0x32, 0x04, 0x08, 0x01, 0x10, 0x01,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0a, 0x8a, 0x9b, 0x81, 0xca, 0x02,
	0x04, 0x6a, 0x02, 0x08, 0x01, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x42,
	0x0a, 0x06, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x73, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x42, 0x13, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0b, 0xf8, 0x42, 0x01,
	0x92, 0x9b, 0x81, 0xca, 0x02, 0x02, 0x08, 0x01, 0x22, 0x66, 0x0a, 0x0c, 0x53, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x13, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x8a, 0x9b, 0x81, 0xca, 0x02, 0x06, 0x32, 0x04, 0x08, 0x01, 0x10, 0x01, 0x52, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x13, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x8a, 0x9b, 0x81, 0xca,
	0x02, 0x06, 0x32, 0x04, 0x08, 0x01, 0x10, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x3e, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x2f, 0x0a, 0x08, 0x64, 0x69, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x13, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x8a, 0x9b, 0x81, 0xca, 0x02, 0x06,
	0x32, 0x04, 0x08, 0x01, 0x10, 0x01, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x22, 0xc3, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x29,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x26, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x2a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x2b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x2a, 0x65, 0x0a, 0x0c, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x49,
	0x44, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x57, 0x49, 0x44,
	0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x49, 0x44, 0x47, 0x45, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x32, 0xb2, 0x02, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x2c, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x95,
	0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12, 0x2e,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x82, 0x9b,
	0x81, 0xca, 0x02, 0x02, 0x08, 0x01, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x66, 0x79, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_listify_services_sample_v1_sample_proto_rawDescOnce sync.Once
	file_listify_services_sample_v1_sample_proto_rawDescData = file_listify_services_sample_v1_sample_proto_rawDesc
)

func file_listify_services_sample_v1_sample_proto_rawDescGZIP() []byte {
	file_listify_services_sample_v1_sample_proto_rawDescOnce.Do(func() {
		file_listify_services_sample_v1_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_listify_services_sample_v1_sample_proto_rawDescData)
	})
	return file_listify_services_sample_v1_sample_proto_rawDescData
}

var file_listify_services_sample_v1_sample_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_listify_services_sample_v1_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_listify_services_sample_v1_sample_proto_goTypes = []interface{}{
	(WidgetStatus)(0),             // 0: listify.services.sample.v1.WidgetStatus
	(*GetWidgetRequest)(nil),      // 1: listify.services.sample.v1.GetWidgetRequest
	(*GetWidgetResponse)(nil),     // 2: listify.services.sample.v1.GetWidgetResponse
	(*Widget)(nil),                // 3: listify.services.sample.v1.Widget
	(*WidgetDetails)(nil),         // 4: listify.services.sample.v1.WidgetDetails
	(*SquareWidget)(nil),          // 5: listify.services.sample.v1.SquareWidget
	(*RoundWidget)(nil),           // 6: listify.services.sample.v1.RoundWidget
	(*ListWidgetsRequest)(nil),    // 7: listify.services.sample.v1.ListWidgetsRequest
	(*ListWidgetsResponse)(nil),   // 8: listify.services.sample.v1.ListWidgetsResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*v1.PageRequest)(nil),        // 10: listify.v1.PageRequest
	(*v1.Search)(nil),             // 11: listify.v1.Search
	(*v1.Sort)(nil),               // 12: listify.v1.Sort
	(*v1.Filter)(nil),             // 13: listify.v1.Filter
	(*v1.PageResponse)(nil),       // 14: listify.v1.PageResponse
}
var file_listify_services_sample_v1_sample_proto_depIdxs = []int32{
	3,  // 0: listify.services.sample.v1.GetWidgetResponse.widget:type_name -> listify.services.sample.v1.Widget
	4,  // 1: listify.services.sample.v1.Widget.details:type_name -> listify.services.sample.v1.WidgetDetails
	0,  // 2: listify.services.sample.v1.Widget.status:type_name -> listify.services.sample.v1.WidgetStatus
	9,  // 3: listify.services.sample.v1.Widget.created:type_name -> google.protobuf.Timestamp
	5,  // 4: listify.services.sample.v1.WidgetDetails.square:type_name -> listify.services.sample.v1.SquareWidget
	6,  // 5: listify.services.sample.v1.WidgetDetails.round:type_name -> listify.services.sample.v1.RoundWidget
	10, // 6: listify.services.sample.v1.ListWidgetsRequest.page:type_name -> listify.v1.PageRequest
	11, // 7: listify.services.sample.v1.ListWidgetsRequest.search:type_name -> listify.v1.Search
	12, // 8: listify.services.sample.v1.ListWidgetsRequest.sorts:type_name -> listify.v1.Sort
	13, // 9: listify.services.sample.v1.ListWidgetsRequest.filters:type_name -> listify.v1.Filter
	3,  // 10: listify.services.sample.v1.ListWidgetsResponse.widgets:type_name -> listify.services.sample.v1.Widget
	14, // 11: listify.services.sample.v1.ListWidgetsResponse.page:type_name -> listify.v1.PageResponse
	1,  // 12: listify.services.sample.v1.SampleService.GetWidget:input_type -> listify.services.sample.v1.GetWidgetRequest
	7,  // 13: listify.services.sample.v1.SampleService.ListWidgets:input_type -> listify.services.sample.v1.ListWidgetsRequest
	2,  // 14: listify.services.sample.v1.SampleService.GetWidget:output_type -> listify.services.sample.v1.GetWidgetResponse
	8,  // 15: listify.services.sample.v1.SampleService.ListWidgets:output_type -> listify.services.sample.v1.ListWidgetsResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_listify_services_sample_v1_sample_proto_init() }
func file_listify_services_sample_v1_sample_proto_init() {
	if File_listify_services_sample_v1_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_listify_services_sample_v1_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWidgetRequest); i {
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
		file_listify_services_sample_v1_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWidgetResponse); i {
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
		file_listify_services_sample_v1_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Widget); i {
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
		file_listify_services_sample_v1_sample_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetDetails); i {
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
		file_listify_services_sample_v1_sample_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareWidget); i {
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
		file_listify_services_sample_v1_sample_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoundWidget); i {
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
		file_listify_services_sample_v1_sample_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWidgetsRequest); i {
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
		file_listify_services_sample_v1_sample_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWidgetsResponse); i {
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
	file_listify_services_sample_v1_sample_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*WidgetDetails_Square)(nil),
		(*WidgetDetails_Round)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_listify_services_sample_v1_sample_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listify_services_sample_v1_sample_proto_goTypes,
		DependencyIndexes: file_listify_services_sample_v1_sample_proto_depIdxs,
		EnumInfos:         file_listify_services_sample_v1_sample_proto_enumTypes,
		MessageInfos:      file_listify_services_sample_v1_sample_proto_msgTypes,
	}.Build()
	File_listify_services_sample_v1_sample_proto = out.File
	file_listify_services_sample_v1_sample_proto_rawDesc = nil
	file_listify_services_sample_v1_sample_proto_goTypes = nil
	file_listify_services_sample_v1_sample_proto_depIdxs = nil
}
