// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: sum/sum.proto

package sum

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

// SumClient is the client API for Sum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SumClient interface {
	SumNumbers(ctx context.Context, opts ...grpc.CallOption) (Sum_SumNumbersClient, error)
	ContinuousSum(ctx context.Context, opts ...grpc.CallOption) (Sum_ContinuousSumClient, error)
}

type sumClient struct {
	cc grpc.ClientConnInterface
}

func NewSumClient(cc grpc.ClientConnInterface) SumClient {
	return &sumClient{cc}
}

func (c *sumClient) SumNumbers(ctx context.Context, opts ...grpc.CallOption) (Sum_SumNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sum_ServiceDesc.Streams[0], "/sum.Sum/sumNumbers", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumSumNumbersClient{stream}
	return x, nil
}

type Sum_SumNumbersClient interface {
	Send(*Number) error
	CloseAndRecv() (*FinalSum, error)
	grpc.ClientStream
}

type sumSumNumbersClient struct {
	grpc.ClientStream
}

func (x *sumSumNumbersClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumSumNumbersClient) CloseAndRecv() (*FinalSum, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FinalSum)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumClient) ContinuousSum(ctx context.Context, opts ...grpc.CallOption) (Sum_ContinuousSumClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sum_ServiceDesc.Streams[1], "/sum.Sum/continuousSum", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumContinuousSumClient{stream}
	return x, nil
}

type Sum_ContinuousSumClient interface {
	Send(*Number) error
	Recv() (*FinalSum, error)
	grpc.ClientStream
}

type sumContinuousSumClient struct {
	grpc.ClientStream
}

func (x *sumContinuousSumClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumContinuousSumClient) Recv() (*FinalSum, error) {
	m := new(FinalSum)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServer is the server API for Sum service.
// All implementations must embed UnimplementedSumServer
// for forward compatibility
type SumServer interface {
	SumNumbers(Sum_SumNumbersServer) error
	ContinuousSum(Sum_ContinuousSumServer) error
	mustEmbedUnimplementedSumServer()
}

// UnimplementedSumServer must be embedded to have forward compatible implementations.
type UnimplementedSumServer struct {
}

func (UnimplementedSumServer) SumNumbers(Sum_SumNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method SumNumbers not implemented")
}
func (UnimplementedSumServer) ContinuousSum(Sum_ContinuousSumServer) error {
	return status.Errorf(codes.Unimplemented, "method ContinuousSum not implemented")
}
func (UnimplementedSumServer) mustEmbedUnimplementedSumServer() {}

// UnsafeSumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SumServer will
// result in compilation errors.
type UnsafeSumServer interface {
	mustEmbedUnimplementedSumServer()
}

func RegisterSumServer(s grpc.ServiceRegistrar, srv SumServer) {
	s.RegisterService(&Sum_ServiceDesc, srv)
}

func _Sum_SumNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServer).SumNumbers(&sumSumNumbersServer{stream})
}

type Sum_SumNumbersServer interface {
	SendAndClose(*FinalSum) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type sumSumNumbersServer struct {
	grpc.ServerStream
}

func (x *sumSumNumbersServer) SendAndClose(m *FinalSum) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumSumNumbersServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Sum_ContinuousSum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServer).ContinuousSum(&sumContinuousSumServer{stream})
}

type Sum_ContinuousSumServer interface {
	Send(*FinalSum) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type sumContinuousSumServer struct {
	grpc.ServerStream
}

func (x *sumContinuousSumServer) Send(m *FinalSum) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumContinuousSumServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Sum_ServiceDesc is the grpc.ServiceDesc for Sum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sum.Sum",
	HandlerType: (*SumServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sumNumbers",
			Handler:       _Sum_SumNumbers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "continuousSum",
			Handler:       _Sum_ContinuousSum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sum/sum.proto",
}
