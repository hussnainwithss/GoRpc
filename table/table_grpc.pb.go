// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: table/table.proto

package table

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TableClient is the client API for Table service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TableClient interface {
	Times(ctx context.Context, in *Request, opts ...grpc.CallOption) (Table_TimesClient, error)
}

type tableClient struct {
	cc grpc.ClientConnInterface
}

func NewTableClient(cc grpc.ClientConnInterface) TableClient {
	return &tableClient{cc}
}

func (c *tableClient) Times(ctx context.Context, in *Request, opts ...grpc.CallOption) (Table_TimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Table_ServiceDesc.Streams[0], "/table.Table/times", opts...)
	if err != nil {
		return nil, err
	}
	x := &tableTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Table_TimesClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type tableTimesClient struct {
	grpc.ClientStream
}

func (x *tableTimesClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TableServer is the server API for Table service.
// All implementations must embed UnimplementedTableServer
// for forward compatibility
type TableServer interface {
	Times(*Request, Table_TimesServer) error
	mustEmbedUnimplementedTableServer()
}

// UnimplementedTableServer must be embedded to have forward compatible implementations.
type UnimplementedTableServer struct {
}

func (UnimplementedTableServer) Times(*Request, Table_TimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Times not implemented")
}
func (UnimplementedTableServer) mustEmbedUnimplementedTableServer() {}

// UnsafeTableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TableServer will
// result in compilation errors.
type UnsafeTableServer interface {
	mustEmbedUnimplementedTableServer()
}

func RegisterTableServer(s grpc.ServiceRegistrar, srv TableServer) {
	s.RegisterService(&Table_ServiceDesc, srv)
}

func _Table_Times_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TableServer).Times(m, &tableTimesServer{stream})
}

type Table_TimesServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type tableTimesServer struct {
	grpc.ServerStream
}

func (x *tableTimesServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

// Table_ServiceDesc is the grpc.ServiceDesc for Table service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Table_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "table.Table",
	HandlerType: (*TableServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "times",
			Handler:       _Table_Times_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "table/table.proto",
}
