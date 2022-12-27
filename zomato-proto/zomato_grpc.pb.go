// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: zomato.proto

package proto

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

// ZomatoDatabaseCrudClient is the client API for ZomatoDatabaseCrud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZomatoDatabaseCrudClient interface {
	// user rpc
	CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *VoidUserRequest, opts ...grpc.CallOption) (*AllUsers, error)
	GetUserOrders(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserOrder, error)
	PlaceOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error)
	// restaurant rpc
	CreateRestaurant(ctx context.Context, in *NewRestaurant, opts ...grpc.CallOption) (*Restaurant, error)
	GetRestaurantMenu(ctx context.Context, in *RestaurantMenuRequest, opts ...grpc.CallOption) (*Menu, error)
	UpdateRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Restaurant, error)
	// menu rpc
	CreateDish(ctx context.Context, in *NewDish, opts ...grpc.CallOption) (*Dish, error)
	UpdateDish(ctx context.Context, in *Dish, opts ...grpc.CallOption) (*Dish, error)
	DeleteDish(ctx context.Context, in *Dish, opts ...grpc.CallOption) (*VoidDishResponse, error)
	// agent rpc
	CreateAgent(ctx context.Context, in *NewAgent, opts ...grpc.CallOption) (*Agent, error)
	UpdateAgentStatus(ctx context.Context, in *AgentStatus, opts ...grpc.CallOption) (*AgentStatus, error)
}

type zomatoDatabaseCrudClient struct {
	cc grpc.ClientConnInterface
}

func NewZomatoDatabaseCrudClient(cc grpc.ClientConnInterface) ZomatoDatabaseCrudClient {
	return &zomatoDatabaseCrudClient{cc}
}

func (c *zomatoDatabaseCrudClient) CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) GetUsers(ctx context.Context, in *VoidUserRequest, opts ...grpc.CallOption) (*AllUsers, error) {
	out := new(AllUsers)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) GetUserOrders(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserOrder, error) {
	out := new(UserOrder)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/GetUserOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) PlaceOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) CreateRestaurant(ctx context.Context, in *NewRestaurant, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/CreateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) GetRestaurantMenu(ctx context.Context, in *RestaurantMenuRequest, opts ...grpc.CallOption) (*Menu, error) {
	out := new(Menu)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/GetRestaurantMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) UpdateRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/UpdateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) CreateDish(ctx context.Context, in *NewDish, opts ...grpc.CallOption) (*Dish, error) {
	out := new(Dish)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/CreateDish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) UpdateDish(ctx context.Context, in *Dish, opts ...grpc.CallOption) (*Dish, error) {
	out := new(Dish)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/UpdateDish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) DeleteDish(ctx context.Context, in *Dish, opts ...grpc.CallOption) (*VoidDishResponse, error) {
	out := new(VoidDishResponse)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/DeleteDish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) CreateAgent(ctx context.Context, in *NewAgent, opts ...grpc.CallOption) (*Agent, error) {
	out := new(Agent)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/CreateAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zomatoDatabaseCrudClient) UpdateAgentStatus(ctx context.Context, in *AgentStatus, opts ...grpc.CallOption) (*AgentStatus, error) {
	out := new(AgentStatus)
	err := c.cc.Invoke(ctx, "/zomatoDB.ZomatoDatabaseCrud/UpdateAgentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZomatoDatabaseCrudServer is the server API for ZomatoDatabaseCrud service.
// All implementations must embed UnimplementedZomatoDatabaseCrudServer
// for forward compatibility
type ZomatoDatabaseCrudServer interface {
	// user rpc
	CreateUser(context.Context, *NewUser) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	GetUsers(context.Context, *VoidUserRequest) (*AllUsers, error)
	GetUserOrders(context.Context, *User) (*UserOrder, error)
	PlaceOrder(context.Context, *OrderRequest) (*Order, error)
	// restaurant rpc
	CreateRestaurant(context.Context, *NewRestaurant) (*Restaurant, error)
	GetRestaurantMenu(context.Context, *RestaurantMenuRequest) (*Menu, error)
	UpdateRestaurant(context.Context, *Restaurant) (*Restaurant, error)
	// menu rpc
	CreateDish(context.Context, *NewDish) (*Dish, error)
	UpdateDish(context.Context, *Dish) (*Dish, error)
	DeleteDish(context.Context, *Dish) (*VoidDishResponse, error)
	// agent rpc
	CreateAgent(context.Context, *NewAgent) (*Agent, error)
	UpdateAgentStatus(context.Context, *AgentStatus) (*AgentStatus, error)
	mustEmbedUnimplementedZomatoDatabaseCrudServer()
}

// UnimplementedZomatoDatabaseCrudServer must be embedded to have forward compatible implementations.
type UnimplementedZomatoDatabaseCrudServer struct {
}

func (UnimplementedZomatoDatabaseCrudServer) CreateUser(context.Context, *NewUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) UpdateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) GetUsers(context.Context, *VoidUserRequest) (*AllUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) GetUserOrders(context.Context, *User) (*UserOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserOrders not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) PlaceOrder(context.Context, *OrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) CreateRestaurant(context.Context, *NewRestaurant) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) GetRestaurantMenu(context.Context, *RestaurantMenuRequest) (*Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantMenu not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) UpdateRestaurant(context.Context, *Restaurant) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestaurant not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) CreateDish(context.Context, *NewDish) (*Dish, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDish not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) UpdateDish(context.Context, *Dish) (*Dish, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDish not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) DeleteDish(context.Context, *Dish) (*VoidDishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDish not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) CreateAgent(context.Context, *NewAgent) (*Agent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgent not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) UpdateAgentStatus(context.Context, *AgentStatus) (*AgentStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentStatus not implemented")
}
func (UnimplementedZomatoDatabaseCrudServer) mustEmbedUnimplementedZomatoDatabaseCrudServer() {}

// UnsafeZomatoDatabaseCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZomatoDatabaseCrudServer will
// result in compilation errors.
type UnsafeZomatoDatabaseCrudServer interface {
	mustEmbedUnimplementedZomatoDatabaseCrudServer()
}

func RegisterZomatoDatabaseCrudServer(s grpc.ServiceRegistrar, srv ZomatoDatabaseCrudServer) {
	s.RegisterService(&ZomatoDatabaseCrud_ServiceDesc, srv)
}

func _ZomatoDatabaseCrud_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).CreateUser(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).GetUsers(ctx, req.(*VoidUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_GetUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).GetUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/GetUserOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).GetUserOrders(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).PlaceOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRestaurant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/CreateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).CreateRestaurant(ctx, req.(*NewRestaurant))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_GetRestaurantMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).GetRestaurantMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/GetRestaurantMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).GetRestaurantMenu(ctx, req.(*RestaurantMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_UpdateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Restaurant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).UpdateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/UpdateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).UpdateRestaurant(ctx, req.(*Restaurant))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_CreateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewDish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).CreateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/CreateDish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).CreateDish(ctx, req.(*NewDish))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_UpdateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).UpdateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/UpdateDish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).UpdateDish(ctx, req.(*Dish))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_DeleteDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).DeleteDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/DeleteDish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).DeleteDish(ctx, req.(*Dish))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_CreateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAgent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).CreateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/CreateAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).CreateAgent(ctx, req.(*NewAgent))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZomatoDatabaseCrud_UpdateAgentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZomatoDatabaseCrudServer).UpdateAgentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zomatoDB.ZomatoDatabaseCrud/UpdateAgentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZomatoDatabaseCrudServer).UpdateAgentStatus(ctx, req.(*AgentStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// ZomatoDatabaseCrud_ServiceDesc is the grpc.ServiceDesc for ZomatoDatabaseCrud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZomatoDatabaseCrud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zomatoDB.ZomatoDatabaseCrud",
	HandlerType: (*ZomatoDatabaseCrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ZomatoDatabaseCrud_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ZomatoDatabaseCrud_UpdateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _ZomatoDatabaseCrud_GetUsers_Handler,
		},
		{
			MethodName: "GetUserOrders",
			Handler:    _ZomatoDatabaseCrud_GetUserOrders_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _ZomatoDatabaseCrud_PlaceOrder_Handler,
		},
		{
			MethodName: "CreateRestaurant",
			Handler:    _ZomatoDatabaseCrud_CreateRestaurant_Handler,
		},
		{
			MethodName: "GetRestaurantMenu",
			Handler:    _ZomatoDatabaseCrud_GetRestaurantMenu_Handler,
		},
		{
			MethodName: "UpdateRestaurant",
			Handler:    _ZomatoDatabaseCrud_UpdateRestaurant_Handler,
		},
		{
			MethodName: "CreateDish",
			Handler:    _ZomatoDatabaseCrud_CreateDish_Handler,
		},
		{
			MethodName: "UpdateDish",
			Handler:    _ZomatoDatabaseCrud_UpdateDish_Handler,
		},
		{
			MethodName: "DeleteDish",
			Handler:    _ZomatoDatabaseCrud_DeleteDish_Handler,
		},
		{
			MethodName: "CreateAgent",
			Handler:    _ZomatoDatabaseCrud_CreateAgent_Handler,
		},
		{
			MethodName: "UpdateAgentStatus",
			Handler:    _ZomatoDatabaseCrud_UpdateAgentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zomato.proto",
}
