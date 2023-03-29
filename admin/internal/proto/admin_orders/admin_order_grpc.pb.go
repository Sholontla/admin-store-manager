// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin_order.proto

package admin_orders

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

// RequestOrderStoreInventoryServiceClient is the client API for RequestOrderStoreInventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestOrderStoreInventoryServiceClient interface {
	StreamOrderStoreInventory(ctx context.Context, opts ...grpc.CallOption) (RequestOrderStoreInventoryService_StreamOrderStoreInventoryClient, error)
}

type requestOrderStoreInventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestOrderStoreInventoryServiceClient(cc grpc.ClientConnInterface) RequestOrderStoreInventoryServiceClient {
	return &requestOrderStoreInventoryServiceClient{cc}
}

func (c *requestOrderStoreInventoryServiceClient) StreamOrderStoreInventory(ctx context.Context, opts ...grpc.CallOption) (RequestOrderStoreInventoryService_StreamOrderStoreInventoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &RequestOrderStoreInventoryService_ServiceDesc.Streams[0], "/stores.RequestOrderStoreInventoryService/StreamOrderStoreInventory", opts...)
	if err != nil {
		return nil, err
	}
	x := &requestOrderStoreInventoryServiceStreamOrderStoreInventoryClient{stream}
	return x, nil
}

type RequestOrderStoreInventoryService_StreamOrderStoreInventoryClient interface {
	Send(*RequestOrderStoreInventory) error
	Recv() (*ResponseOrderStoreInventory, error)
	grpc.ClientStream
}

type requestOrderStoreInventoryServiceStreamOrderStoreInventoryClient struct {
	grpc.ClientStream
}

func (x *requestOrderStoreInventoryServiceStreamOrderStoreInventoryClient) Send(m *RequestOrderStoreInventory) error {
	return x.ClientStream.SendMsg(m)
}

func (x *requestOrderStoreInventoryServiceStreamOrderStoreInventoryClient) Recv() (*ResponseOrderStoreInventory, error) {
	m := new(ResponseOrderStoreInventory)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RequestOrderStoreInventoryServiceServer is the server API for RequestOrderStoreInventoryService service.
// All implementations must embed UnimplementedRequestOrderStoreInventoryServiceServer
// for forward compatibility
type RequestOrderStoreInventoryServiceServer interface {
	StreamOrderStoreInventory(RequestOrderStoreInventoryService_StreamOrderStoreInventoryServer) error
	mustEmbedUnimplementedRequestOrderStoreInventoryServiceServer()
}

// UnimplementedRequestOrderStoreInventoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRequestOrderStoreInventoryServiceServer struct {
}

func (UnimplementedRequestOrderStoreInventoryServiceServer) StreamOrderStoreInventory(RequestOrderStoreInventoryService_StreamOrderStoreInventoryServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamOrderStoreInventory not implemented")
}
func (UnimplementedRequestOrderStoreInventoryServiceServer) mustEmbedUnimplementedRequestOrderStoreInventoryServiceServer() {
}

// UnsafeRequestOrderStoreInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestOrderStoreInventoryServiceServer will
// result in compilation errors.
type UnsafeRequestOrderStoreInventoryServiceServer interface {
	mustEmbedUnimplementedRequestOrderStoreInventoryServiceServer()
}

func RegisterRequestOrderStoreInventoryServiceServer(s grpc.ServiceRegistrar, srv RequestOrderStoreInventoryServiceServer) {
	s.RegisterService(&RequestOrderStoreInventoryService_ServiceDesc, srv)
}

func _RequestOrderStoreInventoryService_StreamOrderStoreInventory_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RequestOrderStoreInventoryServiceServer).StreamOrderStoreInventory(&requestOrderStoreInventoryServiceStreamOrderStoreInventoryServer{stream})
}

type RequestOrderStoreInventoryService_StreamOrderStoreInventoryServer interface {
	Send(*ResponseOrderStoreInventory) error
	Recv() (*RequestOrderStoreInventory, error)
	grpc.ServerStream
}

type requestOrderStoreInventoryServiceStreamOrderStoreInventoryServer struct {
	grpc.ServerStream
}

func (x *requestOrderStoreInventoryServiceStreamOrderStoreInventoryServer) Send(m *ResponseOrderStoreInventory) error {
	return x.ServerStream.SendMsg(m)
}

func (x *requestOrderStoreInventoryServiceStreamOrderStoreInventoryServer) Recv() (*RequestOrderStoreInventory, error) {
	m := new(RequestOrderStoreInventory)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RequestOrderStoreInventoryService_ServiceDesc is the grpc.ServiceDesc for RequestOrderStoreInventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestOrderStoreInventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stores.RequestOrderStoreInventoryService",
	HandlerType: (*RequestOrderStoreInventoryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamOrderStoreInventory",
			Handler:       _RequestOrderStoreInventoryService_StreamOrderStoreInventory_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "admin_order.proto",
}
