// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: edges/data_service.proto

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
	DataService_Upload_FullMethodName   = "/edges.DataService/Upload"
	DataService_Compile_FullMethodName  = "/edges.DataService/Compile"
	DataService_Query_FullMethodName    = "/edges.DataService/Query"
	DataService_QueryTag_FullMethodName = "/edges.DataService/QueryTag"
)

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	Upload(ctx context.Context, in *DataUploadRequest, opts ...grpc.CallOption) (*DataUploadResponse, error)
	Compile(ctx context.Context, in *DataQueryRequest, opts ...grpc.CallOption) (*pb.Message, error)
	Query(ctx context.Context, in *DataQueryRequest, opts ...grpc.CallOption) (DataService_QueryClient, error)
	QueryTag(ctx context.Context, in *DataQueryTagRequest, opts ...grpc.CallOption) (DataService_QueryTagClient, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) Upload(ctx context.Context, in *DataUploadRequest, opts ...grpc.CallOption) (*DataUploadResponse, error) {
	out := new(DataUploadResponse)
	err := c.cc.Invoke(ctx, DataService_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) Compile(ctx context.Context, in *DataQueryRequest, opts ...grpc.CallOption) (*pb.Message, error) {
	out := new(pb.Message)
	err := c.cc.Invoke(ctx, DataService_Compile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) Query(ctx context.Context, in *DataQueryRequest, opts ...grpc.CallOption) (DataService_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[0], DataService_Query_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_QueryClient interface {
	Recv() (*pb.Message, error)
	grpc.ClientStream
}

type dataServiceQueryClient struct {
	grpc.ClientStream
}

func (x *dataServiceQueryClient) Recv() (*pb.Message, error) {
	m := new(pb.Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) QueryTag(ctx context.Context, in *DataQueryTagRequest, opts ...grpc.CallOption) (DataService_QueryTagClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[1], DataService_QueryTag_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceQueryTagClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_QueryTagClient interface {
	Recv() (*pb.Message, error)
	grpc.ClientStream
}

type dataServiceQueryTagClient struct {
	grpc.ClientStream
}

func (x *dataServiceQueryTagClient) Recv() (*pb.Message, error) {
	m := new(pb.Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	Upload(context.Context, *DataUploadRequest) (*DataUploadResponse, error)
	Compile(context.Context, *DataQueryRequest) (*pb.Message, error)
	Query(*DataQueryRequest, DataService_QueryServer) error
	QueryTag(*DataQueryTagRequest, DataService_QueryTagServer) error
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) Upload(context.Context, *DataUploadRequest) (*DataUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedDataServiceServer) Compile(context.Context, *DataQueryRequest) (*pb.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}
func (UnimplementedDataServiceServer) Query(*DataQueryRequest, DataService_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDataServiceServer) QueryTag(*DataQueryTagRequest, DataService_QueryTagServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryTag not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).Upload(ctx, req.(*DataUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_Compile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).Compile(ctx, req.(*DataQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).Query(m, &dataServiceQueryServer{stream})
}

type DataService_QueryServer interface {
	Send(*pb.Message) error
	grpc.ServerStream
}

type dataServiceQueryServer struct {
	grpc.ServerStream
}

func (x *dataServiceQueryServer) Send(m *pb.Message) error {
	return x.ServerStream.SendMsg(m)
}

func _DataService_QueryTag_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataQueryTagRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).QueryTag(m, &dataServiceQueryTagServer{stream})
}

type DataService_QueryTagServer interface {
	Send(*pb.Message) error
	grpc.ServerStream
}

type dataServiceQueryTagServer struct {
	grpc.ServerStream
}

func (x *dataServiceQueryTagServer) Send(m *pb.Message) error {
	return x.ServerStream.SendMsg(m)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edges.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _DataService_Upload_Handler,
		},
		{
			MethodName: "Compile",
			Handler:    _DataService_Compile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _DataService_Query_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryTag",
			Handler:       _DataService_QueryTag_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "edges/data_service.proto",
}
