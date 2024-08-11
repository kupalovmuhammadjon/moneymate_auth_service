// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: goals.proto

package budgeting_service

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

// GoalServiceClient is the client API for GoalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoalServiceClient interface {
	Create(ctx context.Context, in *CreateGoal, opts ...grpc.CallOption) (*Goal, error)
	GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Goal, error)
	GetAll(ctx context.Context, in *GoalFilter, opts ...grpc.CallOption) (*Goals, error)
	Update(ctx context.Context, in *Goal, opts ...grpc.CallOption) (*Goal, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error)
}

type goalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoalServiceClient(cc grpc.ClientConnInterface) GoalServiceClient {
	return &goalServiceClient{cc}
}

func (c *goalServiceClient) Create(ctx context.Context, in *CreateGoal, opts ...grpc.CallOption) (*Goal, error) {
	out := new(Goal)
	err := c.cc.Invoke(ctx, "/budgeting_service.GoalService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Goal, error) {
	out := new(Goal)
	err := c.cc.Invoke(ctx, "/budgeting_service.GoalService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) GetAll(ctx context.Context, in *GoalFilter, opts ...grpc.CallOption) (*Goals, error) {
	out := new(Goals)
	err := c.cc.Invoke(ctx, "/budgeting_service.GoalService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) Update(ctx context.Context, in *Goal, opts ...grpc.CallOption) (*Goal, error) {
	out := new(Goal)
	err := c.cc.Invoke(ctx, "/budgeting_service.GoalService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goalServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/budgeting_service.GoalService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoalServiceServer is the server API for GoalService service.
// All implementations must embed UnimplementedGoalServiceServer
// for forward compatibility
type GoalServiceServer interface {
	Create(context.Context, *CreateGoal) (*Goal, error)
	GetById(context.Context, *PrimaryKey) (*Goal, error)
	GetAll(context.Context, *GoalFilter) (*Goals, error)
	Update(context.Context, *Goal) (*Goal, error)
	Delete(context.Context, *PrimaryKey) (*Void, error)
	mustEmbedUnimplementedGoalServiceServer()
}

// UnimplementedGoalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoalServiceServer struct {
}

func (UnimplementedGoalServiceServer) Create(context.Context, *CreateGoal) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGoalServiceServer) GetById(context.Context, *PrimaryKey) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedGoalServiceServer) GetAll(context.Context, *GoalFilter) (*Goals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGoalServiceServer) Update(context.Context, *Goal) (*Goal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGoalServiceServer) Delete(context.Context, *PrimaryKey) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGoalServiceServer) mustEmbedUnimplementedGoalServiceServer() {}

// UnsafeGoalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoalServiceServer will
// result in compilation errors.
type UnsafeGoalServiceServer interface {
	mustEmbedUnimplementedGoalServiceServer()
}

func RegisterGoalServiceServer(s grpc.ServiceRegistrar, srv GoalServiceServer) {
	s.RegisterService(&GoalService_ServiceDesc, srv)
}

func _GoalService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.GoalService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).Create(ctx, req.(*CreateGoal))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.GoalService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).GetById(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoalFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.GoalService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).GetAll(ctx, req.(*GoalFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Goal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.GoalService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).Update(ctx, req.(*Goal))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoalService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoalServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/budgeting_service.GoalService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoalServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// GoalService_ServiceDesc is the grpc.ServiceDesc for GoalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "budgeting_service.GoalService",
	HandlerType: (*GoalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GoalService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _GoalService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _GoalService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GoalService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GoalService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goals.proto",
}
