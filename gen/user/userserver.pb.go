// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: userserver.proto

package user_service

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

type AcceptStatus int32

const (
	AcceptStatus_ACCEPTED AcceptStatus = 0
	AcceptStatus_REJECTED AcceptStatus = 1
)

// Enum value maps for AcceptStatus.
var (
	AcceptStatus_name = map[int32]string{
		0: "ACCEPTED",
		1: "REJECTED",
	}
	AcceptStatus_value = map[string]int32{
		"ACCEPTED": 0,
		"REJECTED": 1,
	}
)

func (x AcceptStatus) Enum() *AcceptStatus {
	p := new(AcceptStatus)
	*p = x
	return p
}

func (x AcceptStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcceptStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_userserver_proto_enumTypes[0].Descriptor()
}

func (AcceptStatus) Type() protoreflect.EnumType {
	return &file_userserver_proto_enumTypes[0]
}

func (x AcceptStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcceptStatus.Descriptor instead.
func (AcceptStatus) EnumDescriptor() ([]byte, []int) {
	return file_userserver_proto_rawDescGZIP(), []int{0}
}

type CompleteStatus int32

const (
	CompleteStatus_SUCCESS  CompleteStatus = 0
	CompleteStatus_CANCELED CompleteStatus = 1
)

// Enum value maps for CompleteStatus.
var (
	CompleteStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "CANCELED",
	}
	CompleteStatus_value = map[string]int32{
		"SUCCESS":  0,
		"CANCELED": 1,
	}
)

func (x CompleteStatus) Enum() *CompleteStatus {
	p := new(CompleteStatus)
	*p = x
	return p
}

func (x CompleteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompleteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_userserver_proto_enumTypes[1].Descriptor()
}

func (CompleteStatus) Type() protoreflect.EnumType {
	return &file_userserver_proto_enumTypes[1]
}

func (x CompleteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompleteStatus.Descriptor instead.
func (CompleteStatus) EnumDescriptor() ([]byte, []int) {
	return file_userserver_proto_rawDescGZIP(), []int{1}
}

type AcceptOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  string       `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	OrderID string       `protobuf:"bytes,2,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Status  AcceptStatus `protobuf:"varint,3,opt,name=status,proto3,enum=userServer.AcceptStatus" json:"status,omitempty"`
}

func (x *AcceptOrderRequest) Reset() {
	*x = AcceptOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptOrderRequest) ProtoMessage() {}

func (x *AcceptOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptOrderRequest.ProtoReflect.Descriptor instead.
func (*AcceptOrderRequest) Descriptor() ([]byte, []int) {
	return file_userserver_proto_rawDescGZIP(), []int{0}
}

func (x *AcceptOrderRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AcceptOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *AcceptOrderRequest) GetStatus() AcceptStatus {
	if x != nil {
		return x.Status
	}
	return AcceptStatus_ACCEPTED
}

type CompleteOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  string         `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	OrderID string         `protobuf:"bytes,2,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Status  CompleteStatus `protobuf:"varint,3,opt,name=status,proto3,enum=userServer.CompleteStatus" json:"status,omitempty"`
}

func (x *CompleteOrderRequest) Reset() {
	*x = CompleteOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOrderRequest) ProtoMessage() {}

func (x *CompleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOrderRequest.ProtoReflect.Descriptor instead.
func (*CompleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_userserver_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteOrderRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CompleteOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *CompleteOrderRequest) GetStatus() CompleteStatus {
	if x != nil {
		return x.Status
	}
	return CompleteStatus_SUCCESS
}

type AcceptOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Confirmed bool   `protobuf:"varint,2,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
}

func (x *AcceptOrderResponse) Reset() {
	*x = AcceptOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptOrderResponse) ProtoMessage() {}

func (x *AcceptOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptOrderResponse.ProtoReflect.Descriptor instead.
func (*AcceptOrderResponse) Descriptor() ([]byte, []int) {
	return file_userserver_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptOrderResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AcceptOrderResponse) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

type CompleteOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Confirmed bool   `protobuf:"varint,2,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
}

func (x *CompleteOrderResponse) Reset() {
	*x = CompleteOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOrderResponse) ProtoMessage() {}

func (x *CompleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOrderResponse.ProtoReflect.Descriptor instead.
func (*CompleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_userserver_proto_rawDescGZIP(), []int{3}
}

func (x *CompleteOrderResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CompleteOrderResponse) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

var File_userserver_proto protoreflect.FileDescriptor

var file_userserver_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x78,
	0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x7c, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x65, 0x64, 0x2a, 0x2a, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x2b,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0xb3, 0x01, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x48, 0x65, 0x61, 0x64, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x54, 0x61, 0x78,
	0x69, 0x41, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userserver_proto_rawDescOnce sync.Once
	file_userserver_proto_rawDescData = file_userserver_proto_rawDesc
)

func file_userserver_proto_rawDescGZIP() []byte {
	file_userserver_proto_rawDescOnce.Do(func() {
		file_userserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_userserver_proto_rawDescData)
	})
	return file_userserver_proto_rawDescData
}

var file_userserver_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_userserver_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_userserver_proto_goTypes = []interface{}{
	(AcceptStatus)(0),             // 0: userServer.AcceptStatus
	(CompleteStatus)(0),           // 1: userServer.CompleteStatus
	(*AcceptOrderRequest)(nil),    // 2: userServer.AcceptOrderRequest
	(*CompleteOrderRequest)(nil),  // 3: userServer.CompleteOrderRequest
	(*AcceptOrderResponse)(nil),   // 4: userServer.AcceptOrderResponse
	(*CompleteOrderResponse)(nil), // 5: userServer.CompleteOrderResponse
}
var file_userserver_proto_depIdxs = []int32{
	0, // 0: userServer.AcceptOrderRequest.status:type_name -> userServer.AcceptStatus
	1, // 1: userServer.CompleteOrderRequest.status:type_name -> userServer.CompleteStatus
	2, // 2: userServer.UserService.AcceptOrder:input_type -> userServer.AcceptOrderRequest
	3, // 3: userServer.UserService.CompleteOrder:input_type -> userServer.CompleteOrderRequest
	4, // 4: userServer.UserService.AcceptOrder:output_type -> userServer.AcceptOrderResponse
	5, // 5: userServer.UserService.CompleteOrder:output_type -> userServer.CompleteOrderResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userserver_proto_init() }
func file_userserver_proto_init() {
	if File_userserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptOrderRequest); i {
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
		file_userserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOrderRequest); i {
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
		file_userserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptOrderResponse); i {
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
		file_userserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOrderResponse); i {
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
			RawDescriptor: file_userserver_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userserver_proto_goTypes,
		DependencyIndexes: file_userserver_proto_depIdxs,
		EnumInfos:         file_userserver_proto_enumTypes,
		MessageInfos:      file_userserver_proto_msgTypes,
	}.Build()
	File_userserver_proto = out.File
	file_userserver_proto_rawDesc = nil
	file_userserver_proto_goTypes = nil
	file_userserver_proto_depIdxs = nil
}