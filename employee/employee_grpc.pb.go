// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: employee/employee.proto

package employee

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

// EmployeeClient is the client API for Employee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeClient interface {
	GetEmployeeDetails(ctx context.Context, in *EmployeeQuery, opts ...grpc.CallOption) (*EmployeeDetails, error)
	AddEmployeeDetails(ctx context.Context, in *EmployeeDetails, opts ...grpc.CallOption) (*EmployeeResponse, error)
}

type employeeClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeClient(cc grpc.ClientConnInterface) EmployeeClient {
	return &employeeClient{cc}
}

func (c *employeeClient) GetEmployeeDetails(ctx context.Context, in *EmployeeQuery, opts ...grpc.CallOption) (*EmployeeDetails, error) {
	out := new(EmployeeDetails)
	err := c.cc.Invoke(ctx, "/employee.employee/getEmployeeDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeClient) AddEmployeeDetails(ctx context.Context, in *EmployeeDetails, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/employee.employee/addEmployeeDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServer is the server API for Employee service.
// All implementations must embed UnimplementedEmployeeServer
// for forward compatibility
type EmployeeServer interface {
	GetEmployeeDetails(context.Context, *EmployeeQuery) (*EmployeeDetails, error)
	AddEmployeeDetails(context.Context, *EmployeeDetails) (*EmployeeResponse, error)
	mustEmbedUnimplementedEmployeeServer()
}

// UnimplementedEmployeeServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServer struct {
}

func (UnimplementedEmployeeServer) GetEmployeeDetails(context.Context, *EmployeeQuery) (*EmployeeDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployeeDetails not implemented")
}
func (UnimplementedEmployeeServer) AddEmployeeDetails(context.Context, *EmployeeDetails) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployeeDetails not implemented")
}
func (UnimplementedEmployeeServer) mustEmbedUnimplementedEmployeeServer() {}

// UnsafeEmployeeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServer will
// result in compilation errors.
type UnsafeEmployeeServer interface {
	mustEmbedUnimplementedEmployeeServer()
}

func RegisterEmployeeServer(s grpc.ServiceRegistrar, srv EmployeeServer) {
	s.RegisterService(&Employee_ServiceDesc, srv)
}

func _Employee_GetEmployeeDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).GetEmployeeDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.employee/getEmployeeDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).GetEmployeeDetails(ctx, req.(*EmployeeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Employee_AddEmployeeDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServer).AddEmployeeDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.employee/addEmployeeDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServer).AddEmployeeDetails(ctx, req.(*EmployeeDetails))
	}
	return interceptor(ctx, in, info, handler)
}

// Employee_ServiceDesc is the grpc.ServiceDesc for Employee service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Employee_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee.employee",
	HandlerType: (*EmployeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getEmployeeDetails",
			Handler:    _Employee_GetEmployeeDetails_Handler,
		},
		{
			MethodName: "addEmployeeDetails",
			Handler:    _Employee_AddEmployeeDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee/employee.proto",
}
