// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: pkg/protoorders/pb/orders.proto

package pb

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetOrdersByUser(ctx context.Context, in *GetOrdersByUserRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	GetOrderByID(ctx context.Context, in *GetOrderByIDRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/orders.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrdersByUser(ctx context.Context, in *GetOrdersByUserRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/orders.OrderService/GetOrdersByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderByID(ctx context.Context, in *GetOrderByIDRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/orders.OrderService/GetOrderByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/orders.OrderService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateOrder(context.Context, *NewOrderRequest) (*OrderResponse, error)
	GetOrdersByUser(context.Context, *GetOrdersByUserRequest) (*OrdersResponse, error)
	GetOrderByID(context.Context, *GetOrderByIDRequest) (*OrderResponse, error)
	UpdateStatus(context.Context, *UpdateStatusRequest) (*EmptyResponse, error)
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *NewOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrdersByUser(context.Context, *GetOrdersByUserRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByUser not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderByID(context.Context, *GetOrderByIDRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByID not implemented")
}
func (UnimplementedOrderServiceServer) UpdateStatus(context.Context, *UpdateStatusRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrdersByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrdersByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrderService/GetOrdersByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrdersByUser(ctx, req.(*GetOrdersByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrderService/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderByID(ctx, req.(*GetOrderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrderService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrdersByUser",
			Handler:    _OrderService_GetOrdersByUser_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _OrderService_GetOrderByID_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _OrderService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protoorders/pb/orders.proto",
}