// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/v1/testapi/testapi.proto

package testapiv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TestAPIServiceEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestAPIServiceEchoRequest) Reset() {
	*x = TestAPIServiceEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_testapi_testapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAPIServiceEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAPIServiceEchoRequest) ProtoMessage() {}

func (x *TestAPIServiceEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_testapi_testapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAPIServiceEchoRequest.ProtoReflect.Descriptor instead.
func (*TestAPIServiceEchoRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_testapi_testapi_proto_rawDescGZIP(), []int{0}
}

func (x *TestAPIServiceEchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TestAPIServiceEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestAPIServiceEchoResponse) Reset() {
	*x = TestAPIServiceEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_testapi_testapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAPIServiceEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAPIServiceEchoResponse) ProtoMessage() {}

func (x *TestAPIServiceEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_testapi_testapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAPIServiceEchoResponse.ProtoReflect.Descriptor instead.
func (*TestAPIServiceEchoResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_testapi_testapi_proto_rawDescGZIP(), []int{1}
}

func (x *TestAPIServiceEchoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TestAPIServiceRaiseErrorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestAPIServiceRaiseErrorRequest) Reset() {
	*x = TestAPIServiceRaiseErrorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_testapi_testapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAPIServiceRaiseErrorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAPIServiceRaiseErrorRequest) ProtoMessage() {}

func (x *TestAPIServiceRaiseErrorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_testapi_testapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAPIServiceRaiseErrorRequest.ProtoReflect.Descriptor instead.
func (*TestAPIServiceRaiseErrorRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_testapi_testapi_proto_rawDescGZIP(), []int{2}
}

func (x *TestAPIServiceRaiseErrorRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TestAPIServiceRaiseErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TestAPIServiceRaiseErrorResponse) Reset() {
	*x = TestAPIServiceRaiseErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_testapi_testapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAPIServiceRaiseErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAPIServiceRaiseErrorResponse) ProtoMessage() {}

func (x *TestAPIServiceRaiseErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_testapi_testapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAPIServiceRaiseErrorResponse.ProtoReflect.Descriptor instead.
func (*TestAPIServiceRaiseErrorResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_testapi_testapi_proto_rawDescGZIP(), []int{3}
}

func (x *TestAPIServiceRaiseErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_v1_testapi_testapi_proto protoreflect.FileDescriptor

var file_proto_v1_testapi_testapi_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x19, 0x54, 0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x36, 0x0a, 0x1a, 0x54, 0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x1f, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x61, 0x69,
	0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x20, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x61, 0x69, 0x73, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x97, 0x02, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x04, 0x45, 0x63, 0x68,
	0x6f, 0x12, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01,
	0x2a, 0x12, 0x8c, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x69, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x2b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x61, 0x69, 0x73,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x61, 0x69, 0x73, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x69, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x69, 0x6e, 0x6f, 0x6b, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x65, 0x72,
	0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x3b,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_v1_testapi_testapi_proto_rawDescOnce sync.Once
	file_proto_v1_testapi_testapi_proto_rawDescData = file_proto_v1_testapi_testapi_proto_rawDesc
)

func file_proto_v1_testapi_testapi_proto_rawDescGZIP() []byte {
	file_proto_v1_testapi_testapi_proto_rawDescOnce.Do(func() {
		file_proto_v1_testapi_testapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_testapi_testapi_proto_rawDescData)
	})
	return file_proto_v1_testapi_testapi_proto_rawDescData
}

var file_proto_v1_testapi_testapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1_testapi_testapi_proto_goTypes = []interface{}{
	(*TestAPIServiceEchoRequest)(nil),        // 0: testapi.v1.TestAPIServiceEchoRequest
	(*TestAPIServiceEchoResponse)(nil),       // 1: testapi.v1.TestAPIServiceEchoResponse
	(*TestAPIServiceRaiseErrorRequest)(nil),  // 2: testapi.v1.TestAPIServiceRaiseErrorRequest
	(*TestAPIServiceRaiseErrorResponse)(nil), // 3: testapi.v1.TestAPIServiceRaiseErrorResponse
}
var file_proto_v1_testapi_testapi_proto_depIdxs = []int32{
	0, // 0: testapi.v1.TestAPIService.Echo:input_type -> testapi.v1.TestAPIServiceEchoRequest
	2, // 1: testapi.v1.TestAPIService.RaiseError:input_type -> testapi.v1.TestAPIServiceRaiseErrorRequest
	1, // 2: testapi.v1.TestAPIService.Echo:output_type -> testapi.v1.TestAPIServiceEchoResponse
	3, // 3: testapi.v1.TestAPIService.RaiseError:output_type -> testapi.v1.TestAPIServiceRaiseErrorResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_testapi_testapi_proto_init() }
func file_proto_v1_testapi_testapi_proto_init() {
	if File_proto_v1_testapi_testapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_testapi_testapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAPIServiceEchoRequest); i {
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
		file_proto_v1_testapi_testapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAPIServiceEchoResponse); i {
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
		file_proto_v1_testapi_testapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAPIServiceRaiseErrorRequest); i {
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
		file_proto_v1_testapi_testapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAPIServiceRaiseErrorResponse); i {
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
			RawDescriptor: file_proto_v1_testapi_testapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_testapi_testapi_proto_goTypes,
		DependencyIndexes: file_proto_v1_testapi_testapi_proto_depIdxs,
		MessageInfos:      file_proto_v1_testapi_testapi_proto_msgTypes,
	}.Build()
	File_proto_v1_testapi_testapi_proto = out.File
	file_proto_v1_testapi_testapi_proto_rawDesc = nil
	file_proto_v1_testapi_testapi_proto_goTypes = nil
	file_proto_v1_testapi_testapi_proto_depIdxs = nil
}
