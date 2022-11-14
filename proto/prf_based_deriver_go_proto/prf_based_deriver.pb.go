// Copyright 2019 Google LLC
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
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: third_party/tink/proto/prf_based_deriver.proto

package prf_based_deriver_go_proto

import (
	tink_go_proto "github.com/tink-crypto/tink-go/proto/tink_go_proto"
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

type PrfBasedDeriverParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DerivedKeyTemplate *tink_go_proto.KeyTemplate `protobuf:"bytes,1,opt,name=derived_key_template,json=derivedKeyTemplate,proto3" json:"derived_key_template,omitempty"`
}

func (x *PrfBasedDeriverParams) Reset() {
	*x = PrfBasedDeriverParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrfBasedDeriverParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrfBasedDeriverParams) ProtoMessage() {}

func (x *PrfBasedDeriverParams) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrfBasedDeriverParams.ProtoReflect.Descriptor instead.
func (*PrfBasedDeriverParams) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_prf_based_deriver_proto_rawDescGZIP(), []int{0}
}

func (x *PrfBasedDeriverParams) GetDerivedKeyTemplate() *tink_go_proto.KeyTemplate {
	if x != nil {
		return x.DerivedKeyTemplate
	}
	return nil
}

type PrfBasedDeriverKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrfKeyTemplate *tink_go_proto.KeyTemplate `protobuf:"bytes,1,opt,name=prf_key_template,json=prfKeyTemplate,proto3" json:"prf_key_template,omitempty"`
	Params         *PrfBasedDeriverParams     `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *PrfBasedDeriverKeyFormat) Reset() {
	*x = PrfBasedDeriverKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrfBasedDeriverKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrfBasedDeriverKeyFormat) ProtoMessage() {}

func (x *PrfBasedDeriverKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrfBasedDeriverKeyFormat.ProtoReflect.Descriptor instead.
func (*PrfBasedDeriverKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_prf_based_deriver_proto_rawDescGZIP(), []int{1}
}

func (x *PrfBasedDeriverKeyFormat) GetPrfKeyTemplate() *tink_go_proto.KeyTemplate {
	if x != nil {
		return x.PrfKeyTemplate
	}
	return nil
}

func (x *PrfBasedDeriverKeyFormat) GetParams() *PrfBasedDeriverParams {
	if x != nil {
		return x.Params
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.PrfBasedDeriverKey
type PrfBasedDeriverKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PrfKey  *tink_go_proto.KeyData `protobuf:"bytes,2,opt,name=prf_key,json=prfKey,proto3" json:"prf_key,omitempty"`
	Params  *PrfBasedDeriverParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *PrfBasedDeriverKey) Reset() {
	*x = PrfBasedDeriverKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrfBasedDeriverKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrfBasedDeriverKey) ProtoMessage() {}

func (x *PrfBasedDeriverKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrfBasedDeriverKey.ProtoReflect.Descriptor instead.
func (*PrfBasedDeriverKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_prf_based_deriver_proto_rawDescGZIP(), []int{2}
}

func (x *PrfBasedDeriverKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PrfBasedDeriverKey) GetPrfKey() *tink_go_proto.KeyData {
	if x != nil {
		return x.PrfKey
	}
	return nil
}

func (x *PrfBasedDeriverKey) GetParams() *PrfBasedDeriverParams {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_third_party_tink_proto_prf_based_deriver_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_prf_based_deriver_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x66, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x64, 0x5f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x1a, 0x21, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x15, 0x50, 0x72, 0x66, 0x42, 0x61,
	0x73, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x51, 0x0a, 0x14, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x12, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x66, 0x42, 0x61, 0x73, 0x65, 0x64,
	0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x49, 0x0a, 0x10, 0x70, 0x72, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e,
	0x4b, 0x65, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x66,
	0x4b, 0x65, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x2e, 0x50, 0x72, 0x66, 0x42, 0x61, 0x73, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xa7,
	0x01, 0x0a, 0x12, 0x50, 0x72, 0x66, 0x42, 0x61, 0x73, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x07, 0x70, 0x72, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x70,
	0x72, 0x66, 0x4b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x66, 0x42, 0x61,
	0x73, 0x65, 0x64, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x59, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x66, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x64, 0x5f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_prf_based_deriver_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_prf_based_deriver_proto_rawDescData = file_third_party_tink_proto_prf_based_deriver_proto_rawDesc
)

func file_third_party_tink_proto_prf_based_deriver_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_prf_based_deriver_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_prf_based_deriver_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_prf_based_deriver_proto_rawDescData)
	})
	return file_third_party_tink_proto_prf_based_deriver_proto_rawDescData
}

var file_third_party_tink_proto_prf_based_deriver_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_third_party_tink_proto_prf_based_deriver_proto_goTypes = []interface{}{
	(*PrfBasedDeriverParams)(nil),     // 0: google.crypto.tink.PrfBasedDeriverParams
	(*PrfBasedDeriverKeyFormat)(nil),  // 1: google.crypto.tink.PrfBasedDeriverKeyFormat
	(*PrfBasedDeriverKey)(nil),        // 2: google.crypto.tink.PrfBasedDeriverKey
	(*tink_go_proto.KeyTemplate)(nil), // 3: google.crypto.tink.KeyTemplate
	(*tink_go_proto.KeyData)(nil),     // 4: google.crypto.tink.KeyData
}
var file_third_party_tink_proto_prf_based_deriver_proto_depIdxs = []int32{
	3, // 0: google.crypto.tink.PrfBasedDeriverParams.derived_key_template:type_name -> google.crypto.tink.KeyTemplate
	3, // 1: google.crypto.tink.PrfBasedDeriverKeyFormat.prf_key_template:type_name -> google.crypto.tink.KeyTemplate
	0, // 2: google.crypto.tink.PrfBasedDeriverKeyFormat.params:type_name -> google.crypto.tink.PrfBasedDeriverParams
	4, // 3: google.crypto.tink.PrfBasedDeriverKey.prf_key:type_name -> google.crypto.tink.KeyData
	0, // 4: google.crypto.tink.PrfBasedDeriverKey.params:type_name -> google.crypto.tink.PrfBasedDeriverParams
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_prf_based_deriver_proto_init() }
func file_third_party_tink_proto_prf_based_deriver_proto_init() {
	if File_third_party_tink_proto_prf_based_deriver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrfBasedDeriverParams); i {
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
		file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrfBasedDeriverKeyFormat); i {
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
		file_third_party_tink_proto_prf_based_deriver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrfBasedDeriverKey); i {
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
			RawDescriptor: file_third_party_tink_proto_prf_based_deriver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_prf_based_deriver_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_prf_based_deriver_proto_depIdxs,
		MessageInfos:      file_third_party_tink_proto_prf_based_deriver_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_prf_based_deriver_proto = out.File
	file_third_party_tink_proto_prf_based_deriver_proto_rawDesc = nil
	file_third_party_tink_proto_prf_based_deriver_proto_goTypes = nil
	file_third_party_tink_proto_prf_based_deriver_proto_depIdxs = nil
}
