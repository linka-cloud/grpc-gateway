// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: examples/internal/proto/oneofenum/oneof_enum.proto

package oneofenum

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

type ExampleEnum int32

const (
	ExampleEnum_EXAMPLE_ENUM_UNSPECIFIED ExampleEnum = 0
	ExampleEnum_EXAMPLE_ENUM_FIRST       ExampleEnum = 1
)

// Enum value maps for ExampleEnum.
var (
	ExampleEnum_name = map[int32]string{
		0: "EXAMPLE_ENUM_UNSPECIFIED",
		1: "EXAMPLE_ENUM_FIRST",
	}
	ExampleEnum_value = map[string]int32{
		"EXAMPLE_ENUM_UNSPECIFIED": 0,
		"EXAMPLE_ENUM_FIRST":       1,
	}
)

func (x ExampleEnum) Enum() *ExampleEnum {
	p := new(ExampleEnum)
	*p = x
	return p
}

func (x ExampleEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExampleEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_examples_internal_proto_oneofenum_oneof_enum_proto_enumTypes[0].Descriptor()
}

func (ExampleEnum) Type() protoreflect.EnumType {
	return &file_examples_internal_proto_oneofenum_oneof_enum_proto_enumTypes[0]
}

func (x ExampleEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExampleEnum.Descriptor instead.
func (ExampleEnum) EnumDescriptor() ([]byte, []int) {
	return file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescGZIP(), []int{0}
}

type OneofEnumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to One:
	//
	//	*OneofEnumMessage_ExampleEnum
	One isOneofEnumMessage_One `protobuf_oneof:"one"`
}

func (x *OneofEnumMessage) Reset() {
	*x = OneofEnumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_oneofenum_oneof_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofEnumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofEnumMessage) ProtoMessage() {}

func (x *OneofEnumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_oneofenum_oneof_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofEnumMessage.ProtoReflect.Descriptor instead.
func (*OneofEnumMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescGZIP(), []int{0}
}

func (m *OneofEnumMessage) GetOne() isOneofEnumMessage_One {
	if m != nil {
		return m.One
	}
	return nil
}

func (x *OneofEnumMessage) GetExampleEnum() ExampleEnum {
	if x, ok := x.GetOne().(*OneofEnumMessage_ExampleEnum); ok {
		return x.ExampleEnum
	}
	return ExampleEnum_EXAMPLE_ENUM_UNSPECIFIED
}

type isOneofEnumMessage_One interface {
	isOneofEnumMessage_One()
}

type OneofEnumMessage_ExampleEnum struct {
	ExampleEnum ExampleEnum `protobuf:"varint,1,opt,name=example_enum,json=exampleEnum,proto3,enum=grpc.gateway.examples.internal.proto.oneofenum.ExampleEnum,oneof"`
}

func (*OneofEnumMessage_ExampleEnum) isOneofEnumMessage_One() {}

var File_examples_internal_proto_oneofenum_oneof_enum_proto protoreflect.FileDescriptor

var file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x65,
	0x6e, 0x75, 0x6d, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x65, 0x6e, 0x75, 0x6d, 0x22, 0x7b, 0x0a, 0x10, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x45, 0x6e, 0x75,
	0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x0b, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x42, 0x05, 0x0a, 0x03, 0x6f, 0x6e,
	0x65, 0x2a, 0x43, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16,
	0x0a, 0x12, 0x45, 0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x46,
	0x49, 0x52, 0x53, 0x54, 0x10, 0x01, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescOnce sync.Once
	file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescData = file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDesc
)

func file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescData)
	})
	return file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDescData
}

var file_examples_internal_proto_oneofenum_oneof_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_examples_internal_proto_oneofenum_oneof_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_examples_internal_proto_oneofenum_oneof_enum_proto_goTypes = []interface{}{
	(ExampleEnum)(0),         // 0: grpc.gateway.examples.internal.proto.oneofenum.ExampleEnum
	(*OneofEnumMessage)(nil), // 1: grpc.gateway.examples.internal.proto.oneofenum.OneofEnumMessage
}
var file_examples_internal_proto_oneofenum_oneof_enum_proto_depIdxs = []int32{
	0, // 0: grpc.gateway.examples.internal.proto.oneofenum.OneofEnumMessage.example_enum:type_name -> grpc.gateway.examples.internal.proto.oneofenum.ExampleEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_oneofenum_oneof_enum_proto_init() }
func file_examples_internal_proto_oneofenum_oneof_enum_proto_init() {
	if File_examples_internal_proto_oneofenum_oneof_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_oneofenum_oneof_enum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofEnumMessage); i {
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
	file_examples_internal_proto_oneofenum_oneof_enum_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OneofEnumMessage_ExampleEnum)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examples_internal_proto_oneofenum_oneof_enum_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_oneofenum_oneof_enum_proto_depIdxs,
		EnumInfos:         file_examples_internal_proto_oneofenum_oneof_enum_proto_enumTypes,
		MessageInfos:      file_examples_internal_proto_oneofenum_oneof_enum_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_oneofenum_oneof_enum_proto = out.File
	file_examples_internal_proto_oneofenum_oneof_enum_proto_rawDesc = nil
	file_examples_internal_proto_oneofenum_oneof_enum_proto_goTypes = nil
	file_examples_internal_proto_oneofenum_oneof_enum_proto_depIdxs = nil
}
