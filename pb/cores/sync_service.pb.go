// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: cores/sync_service.proto

package cores

import (
	pb "github.com/snple/kokomi/pb"
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

type SyncUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Updated int64  `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *SyncUpdated) Reset() {
	*x = SyncUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_sync_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncUpdated) ProtoMessage() {}

func (x *SyncUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_cores_sync_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncUpdated.ProtoReflect.Descriptor instead.
func (*SyncUpdated) Descriptor() ([]byte, []int) {
	return file_cores_sync_service_proto_rawDescGZIP(), []int{0}
}

func (x *SyncUpdated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SyncUpdated) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_cores_sync_service_proto protoreflect.FileDescriptor

var file_cores_sync_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x32, 0xdb, 0x03, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x11, 0x57, 0x61, 0x69,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x06,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x54, 0x61, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x13, 0x57, 0x61, 0x69, 0x74, 0x54, 0x61, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x37, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x57, 0x69, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x57, 0x69, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x00, 0x12,
	0x2e, 0x0a, 0x14, 0x57, 0x61, 0x69, 0x74, 0x57, 0x69, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x30, 0x01, 0x32,
	0x13, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f,
	0x70, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cores_sync_service_proto_rawDescOnce sync.Once
	file_cores_sync_service_proto_rawDescData = file_cores_sync_service_proto_rawDesc
)

func file_cores_sync_service_proto_rawDescGZIP() []byte {
	file_cores_sync_service_proto_rawDescOnce.Do(func() {
		file_cores_sync_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cores_sync_service_proto_rawDescData)
	})
	return file_cores_sync_service_proto_rawDescData
}

var file_cores_sync_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cores_sync_service_proto_goTypes = []interface{}{
	(*SyncUpdated)(nil), // 0: cores.SyncUpdated
	(*pb.Id)(nil),       // 1: pb.Id
	(*pb.MyBool)(nil),   // 2: pb.MyBool
}
var file_cores_sync_service_proto_depIdxs = []int32{
	0, // 0: cores.SyncService.SetDeviceUpdated:input_type -> cores.SyncUpdated
	1, // 1: cores.SyncService.GetDeviceUpdated:input_type -> pb.Id
	1, // 2: cores.SyncService.WaitDeviceUpdated:input_type -> pb.Id
	0, // 3: cores.SyncService.SetTagValueUpdated:input_type -> cores.SyncUpdated
	1, // 4: cores.SyncService.GetTagValueUpdated:input_type -> pb.Id
	1, // 5: cores.SyncService.WaitTagValueUpdated:input_type -> pb.Id
	0, // 6: cores.SyncService.SetWireValueUpdated:input_type -> cores.SyncUpdated
	1, // 7: cores.SyncService.GetWireValueUpdated:input_type -> pb.Id
	1, // 8: cores.SyncService.WaitWireValueUpdated:input_type -> pb.Id
	2, // 9: cores.SyncService.SetDeviceUpdated:output_type -> pb.MyBool
	0, // 10: cores.SyncService.GetDeviceUpdated:output_type -> cores.SyncUpdated
	2, // 11: cores.SyncService.WaitDeviceUpdated:output_type -> pb.MyBool
	2, // 12: cores.SyncService.SetTagValueUpdated:output_type -> pb.MyBool
	0, // 13: cores.SyncService.GetTagValueUpdated:output_type -> cores.SyncUpdated
	2, // 14: cores.SyncService.WaitTagValueUpdated:output_type -> pb.MyBool
	2, // 15: cores.SyncService.SetWireValueUpdated:output_type -> pb.MyBool
	0, // 16: cores.SyncService.GetWireValueUpdated:output_type -> cores.SyncUpdated
	2, // 17: cores.SyncService.WaitWireValueUpdated:output_type -> pb.MyBool
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cores_sync_service_proto_init() }
func file_cores_sync_service_proto_init() {
	if File_cores_sync_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cores_sync_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncUpdated); i {
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
			RawDescriptor: file_cores_sync_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_cores_sync_service_proto_goTypes,
		DependencyIndexes: file_cores_sync_service_proto_depIdxs,
		MessageInfos:      file_cores_sync_service_proto_msgTypes,
	}.Build()
	File_cores_sync_service_proto = out.File
	file_cores_sync_service_proto_rawDesc = nil
	file_cores_sync_service_proto_goTypes = nil
	file_cores_sync_service_proto_depIdxs = nil
}
