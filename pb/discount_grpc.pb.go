// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package discount

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

// DiscountClient is the client API for Discount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscountClient interface {
	GetDiscount(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*GetDiscountResponse, error)
}

type discountClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountClient(cc grpc.ClientConnInterface) DiscountClient {
	return &discountClient{cc}
}

func (c *discountClient) GetDiscount(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*GetDiscountResponse, error) {
	out := new(GetDiscountResponse)
	err := c.cc.Invoke(ctx, "/discount.Discount/GetDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountServer is the server API for Discount service.
// All implementations must embed UnimplementedDiscountServer
// for forward compatibility
type DiscountServer interface {
	GetDiscount(context.Context, *GetDiscountRequest) (*GetDiscountResponse, error)
	mustEmbedUnimplementedDiscountServer()
}

// UnimplementedDiscountServer must be embedded to have forward compatible implementations.
type UnimplementedDiscountServer struct {
}

func (UnimplementedDiscountServer) GetDiscount(context.Context, *GetDiscountRequest) (*GetDiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscount not implemented")
}
func (UnimplementedDiscountServer) mustEmbedUnimplementedDiscountServer() {}

// UnsafeDiscountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscountServer will
// result in compilation errors.
type UnsafeDiscountServer interface {
	mustEmbedUnimplementedDiscountServer()
}

func RegisterDiscountServer(s grpc.ServiceRegistrar, srv DiscountServer) {
	s.RegisterService(&Discount_ServiceDesc, srv)
}

func _Discount_GetDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServer).GetDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.Discount/GetDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServer).GetDiscount(ctx, req.(*GetDiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Discount_ServiceDesc is the grpc.ServiceDesc for Discount service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Discount_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discount.Discount",
	HandlerType: (*DiscountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiscount",
			Handler:    _Discount_GetDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/discount.proto",
}
