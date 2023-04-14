// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: cores/route_service.proto

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
	RouteService_Create_FullMethodName     = "/cores.RouteService/Create"
	RouteService_Update_FullMethodName     = "/cores.RouteService/Update"
	RouteService_View_FullMethodName       = "/cores.RouteService/View"
	RouteService_ViewByName_FullMethodName = "/cores.RouteService/ViewByName"
	RouteService_Delete_FullMethodName     = "/cores.RouteService/Delete"
	RouteService_List_FullMethodName       = "/cores.RouteService/List"
)

// RouteServiceClient is the client API for RouteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteServiceClient interface {
	Create(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Route, error)
	Update(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Route, error)
	View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*Route, error)
	ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*Route, error)
	Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error)
	List(ctx context.Context, in *ListRouteRequest, opts ...grpc.CallOption) (*ListRouteResponse, error)
}

type routeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteServiceClient(cc grpc.ClientConnInterface) RouteServiceClient {
	return &routeServiceClient{cc}
}

func (c *routeServiceClient) Create(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Route, error) {
	out := new(Route)
	err := c.cc.Invoke(ctx, RouteService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) Update(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Route, error) {
	out := new(Route)
	err := c.cc.Invoke(ctx, RouteService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) View(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*Route, error) {
	out := new(Route)
	err := c.cc.Invoke(ctx, RouteService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) ViewByName(ctx context.Context, in *pb.Name, opts ...grpc.CallOption) (*Route, error) {
	out := new(Route)
	err := c.cc.Invoke(ctx, RouteService_ViewByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) Delete(ctx context.Context, in *pb.Id, opts ...grpc.CallOption) (*pb.MyBool, error) {
	out := new(pb.MyBool)
	err := c.cc.Invoke(ctx, RouteService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) List(ctx context.Context, in *ListRouteRequest, opts ...grpc.CallOption) (*ListRouteResponse, error) {
	out := new(ListRouteResponse)
	err := c.cc.Invoke(ctx, RouteService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServiceServer is the server API for RouteService service.
// All implementations must embed UnimplementedRouteServiceServer
// for forward compatibility
type RouteServiceServer interface {
	Create(context.Context, *Route) (*Route, error)
	Update(context.Context, *Route) (*Route, error)
	View(context.Context, *pb.Id) (*Route, error)
	ViewByName(context.Context, *pb.Name) (*Route, error)
	Delete(context.Context, *pb.Id) (*pb.MyBool, error)
	List(context.Context, *ListRouteRequest) (*ListRouteResponse, error)
	mustEmbedUnimplementedRouteServiceServer()
}

// UnimplementedRouteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServiceServer struct {
}

func (UnimplementedRouteServiceServer) Create(context.Context, *Route) (*Route, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRouteServiceServer) Update(context.Context, *Route) (*Route, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRouteServiceServer) View(context.Context, *pb.Id) (*Route, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedRouteServiceServer) ViewByName(context.Context, *pb.Name) (*Route, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewByName not implemented")
}
func (UnimplementedRouteServiceServer) Delete(context.Context, *pb.Id) (*pb.MyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRouteServiceServer) List(context.Context, *ListRouteRequest) (*ListRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRouteServiceServer) mustEmbedUnimplementedRouteServiceServer() {}

// UnsafeRouteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteServiceServer will
// result in compilation errors.
type UnsafeRouteServiceServer interface {
	mustEmbedUnimplementedRouteServiceServer()
}

func RegisterRouteServiceServer(s grpc.ServiceRegistrar, srv RouteServiceServer) {
	s.RegisterService(&RouteService_ServiceDesc, srv)
}

func _RouteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).Create(ctx, req.(*Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).Update(ctx, req.(*Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).View(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_ViewByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).ViewByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_ViewByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).ViewByName(ctx, req.(*pb.Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).Delete(ctx, req.(*pb.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RouteService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).List(ctx, req.(*ListRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteService_ServiceDesc is the grpc.ServiceDesc for RouteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cores.RouteService",
	HandlerType: (*RouteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RouteService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RouteService_Update_Handler,
		},
		{
			MethodName: "View",
			Handler:    _RouteService_View_Handler,
		},
		{
			MethodName: "ViewByName",
			Handler:    _RouteService_ViewByName_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RouteService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RouteService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cores/route_service.proto",
}
