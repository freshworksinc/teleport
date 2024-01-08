// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/scim/v1/scim_service.proto

package scimv1

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
	ScimService_Get_FullMethodName = "/teleport.scim.v1.ScimService/Get"
)

// ScimServiceClient is the client API for ScimService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScimServiceClient interface {
	// Handle handles an inbound SCIM GET request from an IDP
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ScimResponse, error)
}

type scimServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScimServiceClient(cc grpc.ClientConnInterface) ScimServiceClient {
	return &scimServiceClient{cc}
}

func (c *scimServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ScimResponse, error) {
	out := new(ScimResponse)
	err := c.cc.Invoke(ctx, ScimService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScimServiceServer is the server API for ScimService service.
// All implementations must embed UnimplementedScimServiceServer
// for forward compatibility
type ScimServiceServer interface {
	// Handle handles an inbound SCIM GET request from an IDP
	Get(context.Context, *GetRequest) (*ScimResponse, error)
	mustEmbedUnimplementedScimServiceServer()
}

// UnimplementedScimServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScimServiceServer struct {
}

func (UnimplementedScimServiceServer) Get(context.Context, *GetRequest) (*ScimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedScimServiceServer) mustEmbedUnimplementedScimServiceServer() {}

// UnsafeScimServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScimServiceServer will
// result in compilation errors.
type UnsafeScimServiceServer interface {
	mustEmbedUnimplementedScimServiceServer()
}

func RegisterScimServiceServer(s grpc.ServiceRegistrar, srv ScimServiceServer) {
	s.RegisterService(&ScimService_ServiceDesc, srv)
}

func _ScimService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScimServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScimService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScimServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScimService_ServiceDesc is the grpc.ServiceDesc for ScimService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScimService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.scim.v1.ScimService",
	HandlerType: (*ScimServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ScimService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/scim/v1/scim_service.proto",
}
