// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: cores/data_service.proto

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

type DataUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContentType int32  `protobuf:"varint,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Content     []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	DeviceId    string `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Cache       bool   `protobuf:"varint,5,opt,name=cache,proto3" json:"cache,omitempty"`
	Save        bool   `protobuf:"varint,6,opt,name=save,proto3" json:"save,omitempty"`
}

func (x *DataUploadRequest) Reset() {
	*x = DataUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_data_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUploadRequest) ProtoMessage() {}

func (x *DataUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_data_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUploadRequest.ProtoReflect.Descriptor instead.
func (*DataUploadRequest) Descriptor() ([]byte, []int) {
	return file_cores_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *DataUploadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataUploadRequest) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *DataUploadRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *DataUploadRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DataUploadRequest) GetCache() bool {
	if x != nil {
		return x.Cache
	}
	return false
}

func (x *DataUploadRequest) GetSave() bool {
	if x != nil {
		return x.Save
	}
	return false
}

type DataUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DataUploadResponse) Reset() {
	*x = DataUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_data_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUploadResponse) ProtoMessage() {}

func (x *DataUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cores_data_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUploadResponse.ProtoReflect.Descriptor instead.
func (*DataUploadResponse) Descriptor() ([]byte, []int) {
	return file_cores_data_service_proto_rawDescGZIP(), []int{1}
}

func (x *DataUploadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataUploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DataQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flux string            `protobuf:"bytes,1,opt,name=flux,proto3" json:"flux,omitempty"`
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataQueryRequest) Reset() {
	*x = DataQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_data_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataQueryRequest) ProtoMessage() {}

func (x *DataQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_data_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataQueryRequest.ProtoReflect.Descriptor instead.
func (*DataQueryRequest) Descriptor() ([]byte, []int) {
	return file_cores_data_service_proto_rawDescGZIP(), []int{2}
}

func (x *DataQueryRequest) GetFlux() string {
	if x != nil {
		return x.Flux
	}
	return ""
}

func (x *DataQueryRequest) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

type DataQueryTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataQueryTagRequest) Reset() {
	*x = DataQueryTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cores_data_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataQueryTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataQueryTagRequest) ProtoMessage() {}

func (x *DataQueryTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cores_data_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataQueryTagRequest.ProtoReflect.Descriptor instead.
func (*DataQueryTagRequest) Descriptor() ([]byte, []int) {
	return file_cores_data_service_proto_rawDescGZIP(), []int{3}
}

func (x *DataQueryTagRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataQueryTagRequest) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

var File_cores_data_service_proto protoreflect.FileDescriptor

var file_cores_data_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x74,
	0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x61, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x61,
	0x76, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x75, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x75, 0x78, 0x12, 0x35, 0x0a, 0x04, 0x76,
	0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x76, 0x61,
	0x72, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x98, 0x01, 0x0a, 0x13,
	0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x76, 0x61, 0x72, 0x73, 0x1a, 0x37, 0x0a,
	0x09, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xed, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a,
	0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d,
	0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cores_data_service_proto_rawDescOnce sync.Once
	file_cores_data_service_proto_rawDescData = file_cores_data_service_proto_rawDesc
)

func file_cores_data_service_proto_rawDescGZIP() []byte {
	file_cores_data_service_proto_rawDescOnce.Do(func() {
		file_cores_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cores_data_service_proto_rawDescData)
	})
	return file_cores_data_service_proto_rawDescData
}

var file_cores_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cores_data_service_proto_goTypes = []interface{}{
	(*DataUploadRequest)(nil),   // 0: cores.DataUploadRequest
	(*DataUploadResponse)(nil),  // 1: cores.DataUploadResponse
	(*DataQueryRequest)(nil),    // 2: cores.DataQueryRequest
	(*DataQueryTagRequest)(nil), // 3: cores.DataQueryTagRequest
	nil,                         // 4: cores.DataQueryRequest.VarsEntry
	nil,                         // 5: cores.DataQueryTagRequest.VarsEntry
	(*pb.Message)(nil),          // 6: pb.Message
}
var file_cores_data_service_proto_depIdxs = []int32{
	4, // 0: cores.DataQueryRequest.vars:type_name -> cores.DataQueryRequest.VarsEntry
	5, // 1: cores.DataQueryTagRequest.vars:type_name -> cores.DataQueryTagRequest.VarsEntry
	0, // 2: cores.DataService.Upload:input_type -> cores.DataUploadRequest
	2, // 3: cores.DataService.Compile:input_type -> cores.DataQueryRequest
	2, // 4: cores.DataService.Query:input_type -> cores.DataQueryRequest
	3, // 5: cores.DataService.QueryTag:input_type -> cores.DataQueryTagRequest
	1, // 6: cores.DataService.Upload:output_type -> cores.DataUploadResponse
	6, // 7: cores.DataService.Compile:output_type -> pb.Message
	6, // 8: cores.DataService.Query:output_type -> pb.Message
	6, // 9: cores.DataService.QueryTag:output_type -> pb.Message
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cores_data_service_proto_init() }
func file_cores_data_service_proto_init() {
	if File_cores_data_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cores_data_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUploadRequest); i {
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
		file_cores_data_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUploadResponse); i {
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
		file_cores_data_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataQueryRequest); i {
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
		file_cores_data_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataQueryTagRequest); i {
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
			RawDescriptor: file_cores_data_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cores_data_service_proto_goTypes,
		DependencyIndexes: file_cores_data_service_proto_depIdxs,
		MessageInfos:      file_cores_data_service_proto_msgTypes,
	}.Build()
	File_cores_data_service_proto = out.File
	file_cores_data_service_proto_rawDesc = nil
	file_cores_data_service_proto_goTypes = nil
	file_cores_data_service_proto_depIdxs = nil
}
