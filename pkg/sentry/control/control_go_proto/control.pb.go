// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: pkg/sentry/control/control.proto

package control_go_proto

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

type ControlConfig_Endpoint int32

const (
	ControlConfig_UNKNOWN   ControlConfig_Endpoint = 0
	ControlConfig_EVENTS    ControlConfig_Endpoint = 1
	ControlConfig_FS        ControlConfig_Endpoint = 2
	ControlConfig_LIFECYCLE ControlConfig_Endpoint = 3
	ControlConfig_LOGGING   ControlConfig_Endpoint = 4
	ControlConfig_PROFILE   ControlConfig_Endpoint = 5
	ControlConfig_USAGE     ControlConfig_Endpoint = 6
	ControlConfig_PROC      ControlConfig_Endpoint = 7
	ControlConfig_STATE     ControlConfig_Endpoint = 8
	ControlConfig_DEBUG     ControlConfig_Endpoint = 9
)

// Enum value maps for ControlConfig_Endpoint.
var (
	ControlConfig_Endpoint_name = map[int32]string{
		0: "UNKNOWN",
		1: "EVENTS",
		2: "FS",
		3: "LIFECYCLE",
		4: "LOGGING",
		5: "PROFILE",
		6: "USAGE",
		7: "PROC",
		8: "STATE",
		9: "DEBUG",
	}
	ControlConfig_Endpoint_value = map[string]int32{
		"UNKNOWN":   0,
		"EVENTS":    1,
		"FS":        2,
		"LIFECYCLE": 3,
		"LOGGING":   4,
		"PROFILE":   5,
		"USAGE":     6,
		"PROC":      7,
		"STATE":     8,
		"DEBUG":     9,
	}
)

func (x ControlConfig_Endpoint) Enum() *ControlConfig_Endpoint {
	p := new(ControlConfig_Endpoint)
	*p = x
	return p
}

func (x ControlConfig_Endpoint) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlConfig_Endpoint) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_sentry_control_control_proto_enumTypes[0].Descriptor()
}

func (ControlConfig_Endpoint) Type() protoreflect.EnumType {
	return &file_pkg_sentry_control_control_proto_enumTypes[0]
}

func (x ControlConfig_Endpoint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlConfig_Endpoint.Descriptor instead.
func (ControlConfig_Endpoint) EnumDescriptor() ([]byte, []int) {
	return file_pkg_sentry_control_control_proto_rawDescGZIP(), []int{0, 0}
}

type ControlConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowedControls []ControlConfig_Endpoint `protobuf:"varint,1,rep,packed,name=allowed_controls,json=allowedControls,proto3,enum=gvisor.ControlConfig_Endpoint" json:"allowed_controls,omitempty"`
}

func (x *ControlConfig) Reset() {
	*x = ControlConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sentry_control_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlConfig) ProtoMessage() {}

func (x *ControlConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sentry_control_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlConfig.ProtoReflect.Descriptor instead.
func (*ControlConfig) Descriptor() ([]byte, []int) {
	return file_pkg_sentry_control_control_proto_rawDescGZIP(), []int{0}
}

func (x *ControlConfig) GetAllowedControls() []ControlConfig_Endpoint {
	if x != nil {
		return x.AllowedControls
	}
	return nil
}

var File_pkg_sentry_control_control_proto protoreflect.FileDescriptor

var file_pkg_sentry_control_control_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x22, 0xdb, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x10,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x22, 0x7f, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02,
	0x46, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c,
	0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x47, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x04,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x09, 0x0a,
	0x05, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x52, 0x4f, 0x43,
	0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x08, 0x12, 0x09, 0x0a,
	0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x09, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sentry_control_control_proto_rawDescOnce sync.Once
	file_pkg_sentry_control_control_proto_rawDescData = file_pkg_sentry_control_control_proto_rawDesc
)

func file_pkg_sentry_control_control_proto_rawDescGZIP() []byte {
	file_pkg_sentry_control_control_proto_rawDescOnce.Do(func() {
		file_pkg_sentry_control_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sentry_control_control_proto_rawDescData)
	})
	return file_pkg_sentry_control_control_proto_rawDescData
}

var file_pkg_sentry_control_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_sentry_control_control_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_sentry_control_control_proto_goTypes = []interface{}{
	(ControlConfig_Endpoint)(0), // 0: gvisor.ControlConfig.Endpoint
	(*ControlConfig)(nil),       // 1: gvisor.ControlConfig
}
var file_pkg_sentry_control_control_proto_depIdxs = []int32{
	0, // 0: gvisor.ControlConfig.allowed_controls:type_name -> gvisor.ControlConfig.Endpoint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_sentry_control_control_proto_init() }
func file_pkg_sentry_control_control_proto_init() {
	if File_pkg_sentry_control_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sentry_control_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlConfig); i {
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
			RawDescriptor: file_pkg_sentry_control_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sentry_control_control_proto_goTypes,
		DependencyIndexes: file_pkg_sentry_control_control_proto_depIdxs,
		EnumInfos:         file_pkg_sentry_control_control_proto_enumTypes,
		MessageInfos:      file_pkg_sentry_control_control_proto_msgTypes,
	}.Build()
	File_pkg_sentry_control_control_proto = out.File
	file_pkg_sentry_control_control_proto_rawDesc = nil
	file_pkg_sentry_control_control_proto_goTypes = nil
	file_pkg_sentry_control_control_proto_depIdxs = nil
}
