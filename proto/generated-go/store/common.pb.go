// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/common.proto

package store

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

type Engine int32

const (
	Engine_ENGINE_UNSPECIFIED Engine = 0
	Engine_CLICKHOUSE         Engine = 1
	Engine_MYSQL              Engine = 2
	Engine_POSTGRES           Engine = 3
	Engine_SNOWFLAKE          Engine = 4
	Engine_SQLITE             Engine = 5
	Engine_TIDB               Engine = 6
	Engine_MONGODB            Engine = 7
	Engine_REDIS              Engine = 8
	Engine_ORACLE             Engine = 9
	Engine_SPANNER            Engine = 10
	Engine_MSSQL              Engine = 11
	Engine_REDSHIFT           Engine = 12
	Engine_MARIADB            Engine = 13
	Engine_OCEANBASE          Engine = 14
	Engine_DM                 Engine = 15
)

// Enum value maps for Engine.
var (
	Engine_name = map[int32]string{
		0:  "ENGINE_UNSPECIFIED",
		1:  "CLICKHOUSE",
		2:  "MYSQL",
		3:  "POSTGRES",
		4:  "SNOWFLAKE",
		5:  "SQLITE",
		6:  "TIDB",
		7:  "MONGODB",
		8:  "REDIS",
		9:  "ORACLE",
		10: "SPANNER",
		11: "MSSQL",
		12: "REDSHIFT",
		13: "MARIADB",
		14: "OCEANBASE",
		15: "DM",
	}
	Engine_value = map[string]int32{
		"ENGINE_UNSPECIFIED": 0,
		"CLICKHOUSE":         1,
		"MYSQL":              2,
		"POSTGRES":           3,
		"SNOWFLAKE":          4,
		"SQLITE":             5,
		"TIDB":               6,
		"MONGODB":            7,
		"REDIS":              8,
		"ORACLE":             9,
		"SPANNER":            10,
		"MSSQL":              11,
		"REDSHIFT":           12,
		"MARIADB":            13,
		"OCEANBASE":          14,
		"DM":                 15,
	}
)

func (x Engine) Enum() *Engine {
	p := new(Engine)
	*p = x
	return p
}

func (x Engine) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Engine) Descriptor() protoreflect.EnumDescriptor {
	return file_store_common_proto_enumTypes[0].Descriptor()
}

func (Engine) Type() protoreflect.EnumType {
	return &file_store_common_proto_enumTypes[0]
}

func (x Engine) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Engine.Descriptor instead.
func (Engine) EnumDescriptor() ([]byte, []int) {
	return file_store_common_proto_rawDescGZIP(), []int{0}
}

// Used internally for obfuscating the page token.
type PageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PageToken) Reset() {
	*x = PageToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageToken) ProtoMessage() {}

func (x *PageToken) ProtoReflect() protoreflect.Message {
	mi := &file_store_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageToken.ProtoReflect.Descriptor instead.
func (*PageToken) Descriptor() ([]byte, []int) {
	return file_store_common_proto_rawDescGZIP(), []int{0}
}

func (x *PageToken) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageToken) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_store_common_proto protoreflect.FileDescriptor

var file_store_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2a,
	0xdc, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e,
	0x47, 0x49, 0x4e, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x48, 0x4f, 0x55, 0x53, 0x45,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x4e, 0x4f, 0x57, 0x46, 0x4c, 0x41, 0x4b, 0x45, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x51,
	0x4c, 0x49, 0x54, 0x45, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x49, 0x44, 0x42, 0x10, 0x06,
	0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x47, 0x4f, 0x44, 0x42, 0x10, 0x07, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x45, 0x44, 0x49, 0x53, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x52, 0x41, 0x43,
	0x4c, 0x45, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x10,
	0x0a, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x53, 0x53, 0x51, 0x4c, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x44, 0x53, 0x48, 0x49, 0x46, 0x54, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41,
	0x52, 0x49, 0x41, 0x44, 0x42, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x43, 0x45, 0x41, 0x4e,
	0x42, 0x41, 0x53, 0x45, 0x10, 0x0e, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x4d, 0x10, 0x0f, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_common_proto_rawDescOnce sync.Once
	file_store_common_proto_rawDescData = file_store_common_proto_rawDesc
)

func file_store_common_proto_rawDescGZIP() []byte {
	file_store_common_proto_rawDescOnce.Do(func() {
		file_store_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_common_proto_rawDescData)
	})
	return file_store_common_proto_rawDescData
}

var file_store_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_common_proto_goTypes = []interface{}{
	(Engine)(0),       // 0: bytebase.store.Engine
	(*PageToken)(nil), // 1: bytebase.store.PageToken
}
var file_store_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_common_proto_init() }
func file_store_common_proto_init() {
	if File_store_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageToken); i {
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
			RawDescriptor: file_store_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_common_proto_goTypes,
		DependencyIndexes: file_store_common_proto_depIdxs,
		EnumInfos:         file_store_common_proto_enumTypes,
		MessageInfos:      file_store_common_proto_msgTypes,
	}.Build()
	File_store_common_proto = out.File
	file_store_common_proto_rawDesc = nil
	file_store_common_proto_goTypes = nil
	file_store_common_proto_depIdxs = nil
}
