// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pack2.proto

package b

import (
	pack1 "./pack1"
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

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	M    *pack1.Model `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pack2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_pack2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_pack2_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetM() *pack1.Model {
	if x != nil {
		return x.M
	}
	return nil
}

var File_pack2_proto protoreflect.FileDescriptor

var file_pack2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x32, 0x1a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34,
	0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x01, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x01, 0x6d, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x32, 0x3b,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pack2_proto_rawDescOnce sync.Once
	file_pack2_proto_rawDescData = file_pack2_proto_rawDesc
)

func file_pack2_proto_rawDescGZIP() []byte {
	file_pack2_proto_rawDescOnce.Do(func() {
		file_pack2_proto_rawDescData = protoimpl.X.CompressGZIP(file_pack2_proto_rawDescData)
	})
	return file_pack2_proto_rawDescData
}

var file_pack2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pack2_proto_goTypes = []interface{}{
	(*Model)(nil),       // 0: p2.Model
	(*pack1.Model)(nil), // 1: p1.Model
}
var file_pack2_proto_depIdxs = []int32{
	1, // 0: p2.Model.m:type_name -> p1.Model
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pack2_proto_init() }
func file_pack2_proto_init() {
	if File_pack2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pack2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
			RawDescriptor: file_pack2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pack2_proto_goTypes,
		DependencyIndexes: file_pack2_proto_depIdxs,
		MessageInfos:      file_pack2_proto_msgTypes,
	}.Build()
	File_pack2_proto = out.File
	file_pack2_proto_rawDesc = nil
	file_pack2_proto_goTypes = nil
	file_pack2_proto_depIdxs = nil
}
