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
// source: integrations/scheduler/v2/definition/chart.proto

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

// Chart Info
type SchedulerV2ChartInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chart Name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Chart Version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Keeps the Platform details from the output
	Platform *SchedulerV2ChartPlatform `protobuf:"bytes,3,opt,name=platform,proto3,oneof" json:"platform,omitempty"`
}

func (x *SchedulerV2ChartInfo) Reset() {
	*x = SchedulerV2ChartInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2ChartInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2ChartInfo) ProtoMessage() {}

func (x *SchedulerV2ChartInfo) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2ChartInfo.ProtoReflect.Descriptor instead.
func (*SchedulerV2ChartInfo) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_chart_proto_rawDescGZIP(), []int{0}
}

func (x *SchedulerV2ChartInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SchedulerV2ChartInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SchedulerV2ChartInfo) GetPlatform() *SchedulerV2ChartPlatform {
	if x != nil {
		return x.Platform
	}
	return nil
}

// Chart Platform Details
type SchedulerV2ChartPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of the requirements
	Requirements map[string]string `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SchedulerV2ChartPlatform) Reset() {
	*x = SchedulerV2ChartPlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2ChartPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2ChartPlatform) ProtoMessage() {}

func (x *SchedulerV2ChartPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2ChartPlatform.ProtoReflect.Descriptor instead.
func (*SchedulerV2ChartPlatform) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_chart_proto_rawDescGZIP(), []int{1}
}

func (x *SchedulerV2ChartPlatform) GetRequirements() map[string]string {
	if x != nil {
		return x.Requirements
	}
	return nil
}

// SchedulerV2 ListCharts Request
type SchedulerV2ListChartsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum items per batch
	Items *int64 `protobuf:"varint,1,opt,name=items,proto3,oneof" json:"items,omitempty"`
}

func (x *SchedulerV2ListChartsRequest) Reset() {
	*x = SchedulerV2ListChartsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2ListChartsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2ListChartsRequest) ProtoMessage() {}

func (x *SchedulerV2ListChartsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2ListChartsRequest.ProtoReflect.Descriptor instead.
func (*SchedulerV2ListChartsRequest) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_chart_proto_rawDescGZIP(), []int{2}
}

func (x *SchedulerV2ListChartsRequest) GetItems() int64 {
	if x != nil && x.Items != nil {
		return *x.Items
	}
	return 0
}

// SchedulerV2 ListCharts Response
type SchedulerV2ListChartsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of the charts
	Charts []string `protobuf:"bytes,1,rep,name=charts,proto3" json:"charts,omitempty"`
}

func (x *SchedulerV2ListChartsResponse) Reset() {
	*x = SchedulerV2ListChartsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2ListChartsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2ListChartsResponse) ProtoMessage() {}

func (x *SchedulerV2ListChartsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2ListChartsResponse.ProtoReflect.Descriptor instead.
func (*SchedulerV2ListChartsResponse) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_chart_proto_rawDescGZIP(), []int{3}
}

func (x *SchedulerV2ListChartsResponse) GetCharts() []string {
	if x != nil {
		return x.Charts
	}
	return nil
}

// SchedulerV2 GetChart Request
type SchedulerV2GetChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chart Name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SchedulerV2GetChartRequest) Reset() {
	*x = SchedulerV2GetChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2GetChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2GetChartRequest) ProtoMessage() {}

func (x *SchedulerV2GetChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2GetChartRequest.ProtoReflect.Descriptor instead.
func (*SchedulerV2GetChartRequest) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_chart_proto_rawDescGZIP(), []int{4}
}

func (x *SchedulerV2GetChartRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SchedulerV2 GetChart Response
type SchedulerV2GetChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chart Data
	Chart []byte `protobuf:"bytes,1,opt,name=chart,proto3" json:"chart,omitempty"`
	// Chart Info
	Info *SchedulerV2ChartInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SchedulerV2GetChartResponse) Reset() {
	*x = SchedulerV2GetChartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerV2GetChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerV2GetChartResponse) ProtoMessage() {}

func (x *SchedulerV2GetChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_scheduler_v2_definition_chart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerV2GetChartResponse.ProtoReflect.Descriptor instead.
func (*SchedulerV2GetChartResponse) Descriptor() ([]byte, []int) {
	return file_integrations_scheduler_v2_definition_chart_proto_rawDescGZIP(), []int{5}
}

func (x *SchedulerV2GetChartResponse) GetChart() []byte {
	if x != nil {
		return x.Chart
	}
	return nil
}

func (x *SchedulerV2GetChartResponse) GetInfo() *SchedulerV2ChartInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_integrations_scheduler_v2_definition_chart_proto protoreflect.FileDescriptor

var file_integrations_scheduler_v2_definition_chart_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x97, 0x01,
	0x0a, 0x14, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x00, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0xb6, 0x01, 0x0a, 0x18, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x59, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x56, 0x32, 0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x3f, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x43, 0x0a, 0x1c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x37, 0x0a, 0x1d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x56, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x22, 0x30,
	0x0a, 0x1a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x68, 0x0a, 0x1b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x32, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64,
	0x62, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integrations_scheduler_v2_definition_chart_proto_rawDescOnce sync.Once
	file_integrations_scheduler_v2_definition_chart_proto_rawDescData = file_integrations_scheduler_v2_definition_chart_proto_rawDesc
)

func file_integrations_scheduler_v2_definition_chart_proto_rawDescGZIP() []byte {
	file_integrations_scheduler_v2_definition_chart_proto_rawDescOnce.Do(func() {
		file_integrations_scheduler_v2_definition_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_integrations_scheduler_v2_definition_chart_proto_rawDescData)
	})
	return file_integrations_scheduler_v2_definition_chart_proto_rawDescData
}

var file_integrations_scheduler_v2_definition_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_integrations_scheduler_v2_definition_chart_proto_goTypes = []interface{}{
	(*SchedulerV2ChartInfo)(nil),          // 0: scheduler.SchedulerV2ChartInfo
	(*SchedulerV2ChartPlatform)(nil),      // 1: scheduler.SchedulerV2ChartPlatform
	(*SchedulerV2ListChartsRequest)(nil),  // 2: scheduler.SchedulerV2ListChartsRequest
	(*SchedulerV2ListChartsResponse)(nil), // 3: scheduler.SchedulerV2ListChartsResponse
	(*SchedulerV2GetChartRequest)(nil),    // 4: scheduler.SchedulerV2GetChartRequest
	(*SchedulerV2GetChartResponse)(nil),   // 5: scheduler.SchedulerV2GetChartResponse
	nil,                                   // 6: scheduler.SchedulerV2ChartPlatform.RequirementsEntry
}
var file_integrations_scheduler_v2_definition_chart_proto_depIdxs = []int32{
	1, // 0: scheduler.SchedulerV2ChartInfo.platform:type_name -> scheduler.SchedulerV2ChartPlatform
	6, // 1: scheduler.SchedulerV2ChartPlatform.requirements:type_name -> scheduler.SchedulerV2ChartPlatform.RequirementsEntry
	0, // 2: scheduler.SchedulerV2GetChartResponse.info:type_name -> scheduler.SchedulerV2ChartInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_integrations_scheduler_v2_definition_chart_proto_init() }
func file_integrations_scheduler_v2_definition_chart_proto_init() {
	if File_integrations_scheduler_v2_definition_chart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integrations_scheduler_v2_definition_chart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2ChartInfo); i {
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
		file_integrations_scheduler_v2_definition_chart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2ChartPlatform); i {
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
		file_integrations_scheduler_v2_definition_chart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2ListChartsRequest); i {
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
		file_integrations_scheduler_v2_definition_chart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2ListChartsResponse); i {
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
		file_integrations_scheduler_v2_definition_chart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2GetChartRequest); i {
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
		file_integrations_scheduler_v2_definition_chart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerV2GetChartResponse); i {
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
	file_integrations_scheduler_v2_definition_chart_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_integrations_scheduler_v2_definition_chart_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_integrations_scheduler_v2_definition_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integrations_scheduler_v2_definition_chart_proto_goTypes,
		DependencyIndexes: file_integrations_scheduler_v2_definition_chart_proto_depIdxs,
		MessageInfos:      file_integrations_scheduler_v2_definition_chart_proto_msgTypes,
	}.Build()
	File_integrations_scheduler_v2_definition_chart_proto = out.File
	file_integrations_scheduler_v2_definition_chart_proto_rawDesc = nil
	file_integrations_scheduler_v2_definition_chart_proto_goTypes = nil
	file_integrations_scheduler_v2_definition_chart_proto_depIdxs = nil
}