// Copyright 2021-2022 Gravitational, Inc
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
// source: teleport/samlidp/v1/samlidp.proto

package samlidpv1

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
	SAMLIdPService_ProcessSAMLIdPRequest_FullMethodName       = "/teleport.samlidp.v1.SAMLIdPService/ProcessSAMLIdPRequest"
	SAMLIdPService_TestSAMLIdPAttributeMapping_FullMethodName = "/teleport.samlidp.v1.SAMLIdPService/TestSAMLIdPAttributeMapping"
)

// SAMLIdPServiceClient is the client API for SAMLIdPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SAMLIdPServiceClient interface {
	// ProcessSAMLIdPRequest processes the SAML auth request.
	ProcessSAMLIdPRequest(ctx context.Context, in *ProcessSAMLIdPRequestRequest, opts ...grpc.CallOption) (*ProcessSAMLIdPRequestResponse, error)
	// TestSAMLIdPAttributeMapping creates test TestSAMLIdPServiceProvider
	TestSAMLIdPAttributeMapping(ctx context.Context, in *TestSAMLIdPAttributeMappingRequest, opts ...grpc.CallOption) (*TestSAMLIdPAttributeMappingResponse, error)
}

type sAMLIdPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSAMLIdPServiceClient(cc grpc.ClientConnInterface) SAMLIdPServiceClient {
	return &sAMLIdPServiceClient{cc}
}

func (c *sAMLIdPServiceClient) ProcessSAMLIdPRequest(ctx context.Context, in *ProcessSAMLIdPRequestRequest, opts ...grpc.CallOption) (*ProcessSAMLIdPRequestResponse, error) {
	out := new(ProcessSAMLIdPRequestResponse)
	err := c.cc.Invoke(ctx, SAMLIdPService_ProcessSAMLIdPRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sAMLIdPServiceClient) TestSAMLIdPAttributeMapping(ctx context.Context, in *TestSAMLIdPAttributeMappingRequest, opts ...grpc.CallOption) (*TestSAMLIdPAttributeMappingResponse, error) {
	out := new(TestSAMLIdPAttributeMappingResponse)
	err := c.cc.Invoke(ctx, SAMLIdPService_TestSAMLIdPAttributeMapping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SAMLIdPServiceServer is the server API for SAMLIdPService service.
// All implementations must embed UnimplementedSAMLIdPServiceServer
// for forward compatibility
type SAMLIdPServiceServer interface {
	// ProcessSAMLIdPRequest processes the SAML auth request.
	ProcessSAMLIdPRequest(context.Context, *ProcessSAMLIdPRequestRequest) (*ProcessSAMLIdPRequestResponse, error)
	// TestSAMLIdPAttributeMapping creates test TestSAMLIdPServiceProvider
	TestSAMLIdPAttributeMapping(context.Context, *TestSAMLIdPAttributeMappingRequest) (*TestSAMLIdPAttributeMappingResponse, error)
	mustEmbedUnimplementedSAMLIdPServiceServer()
}

// UnimplementedSAMLIdPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSAMLIdPServiceServer struct {
}

func (UnimplementedSAMLIdPServiceServer) ProcessSAMLIdPRequest(context.Context, *ProcessSAMLIdPRequestRequest) (*ProcessSAMLIdPRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessSAMLIdPRequest not implemented")
}
func (UnimplementedSAMLIdPServiceServer) TestSAMLIdPAttributeMapping(context.Context, *TestSAMLIdPAttributeMappingRequest) (*TestSAMLIdPAttributeMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestSAMLIdPAttributeMapping not implemented")
}
func (UnimplementedSAMLIdPServiceServer) mustEmbedUnimplementedSAMLIdPServiceServer() {}

// UnsafeSAMLIdPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SAMLIdPServiceServer will
// result in compilation errors.
type UnsafeSAMLIdPServiceServer interface {
	mustEmbedUnimplementedSAMLIdPServiceServer()
}

func RegisterSAMLIdPServiceServer(s grpc.ServiceRegistrar, srv SAMLIdPServiceServer) {
	s.RegisterService(&SAMLIdPService_ServiceDesc, srv)
}

func _SAMLIdPService_ProcessSAMLIdPRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessSAMLIdPRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SAMLIdPServiceServer).ProcessSAMLIdPRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SAMLIdPService_ProcessSAMLIdPRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SAMLIdPServiceServer).ProcessSAMLIdPRequest(ctx, req.(*ProcessSAMLIdPRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SAMLIdPService_TestSAMLIdPAttributeMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSAMLIdPAttributeMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SAMLIdPServiceServer).TestSAMLIdPAttributeMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SAMLIdPService_TestSAMLIdPAttributeMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SAMLIdPServiceServer).TestSAMLIdPAttributeMapping(ctx, req.(*TestSAMLIdPAttributeMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SAMLIdPService_ServiceDesc is the grpc.ServiceDesc for SAMLIdPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SAMLIdPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.samlidp.v1.SAMLIdPService",
	HandlerType: (*SAMLIdPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessSAMLIdPRequest",
			Handler:    _SAMLIdPService_ProcessSAMLIdPRequest_Handler,
		},
		{
			MethodName: "TestSAMLIdPAttributeMapping",
			Handler:    _SAMLIdPService_TestSAMLIdPAttributeMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/samlidp/v1/samlidp.proto",
}
