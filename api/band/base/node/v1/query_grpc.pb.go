// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: band/base/node/v1/query.proto

package nodev1

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
	Service_ChainID_FullMethodName       = "/band.base.node.v1.Service/ChainID"
	Service_EVMValidators_FullMethodName = "/band.base.node.v1.Service/EVMValidators"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// ChainID queries the chain ID of this node
	ChainID(ctx context.Context, in *ChainIDRequest, opts ...grpc.CallOption) (*ChainIDResponse, error)
	// EVMValidators queries current list of validator's address and power
	EVMValidators(ctx context.Context, in *EVMValidatorsRequest, opts ...grpc.CallOption) (*EVMValidatorsResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ChainID(ctx context.Context, in *ChainIDRequest, opts ...grpc.CallOption) (*ChainIDResponse, error) {
	out := new(ChainIDResponse)
	err := c.cc.Invoke(ctx, Service_ChainID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) EVMValidators(ctx context.Context, in *EVMValidatorsRequest, opts ...grpc.CallOption) (*EVMValidatorsResponse, error) {
	out := new(EVMValidatorsResponse)
	err := c.cc.Invoke(ctx, Service_EVMValidators_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// ChainID queries the chain ID of this node
	ChainID(context.Context, *ChainIDRequest) (*ChainIDResponse, error)
	// EVMValidators queries current list of validator's address and power
	EVMValidators(context.Context, *EVMValidatorsRequest) (*EVMValidatorsResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) ChainID(context.Context, *ChainIDRequest) (*ChainIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainID not implemented")
}
func (UnimplementedServiceServer) EVMValidators(context.Context, *EVMValidatorsRequest) (*EVMValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EVMValidators not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_ChainID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ChainID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ChainID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ChainID(ctx, req.(*ChainIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_EVMValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EVMValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).EVMValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_EVMValidators_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).EVMValidators(ctx, req.(*EVMValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "band.base.node.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChainID",
			Handler:    _Service_ChainID_Handler,
		},
		{
			MethodName: "EVMValidators",
			Handler:    _Service_EVMValidators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "band/base/node/v1/query.proto",
}
