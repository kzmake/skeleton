// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: greeter/v1/greeter.proto

package greeter

import (
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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeter_v1_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greeter_v1_greeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_greeter_v1_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greeter_v1_greeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greeter_v1_greeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_greeter_v1_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_greeter_v1_greeter_proto protoreflect.FileDescriptor

var file_greeter_v1_greeter_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x6b, 0x65, 0x6c,
	0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a,
	0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x21, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0x77, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12,
	0x6c, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x21, 0x2e, 0x73, 0x6b, 0x65, 0x6c, 0x65,
	0x74, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6b,
	0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x7a, 0x6d, 0x61,
	0x6b, 0x65, 0x2f, 0x73, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greeter_v1_greeter_proto_rawDescOnce sync.Once
	file_greeter_v1_greeter_proto_rawDescData = file_greeter_v1_greeter_proto_rawDesc
)

func file_greeter_v1_greeter_proto_rawDescGZIP() []byte {
	file_greeter_v1_greeter_proto_rawDescOnce.Do(func() {
		file_greeter_v1_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_greeter_v1_greeter_proto_rawDescData)
	})
	return file_greeter_v1_greeter_proto_rawDescData
}

var file_greeter_v1_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greeter_v1_greeter_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: skeleton.greeter.v1.HelloRequest
	(*HelloResponse)(nil), // 1: skeleton.greeter.v1.HelloResponse
}
var file_greeter_v1_greeter_proto_depIdxs = []int32{
	0, // 0: skeleton.greeter.v1.Greeter.Hello:input_type -> skeleton.greeter.v1.HelloRequest
	1, // 1: skeleton.greeter.v1.Greeter.Hello:output_type -> skeleton.greeter.v1.HelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greeter_v1_greeter_proto_init() }
func file_greeter_v1_greeter_proto_init() {
	if File_greeter_v1_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greeter_v1_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_greeter_v1_greeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_greeter_v1_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greeter_v1_greeter_proto_goTypes,
		DependencyIndexes: file_greeter_v1_greeter_proto_depIdxs,
		MessageInfos:      file_greeter_v1_greeter_proto_msgTypes,
	}.Build()
	File_greeter_v1_greeter_proto = out.File
	file_greeter_v1_greeter_proto_rawDesc = nil
	file_greeter_v1_greeter_proto_goTypes = nil
	file_greeter_v1_greeter_proto_depIdxs = nil
}
