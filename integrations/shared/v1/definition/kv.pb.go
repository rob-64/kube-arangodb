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
// Unless required by applicable law or agreed to in writing, software
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
// source: integrations/shared/v1/definition/kv.proto

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

// Key Values Pairs
type KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key of the pair
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Values of the pair
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValuePair) Reset() {
	*x = KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integrations_shared_v1_definition_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValuePair) ProtoMessage() {}

func (x *KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_integrations_shared_v1_definition_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValuePair.ProtoReflect.Descriptor instead.
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return file_integrations_shared_v1_definition_kv_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValuePair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_integrations_shared_v1_definition_kv_proto protoreflect.FileDescriptor

var file_integrations_shared_v1_definition_kv_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x45, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x67,
	0x6f, 0x64, 0x62, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x61, 0x72, 0x61, 0x6e, 0x67, 0x6f, 0x64,
	0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integrations_shared_v1_definition_kv_proto_rawDescOnce sync.Once
	file_integrations_shared_v1_definition_kv_proto_rawDescData = file_integrations_shared_v1_definition_kv_proto_rawDesc
)

func file_integrations_shared_v1_definition_kv_proto_rawDescGZIP() []byte {
	file_integrations_shared_v1_definition_kv_proto_rawDescOnce.Do(func() {
		file_integrations_shared_v1_definition_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_integrations_shared_v1_definition_kv_proto_rawDescData)
	})
	return file_integrations_shared_v1_definition_kv_proto_rawDescData
}

var file_integrations_shared_v1_definition_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_integrations_shared_v1_definition_kv_proto_goTypes = []interface{}{
	(*KeyValuePair)(nil), // 0: shared.KeyValuePair
}
var file_integrations_shared_v1_definition_kv_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_integrations_shared_v1_definition_kv_proto_init() }
func file_integrations_shared_v1_definition_kv_proto_init() {
	if File_integrations_shared_v1_definition_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integrations_shared_v1_definition_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValuePair); i {
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
			RawDescriptor: file_integrations_shared_v1_definition_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integrations_shared_v1_definition_kv_proto_goTypes,
		DependencyIndexes: file_integrations_shared_v1_definition_kv_proto_depIdxs,
		MessageInfos:      file_integrations_shared_v1_definition_kv_proto_msgTypes,
	}.Build()
	File_integrations_shared_v1_definition_kv_proto = out.File
	file_integrations_shared_v1_definition_kv_proto_rawDesc = nil
	file_integrations_shared_v1_definition_kv_proto_goTypes = nil
	file_integrations_shared_v1_definition_kv_proto_depIdxs = nil
}