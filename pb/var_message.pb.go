// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: var_message.proto

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

type Var struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Tags     string `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	Type     string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	DataType string `protobuf:"bytes,7,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Value    string `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
	HValue   string `protobuf:"bytes,9,opt,name=h_value,json=hValue,proto3" json:"h_value,omitempty"`
	LValue   string `protobuf:"bytes,10,opt,name=l_value,json=lValue,proto3" json:"l_value,omitempty"`
	Config   string `protobuf:"bytes,11,opt,name=config,proto3" json:"config,omitempty"`
	Status   int32  `protobuf:"zigzag32,12,opt,name=status,proto3" json:"status,omitempty"`
	Access   int32  `protobuf:"zigzag32,13,opt,name=access,proto3" json:"access,omitempty"`
	Save     int32  `protobuf:"zigzag32,14,opt,name=save,proto3" json:"save,omitempty"`
	Created  int64  `protobuf:"varint,15,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64  `protobuf:"varint,16,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted  int64  `protobuf:"varint,17,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Var) Reset() {
	*x = Var{}
	if protoimpl.UnsafeEnabled {
		mi := &file_var_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Var) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Var) ProtoMessage() {}

func (x *Var) ProtoReflect() protoreflect.Message {
	mi := &file_var_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Var.ProtoReflect.Descriptor instead.
func (*Var) Descriptor() ([]byte, []int) {
	return file_var_message_proto_rawDescGZIP(), []int{0}
}

func (x *Var) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Var) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Var) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Var) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Var) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Var) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Var) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Var) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Var) GetHValue() string {
	if x != nil {
		return x.HValue
	}
	return ""
}

func (x *Var) GetLValue() string {
	if x != nil {
		return x.LValue
	}
	return ""
}

func (x *Var) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Var) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Var) GetAccess() int32 {
	if x != nil {
		return x.Access
	}
	return 0
}

func (x *Var) GetSave() int32 {
	if x != nil {
		return x.Save
	}
	return 0
}

func (x *Var) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Var) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Var) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type VarValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Updated int64  `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *VarValue) Reset() {
	*x = VarValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_var_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VarValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VarValue) ProtoMessage() {}

func (x *VarValue) ProtoReflect() protoreflect.Message {
	mi := &file_var_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VarValue.ProtoReflect.Descriptor instead.
func (*VarValue) Descriptor() ([]byte, []int) {
	return file_var_message_proto_rawDescGZIP(), []int{1}
}

func (x *VarValue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VarValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *VarValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type VarNameValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Updated int64  `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *VarNameValue) Reset() {
	*x = VarNameValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_var_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VarNameValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VarNameValue) ProtoMessage() {}

func (x *VarNameValue) ProtoReflect() protoreflect.Message {
	mi := &file_var_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VarNameValue.ProtoReflect.Descriptor instead.
func (*VarNameValue) Descriptor() ([]byte, []int) {
	return file_var_message_proto_rawDescGZIP(), []int{2}
}

func (x *VarNameValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VarNameValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *VarNameValue) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_var_message_proto protoreflect.FileDescriptor

var file_var_message_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x61, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x91, 0x03, 0x0a, 0x03, 0x56, 0x61, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x76,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x73, 0x61, 0x76, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x4a, 0x0a, 0x08, 0x56,
	0x61, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x52, 0x0a, 0x0c, 0x56, 0x61, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f,
	0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_var_message_proto_rawDescOnce sync.Once
	file_var_message_proto_rawDescData = file_var_message_proto_rawDesc
)

func file_var_message_proto_rawDescGZIP() []byte {
	file_var_message_proto_rawDescOnce.Do(func() {
		file_var_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_var_message_proto_rawDescData)
	})
	return file_var_message_proto_rawDescData
}

var file_var_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_var_message_proto_goTypes = []interface{}{
	(*Var)(nil),          // 0: pb.Var
	(*VarValue)(nil),     // 1: pb.VarValue
	(*VarNameValue)(nil), // 2: pb.VarNameValue
}
var file_var_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_var_message_proto_init() }
func file_var_message_proto_init() {
	if File_var_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_var_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Var); i {
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
		file_var_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VarValue); i {
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
		file_var_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VarNameValue); i {
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
			RawDescriptor: file_var_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_var_message_proto_goTypes,
		DependencyIndexes: file_var_message_proto_depIdxs,
		MessageInfos:      file_var_message_proto_msgTypes,
	}.Build()
	File_var_message_proto = out.File
	file_var_message_proto_rawDesc = nil
	file_var_message_proto_goTypes = nil
	file_var_message_proto_depIdxs = nil
}
