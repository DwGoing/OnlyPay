// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: proto/funds_service/treasury.proto

package treasury_generated

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

// TreasuryClient is the client API for Treasury service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TreasuryClient interface {
	// 创建充值订单
	CreateRechargeOrderRpc(ctx context.Context, in *CreateRechargeOrderRequest, opts ...grpc.CallOption) (*CreateRechargeOrderResponse, error)
	// 提交充值交易
	SubmitRechargeOrderTransactionRpc(ctx context.Context, in *SubmitRechargeOrderTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 取消充值订单
	CancelRechargeOrderRpc(ctx context.Context, in *CancelRechargeOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 手动检查订单状态
	CheckRechargeOrderStatusRpc(ctx context.Context, in *CheckRechargeOrderStatusRequest, opts ...grpc.CallOption) (*CheckRechargeOrderStatusResponse, error)
}

type treasuryClient struct {
	cc grpc.ClientConnInterface
}

func NewTreasuryClient(cc grpc.ClientConnInterface) TreasuryClient {
	return &treasuryClient{cc}
}

func (c *treasuryClient) CreateRechargeOrderRpc(ctx context.Context, in *CreateRechargeOrderRequest, opts ...grpc.CallOption) (*CreateRechargeOrderResponse, error) {
	out := new(CreateRechargeOrderResponse)
	err := c.cc.Invoke(ctx, "/Treasury/CreateRechargeOrderRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treasuryClient) SubmitRechargeOrderTransactionRpc(ctx context.Context, in *SubmitRechargeOrderTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Treasury/SubmitRechargeOrderTransactionRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treasuryClient) CancelRechargeOrderRpc(ctx context.Context, in *CancelRechargeOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Treasury/CancelRechargeOrderRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treasuryClient) CheckRechargeOrderStatusRpc(ctx context.Context, in *CheckRechargeOrderStatusRequest, opts ...grpc.CallOption) (*CheckRechargeOrderStatusResponse, error) {
	out := new(CheckRechargeOrderStatusResponse)
	err := c.cc.Invoke(ctx, "/Treasury/CheckRechargeOrderStatusRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TreasuryServer is the server API for Treasury service.
// All implementations must embed UnimplementedTreasuryServer
// for forward compatibility
type TreasuryServer interface {
	// 创建充值订单
	CreateRechargeOrderRpc(context.Context, *CreateRechargeOrderRequest) (*CreateRechargeOrderResponse, error)
	// 提交充值交易
	SubmitRechargeOrderTransactionRpc(context.Context, *SubmitRechargeOrderTransactionRequest) (*emptypb.Empty, error)
	// 取消充值订单
	CancelRechargeOrderRpc(context.Context, *CancelRechargeOrderRequest) (*emptypb.Empty, error)
	// 手动检查订单状态
	CheckRechargeOrderStatusRpc(context.Context, *CheckRechargeOrderStatusRequest) (*CheckRechargeOrderStatusResponse, error)
	mustEmbedUnimplementedTreasuryServer()
}

// UnimplementedTreasuryServer must be embedded to have forward compatible implementations.
type UnimplementedTreasuryServer struct {
}

func (UnimplementedTreasuryServer) CreateRechargeOrderRpc(context.Context, *CreateRechargeOrderRequest) (*CreateRechargeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRechargeOrderRpc not implemented")
}
func (UnimplementedTreasuryServer) SubmitRechargeOrderTransactionRpc(context.Context, *SubmitRechargeOrderTransactionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRechargeOrderTransactionRpc not implemented")
}
func (UnimplementedTreasuryServer) CancelRechargeOrderRpc(context.Context, *CancelRechargeOrderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRechargeOrderRpc not implemented")
}
func (UnimplementedTreasuryServer) CheckRechargeOrderStatusRpc(context.Context, *CheckRechargeOrderStatusRequest) (*CheckRechargeOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRechargeOrderStatusRpc not implemented")
}
func (UnimplementedTreasuryServer) mustEmbedUnimplementedTreasuryServer() {}

// UnsafeTreasuryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TreasuryServer will
// result in compilation errors.
type UnsafeTreasuryServer interface {
	mustEmbedUnimplementedTreasuryServer()
}

func RegisterTreasuryServer(s grpc.ServiceRegistrar, srv TreasuryServer) {
	s.RegisterService(&Treasury_ServiceDesc, srv)
}

func _Treasury_CreateRechargeOrderRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRechargeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreasuryServer).CreateRechargeOrderRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Treasury/CreateRechargeOrderRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreasuryServer).CreateRechargeOrderRpc(ctx, req.(*CreateRechargeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Treasury_SubmitRechargeOrderTransactionRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRechargeOrderTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreasuryServer).SubmitRechargeOrderTransactionRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Treasury/SubmitRechargeOrderTransactionRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreasuryServer).SubmitRechargeOrderTransactionRpc(ctx, req.(*SubmitRechargeOrderTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Treasury_CancelRechargeOrderRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRechargeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreasuryServer).CancelRechargeOrderRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Treasury/CancelRechargeOrderRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreasuryServer).CancelRechargeOrderRpc(ctx, req.(*CancelRechargeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Treasury_CheckRechargeOrderStatusRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRechargeOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreasuryServer).CheckRechargeOrderStatusRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Treasury/CheckRechargeOrderStatusRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreasuryServer).CheckRechargeOrderStatusRpc(ctx, req.(*CheckRechargeOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Treasury_ServiceDesc is the grpc.ServiceDesc for Treasury service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Treasury_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Treasury",
	HandlerType: (*TreasuryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRechargeOrderRpc",
			Handler:    _Treasury_CreateRechargeOrderRpc_Handler,
		},
		{
			MethodName: "SubmitRechargeOrderTransactionRpc",
			Handler:    _Treasury_SubmitRechargeOrderTransactionRpc_Handler,
		},
		{
			MethodName: "CancelRechargeOrderRpc",
			Handler:    _Treasury_CancelRechargeOrderRpc_Handler,
		},
		{
			MethodName: "CheckRechargeOrderStatusRpc",
			Handler:    _Treasury_CheckRechargeOrderStatusRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/funds_service/treasury.proto",
}