// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pds

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

// ToolGuideClient is the client API for ToolGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToolGuideClient interface {
	// Check up service health.
	Ping(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error)
}

type toolGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewToolGuideClient(cc grpc.ClientConnInterface) ToolGuideClient {
	return &toolGuideClient{cc}
}

func (c *toolGuideClient) Ping(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/pds.ToolGuide/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToolGuideServer is the server API for ToolGuide service.
// All implementations must embed UnimplementedToolGuideServer
// for forward compatibility
type ToolGuideServer interface {
	// Check up service health.
	Ping(context.Context, *Content) (*Content, error)
	mustEmbedUnimplementedToolGuideServer()
}

// UnimplementedToolGuideServer must be embedded to have forward compatible implementations.
type UnimplementedToolGuideServer struct {
}

func (UnimplementedToolGuideServer) Ping(context.Context, *Content) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedToolGuideServer) mustEmbedUnimplementedToolGuideServer() {}

// UnsafeToolGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToolGuideServer will
// result in compilation errors.
type UnsafeToolGuideServer interface {
	mustEmbedUnimplementedToolGuideServer()
}

func RegisterToolGuideServer(s grpc.ServiceRegistrar, srv ToolGuideServer) {
	s.RegisterService(&ToolGuide_ServiceDesc, srv)
}

func _ToolGuide_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolGuideServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pds.ToolGuide/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolGuideServer).Ping(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

// ToolGuide_ServiceDesc is the grpc.ServiceDesc for ToolGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToolGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pds.ToolGuide",
	HandlerType: (*ToolGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ToolGuide_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pds.proto",
}

// PortGuideClient is the client API for PortGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortGuideClient interface {
	// Accepts a stream of Ports and adds them to map.
	RecordList(ctx context.Context, opts ...grpc.CallOption) (PortGuide_RecordListClient, error)
	// Stores Port to map and return associated key.
	SetByKey(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Key, error)
	// Returns Port by associated key.
	GetByKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Port, error)
	// Returns Port by associated name.
	GetByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Port, error)
	// Finds nearest Port to given coordinates.
	FindNearest(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Port, error)
	// Finds all ports in given circle.
	FindInCircle(ctx context.Context, in *Circle, opts ...grpc.CallOption) (*Ports, error)
	// Finds all ports each of which contains given text
	// in one of the fields: name, city, province, country.
	FindText(ctx context.Context, in *Quest, opts ...grpc.CallOption) (*Ports, error)
}

type portGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewPortGuideClient(cc grpc.ClientConnInterface) PortGuideClient {
	return &portGuideClient{cc}
}

func (c *portGuideClient) RecordList(ctx context.Context, opts ...grpc.CallOption) (PortGuide_RecordListClient, error) {
	stream, err := c.cc.NewStream(ctx, &PortGuide_ServiceDesc.Streams[0], "/pds.PortGuide/RecordList", opts...)
	if err != nil {
		return nil, err
	}
	x := &portGuideRecordListClient{stream}
	return x, nil
}

type PortGuide_RecordListClient interface {
	Send(*Port) error
	CloseAndRecv() (*Summary, error)
	grpc.ClientStream
}

type portGuideRecordListClient struct {
	grpc.ClientStream
}

func (x *portGuideRecordListClient) Send(m *Port) error {
	return x.ClientStream.SendMsg(m)
}

func (x *portGuideRecordListClient) CloseAndRecv() (*Summary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Summary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *portGuideClient) SetByKey(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := c.cc.Invoke(ctx, "/pds.PortGuide/SetByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portGuideClient) GetByKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/pds.PortGuide/GetByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portGuideClient) GetByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/pds.PortGuide/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portGuideClient) FindNearest(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/pds.PortGuide/FindNearest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portGuideClient) FindInCircle(ctx context.Context, in *Circle, opts ...grpc.CallOption) (*Ports, error) {
	out := new(Ports)
	err := c.cc.Invoke(ctx, "/pds.PortGuide/FindInCircle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portGuideClient) FindText(ctx context.Context, in *Quest, opts ...grpc.CallOption) (*Ports, error) {
	out := new(Ports)
	err := c.cc.Invoke(ctx, "/pds.PortGuide/FindText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortGuideServer is the server API for PortGuide service.
// All implementations must embed UnimplementedPortGuideServer
// for forward compatibility
type PortGuideServer interface {
	// Accepts a stream of Ports and adds them to map.
	RecordList(PortGuide_RecordListServer) error
	// Stores Port to map and return associated key.
	SetByKey(context.Context, *Port) (*Key, error)
	// Returns Port by associated key.
	GetByKey(context.Context, *Key) (*Port, error)
	// Returns Port by associated name.
	GetByName(context.Context, *Name) (*Port, error)
	// Finds nearest Port to given coordinates.
	FindNearest(context.Context, *Point) (*Port, error)
	// Finds all ports in given circle.
	FindInCircle(context.Context, *Circle) (*Ports, error)
	// Finds all ports each of which contains given text
	// in one of the fields: name, city, province, country.
	FindText(context.Context, *Quest) (*Ports, error)
	mustEmbedUnimplementedPortGuideServer()
}

// UnimplementedPortGuideServer must be embedded to have forward compatible implementations.
type UnimplementedPortGuideServer struct {
}

func (UnimplementedPortGuideServer) RecordList(PortGuide_RecordListServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordList not implemented")
}
func (UnimplementedPortGuideServer) SetByKey(context.Context, *Port) (*Key, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetByKey not implemented")
}
func (UnimplementedPortGuideServer) GetByKey(context.Context, *Key) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByKey not implemented")
}
func (UnimplementedPortGuideServer) GetByName(context.Context, *Name) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedPortGuideServer) FindNearest(context.Context, *Point) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNearest not implemented")
}
func (UnimplementedPortGuideServer) FindInCircle(context.Context, *Circle) (*Ports, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindInCircle not implemented")
}
func (UnimplementedPortGuideServer) FindText(context.Context, *Quest) (*Ports, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindText not implemented")
}
func (UnimplementedPortGuideServer) mustEmbedUnimplementedPortGuideServer() {}

// UnsafePortGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortGuideServer will
// result in compilation errors.
type UnsafePortGuideServer interface {
	mustEmbedUnimplementedPortGuideServer()
}

func RegisterPortGuideServer(s grpc.ServiceRegistrar, srv PortGuideServer) {
	s.RegisterService(&PortGuide_ServiceDesc, srv)
}

func _PortGuide_RecordList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PortGuideServer).RecordList(&portGuideRecordListServer{stream})
}

type PortGuide_RecordListServer interface {
	SendAndClose(*Summary) error
	Recv() (*Port, error)
	grpc.ServerStream
}

type portGuideRecordListServer struct {
	grpc.ServerStream
}

func (x *portGuideRecordListServer) SendAndClose(m *Summary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *portGuideRecordListServer) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PortGuide_SetByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortGuideServer).SetByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pds.PortGuide/SetByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortGuideServer).SetByKey(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortGuide_GetByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortGuideServer).GetByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pds.PortGuide/GetByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortGuideServer).GetByKey(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortGuide_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortGuideServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pds.PortGuide/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortGuideServer).GetByName(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortGuide_FindNearest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortGuideServer).FindNearest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pds.PortGuide/FindNearest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortGuideServer).FindNearest(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortGuide_FindInCircle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Circle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortGuideServer).FindInCircle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pds.PortGuide/FindInCircle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortGuideServer).FindInCircle(ctx, req.(*Circle))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortGuide_FindText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Quest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortGuideServer).FindText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pds.PortGuide/FindText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortGuideServer).FindText(ctx, req.(*Quest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortGuide_ServiceDesc is the grpc.ServiceDesc for PortGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pds.PortGuide",
	HandlerType: (*PortGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetByKey",
			Handler:    _PortGuide_SetByKey_Handler,
		},
		{
			MethodName: "GetByKey",
			Handler:    _PortGuide_GetByKey_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _PortGuide_GetByName_Handler,
		},
		{
			MethodName: "FindNearest",
			Handler:    _PortGuide_FindNearest_Handler,
		},
		{
			MethodName: "FindInCircle",
			Handler:    _PortGuide_FindInCircle_Handler,
		},
		{
			MethodName: "FindText",
			Handler:    _PortGuide_FindText_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordList",
			Handler:       _PortGuide_RecordList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pds.proto",
}
