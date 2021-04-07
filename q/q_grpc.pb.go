// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package q

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

// QueueClient is the client API for Queue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueClient interface {
	Create(ctx context.Context, in *QueueDetails, opts ...grpc.CallOption) (*CreateResponse, error)
	Enqueue(ctx context.Context, in *Message, opts ...grpc.CallOption) (*WriteResult, error)
	Dequeue(ctx context.Context, in *QueueDetails, opts ...grpc.CallOption) (*Message, error)
	Events(ctx context.Context, in *QueueDetails, opts ...grpc.CallOption) (Queue_EventsClient, error)
}

type queueClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueClient(cc grpc.ClientConnInterface) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) Create(ctx context.Context, in *QueueDetails, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/q.Queue/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Enqueue(ctx context.Context, in *Message, opts ...grpc.CallOption) (*WriteResult, error) {
	out := new(WriteResult)
	err := c.cc.Invoke(ctx, "/q.Queue/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Dequeue(ctx context.Context, in *QueueDetails, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/q.Queue/Dequeue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Events(ctx context.Context, in *QueueDetails, opts ...grpc.CallOption) (Queue_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Queue_ServiceDesc.Streams[0], "/q.Queue/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &queueEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Queue_EventsClient interface {
	Recv() (*QueueEvent, error)
	grpc.ClientStream
}

type queueEventsClient struct {
	grpc.ClientStream
}

func (x *queueEventsClient) Recv() (*QueueEvent, error) {
	m := new(QueueEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueueServer is the server API for Queue service.
// All implementations must embed UnimplementedQueueServer
// for forward compatibility
type QueueServer interface {
	Create(context.Context, *QueueDetails) (*CreateResponse, error)
	Enqueue(context.Context, *Message) (*WriteResult, error)
	Dequeue(context.Context, *QueueDetails) (*Message, error)
	Events(*QueueDetails, Queue_EventsServer) error
	mustEmbedUnimplementedQueueServer()
}

// UnimplementedQueueServer must be embedded to have forward compatible implementations.
type UnimplementedQueueServer struct {
}

func (UnimplementedQueueServer) Create(context.Context, *QueueDetails) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedQueueServer) Enqueue(context.Context, *Message) (*WriteResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (UnimplementedQueueServer) Dequeue(context.Context, *QueueDetails) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dequeue not implemented")
}
func (UnimplementedQueueServer) Events(*QueueDetails, Queue_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}
func (UnimplementedQueueServer) mustEmbedUnimplementedQueueServer() {}

// UnsafeQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServer will
// result in compilation errors.
type UnsafeQueueServer interface {
	mustEmbedUnimplementedQueueServer()
}

func RegisterQueueServer(s grpc.ServiceRegistrar, srv QueueServer) {
	s.RegisterService(&Queue_ServiceDesc, srv)
}

func _Queue_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/q.Queue/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Create(ctx, req.(*QueueDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/q.Queue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Enqueue(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Dequeue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Dequeue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/q.Queue/Dequeue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Dequeue(ctx, req.(*QueueDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueueDetails)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueueServer).Events(m, &queueEventsServer{stream})
}

type Queue_EventsServer interface {
	Send(*QueueEvent) error
	grpc.ServerStream
}

type queueEventsServer struct {
	grpc.ServerStream
}

func (x *queueEventsServer) Send(m *QueueEvent) error {
	return x.ServerStream.SendMsg(m)
}

// Queue_ServiceDesc is the grpc.ServiceDesc for Queue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Queue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "q.Queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Queue_Create_Handler,
		},
		{
			MethodName: "Enqueue",
			Handler:    _Queue_Enqueue_Handler,
		},
		{
			MethodName: "Dequeue",
			Handler:    _Queue_Dequeue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _Queue_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "q/q.proto",
}