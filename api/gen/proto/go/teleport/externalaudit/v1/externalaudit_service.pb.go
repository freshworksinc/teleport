// Copyright 2023 Gravitational, Inc
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: teleport/externalaudit/v1/externalaudit_service.proto

package externalauditv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetExternalAuditRequest is a request to get an external audit.
type GetExternalAuditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of resource to get.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetExternalAuditRequest) Reset() {
	*x = GetExternalAuditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExternalAuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExternalAuditRequest) ProtoMessage() {}

func (x *GetExternalAuditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExternalAuditRequest.ProtoReflect.Descriptor instead.
func (*GetExternalAuditRequest) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetExternalAuditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GetExternalAuditResponse is a response to get an external audit.
type GetExternalAuditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ExternalAudit is the ExternalAudit to be created.
	ExternalAudit *ExternalAudit `protobuf:"bytes,1,opt,name=external_audit,json=externalAudit,proto3" json:"external_audit,omitempty"`
}

func (x *GetExternalAuditResponse) Reset() {
	*x = GetExternalAuditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExternalAuditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExternalAuditResponse) ProtoMessage() {}

func (x *GetExternalAuditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExternalAuditResponse.ProtoReflect.Descriptor instead.
func (*GetExternalAuditResponse) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetExternalAuditResponse) GetExternalAudit() *ExternalAudit {
	if x != nil {
		return x.ExternalAudit
	}
	return nil
}

// CreateExternalAuditRequest is the request to create the provided external audit.
type CreateExternalAuditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ExternalAudit is the ExternalAudit to be created.
	ExternalAudit *ExternalAudit `protobuf:"bytes,1,opt,name=external_audit,json=externalAudit,proto3" json:"external_audit,omitempty"`
}

func (x *CreateExternalAuditRequest) Reset() {
	*x = CreateExternalAuditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExternalAuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExternalAuditRequest) ProtoMessage() {}

func (x *CreateExternalAuditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExternalAuditRequest.ProtoReflect.Descriptor instead.
func (*CreateExternalAuditRequest) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateExternalAuditRequest) GetExternalAudit() *ExternalAudit {
	if x != nil {
		return x.ExternalAudit
	}
	return nil
}

// CreateExternalAuditRequest is the request to create the provided external audit.
type CreateExternalAuditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ExternalAudit is the ExternalAudit that was created.
	ExternalAudit *ExternalAudit `protobuf:"bytes,1,opt,name=external_audit,json=externalAudit,proto3" json:"external_audit,omitempty"`
}

func (x *CreateExternalAuditResponse) Reset() {
	*x = CreateExternalAuditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExternalAuditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExternalAuditResponse) ProtoMessage() {}

func (x *CreateExternalAuditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExternalAuditResponse.ProtoReflect.Descriptor instead.
func (*CreateExternalAuditResponse) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateExternalAuditResponse) GetExternalAudit() *ExternalAudit {
	if x != nil {
		return x.ExternalAudit
	}
	return nil
}

// DeleteExternalAuditRequest is a request for deleting an external audit.
type DeleteExternalAuditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteExternalAuditRequest) Reset() {
	*x = DeleteExternalAuditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExternalAuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExternalAuditRequest) ProtoMessage() {}

func (x *DeleteExternalAuditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExternalAuditRequest.ProtoReflect.Descriptor instead.
func (*DeleteExternalAuditRequest) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteExternalAuditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SetClusterExternalAuditRequest is a request for setting an cluster external audit status.
type SetClusterExternalAuditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of external audit to set as cluster external audit.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetClusterExternalAuditRequest) Reset() {
	*x = SetClusterExternalAuditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetClusterExternalAuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetClusterExternalAuditRequest) ProtoMessage() {}

func (x *SetClusterExternalAuditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetClusterExternalAuditRequest.ProtoReflect.Descriptor instead.
func (*SetClusterExternalAuditRequest) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{5}
}

func (x *SetClusterExternalAuditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SetClusterExternalAuditResponse is a response for setting an cluster external audit status.
type SetClusterExternalAuditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetClusterExternalAuditResponse) Reset() {
	*x = SetClusterExternalAuditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetClusterExternalAuditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetClusterExternalAuditResponse) ProtoMessage() {}

func (x *SetClusterExternalAuditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetClusterExternalAuditResponse.ProtoReflect.Descriptor instead.
func (*SetClusterExternalAuditResponse) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{6}
}

// DeleteClusterExternalAuditRequest is a request for deleting cluster external audit status.
type DeleteClusterExternalAuditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClusterExternalAuditRequest) Reset() {
	*x = DeleteClusterExternalAuditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterExternalAuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterExternalAuditRequest) ProtoMessage() {}

func (x *DeleteClusterExternalAuditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterExternalAuditRequest.ProtoReflect.Descriptor instead.
func (*DeleteClusterExternalAuditRequest) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{7}
}

// DeleteClusterExternalAuditResponse is a response for deleting cluster external audit status.
type DeleteClusterExternalAuditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClusterExternalAuditResponse) Reset() {
	*x = DeleteClusterExternalAuditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterExternalAuditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterExternalAuditResponse) ProtoMessage() {}

func (x *DeleteClusterExternalAuditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterExternalAuditResponse.ProtoReflect.Descriptor instead.
func (*DeleteClusterExternalAuditResponse) Descriptor() ([]byte, []int) {
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP(), []int{8}
}

var File_teleport_externalaudit_v1_externalaudit_service_proto protoreflect.FileDescriptor

var file_teleport_externalaudit_v1_externalaudit_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x0d, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x22, 0x6d, 0x0a, 0x1a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x22, 0x6e, 0x0a, 0x1b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x22, 0x30, 0x0a, 0x1a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x1e, 0x53,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x21, 0x0a, 0x1f, 0x53, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x21, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x22, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xaf, 0x05, 0x0a, 0x14, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x32, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x35, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x12, 0x35, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x90, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x39,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x12, 0x3c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x5e, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescOnce sync.Once
	file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescData = file_teleport_externalaudit_v1_externalaudit_service_proto_rawDesc
)

func file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescGZIP() []byte {
	file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescOnce.Do(func() {
		file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescData)
	})
	return file_teleport_externalaudit_v1_externalaudit_service_proto_rawDescData
}

var file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_teleport_externalaudit_v1_externalaudit_service_proto_goTypes = []interface{}{
	(*GetExternalAuditRequest)(nil),            // 0: teleport.externalaudit.v1.GetExternalAuditRequest
	(*GetExternalAuditResponse)(nil),           // 1: teleport.externalaudit.v1.GetExternalAuditResponse
	(*CreateExternalAuditRequest)(nil),         // 2: teleport.externalaudit.v1.CreateExternalAuditRequest
	(*CreateExternalAuditResponse)(nil),        // 3: teleport.externalaudit.v1.CreateExternalAuditResponse
	(*DeleteExternalAuditRequest)(nil),         // 4: teleport.externalaudit.v1.DeleteExternalAuditRequest
	(*SetClusterExternalAuditRequest)(nil),     // 5: teleport.externalaudit.v1.SetClusterExternalAuditRequest
	(*SetClusterExternalAuditResponse)(nil),    // 6: teleport.externalaudit.v1.SetClusterExternalAuditResponse
	(*DeleteClusterExternalAuditRequest)(nil),  // 7: teleport.externalaudit.v1.DeleteClusterExternalAuditRequest
	(*DeleteClusterExternalAuditResponse)(nil), // 8: teleport.externalaudit.v1.DeleteClusterExternalAuditResponse
	(*ExternalAudit)(nil),                      // 9: teleport.externalaudit.v1.ExternalAudit
	(*emptypb.Empty)(nil),                      // 10: google.protobuf.Empty
}
var file_teleport_externalaudit_v1_externalaudit_service_proto_depIdxs = []int32{
	9,  // 0: teleport.externalaudit.v1.GetExternalAuditResponse.external_audit:type_name -> teleport.externalaudit.v1.ExternalAudit
	9,  // 1: teleport.externalaudit.v1.CreateExternalAuditRequest.external_audit:type_name -> teleport.externalaudit.v1.ExternalAudit
	9,  // 2: teleport.externalaudit.v1.CreateExternalAuditResponse.external_audit:type_name -> teleport.externalaudit.v1.ExternalAudit
	0,  // 3: teleport.externalaudit.v1.ExternalAuditService.GetExternalAudit:input_type -> teleport.externalaudit.v1.GetExternalAuditRequest
	2,  // 4: teleport.externalaudit.v1.ExternalAuditService.CreateExternalAudit:input_type -> teleport.externalaudit.v1.CreateExternalAuditRequest
	4,  // 5: teleport.externalaudit.v1.ExternalAuditService.DeleteExternalAudit:input_type -> teleport.externalaudit.v1.DeleteExternalAuditRequest
	5,  // 6: teleport.externalaudit.v1.ExternalAuditService.SetClusterExternalAudit:input_type -> teleport.externalaudit.v1.SetClusterExternalAuditRequest
	7,  // 7: teleport.externalaudit.v1.ExternalAuditService.DeleteClusterExternalAudit:input_type -> teleport.externalaudit.v1.DeleteClusterExternalAuditRequest
	1,  // 8: teleport.externalaudit.v1.ExternalAuditService.GetExternalAudit:output_type -> teleport.externalaudit.v1.GetExternalAuditResponse
	3,  // 9: teleport.externalaudit.v1.ExternalAuditService.CreateExternalAudit:output_type -> teleport.externalaudit.v1.CreateExternalAuditResponse
	10, // 10: teleport.externalaudit.v1.ExternalAuditService.DeleteExternalAudit:output_type -> google.protobuf.Empty
	6,  // 11: teleport.externalaudit.v1.ExternalAuditService.SetClusterExternalAudit:output_type -> teleport.externalaudit.v1.SetClusterExternalAuditResponse
	8,  // 12: teleport.externalaudit.v1.ExternalAuditService.DeleteClusterExternalAudit:output_type -> teleport.externalaudit.v1.DeleteClusterExternalAuditResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_teleport_externalaudit_v1_externalaudit_service_proto_init() }
func file_teleport_externalaudit_v1_externalaudit_service_proto_init() {
	if File_teleport_externalaudit_v1_externalaudit_service_proto != nil {
		return
	}
	file_teleport_externalaudit_v1_externalaudit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExternalAuditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExternalAuditResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExternalAuditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExternalAuditResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExternalAuditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetClusterExternalAuditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetClusterExternalAuditResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterExternalAuditRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterExternalAuditResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_externalaudit_v1_externalaudit_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_externalaudit_v1_externalaudit_service_proto_goTypes,
		DependencyIndexes: file_teleport_externalaudit_v1_externalaudit_service_proto_depIdxs,
		MessageInfos:      file_teleport_externalaudit_v1_externalaudit_service_proto_msgTypes,
	}.Build()
	File_teleport_externalaudit_v1_externalaudit_service_proto = out.File
	file_teleport_externalaudit_v1_externalaudit_service_proto_rawDesc = nil
	file_teleport_externalaudit_v1_externalaudit_service_proto_goTypes = nil
	file_teleport_externalaudit_v1_externalaudit_service_proto_depIdxs = nil
}
