// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api.proto

package grpcapi

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

const (
	Hub_GetCodeForPrint_FullMethodName  = "/Hub/GetCodeForPrint"
	Hub_ProducePrinted_FullMethodName   = "/Hub/ProducePrinted"
	Hub_DiscardBySerial_FullMethodName  = "/Hub/DiscardBySerial"
	Hub_DiscardByPrintId_FullMethodName = "/Hub/DiscardByPrintId"
	Hub_GetProducedCount_FullMethodName = "/Hub/GetProducedCount"
	Hub_AddCodeForPrint_FullMethodName  = "/Hub/AddCodeForPrint"
	Hub_AddGood_FullMethodName          = "/Hub/AddGood"
	Hub_GetGoodsCodeReq_FullMethodName  = "/Hub/GetGoodsCodeReq"
	Hub_GetCodeForUpload_FullMethodName = "/Hub/GetCodeForUpload"
	Hub_SetCodeUploaded_FullMethodName  = "/Hub/SetCodeUploaded"
)

// HubClient is the client API for Hub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubClient interface {
	// Функции для терминала
	GetCodeForPrint(ctx context.Context, in *GetCodeForPrintReq, opts ...grpc.CallOption) (*GetCodeForPrintResp, error)
	ProducePrinted(ctx context.Context, in *ProducePrintedReq, opts ...grpc.CallOption) (*Empty, error)
	DiscardBySerial(ctx context.Context, in *DiscardBySerialReq, opts ...grpc.CallOption) (*Empty, error)
	DiscardByPrintId(ctx context.Context, in *DiscardByPrintIdReq, opts ...grpc.CallOption) (*DiscardByPrintIdResp, error)
	GetProducedCount(ctx context.Context, in *GetProducedCountReq, opts ...grpc.CallOption) (*GetProducedCountResp, error)
	// Функции загрузки выгрузки кодов в базу
	AddCodeForPrint(ctx context.Context, in *AddCodeForPrintReq, opts ...grpc.CallOption) (*Empty, error)
	// Админские функции
	AddGood(ctx context.Context, in *AddGoodReq, opts ...grpc.CallOption) (*Empty, error)
	GetGoodsCodeReq(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetGoodsCodeResp, error)
	GetCodeForUpload(ctx context.Context, in *GetCodeForUploadReq, opts ...grpc.CallOption) (*GetCodeForUploadResp, error)
	SetCodeUploaded(ctx context.Context, in *SetCodeUploadedReq, opts ...grpc.CallOption) (*Empty, error)
}

type hubClient struct {
	cc grpc.ClientConnInterface
}

func NewHubClient(cc grpc.ClientConnInterface) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) GetCodeForPrint(ctx context.Context, in *GetCodeForPrintReq, opts ...grpc.CallOption) (*GetCodeForPrintResp, error) {
	out := new(GetCodeForPrintResp)
	err := c.cc.Invoke(ctx, Hub_GetCodeForPrint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) ProducePrinted(ctx context.Context, in *ProducePrintedReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Hub_ProducePrinted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) DiscardBySerial(ctx context.Context, in *DiscardBySerialReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Hub_DiscardBySerial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) DiscardByPrintId(ctx context.Context, in *DiscardByPrintIdReq, opts ...grpc.CallOption) (*DiscardByPrintIdResp, error) {
	out := new(DiscardByPrintIdResp)
	err := c.cc.Invoke(ctx, Hub_DiscardByPrintId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) GetProducedCount(ctx context.Context, in *GetProducedCountReq, opts ...grpc.CallOption) (*GetProducedCountResp, error) {
	out := new(GetProducedCountResp)
	err := c.cc.Invoke(ctx, Hub_GetProducedCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) AddCodeForPrint(ctx context.Context, in *AddCodeForPrintReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Hub_AddCodeForPrint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) AddGood(ctx context.Context, in *AddGoodReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Hub_AddGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) GetGoodsCodeReq(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetGoodsCodeResp, error) {
	out := new(GetGoodsCodeResp)
	err := c.cc.Invoke(ctx, Hub_GetGoodsCodeReq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) GetCodeForUpload(ctx context.Context, in *GetCodeForUploadReq, opts ...grpc.CallOption) (*GetCodeForUploadResp, error) {
	out := new(GetCodeForUploadResp)
	err := c.cc.Invoke(ctx, Hub_GetCodeForUpload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) SetCodeUploaded(ctx context.Context, in *SetCodeUploadedReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Hub_SetCodeUploaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubServer is the server API for Hub service.
// All implementations must embed UnimplementedHubServer
// for forward compatibility
type HubServer interface {
	// Функции для терминала
	GetCodeForPrint(context.Context, *GetCodeForPrintReq) (*GetCodeForPrintResp, error)
	ProducePrinted(context.Context, *ProducePrintedReq) (*Empty, error)
	DiscardBySerial(context.Context, *DiscardBySerialReq) (*Empty, error)
	DiscardByPrintId(context.Context, *DiscardByPrintIdReq) (*DiscardByPrintIdResp, error)
	GetProducedCount(context.Context, *GetProducedCountReq) (*GetProducedCountResp, error)
	// Функции загрузки выгрузки кодов в базу
	AddCodeForPrint(context.Context, *AddCodeForPrintReq) (*Empty, error)
	// Админские функции
	AddGood(context.Context, *AddGoodReq) (*Empty, error)
	GetGoodsCodeReq(context.Context, *Empty) (*GetGoodsCodeResp, error)
	GetCodeForUpload(context.Context, *GetCodeForUploadReq) (*GetCodeForUploadResp, error)
	SetCodeUploaded(context.Context, *SetCodeUploadedReq) (*Empty, error)
	mustEmbedUnimplementedHubServer()
}

// UnimplementedHubServer must be embedded to have forward compatible implementations.
type UnimplementedHubServer struct {
}

func (UnimplementedHubServer) GetCodeForPrint(context.Context, *GetCodeForPrintReq) (*GetCodeForPrintResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodeForPrint not implemented")
}
func (UnimplementedHubServer) ProducePrinted(context.Context, *ProducePrintedReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProducePrinted not implemented")
}
func (UnimplementedHubServer) DiscardBySerial(context.Context, *DiscardBySerialReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscardBySerial not implemented")
}
func (UnimplementedHubServer) DiscardByPrintId(context.Context, *DiscardByPrintIdReq) (*DiscardByPrintIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscardByPrintId not implemented")
}
func (UnimplementedHubServer) GetProducedCount(context.Context, *GetProducedCountReq) (*GetProducedCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducedCount not implemented")
}
func (UnimplementedHubServer) AddCodeForPrint(context.Context, *AddCodeForPrintReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCodeForPrint not implemented")
}
func (UnimplementedHubServer) AddGood(context.Context, *AddGoodReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGood not implemented")
}
func (UnimplementedHubServer) GetGoodsCodeReq(context.Context, *Empty) (*GetGoodsCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsCodeReq not implemented")
}
func (UnimplementedHubServer) GetCodeForUpload(context.Context, *GetCodeForUploadReq) (*GetCodeForUploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodeForUpload not implemented")
}
func (UnimplementedHubServer) SetCodeUploaded(context.Context, *SetCodeUploadedReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCodeUploaded not implemented")
}
func (UnimplementedHubServer) mustEmbedUnimplementedHubServer() {}

// UnsafeHubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServer will
// result in compilation errors.
type UnsafeHubServer interface {
	mustEmbedUnimplementedHubServer()
}

func RegisterHubServer(s grpc.ServiceRegistrar, srv HubServer) {
	s.RegisterService(&Hub_ServiceDesc, srv)
}

func _Hub_GetCodeForPrint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodeForPrintReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).GetCodeForPrint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_GetCodeForPrint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).GetCodeForPrint(ctx, req.(*GetCodeForPrintReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_ProducePrinted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProducePrintedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).ProducePrinted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_ProducePrinted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).ProducePrinted(ctx, req.(*ProducePrintedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_DiscardBySerial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscardBySerialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).DiscardBySerial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_DiscardBySerial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).DiscardBySerial(ctx, req.(*DiscardBySerialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_DiscardByPrintId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscardByPrintIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).DiscardByPrintId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_DiscardByPrintId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).DiscardByPrintId(ctx, req.(*DiscardByPrintIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_GetProducedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProducedCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).GetProducedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_GetProducedCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).GetProducedCount(ctx, req.(*GetProducedCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_AddCodeForPrint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCodeForPrintReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).AddCodeForPrint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_AddCodeForPrint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).AddCodeForPrint(ctx, req.(*AddCodeForPrintReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_AddGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).AddGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_AddGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).AddGood(ctx, req.(*AddGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_GetGoodsCodeReq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).GetGoodsCodeReq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_GetGoodsCodeReq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).GetGoodsCodeReq(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_GetCodeForUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodeForUploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).GetCodeForUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_GetCodeForUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).GetCodeForUpload(ctx, req.(*GetCodeForUploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_SetCodeUploaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCodeUploadedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).SetCodeUploaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hub_SetCodeUploaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).SetCodeUploaded(ctx, req.(*SetCodeUploadedReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Hub_ServiceDesc is the grpc.ServiceDesc for Hub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCodeForPrint",
			Handler:    _Hub_GetCodeForPrint_Handler,
		},
		{
			MethodName: "ProducePrinted",
			Handler:    _Hub_ProducePrinted_Handler,
		},
		{
			MethodName: "DiscardBySerial",
			Handler:    _Hub_DiscardBySerial_Handler,
		},
		{
			MethodName: "DiscardByPrintId",
			Handler:    _Hub_DiscardByPrintId_Handler,
		},
		{
			MethodName: "GetProducedCount",
			Handler:    _Hub_GetProducedCount_Handler,
		},
		{
			MethodName: "AddCodeForPrint",
			Handler:    _Hub_AddCodeForPrint_Handler,
		},
		{
			MethodName: "AddGood",
			Handler:    _Hub_AddGood_Handler,
		},
		{
			MethodName: "GetGoodsCodeReq",
			Handler:    _Hub_GetGoodsCodeReq_Handler,
		},
		{
			MethodName: "GetCodeForUpload",
			Handler:    _Hub_GetCodeForUpload_Handler,
		},
		{
			MethodName: "SetCodeUploaded",
			Handler:    _Hub_SetCodeUploaded_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
