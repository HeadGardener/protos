// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: driverserver.proto

package driver_service

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
	return file_driverserver_proto_enumTypes[0].Descriptor()
}

func (AcceptStatus) Type() protoreflect.EnumType {
	return &file_driverserver_proto_enumTypes[0]
}

func (x AcceptStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcceptStatus.Descriptor instead.
func (AcceptStatus) EnumDescriptor() ([]byte, []int) {
	return file_driverserver_proto_rawDescGZIP(), []int{0}
}

type AcceptOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverID string `protobuf:"bytes,1,opt,name=driverID,proto3" json:"driverID,omitempty"`
	UserID   string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	OrderID  string `protobuf:"bytes,3,opt,name=orderID,proto3" json:"orderID,omitempty"`
	From     string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *AcceptOrderRequest) Reset() {
	*x = AcceptOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptOrderRequest) ProtoMessage() {}

func (x *AcceptOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driverserver_proto_msgTypes[0]
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
	return file_driverserver_proto_rawDescGZIP(), []int{0}
}

func (x *AcceptOrderRequest) GetDriverID() string {
	if x != nil {
		return x.DriverID
	}
	return ""
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

func (x *AcceptOrderRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *AcceptOrderRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type AcceptOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverID string       `protobuf:"bytes,1,opt,name=driverID,proto3" json:"driverID,omitempty"`
	OrderID  string       `protobuf:"bytes,2,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Status   AcceptStatus `protobuf:"varint,3,opt,name=status,proto3,enum=driverServer.AcceptStatus" json:"status,omitempty"`
}

func (x *AcceptOrderResponse) Reset() {
	*x = AcceptOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_driverserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptOrderResponse) ProtoMessage() {}

func (x *AcceptOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_driverserver_proto_msgTypes[1]
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
	return file_driverserver_proto_rawDescGZIP(), []int{1}
}

func (x *AcceptOrderResponse) GetDriverID() string {
	if x != nil {
		return x.DriverID
	}
	return ""
}

func (x *AcceptOrderResponse) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *AcceptOrderResponse) GetStatus() AcceptStatus {
	if x != nil {
		return x.Status
	}
	return AcceptStatus_ACCEPTED
}

var File_driverserver_proto protoreflect.FileDescriptor

var file_driverserver_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x13, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x2a, 0x0a, 0x0c,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x01, 0x32, 0x63, 0x0a, 0x0d, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x61, 0x64,
	0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x54, 0x61, 0x78, 0x69, 0x41, 0x70, 0x70,
	0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driverserver_proto_rawDescOnce sync.Once
	file_driverserver_proto_rawDescData = file_driverserver_proto_rawDesc
)

func file_driverserver_proto_rawDescGZIP() []byte {
	file_driverserver_proto_rawDescOnce.Do(func() {
		file_driverserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_driverserver_proto_rawDescData)
	})
	return file_driverserver_proto_rawDescData
}

var file_driverserver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_driverserver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_driverserver_proto_goTypes = []interface{}{
	(AcceptStatus)(0),           // 0: driverServer.AcceptStatus
	(*AcceptOrderRequest)(nil),  // 1: driverServer.AcceptOrderRequest
	(*AcceptOrderResponse)(nil), // 2: driverServer.AcceptOrderResponse
}
var file_driverserver_proto_depIdxs = []int32{
	0, // 0: driverServer.AcceptOrderResponse.status:type_name -> driverServer.AcceptStatus
	1, // 1: driverServer.DriverService.AcceptOrder:input_type -> driverServer.AcceptOrderRequest
	2, // 2: driverServer.DriverService.AcceptOrder:output_type -> driverServer.AcceptOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_driverserver_proto_init() }
func file_driverserver_proto_init() {
	if File_driverserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_driverserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_driverserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_driverserver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driverserver_proto_goTypes,
		DependencyIndexes: file_driverserver_proto_depIdxs,
		EnumInfos:         file_driverserver_proto_enumTypes,
		MessageInfos:      file_driverserver_proto_msgTypes,
	}.Build()
	File_driverserver_proto = out.File
	file_driverserver_proto_rawDesc = nil
	file_driverserver_proto_goTypes = nil
	file_driverserver_proto_depIdxs = nil
}