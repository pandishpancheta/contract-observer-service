// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: pkg/listings/pb/listings.proto

package pb

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

// ListingsServiceClient is the client API for ListingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListingsServiceClient interface {
	CreateListing(ctx context.Context, in *CreateListingRequest, opts ...grpc.CallOption) (*CreateListingResponse, error)
	GetListing(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*GetListingResponse, error)
	GetListings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListingsResponse, error)
	UpdateListing(ctx context.Context, in *UpdateListingRequest, opts ...grpc.CallOption) (*UpdateListingResponse, error)
	UpdateListingStatus(ctx context.Context, in *UpdateListingStatusRequest, opts ...grpc.CallOption) (*UpdateListingStatusResponse, error)
	DeleteListing(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type listingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListingsServiceClient(cc grpc.ClientConnInterface) ListingsServiceClient {
	return &listingsServiceClient{cc}
}

func (c *listingsServiceClient) CreateListing(ctx context.Context, in *CreateListingRequest, opts ...grpc.CallOption) (*CreateListingResponse, error) {
	out := new(CreateListingResponse)
	err := c.cc.Invoke(ctx, "/listings.ListingsService/CreateListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingsServiceClient) GetListing(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*GetListingResponse, error) {
	out := new(GetListingResponse)
	err := c.cc.Invoke(ctx, "/listings.ListingsService/GetListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingsServiceClient) GetListings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetListingsResponse, error) {
	out := new(GetListingsResponse)
	err := c.cc.Invoke(ctx, "/listings.ListingsService/GetListings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingsServiceClient) UpdateListing(ctx context.Context, in *UpdateListingRequest, opts ...grpc.CallOption) (*UpdateListingResponse, error) {
	out := new(UpdateListingResponse)
	err := c.cc.Invoke(ctx, "/listings.ListingsService/UpdateListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingsServiceClient) UpdateListingStatus(ctx context.Context, in *UpdateListingStatusRequest, opts ...grpc.CallOption) (*UpdateListingStatusResponse, error) {
	out := new(UpdateListingStatusResponse)
	err := c.cc.Invoke(ctx, "/listings.ListingsService/UpdateListingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingsServiceClient) DeleteListing(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/listings.ListingsService/DeleteListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListingsServiceServer is the server API for ListingsService service.
// All implementations should embed UnimplementedListingsServiceServer
// for forward compatibility
type ListingsServiceServer interface {
	CreateListing(context.Context, *CreateListingRequest) (*CreateListingResponse, error)
	GetListing(context.Context, *GetListingRequest) (*GetListingResponse, error)
	GetListings(context.Context, *emptypb.Empty) (*GetListingsResponse, error)
	UpdateListing(context.Context, *UpdateListingRequest) (*UpdateListingResponse, error)
	UpdateListingStatus(context.Context, *UpdateListingStatusRequest) (*UpdateListingStatusResponse, error)
	DeleteListing(context.Context, *DeleteListingRequest) (*emptypb.Empty, error)
}

// UnimplementedListingsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedListingsServiceServer struct {
}

func (UnimplementedListingsServiceServer) CreateListing(context.Context, *CreateListingRequest) (*CreateListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateListing not implemented")
}
func (UnimplementedListingsServiceServer) GetListing(context.Context, *GetListingRequest) (*GetListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListing not implemented")
}
func (UnimplementedListingsServiceServer) GetListings(context.Context, *emptypb.Empty) (*GetListingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListings not implemented")
}
func (UnimplementedListingsServiceServer) UpdateListing(context.Context, *UpdateListingRequest) (*UpdateListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListing not implemented")
}
func (UnimplementedListingsServiceServer) UpdateListingStatus(context.Context, *UpdateListingStatusRequest) (*UpdateListingStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListingStatus not implemented")
}
func (UnimplementedListingsServiceServer) DeleteListing(context.Context, *DeleteListingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteListing not implemented")
}

// UnsafeListingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListingsServiceServer will
// result in compilation errors.
type UnsafeListingsServiceServer interface {
	mustEmbedUnimplementedListingsServiceServer()
}

func RegisterListingsServiceServer(s grpc.ServiceRegistrar, srv ListingsServiceServer) {
	s.RegisterService(&ListingsService_ServiceDesc, srv)
}

func _ListingsService_CreateListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingsServiceServer).CreateListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listings.ListingsService/CreateListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingsServiceServer).CreateListing(ctx, req.(*CreateListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingsService_GetListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingsServiceServer).GetListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listings.ListingsService/GetListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingsServiceServer).GetListing(ctx, req.(*GetListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingsService_GetListings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingsServiceServer).GetListings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listings.ListingsService/GetListings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingsServiceServer).GetListings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingsService_UpdateListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingsServiceServer).UpdateListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listings.ListingsService/UpdateListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingsServiceServer).UpdateListing(ctx, req.(*UpdateListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingsService_UpdateListingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListingStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingsServiceServer).UpdateListingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listings.ListingsService/UpdateListingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingsServiceServer).UpdateListingStatus(ctx, req.(*UpdateListingStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingsService_DeleteListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingsServiceServer).DeleteListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listings.ListingsService/DeleteListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingsServiceServer).DeleteListing(ctx, req.(*DeleteListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListingsService_ServiceDesc is the grpc.ServiceDesc for ListingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listings.ListingsService",
	HandlerType: (*ListingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateListing",
			Handler:    _ListingsService_CreateListing_Handler,
		},
		{
			MethodName: "GetListing",
			Handler:    _ListingsService_GetListing_Handler,
		},
		{
			MethodName: "GetListings",
			Handler:    _ListingsService_GetListings_Handler,
		},
		{
			MethodName: "UpdateListing",
			Handler:    _ListingsService_UpdateListing_Handler,
		},
		{
			MethodName: "UpdateListingStatus",
			Handler:    _ListingsService_UpdateListingStatus_Handler,
		},
		{
			MethodName: "DeleteListing",
			Handler:    _ListingsService_DeleteListing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/listings/pb/listings.proto",
}