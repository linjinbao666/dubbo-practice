// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.5
// - protoc             v3.19.4
// source: helloworld.proto

package api

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	fmt "fmt"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// UserProviderClient is the client API for UserProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProviderClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc_go.CallOption) (*User, common.ErrorWithAttachment)
	SayHelloStream(ctx context.Context, opts ...grpc_go.CallOption) (UserProvider_SayHelloStreamClient, error)
}

type userProviderClient struct {
	cc *triple.TripleConn
}

type UserProviderClientImpl struct {
	SayHello       func(ctx context.Context, in *HelloRequest) (*User, error)
	SayHelloStream func(ctx context.Context) (UserProvider_SayHelloStreamClient, error)
}

func (c *UserProviderClientImpl) GetDubboStub(cc *triple.TripleConn) UserProviderClient {
	return NewUserProviderClient(cc)
}

func (c *UserProviderClientImpl) XXX_InterfaceName() string {
	return "org.apache.dubbo.quickstart.samples.UserProvider"
}

func NewUserProviderClient(cc *triple.TripleConn) UserProviderClient {
	return &userProviderClient{cc}
}

func (c *userProviderClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc_go.CallOption) (*User, common.ErrorWithAttachment) {
	out := new(User)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SayHello", in, out)
}

func (c *userProviderClient) SayHelloStream(ctx context.Context, opts ...grpc_go.CallOption) (UserProvider_SayHelloStreamClient, error) {
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	stream, err := c.cc.NewStream(ctx, "/"+interfaceKey+"/SayHelloStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userProviderSayHelloStreamClient{stream}
	return x, nil
}

type UserProvider_SayHelloStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*User, error)
	grpc_go.ClientStream
}

type userProviderSayHelloStreamClient struct {
	grpc_go.ClientStream
}

func (x *userProviderSayHelloStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userProviderSayHelloStreamClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserProviderServer is the server API for UserProvider service.
// All implementations must embed UnimplementedUserProviderServer
// for forward compatibility
type UserProviderServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*User, error)
	SayHelloStream(UserProvider_SayHelloStreamServer) error
	mustEmbedUnimplementedUserProviderServer()
}

// UnimplementedUserProviderServer must be embedded to have forward compatible implementations.
type UnimplementedUserProviderServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedUserProviderServer) SayHello(context.Context, *HelloRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedUserProviderServer) SayHelloStream(UserProvider_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (s *UnimplementedUserProviderServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedUserProviderServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedUserProviderServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &UserProvider_ServiceDesc
}
func (s *UnimplementedUserProviderServer) XXX_InterfaceName() string {
	return "org.apache.dubbo.quickstart.samples.UserProvider"
}

func (UnimplementedUserProviderServer) mustEmbedUnimplementedUserProviderServer() {}

// UnsafeUserProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProviderServer will
// result in compilation errors.
type UnsafeUserProviderServer interface {
	mustEmbedUnimplementedUserProviderServer()
}

func RegisterUserProviderServer(s grpc_go.ServiceRegistrar, srv UserProviderServer) {
	s.RegisterService(&UserProvider_ServiceDesc, srv)
}

func _UserProvider_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SayHello", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProvider_SayHelloStream_Handler(srv interface{}, stream grpc_go.ServerStream) error {
	_, ok := srv.(dubbo3.Dubbo3GrpcService)
	invo := invocation.NewRPCInvocation("SayHelloStream", nil, nil)
	if !ok {
		fmt.Println(invo)
		return nil
	}
	return srv.(UserProviderServer).SayHelloStream(&userProviderSayHelloStreamServer{stream})
}

type UserProvider_SayHelloStreamServer interface {
	Send(*User) error
	Recv() (*HelloRequest, error)
	grpc_go.ServerStream
}

type userProviderSayHelloStreamServer struct {
	grpc_go.ServerStream
}

func (x *userProviderSayHelloStreamServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userProviderSayHelloStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserProvider_ServiceDesc is the grpc_go.ServiceDesc for UserProvider service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProvider_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "org.apache.dubbo.quickstart.samples.UserProvider",
	HandlerType: (*UserProviderServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _UserProvider_SayHello_Handler,
		},
	},
	Streams: []grpc_go.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _UserProvider_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}