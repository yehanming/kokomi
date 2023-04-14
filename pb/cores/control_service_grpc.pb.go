// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: cores/control_service.proto

package cores

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
	ControlService_GetTagValue_FullMethodName = "/cores.ControlService/GetTagValue"
	ControlService_SetTagValue_FullMethodName = "/cores.ControlService/SetTagValue"
)

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlServiceClient interface {
	GetTagValue(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.TagValue, error)
	SetTagValue(ctx context.Context, in *pb.TagValue, opts ...grpc.CallOption) (*pb.TagValue, error)
}

type controlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControlServiceClient(cc grpc.ClientConnInterface) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) GetTagValue(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.TagValue, error) {
	out := new(pb.TagValue)
	err := c.cc.Invoke(ctx, ControlService_GetTagValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) SetTagValue(ctx context.Context, in *pb.TagValue, opts ...grpc.CallOption) (*pb.TagValue, error) {
	out := new(pb.TagValue)
	err := c.cc.Invoke(ctx, ControlService_SetTagValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
// All implementations must embed UnimplementedControlServiceServer
// for forward compatibility
type ControlServiceServer interface {
	GetTagValue(context.Context, *pb.Id) (*pb.TagValue, error)
	SetTagValue(context.Context, *pb.TagValue) (*pb.TagValue, error)
	mustEmbedUnimplementedControlServiceServer()
}

// UnimplementedControlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedControlServiceServer struct {
}

func (UnimplementedControlServiceServer) GetTagValue(context.Context, *pb.Id) (*pb.TagValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagValue not implemented")
}
func (UnimplementedControlServiceServer) SetTagValue(context.Context, *pb.TagValue) (*pb.TagValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTagValue not implemented")
}
func (UnimplementedControlServiceServer) mustEmbedUnimplementedControlServiceServer() {}

// UnsafeControlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlServiceServer will
// result in compilation errors.
type UnsafeControlServiceServer interface {
	mustEmbedUnimplementedControlServiceServer()
}

func RegisterControlServiceServer(s grpc.ServiceRegistrar, srv ControlServiceServer) {
	s.RegisterService(&ControlService_ServiceDesc, srv)
}

func _ControlService_GetTagValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).GetTagValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_GetTagValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).GetTagValue(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_SetTagValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.TagValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).SetTagValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_SetTagValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).SetTagValue(ctx, req.(*pb.TagValue))
	}
	return interceptor(ctx, in, info, handler)
}

// ControlService_ServiceDesc is the grpc.ServiceDesc for ControlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cores.ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTagValue",
			Handler:    _ControlService_GetTagValue_Handler,
		},
		{
			MethodName: "SetTagValue",
			Handler:    _ControlService_SetTagValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cores/control_service.proto",
}
