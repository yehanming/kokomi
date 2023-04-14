// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: edges/slot_service.proto

package edges

import (
	context "context"
	pb "github.com/snple/kokomi/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SlotService_Create_FullMethodName          = "/edges.SlotService/Create"
	SlotService_Update_FullMethodName          = "/edges.SlotService/Update"
	SlotService_View_FullMethodName            = "/edges.SlotService/View"
	SlotService_ViewByName_FullMethodName      = "/edges.SlotService/ViewByName"
	SlotService_Delete_FullMethodName          = "/edges.SlotService/Delete"
	SlotService_List_FullMethodName            = "/edges.SlotService/List"
	SlotService_Link_FullMethodName            = "/edges.SlotService/Link"
	SlotService_Clone_FullMethodName           = "/edges.SlotService/Clone"
	SlotService_ViewWithDeleted_FullMethodName = "/edges.SlotService/ViewWithDeleted"
	SlotService_Pull_FullMethodName            = "/edges.SlotService/Pull"
)

// SlotServiceClient is the client API for SlotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlotServiceClient interface {
	Create(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error)
	Update(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error)
	View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error)
	ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Slot, error)
	Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error)
	List(ctx context.Context, in *ListSlotRequest, opts ...grpc.CallOption) (*ListSlotResponse, error)
	Link(ctx context.Context, in *LinkSlotRequest, opts ...grpc.CallOption) (*pb.MyBool, error)
	Clone(ctx context.Context, in *CloneSlotRequest, opts ...grpc.CallOption) (*pb.MyBool, error)
	ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error)
	Pull(ctx context.Context, in *PullSlotRequest, opts ...grpc.CallOption) (*PullSlotResponse, error)
}

type slotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSlotServiceClient(cc grpc.ClientConnInterface) SlotServiceClient {
	return &slotServiceClient{cc}
}

func (c *slotServiceClient) Create(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error) {
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Update(ctx context.Context, in *pb.Slot, opts ...grpc.CallOption) (*pb.Slot, error) {
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error) {
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*pb.Slot, error) {
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_ViewByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SlotService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) List(ctx context.Context, in *ListSlotRequest, opts ...grpc.CallOption) (*ListSlotResponse, error) {
	out := new(ListSlotResponse)
	err := c.cc.Invoke(ctx, SlotService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Link(ctx context.Context, in *LinkSlotRequest, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SlotService_Link_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Clone(ctx context.Context, in *CloneSlotRequest, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, SlotService_Clone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) ViewWithDeleted(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.Slot, error) {
	out := new(pb.Slot)
	err := c.cc.Invoke(ctx, SlotService_ViewWithDeleted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) Pull(ctx context.Context, in *PullSlotRequest, opts ...grpc.CallOption) (*PullSlotResponse, error) {
	out := new(PullSlotResponse)
	err := c.cc.Invoke(ctx, SlotService_Pull_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlotServiceServer is the server API for SlotService service.
// All implementations must embed UnimplementedSlotServiceServer
// for forward compatibility
type SlotServiceServer interface {
	Create(context.Context, *pb.Slot) (*pb.Slot, error)
	Update(context.Context, *pb.Slot) (*pb.Slot, error)
	View(context.Context, *pb.Id) (*pb.Slot, error)
	ViewByName(context.Context, *pb.Name) (*pb.Slot, error)
	Delete(context.Context, *pb.Id) (*pb.MyBool, error)
	List(context.Context, *ListSlotRequest) (*ListSlotResponse, error)
	Link(context.Context, *LinkSlotRequest) (*pb.MyBool, error)
	Clone(context.Context, *CloneSlotRequest) (*pb.MyBool, error)
	ViewWithDeleted(context.Context, *pb.Id) (*pb.Slot, error)
	Pull(context.Context, *PullSlotRequest) (*PullSlotResponse, error)
	mustEmbedUnimplementedSlotServiceServer()
}

// UnimplementedSlotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSlotServiceServer struct {
}

func (UnimplementedSlotServiceServer) Create(context.Context, *pb.Slot) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSlotServiceServer) Update(context.Context, *pb.Slot) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSlotServiceServer) View(context.Context, *pb.Id) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedSlotServiceServer) ViewByName(context.Context, *pb.Name) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewByName not implemented")
}
func (UnimplementedSlotServiceServer) Delete(context.Context, *pb.Id) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSlotServiceServer) List(context.Context, *ListSlotRequest) (*ListSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSlotServiceServer) Link(context.Context, *LinkSlotRequest) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (UnimplementedSlotServiceServer) Clone(context.Context, *CloneSlotRequest) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clone not implemented")
}
func (UnimplementedSlotServiceServer) ViewWithDeleted(context.Context, *pb.Id) (*pb.Slot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewWithDeleted not implemented")
}
func (UnimplementedSlotServiceServer) Pull(context.Context, *PullSlotRequest) (*PullSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedSlotServiceServer) mustEmbedUnimplementedSlotServiceServer() {}

// UnsafeSlotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlotServiceServer will
// result in compilation errors.
type UnsafeSlotServiceServer interface {
	mustEmbedUnimplementedSlotServiceServer()
}

func RegisterSlotServiceServer(s grpc.ServiceRegistrar, srv SlotServiceServer) {
	s.RegisterService(&SlotService_ServiceDesc, srv)
}

func _SlotService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Slot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Create(ctx, req.(*pb.Slot))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Slot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Update(ctx, req.(*pb.Slot))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).View(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_ViewByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).ViewByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_ViewByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).ViewByName(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Delete(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).List(ctx, req.(*ListSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Link_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Link(ctx, req.(*LinkSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Clone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Clone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Clone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Clone(ctx, req.(*CloneSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_ViewWithDeleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).ViewWithDeleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_ViewWithDeleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).ViewWithDeleted(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SlotService_Pull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).Pull(ctx, req.(*PullSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SlotService_ServiceDesc is the grpc.ServiceDesc for SlotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SlotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edges.SlotService",
	HandlerType: (*SlotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SlotService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SlotService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _SlotService_View_Handler,
		},
		{
			MethodName: "ViewByName",
			Handler:    _SlotService_ViewByName_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SlotService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SlotService_List_Handler,
		},
		{
			MethodName: "Link",
			Handler:    _SlotService_Link_Handler,
		},
		{
			MethodName: "Clone",
			Handler:    _SlotService_Clone_Handler,
		},
		{
			MethodName: "ViewWithDeleted",
			Handler:    _SlotService_ViewWithDeleted_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _SlotService_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edges/slot_service.proto",
}
