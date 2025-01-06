// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: band/tunnel/v1beta1/tx.proto

package tunnelv1beta1

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
	Msg_CreateTunnel_FullMethodName             = "/band.tunnel.v1beta1.Msg/CreateTunnel"
	Msg_UpdateRoute_FullMethodName              = "/band.tunnel.v1beta1.Msg/UpdateRoute"
	Msg_UpdateSignalsAndInterval_FullMethodName = "/band.tunnel.v1beta1.Msg/UpdateSignalsAndInterval"
	Msg_WithdrawFeePayerFunds_FullMethodName    = "/band.tunnel.v1beta1.Msg/WithdrawFeePayerFunds"
	Msg_ActivateTunnel_FullMethodName           = "/band.tunnel.v1beta1.Msg/ActivateTunnel"
	Msg_DeactivateTunnel_FullMethodName         = "/band.tunnel.v1beta1.Msg/DeactivateTunnel"
	Msg_TriggerTunnel_FullMethodName            = "/band.tunnel.v1beta1.Msg/TriggerTunnel"
	Msg_DepositToTunnel_FullMethodName          = "/band.tunnel.v1beta1.Msg/DepositToTunnel"
	Msg_WithdrawFromTunnel_FullMethodName       = "/band.tunnel.v1beta1.Msg/WithdrawFromTunnel"
	Msg_UpdateParams_FullMethodName             = "/band.tunnel.v1beta1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateTunnel is a RPC method to create a new tunnel.
	CreateTunnel(ctx context.Context, in *MsgCreateTunnel, opts ...grpc.CallOption) (*MsgCreateTunnelResponse, error)
	// UpdateRoute is a RPC method to update a route information of the tunnel.
	UpdateRoute(ctx context.Context, in *MsgUpdateRoute, opts ...grpc.CallOption) (*MsgUpdateRouteResponse, error)
	// UpdateSignalsAndInterval is a RPC method to update a signals and interval of the tunnel.
	UpdateSignalsAndInterval(ctx context.Context, in *MsgUpdateSignalsAndInterval, opts ...grpc.CallOption) (*MsgUpdateSignalsAndIntervalResponse, error)
	// WithdrawFeePayerFunds is a RPC method to withdraw fee payer funds to creator.
	WithdrawFeePayerFunds(ctx context.Context, in *MsgWithdrawFeePayerFunds, opts ...grpc.CallOption) (*MsgWithdrawFeePayerFundsResponse, error)
	// ActivateTunnel is a RPC method to activate a tunnel.
	ActivateTunnel(ctx context.Context, in *MsgActivateTunnel, opts ...grpc.CallOption) (*MsgActivateTunnelResponse, error)
	// DeactivateTunnel is a RPC method to deactivate a tunnel.
	DeactivateTunnel(ctx context.Context, in *MsgDeactivateTunnel, opts ...grpc.CallOption) (*MsgDeactivateTunnelResponse, error)
	// TriggerTunnel is a RPC method to manually trigger a tunnel.
	TriggerTunnel(ctx context.Context, in *MsgTriggerTunnel, opts ...grpc.CallOption) (*MsgTriggerTunnelResponse, error)
	// DepositToTunnel is a RPC method to deposit to an existing tunnel.
	DepositToTunnel(ctx context.Context, in *MsgDepositToTunnel, opts ...grpc.CallOption) (*MsgDepositToTunnelResponse, error)
	// WithdrawFromTunnel is a RPC method to withdraw a deposit from an existing tunnel.
	WithdrawFromTunnel(ctx context.Context, in *MsgWithdrawFromTunnel, opts ...grpc.CallOption) (*MsgWithdrawFromTunnelResponse, error)
	// UpdateParams is a RPC method to update parameters
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateTunnel(ctx context.Context, in *MsgCreateTunnel, opts ...grpc.CallOption) (*MsgCreateTunnelResponse, error) {
	out := new(MsgCreateTunnelResponse)
	err := c.cc.Invoke(ctx, Msg_CreateTunnel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRoute(ctx context.Context, in *MsgUpdateRoute, opts ...grpc.CallOption) (*MsgUpdateRouteResponse, error) {
	out := new(MsgUpdateRouteResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateSignalsAndInterval(ctx context.Context, in *MsgUpdateSignalsAndInterval, opts ...grpc.CallOption) (*MsgUpdateSignalsAndIntervalResponse, error) {
	out := new(MsgUpdateSignalsAndIntervalResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateSignalsAndInterval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawFeePayerFunds(ctx context.Context, in *MsgWithdrawFeePayerFunds, opts ...grpc.CallOption) (*MsgWithdrawFeePayerFundsResponse, error) {
	out := new(MsgWithdrawFeePayerFundsResponse)
	err := c.cc.Invoke(ctx, Msg_WithdrawFeePayerFunds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ActivateTunnel(ctx context.Context, in *MsgActivateTunnel, opts ...grpc.CallOption) (*MsgActivateTunnelResponse, error) {
	out := new(MsgActivateTunnelResponse)
	err := c.cc.Invoke(ctx, Msg_ActivateTunnel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeactivateTunnel(ctx context.Context, in *MsgDeactivateTunnel, opts ...grpc.CallOption) (*MsgDeactivateTunnelResponse, error) {
	out := new(MsgDeactivateTunnelResponse)
	err := c.cc.Invoke(ctx, Msg_DeactivateTunnel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TriggerTunnel(ctx context.Context, in *MsgTriggerTunnel, opts ...grpc.CallOption) (*MsgTriggerTunnelResponse, error) {
	out := new(MsgTriggerTunnelResponse)
	err := c.cc.Invoke(ctx, Msg_TriggerTunnel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DepositToTunnel(ctx context.Context, in *MsgDepositToTunnel, opts ...grpc.CallOption) (*MsgDepositToTunnelResponse, error) {
	out := new(MsgDepositToTunnelResponse)
	err := c.cc.Invoke(ctx, Msg_DepositToTunnel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawFromTunnel(ctx context.Context, in *MsgWithdrawFromTunnel, opts ...grpc.CallOption) (*MsgWithdrawFromTunnelResponse, error) {
	out := new(MsgWithdrawFromTunnelResponse)
	err := c.cc.Invoke(ctx, Msg_WithdrawFromTunnel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateTunnel is a RPC method to create a new tunnel.
	CreateTunnel(context.Context, *MsgCreateTunnel) (*MsgCreateTunnelResponse, error)
	// UpdateRoute is a RPC method to update a route information of the tunnel.
	UpdateRoute(context.Context, *MsgUpdateRoute) (*MsgUpdateRouteResponse, error)
	// UpdateSignalsAndInterval is a RPC method to update a signals and interval of the tunnel.
	UpdateSignalsAndInterval(context.Context, *MsgUpdateSignalsAndInterval) (*MsgUpdateSignalsAndIntervalResponse, error)
	// WithdrawFeePayerFunds is a RPC method to withdraw fee payer funds to creator.
	WithdrawFeePayerFunds(context.Context, *MsgWithdrawFeePayerFunds) (*MsgWithdrawFeePayerFundsResponse, error)
	// ActivateTunnel is a RPC method to activate a tunnel.
	ActivateTunnel(context.Context, *MsgActivateTunnel) (*MsgActivateTunnelResponse, error)
	// DeactivateTunnel is a RPC method to deactivate a tunnel.
	DeactivateTunnel(context.Context, *MsgDeactivateTunnel) (*MsgDeactivateTunnelResponse, error)
	// TriggerTunnel is a RPC method to manually trigger a tunnel.
	TriggerTunnel(context.Context, *MsgTriggerTunnel) (*MsgTriggerTunnelResponse, error)
	// DepositToTunnel is a RPC method to deposit to an existing tunnel.
	DepositToTunnel(context.Context, *MsgDepositToTunnel) (*MsgDepositToTunnelResponse, error)
	// WithdrawFromTunnel is a RPC method to withdraw a deposit from an existing tunnel.
	WithdrawFromTunnel(context.Context, *MsgWithdrawFromTunnel) (*MsgWithdrawFromTunnelResponse, error)
	// UpdateParams is a RPC method to update parameters
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateTunnel(context.Context, *MsgCreateTunnel) (*MsgCreateTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTunnel not implemented")
}
func (UnimplementedMsgServer) UpdateRoute(context.Context, *MsgUpdateRoute) (*MsgUpdateRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoute not implemented")
}
func (UnimplementedMsgServer) UpdateSignalsAndInterval(context.Context, *MsgUpdateSignalsAndInterval) (*MsgUpdateSignalsAndIntervalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSignalsAndInterval not implemented")
}
func (UnimplementedMsgServer) WithdrawFeePayerFunds(context.Context, *MsgWithdrawFeePayerFunds) (*MsgWithdrawFeePayerFundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFeePayerFunds not implemented")
}
func (UnimplementedMsgServer) ActivateTunnel(context.Context, *MsgActivateTunnel) (*MsgActivateTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateTunnel not implemented")
}
func (UnimplementedMsgServer) DeactivateTunnel(context.Context, *MsgDeactivateTunnel) (*MsgDeactivateTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateTunnel not implemented")
}
func (UnimplementedMsgServer) TriggerTunnel(context.Context, *MsgTriggerTunnel) (*MsgTriggerTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerTunnel not implemented")
}
func (UnimplementedMsgServer) DepositToTunnel(context.Context, *MsgDepositToTunnel) (*MsgDepositToTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositToTunnel not implemented")
}
func (UnimplementedMsgServer) WithdrawFromTunnel(context.Context, *MsgWithdrawFromTunnel) (*MsgWithdrawFromTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFromTunnel not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateTunnel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateTunnel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateTunnel(ctx, req.(*MsgCreateTunnel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRoute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRoute(ctx, req.(*MsgUpdateRoute))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateSignalsAndInterval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateSignalsAndInterval)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateSignalsAndInterval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateSignalsAndInterval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateSignalsAndInterval(ctx, req.(*MsgUpdateSignalsAndInterval))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawFeePayerFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawFeePayerFunds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawFeePayerFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_WithdrawFeePayerFunds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawFeePayerFunds(ctx, req.(*MsgWithdrawFeePayerFunds))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ActivateTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgActivateTunnel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ActivateTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ActivateTunnel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ActivateTunnel(ctx, req.(*MsgActivateTunnel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeactivateTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeactivateTunnel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeactivateTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeactivateTunnel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeactivateTunnel(ctx, req.(*MsgDeactivateTunnel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TriggerTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTriggerTunnel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TriggerTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TriggerTunnel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TriggerTunnel(ctx, req.(*MsgTriggerTunnel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DepositToTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepositToTunnel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositToTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DepositToTunnel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositToTunnel(ctx, req.(*MsgDepositToTunnel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawFromTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawFromTunnel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawFromTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_WithdrawFromTunnel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawFromTunnel(ctx, req.(*MsgWithdrawFromTunnel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "band.tunnel.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTunnel",
			Handler:    _Msg_CreateTunnel_Handler,
		},
		{
			MethodName: "UpdateRoute",
			Handler:    _Msg_UpdateRoute_Handler,
		},
		{
			MethodName: "UpdateSignalsAndInterval",
			Handler:    _Msg_UpdateSignalsAndInterval_Handler,
		},
		{
			MethodName: "WithdrawFeePayerFunds",
			Handler:    _Msg_WithdrawFeePayerFunds_Handler,
		},
		{
			MethodName: "ActivateTunnel",
			Handler:    _Msg_ActivateTunnel_Handler,
		},
		{
			MethodName: "DeactivateTunnel",
			Handler:    _Msg_DeactivateTunnel_Handler,
		},
		{
			MethodName: "TriggerTunnel",
			Handler:    _Msg_TriggerTunnel_Handler,
		},
		{
			MethodName: "DepositToTunnel",
			Handler:    _Msg_DepositToTunnel_Handler,
		},
		{
			MethodName: "WithdrawFromTunnel",
			Handler:    _Msg_WithdrawFromTunnel_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "band/tunnel/v1beta1/tx.proto",
}
