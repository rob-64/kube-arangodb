//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: integrations/scheduler/v2/definition/kubernetes.proto

package definition

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request
type SchedulerV2DiscoverAPIResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes API Group
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// Kubernetes API Version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SchedulerV2DiscoverAPIResourcesRequest) Reset() {
	*x = SchedulerV2DiscoverAPIResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2DiscoverAPIResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2DiscoverAPIResourcesRequest) ProtoMessage() {}

func (x *SchedulerV2DiscoverAPIResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2DiscoverAPIResourcesRequest.ProtoReflect.Descriptor instead.
func (*SchedulerV2DiscoverAPIResourcesRequest) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP(), []int{0}
}

func (x *SchedulerV2DiscoverAPIResourcesRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResourcesRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Response
type SchedulerV2DiscoverAPIResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes API Resources
	Resources []*SchedulerV2DiscoverAPIResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *SchedulerV2DiscoverAPIResourcesResponse) Reset() {
	*x = SchedulerV2DiscoverAPIResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2DiscoverAPIResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2DiscoverAPIResourcesResponse) ProtoMessage() {}

func (x *SchedulerV2DiscoverAPIResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2DiscoverAPIResourcesResponse.ProtoReflect.Descriptor instead.
func (*SchedulerV2DiscoverAPIResourcesResponse) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP(), []int{1}
}

func (x *SchedulerV2DiscoverAPIResourcesResponse) GetResources() []*SchedulerV2DiscoverAPIResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

// Request
type SchedulerV2DiscoverAPIResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes API Group
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// Kubernetes API Version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Kubernetes API Kind
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *SchedulerV2DiscoverAPIResourceRequest) Reset() {
	*x = SchedulerV2DiscoverAPIResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2DiscoverAPIResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2DiscoverAPIResourceRequest) ProtoMessage() {}

func (x *SchedulerV2DiscoverAPIResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2DiscoverAPIResourceRequest.ProtoReflect.Descriptor instead.
func (*SchedulerV2DiscoverAPIResourceRequest) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP(), []int{2}
}

func (x *SchedulerV2DiscoverAPIResourceRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResourceRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResourceRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

// Response
type SchedulerV2DiscoverAPIResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes API Resource
	Resource *SchedulerV2DiscoverAPIResource `protobuf:"bytes,1,opt,name=resource,proto3,oneof" json:"resource,omitempty"`
}

func (x *SchedulerV2DiscoverAPIResourceResponse) Reset() {
	*x = SchedulerV2DiscoverAPIResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2DiscoverAPIResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2DiscoverAPIResourceResponse) ProtoMessage() {}

func (x *SchedulerV2DiscoverAPIResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2DiscoverAPIResourceResponse.ProtoReflect.Descriptor instead.
func (*SchedulerV2DiscoverAPIResourceResponse) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP(), []int{3}
}

func (x *SchedulerV2DiscoverAPIResourceResponse) GetResource() *SchedulerV2DiscoverAPIResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

// Kubernetes API Resource Definition
type SchedulerV2DiscoverAPIResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes API Resource PluralName
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Kubernetes API Resource SingularName
	SingularName string `protobuf:"bytes,2,opt,name=singular_name,json=singularName,proto3" json:"singular_name,omitempty"`
	// Kubernetes API Resource Namespaced flag
	Namespaced bool `protobuf:"varint,3,opt,name=namespaced,proto3" json:"namespaced,omitempty"`
	// Kubernetes API Group
	Group string `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	// Kubernetes API Version
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// Kubernetes API Kind
	Kind string `protobuf:"bytes,6,opt,name=kind,proto3" json:"kind,omitempty"`
	// Kubernetes API Resource Verbs
	Verbs []string `protobuf:"bytes,7,rep,name=verbs,proto3" json:"verbs,omitempty"`
	// Kubernetes API Resource ShortNames
	ShortNames []string `protobuf:"bytes,8,rep,name=short_names,json=shortNames,proto3" json:"short_names,omitempty"`
	// Kubernetes API Resource Categories
	Categories []string `protobuf:"bytes,9,rep,name=categories,proto3" json:"categories,omitempty"`
	// Kubernetes API Resource StorageVersionHash
	StorageVersionHash string `protobuf:"bytes,10,opt,name=storage_version_hash,json=storageVersionHash,proto3" json:"storage_version_hash,omitempty"`
}

func (x *SchedulerV2DiscoverAPIResource) Reset() {
	*x = SchedulerV2DiscoverAPIResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2DiscoverAPIResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2DiscoverAPIResource) ProtoMessage() {}

func (x *SchedulerV2DiscoverAPIResource) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2DiscoverAPIResource.ProtoReflect.Descriptor instead.
func (*SchedulerV2DiscoverAPIResource) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP(), []int{4}
}

func (x *SchedulerV2DiscoverAPIResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResource) GetSingularName() string {
	if x != nil {
		return x.SingularName
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResource) GetNamespaced() bool {
	if x != nil {
		return x.Namespaced
	}
	return false
}

func (x *SchedulerV2DiscoverAPIResource) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *SchedulerV2DiscoverAPIResource) GetVerbs() []string {
	if x != nil {
		return x.Verbs
	}
	return nil
}

func (x *SchedulerV2DiscoverAPIResource) GetShortNames() []string {
	if x != nil {
		return x.ShortNames
	}
	return nil
}

func (x *SchedulerV2DiscoverAPIResource) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *SchedulerV2DiscoverAPIResource) GetStorageVersionHash() string {
	if x != nil {
		return x.StorageVersionHash
	}
	return ""
}

// Request
type SchedulerV2KubernetesGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes API Resources
	Resources []*SchedulerV2ReleaseInfoResource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *SchedulerV2KubernetesGetRequest) Reset() {
	*x = SchedulerV2KubernetesGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2KubernetesGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2KubernetesGetRequest) ProtoMessage() {}

func (x *SchedulerV2KubernetesGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2KubernetesGetRequest.ProtoReflect.Descriptor instead.
func (*SchedulerV2KubernetesGetRequest) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP(), []int{5}
}

func (x *SchedulerV2KubernetesGetRequest) GetResources() []*SchedulerV2ReleaseInfoResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

// Response
type SchedulerV2KubernetesGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes API Objects
	Objects []*SchedulerV2ReleaseInfoResourceObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *SchedulerV2KubernetesGetResponse) Reset() {
	*x = SchedulerV2KubernetesGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2KubernetesGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2KubernetesGetResponse) ProtoMessage() {}

func (x *SchedulerV2KubernetesGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2KubernetesGetResponse.ProtoReflect.Descriptor instead.
func (*SchedulerV2KubernetesGetResponse) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP(), []int{6}
}

func (x *SchedulerV2KubernetesGetResponse) GetObjects() []*SchedulerV2ReleaseInfoResourceObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_integrations_scheduler_v2_definition_kubernetes_proto protoreflect.FileDescriptor

var file_integrations_scheduler_v2_definition_kubernetes_proto_rawDesc = []byte{
	0x0a, 0x35, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x1a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x26, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x56, 0x32, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x72, 0x0a, 0x27, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x22, 0x6b, 0x0a, 0x25, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x56, 0x32, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x81, 0x01, 0x0a, 0x26, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56,
	0x32, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xc6, 0x02, 0x0a, 0x1e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x56, 0x32, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x65, 0x72, 0x62, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x65, 0x72, 0x62, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x6a,
	0x0a, 0x1f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x47, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x20, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescOnce sync.Once
	file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescData = file_integrations_scheduler_v2_definition_kubernetes_proto_rawDesc
)

func file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescGZIP() []byte {
	file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescOnce.Do(func() {
		file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescData = protoimpl.X.CompressGZIP(file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescData)
	})
	return file_integrations_scheduler_v2_definition_kubernetes_proto_rawDescData
}

var file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_integrations_scheduler_v2_definition_kubernetes_proto_goTypes = []interface{}{
	(*SchedulerV2DiscoverAPIResourcesRequest)(nil),  // 0: scheduler.SchedulerV2DiscoverAPIResourcesRequest
	(*SchedulerV2DiscoverAPIResourcesResponse)(nil), // 1: scheduler.SchedulerV2DiscoverAPIResourcesResponse
	(*SchedulerV2DiscoverAPIResourceRequest)(nil),   // 2: scheduler.SchedulerV2DiscoverAPIResourceRequest
	(*SchedulerV2DiscoverAPIResourceResponse)(nil),  // 3: scheduler.SchedulerV2DiscoverAPIResourceResponse
	(*SchedulerV2DiscoverAPIResource)(nil),          // 4: scheduler.SchedulerV2DiscoverAPIResource
	(*SchedulerV2KubernetesGetRequest)(nil),         // 5: scheduler.SchedulerV2KubernetesGetRequest
	(*SchedulerV2KubernetesGetResponse)(nil),        // 6: scheduler.SchedulerV2KubernetesGetResponse
	(*SchedulerV2ReleaseInfoResource)(nil),          // 7: scheduler.SchedulerV2ReleaseInfoResource
	(*SchedulerV2ReleaseInfoResourceObject)(nil),    // 8: scheduler.SchedulerV2ReleaseInfoResourceObject
}
var file_integrations_scheduler_v2_definition_kubernetes_proto_depIdxs = []int32{
	4, // 0: scheduler.SchedulerV2DiscoverAPIResourcesResponse.resources:type_name -> scheduler.SchedulerV2DiscoverAPIResource
	4, // 1: scheduler.SchedulerV2DiscoverAPIResourceResponse.resource:type_name -> scheduler.SchedulerV2DiscoverAPIResource
	7, // 2: scheduler.SchedulerV2KubernetesGetRequest.resources:type_name -> scheduler.SchedulerV2ReleaseInfoResource
	8, // 3: scheduler.SchedulerV2KubernetesGetResponse.objects:type_name -> scheduler.SchedulerV2ReleaseInfoResourceObject
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_integrations_scheduler_v2_definition_kubernetes_proto_init() }
func file_integrations_scheduler_v2_definition_kubernetes_proto_init() {
	if File_integrations_scheduler_v2_definition_kubernetes_proto != nil {
		return
	}
	file_integrations_scheduler_v2_definition_release_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2DiscoverAPIResourcesRequest); i {
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
		file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2DiscoverAPIResourcesResponse); i {
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
		file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2DiscoverAPIResourceRequest); i {
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
		file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2DiscoverAPIResourceResponse); i {
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
		file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2DiscoverAPIResource); i {
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
		file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2KubernetesGetRequest); i {
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
		file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2KubernetesGetResponse); i {
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
	file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_integrations_scheduler_v2_definition_kubernetes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integrations_scheduler_v2_definition_kubernetes_proto_goTypes,
		DependencyIndexes: file_integrations_scheduler_v2_definition_kubernetes_proto_depIdxs,
		MessageInfos:      file_integrations_scheduler_v2_definition_kubernetes_proto_msgTypes,
	}.Build()
	File_integrations_scheduler_v2_definition_kubernetes_proto = out.File
	file_integrations_scheduler_v2_definition_kubernetes_proto_rawDesc = nil
	file_integrations_scheduler_v2_definition_kubernetes_proto_goTypes = nil
	file_integrations_scheduler_v2_definition_kubernetes_proto_depIdxs = nil
}
