// Copyright 2017 Google Inc.
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

// Definitions for Ed25519 Digital Signature Algorithm.
// See https://ed25519.cr.yp.to/ed25519-20110926.pdf and
// https://tools.ietf.org/html/rfc8032.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: third_party/tink/proto/ed25519.proto

package ed25519_go_proto

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

type Ed25519KeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Ed25519KeyFormat) Reset() {
	*x = Ed25519KeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_ed25519_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ed25519KeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ed25519KeyFormat) ProtoMessage() {}

func (x *Ed25519KeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_ed25519_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ed25519KeyFormat.ProtoReflect.Descriptor instead.
func (*Ed25519KeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ed25519_proto_rawDescGZIP(), []int{0}
}

func (x *Ed25519KeyFormat) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.Ed25519PublicKey
type Ed25519PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// The public key is 32 bytes, encoded according to
	// https://tools.ietf.org/html/rfc8032#section-5.1.2.
	// Required.
	KeyValue []byte `protobuf:"bytes,2,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
}

func (x *Ed25519PublicKey) Reset() {
	*x = Ed25519PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_ed25519_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ed25519PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ed25519PublicKey) ProtoMessage() {}

func (x *Ed25519PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_ed25519_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ed25519PublicKey.ProtoReflect.Descriptor instead.
func (*Ed25519PublicKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ed25519_proto_rawDescGZIP(), []int{1}
}

func (x *Ed25519PublicKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Ed25519PublicKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.Ed25519PrivateKey
type Ed25519PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// The private key is 32 bytes of cryptographically secure random data.
	// See https://tools.ietf.org/html/rfc8032#section-5.1.5.
	// Required.
	KeyValue []byte `protobuf:"bytes,2,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	// The corresponding public key.
	PublicKey *Ed25519PublicKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Ed25519PrivateKey) Reset() {
	*x = Ed25519PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_ed25519_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ed25519PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ed25519PrivateKey) ProtoMessage() {}

func (x *Ed25519PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_ed25519_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ed25519PrivateKey.ProtoReflect.Descriptor instead.
func (*Ed25519PrivateKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_ed25519_proto_rawDescGZIP(), []int{2}
}

func (x *Ed25519PrivateKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Ed25519PrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *Ed25519PrivateKey) GetPublicKey() *Ed25519PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

var File_third_party_tink_proto_ed25519_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_ed25519_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x22, 0x2c, 0x0a, 0x10, 0x45, 0x64,
	0x32, 0x35, 0x35, 0x31, 0x39, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x10, 0x45, 0x64, 0x32, 0x35,
	0x35, 0x31, 0x39, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31,
	0x39, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x52, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_third_party_tink_proto_ed25519_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_ed25519_proto_rawDescData = file_third_party_tink_proto_ed25519_proto_rawDesc
)

func file_third_party_tink_proto_ed25519_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_ed25519_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_ed25519_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_ed25519_proto_rawDescData)
	})
	return file_third_party_tink_proto_ed25519_proto_rawDescData
}

var file_third_party_tink_proto_ed25519_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_third_party_tink_proto_ed25519_proto_goTypes = []interface{}{
	(*Ed25519KeyFormat)(nil),  // 0: google.crypto.tink.Ed25519KeyFormat
	(*Ed25519PublicKey)(nil),  // 1: google.crypto.tink.Ed25519PublicKey
	(*Ed25519PrivateKey)(nil), // 2: google.crypto.tink.Ed25519PrivateKey
}
var file_third_party_tink_proto_ed25519_proto_depIdxs = []int32{
	1, // 0: google.crypto.tink.Ed25519PrivateKey.public_key:type_name -> google.crypto.tink.Ed25519PublicKey
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_ed25519_proto_init() }
func file_third_party_tink_proto_ed25519_proto_init() {
	if File_third_party_tink_proto_ed25519_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_ed25519_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ed25519KeyFormat); i {
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
		file_third_party_tink_proto_ed25519_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ed25519PublicKey); i {
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
		file_third_party_tink_proto_ed25519_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ed25519PrivateKey); i {
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
			RawDescriptor: file_third_party_tink_proto_ed25519_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_ed25519_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_ed25519_proto_depIdxs,
		MessageInfos:      file_third_party_tink_proto_ed25519_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_ed25519_proto = out.File
	file_third_party_tink_proto_ed25519_proto_rawDesc = nil
	file_third_party_tink_proto_ed25519_proto_goTypes = nil
	file_third_party_tink_proto_ed25519_proto_depIdxs = nil
}
