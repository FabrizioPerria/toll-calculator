// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: shared/types/pb/types.proto

package pb

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

type None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *None) Reset() {
	*x = None{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_types_pb_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_shared_types_pb_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_shared_types_pb_types_proto_rawDescGZIP(), []int{0}
}

type AggregateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuId     int32   `protobuf:"varint,1,opt,name=ObuId,proto3" json:"ObuId,omitempty"`
	Value     float64 `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Timestamp int64   `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *AggregateRequest) Reset() {
	*x = AggregateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_types_pb_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateRequest) ProtoMessage() {}

func (x *AggregateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_types_pb_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateRequest.ProtoReflect.Descriptor instead.
func (*AggregateRequest) Descriptor() ([]byte, []int) {
	return file_shared_types_pb_types_proto_rawDescGZIP(), []int{1}
}

func (x *AggregateRequest) GetObuId() int32 {
	if x != nil {
		return x.ObuId
	}
	return 0
}

func (x *AggregateRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AggregateRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type InvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuId int32 `protobuf:"varint,1,opt,name=ObuId,proto3" json:"ObuId,omitempty"`
}

func (x *InvoiceRequest) Reset() {
	*x = InvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_types_pb_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceRequest) ProtoMessage() {}

func (x *InvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_types_pb_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceRequest.ProtoReflect.Descriptor instead.
func (*InvoiceRequest) Descriptor() ([]byte, []int) {
	return file_shared_types_pb_types_proto_rawDescGZIP(), []int{2}
}

func (x *InvoiceRequest) GetObuId() int32 {
	if x != nil {
		return x.ObuId
	}
	return 0
}

type InvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObuId     int32   `protobuf:"varint,1,opt,name=ObuId,proto3" json:"ObuId,omitempty"`
	Amount    float64 `protobuf:"fixed64,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Distance  float64 `protobuf:"fixed64,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
	Timestamp int64   `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *InvoiceResponse) Reset() {
	*x = InvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_types_pb_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceResponse) ProtoMessage() {}

func (x *InvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shared_types_pb_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceResponse.ProtoReflect.Descriptor instead.
func (*InvoiceResponse) Descriptor() ([]byte, []int) {
	return file_shared_types_pb_types_proto_rawDescGZIP(), []int{3}
}

func (x *InvoiceResponse) GetObuId() int32 {
	if x != nil {
		return x.ObuId
	}
	return 0
}

func (x *InvoiceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *InvoiceResponse) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *InvoiceResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_shared_types_pb_types_proto protoreflect.FileDescriptor

var file_shared_types_pb_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a,
	0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x22, 0x5c, 0x0a, 0x10, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x62, 0x75,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x62, 0x75, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x62, 0x75, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f, 0x62, 0x75, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x0f, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x4f, 0x62, 0x75, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4f,
	0x62, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x68, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x4e, 0x6f,
	0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x61, 0x62, 0x72, 0x69, 0x7a, 0x69, 0x6f, 0x70, 0x65, 0x72, 0x72, 0x69, 0x61, 0x2f, 0x74, 0x6f,
	0x6c, 0x6c, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_types_pb_types_proto_rawDescOnce sync.Once
	file_shared_types_pb_types_proto_rawDescData = file_shared_types_pb_types_proto_rawDesc
)

func file_shared_types_pb_types_proto_rawDescGZIP() []byte {
	file_shared_types_pb_types_proto_rawDescOnce.Do(func() {
		file_shared_types_pb_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_types_pb_types_proto_rawDescData)
	})
	return file_shared_types_pb_types_proto_rawDescData
}

var file_shared_types_pb_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shared_types_pb_types_proto_goTypes = []interface{}{
	(*None)(nil),             // 0: None
	(*AggregateRequest)(nil), // 1: AggregateRequest
	(*InvoiceRequest)(nil),   // 2: InvoiceRequest
	(*InvoiceResponse)(nil),  // 3: InvoiceResponse
}
var file_shared_types_pb_types_proto_depIdxs = []int32{
	1, // 0: AggregatorService.Aggregate:input_type -> AggregateRequest
	2, // 1: AggregatorService.Invoice:input_type -> InvoiceRequest
	0, // 2: AggregatorService.Aggregate:output_type -> None
	3, // 3: AggregatorService.Invoice:output_type -> InvoiceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shared_types_pb_types_proto_init() }
func file_shared_types_pb_types_proto_init() {
	if File_shared_types_pb_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_types_pb_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*None); i {
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
		file_shared_types_pb_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateRequest); i {
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
		file_shared_types_pb_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceRequest); i {
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
		file_shared_types_pb_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceResponse); i {
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
			RawDescriptor: file_shared_types_pb_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shared_types_pb_types_proto_goTypes,
		DependencyIndexes: file_shared_types_pb_types_proto_depIdxs,
		MessageInfos:      file_shared_types_pb_types_proto_msgTypes,
	}.Build()
	File_shared_types_pb_types_proto = out.File
	file_shared_types_pb_types_proto_rawDesc = nil
	file_shared_types_pb_types_proto_goTypes = nil
	file_shared_types_pb_types_proto_depIdxs = nil
}
