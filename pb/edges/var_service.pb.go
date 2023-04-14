// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: edges/var_service.proto

package edges

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

type ListVarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *pb.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// string device_id = 2;
	Tags string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListVarRequest) Reset() {
	*x = ListVarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_var_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVarRequest) ProtoMessage() {}

func (x *ListVarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edges_var_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVarRequest.ProtoReflect.Descriptor instead.
func (*ListVarRequest) Descriptor() ([]byte, []int) {
	return file_edges_var_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListVarRequest) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListVarRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ListVarRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListVarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *pb.Page  `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Count uint32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Var   []*pb.Var `protobuf:"bytes,3,rep,name=var,proto3" json:"var,omitempty"`
}

func (x *ListVarResponse) Reset() {
	*x = ListVarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_var_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVarResponse) ProtoMessage() {}

func (x *ListVarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edges_var_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVarResponse.ProtoReflect.Descriptor instead.
func (*ListVarResponse) Descriptor() ([]byte, []int) {
	return file_edges_var_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListVarResponse) GetPage() *pb.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListVarResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListVarResponse) GetVar() []*pb.Var {
	if x != nil {
		return x.Var
	}
	return nil
}

type CloneVarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // string device_id = 2;
}

func (x *CloneVarRequest) Reset() {
	*x = CloneVarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_var_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloneVarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloneVarRequest) ProtoMessage() {}

func (x *CloneVarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edges_var_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloneVarRequest.ProtoReflect.Descriptor instead.
func (*CloneVarRequest) Descriptor() ([]byte, []int) {
	return file_edges_var_service_proto_rawDescGZIP(), []int{2}
}

func (x *CloneVarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PullVarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64  `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// string device_id = 3;
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PullVarRequest) Reset() {
	*x = PullVarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_var_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullVarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullVarRequest) ProtoMessage() {}

func (x *PullVarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edges_var_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullVarRequest.ProtoReflect.Descriptor instead.
func (*PullVarRequest) Descriptor() ([]byte, []int) {
	return file_edges_var_service_proto_rawDescGZIP(), []int{3}
}

func (x *PullVarRequest) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullVarRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullVarRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type PullVarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After int64     `protobuf:"varint,1,opt,name=after,proto3" json:"after,omitempty"`
	Limit uint32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Var   []*pb.Var `protobuf:"bytes,3,rep,name=var,proto3" json:"var,omitempty"`
}

func (x *PullVarResponse) Reset() {
	*x = PullVarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edges_var_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullVarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullVarResponse) ProtoMessage() {}

func (x *PullVarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edges_var_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullVarResponse.ProtoReflect.Descriptor instead.
func (*PullVarResponse) Descriptor() ([]byte, []int) {
	return file_edges_var_service_proto_rawDescGZIP(), []int{4}
}

func (x *PullVarResponse) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *PullVarResponse) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PullVarResponse) GetVar() []*pb.Var {
	if x != nil {
		return x.Var
	}
	return nil
}

var File_edges_var_service_proto protoreflect.FileDescriptor

var file_edges_var_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73,
	0x1a, 0x11, 0x76, 0x61, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x60, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x76, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x52,
	0x03, 0x76, 0x61, 0x72, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x56, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x0e, 0x50, 0x75, 0x6c, 0x6c, 0x56,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x0f, 0x50, 0x75, 0x6c,
	0x6c, 0x56, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x76, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x52, 0x03,
	0x76, 0x61, 0x72, 0x32, 0x87, 0x05, 0x0a, 0x0a, 0x56, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x07, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x61, 0x72, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x22, 0x00,
	0x12, 0x1c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x07, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x61, 0x72, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x22, 0x00, 0x12, 0x19,
	0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x07,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x56, 0x69, 0x65,
	0x77, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x12, 0x16,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x56, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x6e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x24,
	0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x06, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x56,
	0x61, 0x72, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c,
	0x56, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x70, 0x6c,
	0x65, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x6d, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x3b, 0x65, 0x64, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edges_var_service_proto_rawDescOnce sync.Once
	file_edges_var_service_proto_rawDescData = file_edges_var_service_proto_rawDesc
)

func file_edges_var_service_proto_rawDescGZIP() []byte {
	file_edges_var_service_proto_rawDescOnce.Do(func() {
		file_edges_var_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_edges_var_service_proto_rawDescData)
	})
	return file_edges_var_service_proto_rawDescData
}

var file_edges_var_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_edges_var_service_proto_goTypes = []interface{}{
	(*ListVarRequest)(nil),  // 0: edges.ListVarRequest
	(*ListVarResponse)(nil), // 1: edges.ListVarResponse
	(*CloneVarRequest)(nil), // 2: edges.CloneVarRequest
	(*PullVarRequest)(nil),  // 3: edges.PullVarRequest
	(*PullVarResponse)(nil), // 4: edges.PullVarResponse
	(*pb.Page)(nil),         // 5: pb.Page
	(*pb.Var)(nil),          // 6: pb.Var
	(*pb.Id)(nil),           // 7: pb.Id
	(*pb.Name)(nil),         // 8: pb.Name
	(*pb.VarValue)(nil),     // 9: pb.VarValue
	(*pb.VarNameValue)(nil), // 10: pb.VarNameValue
	(*pb.MyBool)(nil),       // 11: pb.MyBool
}
var file_edges_var_service_proto_depIdxs = []int32{
	5,  // 0: edges.ListVarRequest.page:type_name -> pb.Page
	5,  // 1: edges.ListVarResponse.page:type_name -> pb.Page
	6,  // 2: edges.ListVarResponse.var:type_name -> pb.Var
	6,  // 3: edges.PullVarResponse.var:type_name -> pb.Var
	6,  // 4: edges.VarService.Create:input_type -> pb.Var
	6,  // 5: edges.VarService.Update:input_type -> pb.Var
	7,  // 6: edges.VarService.View:input_type -> pb.Id
	8,  // 7: edges.VarService.ViewByName:input_type -> pb.Name
	7,  // 8: edges.VarService.Delete:input_type -> pb.Id
	0,  // 9: edges.VarService.List:input_type -> edges.ListVarRequest
	2,  // 10: edges.VarService.Clone:input_type -> edges.CloneVarRequest
	7,  // 11: edges.VarService.GetValue:input_type -> pb.Id
	9,  // 12: edges.VarService.SetValue:input_type -> pb.VarValue
	9,  // 13: edges.VarService.SetValueUnchecked:input_type -> pb.VarValue
	8,  // 14: edges.VarService.GetValueByName:input_type -> pb.Name
	10, // 15: edges.VarService.SetValueByName:input_type -> pb.VarNameValue
	10, // 16: edges.VarService.SetValueByNameUnchecked:input_type -> pb.VarNameValue
	7,  // 17: edges.VarService.ViewWithDeleted:input_type -> pb.Id
	3,  // 18: edges.VarService.Pull:input_type -> edges.PullVarRequest
	6,  // 19: edges.VarService.Create:output_type -> pb.Var
	6,  // 20: edges.VarService.Update:output_type -> pb.Var
	6,  // 21: edges.VarService.View:output_type -> pb.Var
	6,  // 22: edges.VarService.ViewByName:output_type -> pb.Var
	11, // 23: edges.VarService.Delete:output_type -> pb.MyBool
	1,  // 24: edges.VarService.List:output_type -> edges.ListVarResponse
	11, // 25: edges.VarService.Clone:output_type -> pb.MyBool
	9,  // 26: edges.VarService.GetValue:output_type -> pb.VarValue
	11, // 27: edges.VarService.SetValue:output_type -> pb.MyBool
	11, // 28: edges.VarService.SetValueUnchecked:output_type -> pb.MyBool
	10, // 29: edges.VarService.GetValueByName:output_type -> pb.VarNameValue
	11, // 30: edges.VarService.SetValueByName:output_type -> pb.MyBool
	11, // 31: edges.VarService.SetValueByNameUnchecked:output_type -> pb.MyBool
	6,  // 32: edges.VarService.ViewWithDeleted:output_type -> pb.Var
	4,  // 33: edges.VarService.Pull:output_type -> edges.PullVarResponse
	19, // [19:34] is the sub-list for method output_type
	4,  // [4:19] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_edges_var_service_proto_init() }
func file_edges_var_service_proto_init() {
	if File_edges_var_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_edges_var_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVarRequest); i {
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
		file_edges_var_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVarResponse); i {
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
		file_edges_var_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloneVarRequest); i {
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
		file_edges_var_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullVarRequest); i {
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
		file_edges_var_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullVarResponse); i {
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
			RawDescriptor: file_edges_var_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_edges_var_service_proto_goTypes,
		DependencyIndexes: file_edges_var_service_proto_depIdxs,
		MessageInfos:      file_edges_var_service_proto_msgTypes,
	}.Build()
	File_edges_var_service_proto = out.File
	file_edges_var_service_proto_rawDesc = nil
	file_edges_var_service_proto_goTypes = nil
	file_edges_var_service_proto_depIdxs = nil
}
