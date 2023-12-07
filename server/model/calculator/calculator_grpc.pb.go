// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: server/proto/calculator.proto

package calculator

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Addition(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error)
	Substraction(ctx context.Context, in *SubstractionRequest, opts ...grpc.CallOption) (*SubstractionResponse, error)
	Multiplication(ctx context.Context, in *MultiplicationRequest, opts ...grpc.CallOption) (*MultiplicationResponse, error)
	Division(ctx context.Context, in *DivisionRequest, opts ...grpc.CallOption) (*DivisionResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Addition(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error) {
	out := new(AdditionResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Addition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Substraction(ctx context.Context, in *SubstractionRequest, opts ...grpc.CallOption) (*SubstractionResponse, error) {
	out := new(SubstractionResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Substraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiplication(ctx context.Context, in *MultiplicationRequest, opts ...grpc.CallOption) (*MultiplicationResponse, error) {
	out := new(MultiplicationResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Multiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Division(ctx context.Context, in *DivisionRequest, opts ...grpc.CallOption) (*DivisionResponse, error) {
	out := new(DivisionResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Division", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Addition(context.Context, *AdditionRequest) (*AdditionResponse, error)
	Substraction(context.Context, *SubstractionRequest) (*SubstractionResponse, error)
	Multiplication(context.Context, *MultiplicationRequest) (*MultiplicationResponse, error)
	Division(context.Context, *DivisionRequest) (*DivisionResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Addition(context.Context, *AdditionRequest) (*AdditionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addition not implemented")
}
func (UnimplementedCalculatorServiceServer) Substraction(context.Context, *SubstractionRequest) (*SubstractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Substraction not implemented")
}
func (UnimplementedCalculatorServiceServer) Multiplication(context.Context, *MultiplicationRequest) (*MultiplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiplication not implemented")
}
func (UnimplementedCalculatorServiceServer) Division(context.Context, *DivisionRequest) (*DivisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Division not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Addition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Addition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Addition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Addition(ctx, req.(*AdditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Substraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubstractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Substraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Substraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Substraction(ctx, req.(*SubstractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Multiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiplication(ctx, req.(*MultiplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Division_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Division(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Division",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Division(ctx, req.(*DivisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Addition",
			Handler:    _CalculatorService_Addition_Handler,
		},
		{
			MethodName: "Substraction",
			Handler:    _CalculatorService_Substraction_Handler,
		},
		{
			MethodName: "Multiplication",
			Handler:    _CalculatorService_Multiplication_Handler,
		},
		{
			MethodName: "Division",
			Handler:    _CalculatorService_Division_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/proto/calculator.proto",
}