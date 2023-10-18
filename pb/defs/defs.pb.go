// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: defs.proto

package defs

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

type Status_Type int32

const (
	Status_UNKNOWN Status_Type = 0
)

// Enum value maps for Status_Type.
var (
	Status_Type_name = map[int32]string{
		0: "UNKNOWN",
	}
	Status_Type_value = map[string]int32{
		"UNKNOWN": 0,
	}
)

func (x Status_Type) Enum() *Status_Type {
	p := new(Status_Type)
	*p = x
	return p
}

func (x Status_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_defs_proto_enumTypes[0].Descriptor()
}

func (Status_Type) Type() protoreflect.EnumType {
	return &file_defs_proto_enumTypes[0]
}

func (x Status_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_Type.Descriptor instead.
func (Status_Type) EnumDescriptor() ([]byte, []int) {
	return file_defs_proto_rawDescGZIP(), []int{1, 0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_defs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_defs_proto_rawDescGZIP(), []int{0}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Status_Type `protobuf:"varint,1,opt,name=type,proto3,enum=defs.Status_Type" json:"type,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_defs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_defs_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetType() Status_Type {
	if x != nil {
		return x.Type
	}
	return Status_UNKNOWN
}

type MResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mresult:
	//
	//	*MResult_Status
	//	*MResult_Ok
	Mresult isMResult_Mresult `protobuf_oneof:"mresult"`
}

func (x *MResult) Reset() {
	*x = MResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MResult) ProtoMessage() {}

func (x *MResult) ProtoReflect() protoreflect.Message {
	mi := &file_defs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MResult.ProtoReflect.Descriptor instead.
func (*MResult) Descriptor() ([]byte, []int) {
	return file_defs_proto_rawDescGZIP(), []int{2}
}

func (m *MResult) GetMresult() isMResult_Mresult {
	if m != nil {
		return m.Mresult
	}
	return nil
}

func (x *MResult) GetStatus() *Status {
	if x, ok := x.GetMresult().(*MResult_Status); ok {
		return x.Status
	}
	return nil
}

func (x *MResult) GetOk() *MResult_MOk {
	if x, ok := x.GetMresult().(*MResult_Ok); ok {
		return x.Ok
	}
	return nil
}

type isMResult_Mresult interface {
	isMResult_Mresult()
}

type MResult_Status struct {
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type MResult_Ok struct {
	Ok *MResult_MOk `protobuf:"bytes,2,opt,name=ok,proto3,oneof"`
}

func (*MResult_Status) isMResult_Mresult() {}

func (*MResult_Ok) isMResult_Mresult() {}

type MResult_MOk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MResult_MOk) Reset() {
	*x = MResult_MOk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MResult_MOk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MResult_MOk) ProtoMessage() {}

func (x *MResult_MOk) ProtoReflect() protoreflect.Message {
	mi := &file_defs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MResult_MOk.ProtoReflect.Descriptor instead.
func (*MResult_MOk) Descriptor() ([]byte, []int) {
	return file_defs_proto_rawDescGZIP(), []int{2, 0}
}

var File_defs_proto protoreflect.FileDescriptor

var file_defs_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65,
	0x66, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x44, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x66, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x13, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x22, 0x68, 0x0a, 0x07, 0x4d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64,
	0x65, 0x66, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x65, 0x66, 0x73, 0x2e, 0x4d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x4d, 0x4f, 0x6b, 0x48, 0x00, 0x52, 0x02, 0x6f, 0x6b, 0x1a, 0x05, 0x0a, 0x03, 0x4d, 0x4f, 0x6b,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x20, 0x5a, 0x1e, 0x71,
	0x73, 0x74, 0x2d, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x70, 0x70, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x66, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_defs_proto_rawDescOnce sync.Once
	file_defs_proto_rawDescData = file_defs_proto_rawDesc
)

func file_defs_proto_rawDescGZIP() []byte {
	file_defs_proto_rawDescOnce.Do(func() {
		file_defs_proto_rawDescData = protoimpl.X.CompressGZIP(file_defs_proto_rawDescData)
	})
	return file_defs_proto_rawDescData
}

var file_defs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_defs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_defs_proto_goTypes = []interface{}{
	(Status_Type)(0),    // 0: defs.Status.Type
	(*Empty)(nil),       // 1: defs.Empty
	(*Status)(nil),      // 2: defs.Status
	(*MResult)(nil),     // 3: defs.MResult
	(*MResult_MOk)(nil), // 4: defs.MResult.MOk
}
var file_defs_proto_depIdxs = []int32{
	0, // 0: defs.Status.type:type_name -> defs.Status.Type
	2, // 1: defs.MResult.status:type_name -> defs.Status
	4, // 2: defs.MResult.ok:type_name -> defs.MResult.MOk
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_defs_proto_init() }
func file_defs_proto_init() {
	if File_defs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_defs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_defs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_defs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MResult); i {
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
		file_defs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MResult_MOk); i {
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
	file_defs_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MResult_Status)(nil),
		(*MResult_Ok)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_defs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_defs_proto_goTypes,
		DependencyIndexes: file_defs_proto_depIdxs,
		EnumInfos:         file_defs_proto_enumTypes,
		MessageInfos:      file_defs_proto_msgTypes,
	}.Build()
	File_defs_proto = out.File
	file_defs_proto_rawDesc = nil
	file_defs_proto_goTypes = nil
	file_defs_proto_depIdxs = nil
}