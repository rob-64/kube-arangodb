// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: integrations/scheduler/v2/definition/definition.proto

package definition

import (
	context "context"
	definition "github.com/arangodb/kube-arangodb/integrations/shared/v1/definition"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SchedulerV2Client is the client API for SchedulerV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerV2Client interface {
	// Invalidates local ServiceDiscover cache
	InvalidateCache(ctx context.Context, in *definition.Empty, opts ...grpc.CallOption) (*definition.Empty, error)
	// Ensure that Helm Client is able to communicate with Kubernetes
	Alive(ctx context.Context, in *definition.Empty, opts ...grpc.CallOption) (*definition.Empty, error)
	// Executes Helm List Action
	List(ctx context.Context, in *SchedulerV2ListRequest, opts ...grpc.CallOption) (*SchedulerV2ListResponse, error)
	// Executes Helm Status Action
	Status(ctx context.Context, in *SchedulerV2StatusRequest, opts ...grpc.CallOption) (*SchedulerV2StatusResponse, error)
	// Executes Helm Status Action and fetch Resources from Kubernetes API
	StatusObjects(ctx context.Context, in *SchedulerV2StatusObjectsRequest, opts ...grpc.CallOption) (*SchedulerV2StatusObjectsResponse, error)
	// Executes Helm Install Action
	Install(ctx context.Context, in *SchedulerV2InstallRequest, opts ...grpc.CallOption) (*SchedulerV2InstallResponse, error)
	// Executes Helm Upgrade Action
	Upgrade(ctx context.Context, in *SchedulerV2UpgradeRequest, opts ...grpc.CallOption) (*SchedulerV2UpgradeResponse, error)
	// Executes Helm Uninstall Action
	Uninstall(ctx context.Context, in *SchedulerV2UninstallRequest, opts ...grpc.CallOption) (*SchedulerV2UninstallResponse, error)
	// Executes Helm Test Action
	Test(ctx context.Context, in *SchedulerV2TestRequest, opts ...grpc.CallOption) (*SchedulerV2TestResponse, error)
	// Discovers Kubernetes API Resources for Group
	DiscoverAPIResources(ctx context.Context, in *SchedulerV2DiscoverAPIResourcesRequest, opts ...grpc.CallOption) (*SchedulerV2DiscoverAPIResourcesResponse, error)
	// Discovers Kubernetes API Resources for Kind
	DiscoverAPIResource(ctx context.Context, in *SchedulerV2DiscoverAPIResourceRequest, opts ...grpc.CallOption) (*SchedulerV2DiscoverAPIResourceResponse, error)
	// Gets Kubernetes objects from the API
	KubernetesGet(ctx context.Context, in *SchedulerV2KubernetesGetRequest, opts ...grpc.CallOption) (*SchedulerV2KubernetesGetResponse, error)
}

type schedulerV2Client struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerV2Client(cc grpc.ClientConnInterface) SchedulerV2Client {
	return &schedulerV2Client{cc}
}

func (c *schedulerV2Client) InvalidateCache(ctx context.Context, in *definition.Empty, opts ...grpc.CallOption) (*definition.Empty, error) {
	out := new(definition.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/InvalidateCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) Alive(ctx context.Context, in *definition.Empty, opts ...grpc.CallOption) (*definition.Empty, error) {
	out := new(definition.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/Alive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) List(ctx context.Context, in *SchedulerV2ListRequest, opts ...grpc.CallOption) (*SchedulerV2ListResponse, error) {
	out := new(SchedulerV2ListResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) Status(ctx context.Context, in *SchedulerV2StatusRequest, opts ...grpc.CallOption) (*SchedulerV2StatusResponse, error) {
	out := new(SchedulerV2StatusResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) StatusObjects(ctx context.Context, in *SchedulerV2StatusObjectsRequest, opts ...grpc.CallOption) (*SchedulerV2StatusObjectsResponse, error) {
	out := new(SchedulerV2StatusObjectsResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/StatusObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) Install(ctx context.Context, in *SchedulerV2InstallRequest, opts ...grpc.CallOption) (*SchedulerV2InstallResponse, error) {
	out := new(SchedulerV2InstallResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/Install", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) Upgrade(ctx context.Context, in *SchedulerV2UpgradeRequest, opts ...grpc.CallOption) (*SchedulerV2UpgradeResponse, error) {
	out := new(SchedulerV2UpgradeResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/Upgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) Uninstall(ctx context.Context, in *SchedulerV2UninstallRequest, opts ...grpc.CallOption) (*SchedulerV2UninstallResponse, error) {
	out := new(SchedulerV2UninstallResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/Uninstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) Test(ctx context.Context, in *SchedulerV2TestRequest, opts ...grpc.CallOption) (*SchedulerV2TestResponse, error) {
	out := new(SchedulerV2TestResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) DiscoverAPIResources(ctx context.Context, in *SchedulerV2DiscoverAPIResourcesRequest, opts ...grpc.CallOption) (*SchedulerV2DiscoverAPIResourcesResponse, error) {
	out := new(SchedulerV2DiscoverAPIResourcesResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/DiscoverAPIResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) DiscoverAPIResource(ctx context.Context, in *SchedulerV2DiscoverAPIResourceRequest, opts ...grpc.CallOption) (*SchedulerV2DiscoverAPIResourceResponse, error) {
	out := new(SchedulerV2DiscoverAPIResourceResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/DiscoverAPIResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerV2Client) KubernetesGet(ctx context.Context, in *SchedulerV2KubernetesGetRequest, opts ...grpc.CallOption) (*SchedulerV2KubernetesGetResponse, error) {
	out := new(SchedulerV2KubernetesGetResponse)
	err := c.cc.Invoke(ctx, "/scheduler.SchedulerV2/KubernetesGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerV2Server is the server API for SchedulerV2 service.
// All implementations must embed UnimplementedSchedulerV2Server
// for forward compatibility
type SchedulerV2Server interface {
	// Invalidates local ServiceDiscover cache
	InvalidateCache(context.Context, *definition.Empty) (*definition.Empty, error)
	// Ensure that Helm Client is able to communicate with Kubernetes
	Alive(context.Context, *definition.Empty) (*definition.Empty, error)
	// Executes Helm List Action
	List(context.Context, *SchedulerV2ListRequest) (*SchedulerV2ListResponse, error)
	// Executes Helm Status Action
	Status(context.Context, *SchedulerV2StatusRequest) (*SchedulerV2StatusResponse, error)
	// Executes Helm Status Action and fetch Resources from Kubernetes API
	StatusObjects(context.Context, *SchedulerV2StatusObjectsRequest) (*SchedulerV2StatusObjectsResponse, error)
	// Executes Helm Install Action
	Install(context.Context, *SchedulerV2InstallRequest) (*SchedulerV2InstallResponse, error)
	// Executes Helm Upgrade Action
	Upgrade(context.Context, *SchedulerV2UpgradeRequest) (*SchedulerV2UpgradeResponse, error)
	// Executes Helm Uninstall Action
	Uninstall(context.Context, *SchedulerV2UninstallRequest) (*SchedulerV2UninstallResponse, error)
	// Executes Helm Test Action
	Test(context.Context, *SchedulerV2TestRequest) (*SchedulerV2TestResponse, error)
	// Discovers Kubernetes API Resources for Group
	DiscoverAPIResources(context.Context, *SchedulerV2DiscoverAPIResourcesRequest) (*SchedulerV2DiscoverAPIResourcesResponse, error)
	// Discovers Kubernetes API Resources for Kind
	DiscoverAPIResource(context.Context, *SchedulerV2DiscoverAPIResourceRequest) (*SchedulerV2DiscoverAPIResourceResponse, error)
	// Gets Kubernetes objects from the API
	KubernetesGet(context.Context, *SchedulerV2KubernetesGetRequest) (*SchedulerV2KubernetesGetResponse, error)
	mustEmbedUnimplementedSchedulerV2Server()
}

// UnimplementedSchedulerV2Server must be embedded to have forward compatible implementations.
type UnimplementedSchedulerV2Server struct {
}

func (UnimplementedSchedulerV2Server) InvalidateCache(context.Context, *definition.Empty) (*definition.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateCache not implemented")
}
func (UnimplementedSchedulerV2Server) Alive(context.Context, *definition.Empty) (*definition.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (UnimplementedSchedulerV2Server) List(context.Context, *SchedulerV2ListRequest) (*SchedulerV2ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSchedulerV2Server) Status(context.Context, *SchedulerV2StatusRequest) (*SchedulerV2StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedSchedulerV2Server) StatusObjects(context.Context, *SchedulerV2StatusObjectsRequest) (*SchedulerV2StatusObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusObjects not implemented")
}
func (UnimplementedSchedulerV2Server) Install(context.Context, *SchedulerV2InstallRequest) (*SchedulerV2InstallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedSchedulerV2Server) Upgrade(context.Context, *SchedulerV2UpgradeRequest) (*SchedulerV2UpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upgrade not implemented")
}
func (UnimplementedSchedulerV2Server) Uninstall(context.Context, *SchedulerV2UninstallRequest) (*SchedulerV2UninstallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uninstall not implemented")
}
func (UnimplementedSchedulerV2Server) Test(context.Context, *SchedulerV2TestRequest) (*SchedulerV2TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedSchedulerV2Server) DiscoverAPIResources(context.Context, *SchedulerV2DiscoverAPIResourcesRequest) (*SchedulerV2DiscoverAPIResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverAPIResources not implemented")
}
func (UnimplementedSchedulerV2Server) DiscoverAPIResource(context.Context, *SchedulerV2DiscoverAPIResourceRequest) (*SchedulerV2DiscoverAPIResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverAPIResource not implemented")
}
func (UnimplementedSchedulerV2Server) KubernetesGet(context.Context, *SchedulerV2KubernetesGetRequest) (*SchedulerV2KubernetesGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KubernetesGet not implemented")
}
func (UnimplementedSchedulerV2Server) mustEmbedUnimplementedSchedulerV2Server() {}

// UnsafeSchedulerV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerV2Server will
// result in compilation errors.
type UnsafeSchedulerV2Server interface {
	mustEmbedUnimplementedSchedulerV2Server()
}

func RegisterSchedulerV2Server(s grpc.ServiceRegistrar, srv SchedulerV2Server) {
	s.RegisterService(&SchedulerV2_ServiceDesc, srv)
}

func _SchedulerV2_InvalidateCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(definition.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).InvalidateCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/InvalidateCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).InvalidateCache(ctx, req.(*definition.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(definition.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/Alive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).Alive(ctx, req.(*definition.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).List(ctx, req.(*SchedulerV2ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).Status(ctx, req.(*SchedulerV2StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_StatusObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2StatusObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).StatusObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/StatusObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).StatusObjects(ctx, req.(*SchedulerV2StatusObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2InstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/Install",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).Install(ctx, req.(*SchedulerV2InstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_Upgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2UpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).Upgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/Upgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).Upgrade(ctx, req.(*SchedulerV2UpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_Uninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2UninstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).Uninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/Uninstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).Uninstall(ctx, req.(*SchedulerV2UninstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).Test(ctx, req.(*SchedulerV2TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_DiscoverAPIResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2DiscoverAPIResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).DiscoverAPIResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/DiscoverAPIResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).DiscoverAPIResources(ctx, req.(*SchedulerV2DiscoverAPIResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_DiscoverAPIResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2DiscoverAPIResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).DiscoverAPIResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/DiscoverAPIResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).DiscoverAPIResource(ctx, req.(*SchedulerV2DiscoverAPIResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerV2_KubernetesGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerV2KubernetesGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerV2Server).KubernetesGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerV2/KubernetesGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerV2Server).KubernetesGet(ctx, req.(*SchedulerV2KubernetesGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchedulerV2_ServiceDesc is the grpc.ServiceDesc for SchedulerV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchedulerV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.SchedulerV2",
	HandlerType: (*SchedulerV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InvalidateCache",
			Handler:    _SchedulerV2_InvalidateCache_Handler,
		},
		{
			MethodName: "Alive",
			Handler:    _SchedulerV2_Alive_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SchedulerV2_List_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _SchedulerV2_Status_Handler,
		},
		{
			MethodName: "StatusObjects",
			Handler:    _SchedulerV2_StatusObjects_Handler,
		},
		{
			MethodName: "Install",
			Handler:    _SchedulerV2_Install_Handler,
		},
		{
			MethodName: "Upgrade",
			Handler:    _SchedulerV2_Upgrade_Handler,
		},
		{
			MethodName: "Uninstall",
			Handler:    _SchedulerV2_Uninstall_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _SchedulerV2_Test_Handler,
		},
		{
			MethodName: "DiscoverAPIResources",
			Handler:    _SchedulerV2_DiscoverAPIResources_Handler,
		},
		{
			MethodName: "DiscoverAPIResource",
			Handler:    _SchedulerV2_DiscoverAPIResource_Handler,
		},
		{
			MethodName: "KubernetesGet",
			Handler:    _SchedulerV2_KubernetesGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrations/scheduler/v2/definition/definition.proto",
}
