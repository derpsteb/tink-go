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

// Definitions for Cloud Crypto SDK (Tink) library.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: third_party/tink/proto/tink.proto

package tink_go_proto

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

type KeyStatusType int32

const (
	KeyStatusType_UNKNOWN_STATUS KeyStatusType = 0
	KeyStatusType_ENABLED        KeyStatusType = 1 // Can be used for crypto operations.
	KeyStatusType_DISABLED       KeyStatusType = 2 // Cannot be used, but exists and can become ENABLED.
	KeyStatusType_DESTROYED      KeyStatusType = 3 // Key data does not exist in this Keyset any more.
)

// Enum value maps for KeyStatusType.
var (
	KeyStatusType_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "ENABLED",
		2: "DISABLED",
		3: "DESTROYED",
	}
	KeyStatusType_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"ENABLED":        1,
		"DISABLED":       2,
		"DESTROYED":      3,
	}
)

func (x KeyStatusType) Enum() *KeyStatusType {
	p := new(KeyStatusType)
	*p = x
	return p
}

func (x KeyStatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyStatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_tink_proto_enumTypes[0].Descriptor()
}

func (KeyStatusType) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_tink_proto_enumTypes[0]
}

func (x KeyStatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyStatusType.Descriptor instead.
func (KeyStatusType) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{0}
}

// Tink produces and accepts ciphertexts or signatures that consist
// of a prefix and a payload. The payload and its format is determined
// entirely by the primitive, but the prefix has to be one of the following
// 4 types:
//   - Legacy: prefix is 5 bytes, starts with \x00 and followed by a 4-byte
//     key id that is computed from the key material. In addition to
//     that, signature schemes and MACs will add a \x00 byte to the
//     end of the data being signed / MACed when operating on keys
//     with this OutputPrefixType.
//   - Crunchy: prefix is 5 bytes, starts with \x00 and followed by a 4-byte
//     key id that is generated randomly.
//   - Tink  : prefix is 5 bytes, starts with \x01 and followed by 4-byte
//     key id that is generated randomly.
//   - Raw   : prefix is 0 byte, i.e., empty.
type OutputPrefixType int32

const (
	OutputPrefixType_UNKNOWN_PREFIX OutputPrefixType = 0
	OutputPrefixType_TINK           OutputPrefixType = 1
	OutputPrefixType_LEGACY         OutputPrefixType = 2
	OutputPrefixType_RAW            OutputPrefixType = 3
	OutputPrefixType_CRUNCHY        OutputPrefixType = 4
)

// Enum value maps for OutputPrefixType.
var (
	OutputPrefixType_name = map[int32]string{
		0: "UNKNOWN_PREFIX",
		1: "TINK",
		2: "LEGACY",
		3: "RAW",
		4: "CRUNCHY",
	}
	OutputPrefixType_value = map[string]int32{
		"UNKNOWN_PREFIX": 0,
		"TINK":           1,
		"LEGACY":         2,
		"RAW":            3,
		"CRUNCHY":        4,
	}
)

func (x OutputPrefixType) Enum() *OutputPrefixType {
	p := new(OutputPrefixType)
	*p = x
	return p
}

func (x OutputPrefixType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OutputPrefixType) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_tink_proto_enumTypes[1].Descriptor()
}

func (OutputPrefixType) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_tink_proto_enumTypes[1]
}

func (x OutputPrefixType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OutputPrefixType.Descriptor instead.
func (OutputPrefixType) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{1}
}

type KeyData_KeyMaterialType int32

const (
	KeyData_UNKNOWN_KEYMATERIAL KeyData_KeyMaterialType = 0
	KeyData_SYMMETRIC           KeyData_KeyMaterialType = 1
	KeyData_ASYMMETRIC_PRIVATE  KeyData_KeyMaterialType = 2
	KeyData_ASYMMETRIC_PUBLIC   KeyData_KeyMaterialType = 3
	KeyData_REMOTE              KeyData_KeyMaterialType = 4 // points to a remote key, i.e., in a KMS.
)

// Enum value maps for KeyData_KeyMaterialType.
var (
	KeyData_KeyMaterialType_name = map[int32]string{
		0: "UNKNOWN_KEYMATERIAL",
		1: "SYMMETRIC",
		2: "ASYMMETRIC_PRIVATE",
		3: "ASYMMETRIC_PUBLIC",
		4: "REMOTE",
	}
	KeyData_KeyMaterialType_value = map[string]int32{
		"UNKNOWN_KEYMATERIAL": 0,
		"SYMMETRIC":           1,
		"ASYMMETRIC_PRIVATE":  2,
		"ASYMMETRIC_PUBLIC":   3,
		"REMOTE":              4,
	}
)

func (x KeyData_KeyMaterialType) Enum() *KeyData_KeyMaterialType {
	p := new(KeyData_KeyMaterialType)
	*p = x
	return p
}

func (x KeyData_KeyMaterialType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyData_KeyMaterialType) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_tink_proto_enumTypes[2].Descriptor()
}

func (KeyData_KeyMaterialType) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_tink_proto_enumTypes[2]
}

func (x KeyData_KeyMaterialType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyData_KeyMaterialType.Descriptor instead.
func (KeyData_KeyMaterialType) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{1, 0}
}

type KeyTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The type_url of the key type in format
	// type.googleapis.com/packagename.messagename -- see above for details.
	// This is typically the protobuf type URL of the *Key proto. In particular,
	// this is different of the protobuf type URL of the *KeyFormat proto.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Required. The serialized *KeyFormat proto.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Required. The type of prefix used when computing some primitives to
	// identify the ciphertext/signature, etc.
	OutputPrefixType OutputPrefixType `protobuf:"varint,3,opt,name=output_prefix_type,json=outputPrefixType,proto3,enum=google.crypto.tink.OutputPrefixType" json:"output_prefix_type,omitempty"`
}

func (x *KeyTemplate) Reset() {
	*x = KeyTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_tink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyTemplate) ProtoMessage() {}

func (x *KeyTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_tink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyTemplate.ProtoReflect.Descriptor instead.
func (*KeyTemplate) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{0}
}

func (x *KeyTemplate) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *KeyTemplate) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyTemplate) GetOutputPrefixType() OutputPrefixType {
	if x != nil {
		return x.OutputPrefixType
	}
	return OutputPrefixType_UNKNOWN_PREFIX
}

// The actual *Key-proto is wrapped in a KeyData message, which in addition
// to this serialized proto contains also type_url identifying the
// definition of *Key-proto (as in KeyFormat-message), and some extra metadata
// about the type key material.
type KeyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"` // In format type.googleapis.com/packagename.messagename
	// Required.
	// Contains specific serialized *Key proto
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Required.
	KeyMaterialType KeyData_KeyMaterialType `protobuf:"varint,3,opt,name=key_material_type,json=keyMaterialType,proto3,enum=google.crypto.tink.KeyData_KeyMaterialType" json:"key_material_type,omitempty"`
}

func (x *KeyData) Reset() {
	*x = KeyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_tink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyData) ProtoMessage() {}

func (x *KeyData) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_tink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyData.ProtoReflect.Descriptor instead.
func (*KeyData) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{1}
}

func (x *KeyData) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *KeyData) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyData) GetKeyMaterialType() KeyData_KeyMaterialType {
	if x != nil {
		return x.KeyMaterialType
	}
	return KeyData_UNKNOWN_KEYMATERIAL
}

// A Tink user works usually not with single keys, but with keysets,
// to enable key rotation.  The keys in a keyset can belong to different
// implementations/key types, but must all implement the same primitive.
// Any given keyset (and any given key) can be used for one primitive only.
type Keyset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies key used to generate new crypto data (encrypt, sign).
	// Required.
	PrimaryKeyId uint32 `protobuf:"varint,1,opt,name=primary_key_id,json=primaryKeyId,proto3" json:"primary_key_id,omitempty"`
	// Actual keys in the Keyset.
	// Required.
	Key []*Keyset_Key `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty"`
}

func (x *Keyset) Reset() {
	*x = Keyset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_tink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyset) ProtoMessage() {}

func (x *Keyset) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_tink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyset.ProtoReflect.Descriptor instead.
func (*Keyset) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{2}
}

func (x *Keyset) GetPrimaryKeyId() uint32 {
	if x != nil {
		return x.PrimaryKeyId
	}
	return 0
}

func (x *Keyset) GetKey() []*Keyset_Key {
	if x != nil {
		return x.Key
	}
	return nil
}

// Represents a "safe" Keyset that doesn't contain any actual key material,
// thus can be used for logging or monitoring. Most fields are copied from
// Keyset.
type KeysetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// See Keyset.primary_key_id.
	PrimaryKeyId uint32 `protobuf:"varint,1,opt,name=primary_key_id,json=primaryKeyId,proto3" json:"primary_key_id,omitempty"`
	// KeyInfos in the KeysetInfo.
	// Each KeyInfo is corresponding to a Key in the corresponding Keyset.
	KeyInfo []*KeysetInfo_KeyInfo `protobuf:"bytes,2,rep,name=key_info,json=keyInfo,proto3" json:"key_info,omitempty"`
}

func (x *KeysetInfo) Reset() {
	*x = KeysetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_tink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeysetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeysetInfo) ProtoMessage() {}

func (x *KeysetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_tink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeysetInfo.ProtoReflect.Descriptor instead.
func (*KeysetInfo) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{3}
}

func (x *KeysetInfo) GetPrimaryKeyId() uint32 {
	if x != nil {
		return x.PrimaryKeyId
	}
	return 0
}

func (x *KeysetInfo) GetKeyInfo() []*KeysetInfo_KeyInfo {
	if x != nil {
		return x.KeyInfo
	}
	return nil
}

// Represents a keyset that is encrypted with a master key.
type EncryptedKeyset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	EncryptedKeyset []byte `protobuf:"bytes,2,opt,name=encrypted_keyset,json=encryptedKeyset,proto3" json:"encrypted_keyset,omitempty"`
	// Optional.
	KeysetInfo *KeysetInfo `protobuf:"bytes,3,opt,name=keyset_info,json=keysetInfo,proto3" json:"keyset_info,omitempty"`
}

func (x *EncryptedKeyset) Reset() {
	*x = EncryptedKeyset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_tink_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedKeyset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedKeyset) ProtoMessage() {}

func (x *EncryptedKeyset) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_tink_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedKeyset.ProtoReflect.Descriptor instead.
func (*EncryptedKeyset) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptedKeyset) GetEncryptedKeyset() []byte {
	if x != nil {
		return x.EncryptedKeyset
	}
	return nil
}

func (x *EncryptedKeyset) GetKeysetInfo() *KeysetInfo {
	if x != nil {
		return x.KeysetInfo
	}
	return nil
}

type Keyset_Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains the actual, instantiation specific key proto.
	// By convention, each key proto contains a version field.
	KeyData *KeyData      `protobuf:"bytes,1,opt,name=key_data,json=keyData,proto3" json:"key_data,omitempty"`
	Status  KeyStatusType `protobuf:"varint,2,opt,name=status,proto3,enum=google.crypto.tink.KeyStatusType" json:"status,omitempty"`
	// Identifies a key within a keyset, is a part of metadata
	// of a ciphertext/signature.
	KeyId uint32 `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Determines the prefix of the ciphertexts/signatures produced by this key.
	// This value is copied verbatim from the key template.
	OutputPrefixType OutputPrefixType `protobuf:"varint,4,opt,name=output_prefix_type,json=outputPrefixType,proto3,enum=google.crypto.tink.OutputPrefixType" json:"output_prefix_type,omitempty"`
}

func (x *Keyset_Key) Reset() {
	*x = Keyset_Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_tink_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyset_Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyset_Key) ProtoMessage() {}

func (x *Keyset_Key) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_tink_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyset_Key.ProtoReflect.Descriptor instead.
func (*Keyset_Key) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Keyset_Key) GetKeyData() *KeyData {
	if x != nil {
		return x.KeyData
	}
	return nil
}

func (x *Keyset_Key) GetStatus() KeyStatusType {
	if x != nil {
		return x.Status
	}
	return KeyStatusType_UNKNOWN_STATUS
}

func (x *Keyset_Key) GetKeyId() uint32 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *Keyset_Key) GetOutputPrefixType() OutputPrefixType {
	if x != nil {
		return x.OutputPrefixType
	}
	return OutputPrefixType_UNKNOWN_PREFIX
}

type KeysetInfo_KeyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the type url of this key,
	// e.g., type.googleapis.com/google.crypto.tink.HmacKey.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// See Keyset.Key.status.
	Status KeyStatusType `protobuf:"varint,2,opt,name=status,proto3,enum=google.crypto.tink.KeyStatusType" json:"status,omitempty"`
	// See Keyset.Key.key_id.
	KeyId uint32 `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// See Keyset.Key.output_prefix_type.
	OutputPrefixType OutputPrefixType `protobuf:"varint,4,opt,name=output_prefix_type,json=outputPrefixType,proto3,enum=google.crypto.tink.OutputPrefixType" json:"output_prefix_type,omitempty"`
}

func (x *KeysetInfo_KeyInfo) Reset() {
	*x = KeysetInfo_KeyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_tink_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeysetInfo_KeyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeysetInfo_KeyInfo) ProtoMessage() {}

func (x *KeysetInfo_KeyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_tink_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeysetInfo_KeyInfo.ProtoReflect.Descriptor instead.
func (*KeysetInfo_KeyInfo) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_tink_proto_rawDescGZIP(), []int{3, 0}
}

func (x *KeysetInfo_KeyInfo) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *KeysetInfo_KeyInfo) GetStatus() KeyStatusType {
	if x != nil {
		return x.Status
	}
	return KeyStatusType_UNKNOWN_STATUS
}

func (x *KeysetInfo_KeyInfo) GetKeyId() uint32 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *KeysetInfo_KeyInfo) GetOutputPrefixType() OutputPrefixType {
	if x != nil {
		return x.OutputPrefixType
	}
	return OutputPrefixType_UNKNOWN_PREFIX
}

var File_third_party_tink_proto_tink_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_tink_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x52, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8d, 0x02, 0x0a,
	0x07, 0x4b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x02, 0x08, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x57, 0x0a,
	0x11, 0x6b, 0x65, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4b, 0x65,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x6b, 0x65, 0x79, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x74, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4b, 0x45, 0x59, 0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x59, 0x4d, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10,
	0x01, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x53, 0x59, 0x4d, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f,
	0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x53, 0x59,
	0x4d, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x10, 0x04, 0x22, 0xc6, 0x02, 0x0a,
	0x06, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e,
	0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a,
	0xe3, 0x01, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4b,
	0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49,
	0x64, 0x12, 0x52, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc2, 0x02, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4b, 0x65, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0xca, 0x01,
	0x0a, 0x07, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x54, 0x79, 0x70, 0x65, 0x22, 0x7d, 0x0a, 0x0f, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6b,
	0x65, 0x79, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x4d, 0x0a, 0x0d, 0x4b, 0x65, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x53,
	0x54, 0x52, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x52, 0x0a, 0x10, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x49, 0x4e, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45,
	0x47, 0x41, 0x43, 0x59, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41, 0x57, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x52, 0x55, 0x4e, 0x43, 0x48, 0x59, 0x10, 0x04, 0x42, 0x58, 0x0a, 0x1c,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x69, 0x6e, 0x6b, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x06,
	0x54, 0x49, 0x4e, 0x4b, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_tink_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_tink_proto_rawDescData = file_third_party_tink_proto_tink_proto_rawDesc
)

func file_third_party_tink_proto_tink_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_tink_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_tink_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_tink_proto_rawDescData)
	})
	return file_third_party_tink_proto_tink_proto_rawDescData
}

var file_third_party_tink_proto_tink_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_third_party_tink_proto_tink_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_third_party_tink_proto_tink_proto_goTypes = []interface{}{
	(KeyStatusType)(0),           // 0: google.crypto.tink.KeyStatusType
	(OutputPrefixType)(0),        // 1: google.crypto.tink.OutputPrefixType
	(KeyData_KeyMaterialType)(0), // 2: google.crypto.tink.KeyData.KeyMaterialType
	(*KeyTemplate)(nil),          // 3: google.crypto.tink.KeyTemplate
	(*KeyData)(nil),              // 4: google.crypto.tink.KeyData
	(*Keyset)(nil),               // 5: google.crypto.tink.Keyset
	(*KeysetInfo)(nil),           // 6: google.crypto.tink.KeysetInfo
	(*EncryptedKeyset)(nil),      // 7: google.crypto.tink.EncryptedKeyset
	(*Keyset_Key)(nil),           // 8: google.crypto.tink.Keyset.Key
	(*KeysetInfo_KeyInfo)(nil),   // 9: google.crypto.tink.KeysetInfo.KeyInfo
}
var file_third_party_tink_proto_tink_proto_depIdxs = []int32{
	1,  // 0: google.crypto.tink.KeyTemplate.output_prefix_type:type_name -> google.crypto.tink.OutputPrefixType
	2,  // 1: google.crypto.tink.KeyData.key_material_type:type_name -> google.crypto.tink.KeyData.KeyMaterialType
	8,  // 2: google.crypto.tink.Keyset.key:type_name -> google.crypto.tink.Keyset.Key
	9,  // 3: google.crypto.tink.KeysetInfo.key_info:type_name -> google.crypto.tink.KeysetInfo.KeyInfo
	6,  // 4: google.crypto.tink.EncryptedKeyset.keyset_info:type_name -> google.crypto.tink.KeysetInfo
	4,  // 5: google.crypto.tink.Keyset.Key.key_data:type_name -> google.crypto.tink.KeyData
	0,  // 6: google.crypto.tink.Keyset.Key.status:type_name -> google.crypto.tink.KeyStatusType
	1,  // 7: google.crypto.tink.Keyset.Key.output_prefix_type:type_name -> google.crypto.tink.OutputPrefixType
	0,  // 8: google.crypto.tink.KeysetInfo.KeyInfo.status:type_name -> google.crypto.tink.KeyStatusType
	1,  // 9: google.crypto.tink.KeysetInfo.KeyInfo.output_prefix_type:type_name -> google.crypto.tink.OutputPrefixType
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_tink_proto_init() }
func file_third_party_tink_proto_tink_proto_init() {
	if File_third_party_tink_proto_tink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_tink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyTemplate); i {
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
		file_third_party_tink_proto_tink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyData); i {
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
		file_third_party_tink_proto_tink_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyset); i {
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
		file_third_party_tink_proto_tink_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeysetInfo); i {
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
		file_third_party_tink_proto_tink_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedKeyset); i {
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
		file_third_party_tink_proto_tink_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyset_Key); i {
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
		file_third_party_tink_proto_tink_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeysetInfo_KeyInfo); i {
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
			RawDescriptor: file_third_party_tink_proto_tink_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_tink_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_tink_proto_depIdxs,
		EnumInfos:         file_third_party_tink_proto_tink_proto_enumTypes,
		MessageInfos:      file_third_party_tink_proto_tink_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_tink_proto = out.File
	file_third_party_tink_proto_tink_proto_rawDesc = nil
	file_third_party_tink_proto_tink_proto_goTypes = nil
	file_third_party_tink_proto_tink_proto_depIdxs = nil
}
