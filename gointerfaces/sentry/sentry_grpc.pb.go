// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sentry

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SentryClient is the client API for Sentry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SentryClient interface {
	PenalizePeer(ctx context.Context, in *PenalizePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PeerMinBlock(ctx context.Context, in *PeerMinBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendMessageByMinBlock(ctx context.Context, in *SendMessageByMinBlockRequest, opts ...grpc.CallOption) (*SentPeers, error)
	SendMessageById(ctx context.Context, in *SendMessageByIdRequest, opts ...grpc.CallOption) (*SentPeers, error)
	SendMessageToRandomPeers(ctx context.Context, in *SendMessageToRandomPeersRequest, opts ...grpc.CallOption) (*SentPeers, error)
	SendMessageToAll(ctx context.Context, in *OutboundMessageData, opts ...grpc.CallOption) (*SentPeers, error)
	SetStatus(ctx context.Context, in *StatusData, opts ...grpc.CallOption) (*SetStatusReply, error)
	// Subscribe to receive messages.
	// Calling multiple times with a different set of ids starts separate streams.
	// It is possible to subscribe to the same set if ids more than once.
	Messages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (Sentry_MessagesClient, error)
	PeerCount(ctx context.Context, in *PeerCountRequest, opts ...grpc.CallOption) (*PeerCountReply, error)
	// Notifications about connected (after sub-protocol handshake) or lost peer
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (Sentry_PeersClient, error)
}

type sentryClient struct {
	cc grpc.ClientConnInterface
}

func NewSentryClient(cc grpc.ClientConnInterface) SentryClient {
	return &sentryClient{cc}
}

func (c *sentryClient) PenalizePeer(ctx context.Context, in *PenalizePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/PenalizePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) PeerMinBlock(ctx context.Context, in *PeerMinBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/PeerMinBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageByMinBlock(ctx context.Context, in *SendMessageByMinBlockRequest, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageByMinBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageById(ctx context.Context, in *SendMessageByIdRequest, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageToRandomPeers(ctx context.Context, in *SendMessageToRandomPeersRequest, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageToRandomPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SendMessageToAll(ctx context.Context, in *OutboundMessageData, opts ...grpc.CallOption) (*SentPeers, error) {
	out := new(SentPeers)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SendMessageToAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) SetStatus(ctx context.Context, in *StatusData, opts ...grpc.CallOption) (*SetStatusReply, error) {
	out := new(SetStatusReply)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) Messages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (Sentry_MessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sentry_ServiceDesc.Streams[0], "/sentry.Sentry/Messages", opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sentry_MessagesClient interface {
	Recv() (*InboundMessage, error)
	grpc.ClientStream
}

type sentryMessagesClient struct {
	grpc.ClientStream
}

func (x *sentryMessagesClient) Recv() (*InboundMessage, error) {
	m := new(InboundMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sentryClient) PeerCount(ctx context.Context, in *PeerCountRequest, opts ...grpc.CallOption) (*PeerCountReply, error) {
	out := new(PeerCountReply)
	err := c.cc.Invoke(ctx, "/sentry.Sentry/PeerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentryClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (Sentry_PeersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Sentry_ServiceDesc.Streams[1], "/sentry.Sentry/Peers", opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryPeersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sentry_PeersClient interface {
	Recv() (*PeersReply, error)
	grpc.ClientStream
}

type sentryPeersClient struct {
	grpc.ClientStream
}

func (x *sentryPeersClient) Recv() (*PeersReply, error) {
	m := new(PeersReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SentryServer is the server API for Sentry service.
// All implementations must embed UnimplementedSentryServer
// for forward compatibility
type SentryServer interface {
	PenalizePeer(context.Context, *PenalizePeerRequest) (*emptypb.Empty, error)
	PeerMinBlock(context.Context, *PeerMinBlockRequest) (*emptypb.Empty, error)
	SendMessageByMinBlock(context.Context, *SendMessageByMinBlockRequest) (*SentPeers, error)
	SendMessageById(context.Context, *SendMessageByIdRequest) (*SentPeers, error)
	SendMessageToRandomPeers(context.Context, *SendMessageToRandomPeersRequest) (*SentPeers, error)
	SendMessageToAll(context.Context, *OutboundMessageData) (*SentPeers, error)
	SetStatus(context.Context, *StatusData) (*SetStatusReply, error)
	// Subscribe to receive messages.
	// Calling multiple times with a different set of ids starts separate streams.
	// It is possible to subscribe to the same set if ids more than once.
	Messages(*MessagesRequest, Sentry_MessagesServer) error
	PeerCount(context.Context, *PeerCountRequest) (*PeerCountReply, error)
	// Notifications about connected (after sub-protocol handshake) or lost peer
	Peers(*PeersRequest, Sentry_PeersServer) error
	mustEmbedUnimplementedSentryServer()
}

// UnimplementedSentryServer must be embedded to have forward compatible implementations.
type UnimplementedSentryServer struct {
}

func (UnimplementedSentryServer) PenalizePeer(context.Context, *PenalizePeerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PenalizePeer not implemented")
}
func (UnimplementedSentryServer) PeerMinBlock(context.Context, *PeerMinBlockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeerMinBlock not implemented")
}
func (UnimplementedSentryServer) SendMessageByMinBlock(context.Context, *SendMessageByMinBlockRequest) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageByMinBlock not implemented")
}
func (UnimplementedSentryServer) SendMessageById(context.Context, *SendMessageByIdRequest) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageById not implemented")
}
func (UnimplementedSentryServer) SendMessageToRandomPeers(context.Context, *SendMessageToRandomPeersRequest) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToRandomPeers not implemented")
}
func (UnimplementedSentryServer) SendMessageToAll(context.Context, *OutboundMessageData) (*SentPeers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToAll not implemented")
}
func (UnimplementedSentryServer) SetStatus(context.Context, *StatusData) (*SetStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (UnimplementedSentryServer) Messages(*MessagesRequest, Sentry_MessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method Messages not implemented")
}
func (UnimplementedSentryServer) PeerCount(context.Context, *PeerCountRequest) (*PeerCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeerCount not implemented")
}
func (UnimplementedSentryServer) Peers(*PeersRequest, Sentry_PeersServer) error {
	return status.Errorf(codes.Unimplemented, "method Peers not implemented")
}
func (UnimplementedSentryServer) mustEmbedUnimplementedSentryServer() {}

// UnsafeSentryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SentryServer will
// result in compilation errors.
type UnsafeSentryServer interface {
	mustEmbedUnimplementedSentryServer()
}

func RegisterSentryServer(s grpc.ServiceRegistrar, srv SentryServer) {
	s.RegisterService(&Sentry_ServiceDesc, srv)
}

func _Sentry_PenalizePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PenalizePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).PenalizePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/PenalizePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).PenalizePeer(ctx, req.(*PenalizePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_PeerMinBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerMinBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).PeerMinBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/PeerMinBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).PeerMinBlock(ctx, req.(*PeerMinBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageByMinBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageByMinBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageByMinBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageByMinBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageByMinBlock(ctx, req.(*SendMessageByMinBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageById(ctx, req.(*SendMessageByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageToRandomPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToRandomPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageToRandomPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageToRandomPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageToRandomPeers(ctx, req.(*SendMessageToRandomPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SendMessageToAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutboundMessageData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SendMessageToAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SendMessageToAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SendMessageToAll(ctx, req.(*OutboundMessageData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).SetStatus(ctx, req.(*StatusData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_Messages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryServer).Messages(m, &sentryMessagesServer{stream})
}

type Sentry_MessagesServer interface {
	Send(*InboundMessage) error
	grpc.ServerStream
}

type sentryMessagesServer struct {
	grpc.ServerStream
}

func (x *sentryMessagesServer) Send(m *InboundMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Sentry_PeerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentryServer).PeerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentry.Sentry/PeerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentryServer).PeerCount(ctx, req.(*PeerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentry_Peers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PeersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryServer).Peers(m, &sentryPeersServer{stream})
}

type Sentry_PeersServer interface {
	Send(*PeersReply) error
	grpc.ServerStream
}

type sentryPeersServer struct {
	grpc.ServerStream
}

func (x *sentryPeersServer) Send(m *PeersReply) error {
	return x.ServerStream.SendMsg(m)
}

// Sentry_ServiceDesc is the grpc.ServiceDesc for Sentry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sentry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sentry.Sentry",
	HandlerType: (*SentryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PenalizePeer",
			Handler:    _Sentry_PenalizePeer_Handler,
		},
		{
			MethodName: "PeerMinBlock",
			Handler:    _Sentry_PeerMinBlock_Handler,
		},
		{
			MethodName: "SendMessageByMinBlock",
			Handler:    _Sentry_SendMessageByMinBlock_Handler,
		},
		{
			MethodName: "SendMessageById",
			Handler:    _Sentry_SendMessageById_Handler,
		},
		{
			MethodName: "SendMessageToRandomPeers",
			Handler:    _Sentry_SendMessageToRandomPeers_Handler,
		},
		{
			MethodName: "SendMessageToAll",
			Handler:    _Sentry_SendMessageToAll_Handler,
		},
		{
			MethodName: "SetStatus",
			Handler:    _Sentry_SetStatus_Handler,
		},
		{
			MethodName: "PeerCount",
			Handler:    _Sentry_PeerCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Messages",
			Handler:       _Sentry_Messages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Peers",
			Handler:       _Sentry_Peers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "p2psentry/sentry.proto",
}
