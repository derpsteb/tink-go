// Copyright 2022 Google LLC
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

package keyderivation_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/tink-crypto/tink-go/aead"
	"github.com/tink-crypto/tink-go/core/registry"
	"github.com/tink-crypto/tink-go/keyderivation/internal/streamingprf"
	"github.com/tink-crypto/tink-go/keyderivation"
	"github.com/tink-crypto/tink-go/subtle/random"
	aesgcmpb "github.com/tink-crypto/tink-go/proto/aes_gcm_go_proto"
	prfderpb "github.com/tink-crypto/tink-go/proto/prf_based_deriver_go_proto"
	tinkpb "github.com/tink-crypto/tink-go/proto/tink_go_proto"
)

const (
	prfBasedDeriverKeyVersion = 0
	prfBasedDeriverTypeURL    = "type.googleapis.com/google.crypto.tink.PrfBasedDeriverKey"
)

type namedTemplate struct {
	name     string
	template *tinkpb.KeyTemplate
}

func TestPRFBasedDeriverKeyManagerPrimitive(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}
	prfs := []namedTemplate{
		{"HKDF-SHA256", streamingprf.HKDFSHA256RawKeyTemplate()},
		{"HKDF-SHA512", streamingprf.HKDFSHA512RawKeyTemplate()},
	}
	derivations := []namedTemplate{
		{"AES128GCM", aead.AES128GCMKeyTemplate()},
		{"AES256GCM", aead.AES256GCMKeyTemplate()},
		{"AES256GCMNoPrefix", aead.AES256GCMNoPrefixKeyTemplate()},
	}
	for _, prf := range prfs {
		for _, der := range derivations {
			for _, salt := range [][]byte{nil, []byte("salt")} {
				name := fmt.Sprintf("%s/%s", prf.name, der.name)
				if salt != nil {
					name += "/salt"
				}
				t.Run(name, func(t *testing.T) {
					prfKey, err := registry.NewKeyData(prf.template)
					if err != nil {
						t.Fatalf("registry.NewKeyData() err = %v, want nil", err)
					}
					key := &prfderpb.PrfBasedDeriverKey{
						Version: 0,
						PrfKey:  prfKey,
						Params: &prfderpb.PrfBasedDeriverParams{
							DerivedKeyTemplate: der.template,
						},
					}
					serializedKey, err := proto.Marshal(key)
					if err != nil {
						t.Fatalf("proto.Marshal(%v) err = %v, want nil", key, err)
					}
					p, err := km.Primitive(serializedKey)
					if err != nil {
						t.Fatalf("Primitive() err = %v, want nil", err)
					}
					d, ok := p.(keyderivation.KeysetDeriver)
					if !ok {
						t.Fatal("primitive is not KeysetDeriver")
					}
					if _, err := d.DeriveKeyset(salt); err != nil {
						t.Fatalf("DeriveKeyset() err = %v, want nil", err)
					}
					// We cannot test the derived keyset handle because, at this point, it
					// is filled with placeholder values for the key ID, status, and
					// output prefix type fields.
				})
			}
		}
	}
}

func TestPRFBasedDeriverKeyManagerPrimitiveRejectsIncorrectKeys(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}
	prfKey, err := registry.NewKeyData(streamingprf.HKDFSHA256RawKeyTemplate())
	if err != nil {
		t.Fatalf("registry.NewKeyData() err = %v, want nil", err)
	}
	missingParamsKey := &prfderpb.PrfBasedDeriverKey{
		Version: prfBasedDeriverKeyVersion,
		PrfKey:  prfKey,
	}
	serializedMissingParamsKey, err := proto.Marshal(missingParamsKey)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) err = %v, want nil", serializedMissingParamsKey, err)
	}
	aesGCMKey := &aesgcmpb.AesGcmKey{Version: 0, KeyValue: random.GetRandomBytes(32)}
	serializedAESGCMKey, err := proto.Marshal(aesGCMKey)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) err = %v, want nil", aesGCMKey, err)
	}
	for _, test := range []struct {
		name          string
		serializedKey []byte
	}{
		{"nil key", nil},
		{"zero-length key", []byte{}},
		{"missing params", serializedMissingParamsKey},
		{"wrong key type", serializedAESGCMKey},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := km.Primitive(test.serializedKey); err == nil {
				t.Error("Primitive() err = nil, want non-nil")
			}
		})
	}
}

func TestPRFBasedDeriverKeyManagerPrimitiveRejectsInvalidKeys(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}

	validPRFKey, err := registry.NewKeyData(streamingprf.HKDFSHA256RawKeyTemplate())
	if err != nil {
		t.Fatalf("registry.NewKeyData() err = %v, want nil", err)
	}
	validKey := &prfderpb.PrfBasedDeriverKey{
		Version: 0,
		PrfKey:  validPRFKey,
		Params: &prfderpb.PrfBasedDeriverParams{
			DerivedKeyTemplate: aead.AES128GCMKeyTemplate(),
		},
	}
	serializedValidKey, err := proto.Marshal(validKey)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) err = %v, want nil", validKey, err)
	}
	if _, err := km.Primitive(serializedValidKey); err != nil {
		t.Errorf("Primitive() err = %v, want nil", err)
	}

	invalidPRFKey, err := registry.NewKeyData(aead.AES128GCMKeyTemplate())
	if err != nil {
		t.Fatalf("registry.NewKeyData() err = %v, want nil", err)
	}

	for _, test := range []struct {
		name           string
		version        uint32
		prfKey         *tinkpb.KeyData
		derKeyTemplate *tinkpb.KeyTemplate
	}{
		{
			"invalid version",
			100,
			validKey.GetPrfKey(),
			validKey.GetParams().GetDerivedKeyTemplate(),
		},
		{
			"invalid PRF key",
			validKey.GetVersion(),
			invalidPRFKey,
			validKey.GetParams().GetDerivedKeyTemplate(),
		},
		{
			"invalid derived key template",
			validKey.GetVersion(),
			validKey.GetPrfKey(),
			aead.AES128CTRHMACSHA256KeyTemplate(),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			key := &prfderpb.PrfBasedDeriverKey{
				Version: test.version,
				PrfKey:  test.prfKey,
				Params: &prfderpb.PrfBasedDeriverParams{
					DerivedKeyTemplate: test.derKeyTemplate,
				},
			}
			serializedKey, err := proto.Marshal(key)
			if err != nil {
				t.Fatalf("proto.Marshal(%v) err = %v, want nil", key, err)
			}
			if _, err := km.Primitive(serializedKey); err == nil {
				t.Error("Primitive() err = nil, want non-nil")
			}
		})
	}
}

func TestPRFBasedDeriverKeyManagerNewKey(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}
	prfs := []namedTemplate{
		{"HKDF-SHA256", streamingprf.HKDFSHA256RawKeyTemplate()},
		{"HKDF-SHA512", streamingprf.HKDFSHA512RawKeyTemplate()},
	}
	derivations := []namedTemplate{
		{"AES128GCM", aead.AES128GCMKeyTemplate()},
		{"AES256GCM", aead.AES256GCMKeyTemplate()},
		{"AES256GCMNoPrefix", aead.AES256GCMNoPrefixKeyTemplate()},
	}
	for _, prf := range prfs {
		for _, der := range derivations {
			for _, salt := range [][]byte{nil, []byte("salt")} {
				name := fmt.Sprintf("%s/%s", prf.name, der.name)
				if salt != nil {
					name += "/salt"
				}
				t.Run(name, func(t *testing.T) {
					keyFormat := &prfderpb.PrfBasedDeriverKeyFormat{
						PrfKeyTemplate: prf.template,
						Params: &prfderpb.PrfBasedDeriverParams{
							DerivedKeyTemplate: der.template,
						},
					}
					serializedKeyFormat, err := proto.Marshal(keyFormat)
					if err != nil {
						t.Fatalf("proto.Marshal(%v) err = %v, want nil", keyFormat, err)
					}
					k, err := km.NewKey(serializedKeyFormat)
					if err != nil {
						t.Errorf("NewKey() err = %v, want nil", err)
					}
					key, ok := k.(*prfderpb.PrfBasedDeriverKey)
					if !ok {
						t.Fatal("key is not PrfBasedDeriverKey")
					}
					if key.GetVersion() != prfBasedDeriverKeyVersion {
						t.Errorf("GetVersion() = %d, want 0", key.GetVersion())
					}
					prfKeyData := key.GetPrfKey()
					if got, want := prfKeyData.GetTypeUrl(), prf.template.GetTypeUrl(); got != want {
						t.Errorf("GetTypeUrl() = %q, want %q", got, want)
					}
					if got, want := prfKeyData.GetKeyMaterialType(), tinkpb.KeyData_SYMMETRIC; got != want {
						t.Errorf("GetKeyMaterialType() = %s, want %s", got, want)
					}
					if diff := cmp.Diff(key.GetParams().GetDerivedKeyTemplate(), der.template, protocmp.Transform()); diff != "" {
						t.Errorf("GetDerivedKeyTemplate() diff = %s", diff)
					}
				})
			}
		}
	}
}

func TestPRFBasedDeriverKeyManagerNewKeyData(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}
	prfs := []namedTemplate{
		{"HKDF-SHA256", streamingprf.HKDFSHA256RawKeyTemplate()},
		{"HKDF-SHA512", streamingprf.HKDFSHA512RawKeyTemplate()},
	}
	derivations := []namedTemplate{
		{"AES128GCM", aead.AES128GCMKeyTemplate()},
		{"AES256GCM", aead.AES256GCMKeyTemplate()},
		{"AES256GCMNoPrefix", aead.AES256GCMNoPrefixKeyTemplate()},
	}
	for _, prf := range prfs {
		for _, der := range derivations {
			for _, salt := range [][]byte{nil, []byte("salt")} {
				name := fmt.Sprintf("%s/%s", prf.name, der.name)
				if salt != nil {
					name += "/salt"
				}
				t.Run(name, func(t *testing.T) {
					keyFormat := &prfderpb.PrfBasedDeriverKeyFormat{
						PrfKeyTemplate: prf.template,
						Params: &prfderpb.PrfBasedDeriverParams{
							DerivedKeyTemplate: der.template,
						},
					}
					serializedKeyFormat, err := proto.Marshal(keyFormat)
					if err != nil {
						t.Fatalf("proto.Marshal(%v) err = %v, want nil", keyFormat, err)
					}
					keyData, err := km.NewKeyData(serializedKeyFormat)
					if err != nil {
						t.Errorf("NewKeyData() err = %v, want nil", err)
					}
					if keyData.GetTypeUrl() != prfBasedDeriverTypeURL {
						t.Errorf("GetTypeUrl() = %s, want %s", keyData.GetTypeUrl(), prfBasedDeriverTypeURL)
					}
					if keyData.GetKeyMaterialType() != tinkpb.KeyData_SYMMETRIC {
						t.Errorf("GetKeyMaterialType() = %s, want %s", keyData.GetKeyMaterialType(), tinkpb.KeyData_SYMMETRIC)
					}
					key := &prfderpb.PrfBasedDeriverKey{}
					if err := proto.Unmarshal(keyData.GetValue(), key); err != nil {
						t.Fatalf("proto.Unmarshal() err = %v, want nil", err)
					}
					if key.GetVersion() != prfBasedDeriverKeyVersion {
						t.Errorf("GetVersion() = %d, want %d", key.GetVersion(), prfBasedDeriverKeyVersion)
					}
					prfKeyData := key.GetPrfKey()
					if got, want := prfKeyData.GetTypeUrl(), prf.template.GetTypeUrl(); got != want {
						t.Errorf("GetTypeUrl() = %q, want %q", got, want)
					}
					if got, want := prfKeyData.GetKeyMaterialType(), tinkpb.KeyData_SYMMETRIC; got != want {
						t.Errorf("GetKeyMaterialType() = %s, want %s", got, want)
					}
					if diff := cmp.Diff(key.GetParams().GetDerivedKeyTemplate(), der.template, protocmp.Transform()); diff != "" {
						t.Errorf("GetDerivedKeyTemplate() diff = %s", diff)
					}
				})
			}
		}
	}
}

func TestPRFBasedDeriverKeyManagerNewKeyAndNewKeyDataRejectsIncorrectKeyFormats(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}
	missingParamsKeyFormat := &prfderpb.PrfBasedDeriverKeyFormat{
		PrfKeyTemplate: streamingprf.HKDFSHA256RawKeyTemplate(),
	}
	serializedMissingParamsKeyFormat, err := proto.Marshal(missingParamsKeyFormat)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) err = %v, want nil", missingParamsKeyFormat, err)
	}
	aesGCMKeyFormat := &aesgcmpb.AesGcmKeyFormat{KeySize: 32, Version: 0}
	serializedAESGCMKeyFormat, err := proto.Marshal(aesGCMKeyFormat)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) err = %v, want nil", aesGCMKeyFormat, err)
	}
	for _, test := range []struct {
		name                string
		serializedKeyFormat []byte
	}{
		{"nil key", nil},
		{"zero-length key", []byte{}},
		{"missing params", serializedMissingParamsKeyFormat},
		{"wrong key type", serializedAESGCMKeyFormat},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := km.NewKey(test.serializedKeyFormat); err == nil {
				t.Error("NewKey() err = nil, want non-nil")
			}
			if _, err := km.NewKeyData(test.serializedKeyFormat); err == nil {
				t.Error("NewKeyData() err = nil, want non-nil")
			}
		})
	}
}

func TestPRFBasedDeriverKeyManagerNewKeyAndNewKeyDataRejectsInvalidKeyFormats(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}

	validKeyFormat := &prfderpb.PrfBasedDeriverKeyFormat{
		PrfKeyTemplate: streamingprf.HKDFSHA256RawKeyTemplate(),
		Params: &prfderpb.PrfBasedDeriverParams{
			DerivedKeyTemplate: aead.AES128GCMKeyTemplate(),
		},
	}
	serializedValidKeyFormat, err := proto.Marshal(validKeyFormat)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) err = %v, want nil", validKeyFormat, err)
	}
	if _, err := km.NewKey(serializedValidKeyFormat); err != nil {
		t.Errorf("Primitive() err = %v, want nil", err)
	}

	for _, test := range []struct {
		name           string
		prfKeyTemplate *tinkpb.KeyTemplate
		derKeyTemplate *tinkpb.KeyTemplate
	}{
		{
			"invalid PRF key template",
			aead.AES128GCMKeyTemplate(),
			validKeyFormat.GetParams().GetDerivedKeyTemplate(),
		},
		{
			"invalid derived key template",
			validKeyFormat.GetPrfKeyTemplate(),
			aead.AES128CTRHMACSHA256KeyTemplate(),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			keyFormat := &prfderpb.PrfBasedDeriverKeyFormat{
				PrfKeyTemplate: test.prfKeyTemplate,
				Params: &prfderpb.PrfBasedDeriverParams{
					DerivedKeyTemplate: test.derKeyTemplate,
				},
			}
			serializedKeyFormat, err := proto.Marshal(keyFormat)
			if err != nil {
				t.Fatalf("proto.Marshal(%v) err = %v, want nil", keyFormat, err)
			}
			if _, err := km.NewKey(serializedKeyFormat); err == nil {
				t.Error("NewKey() err = nil, want non-nil")
			}
			if _, err := km.NewKeyData(serializedKeyFormat); err == nil {
				t.Error("NewKeyData() err = nil, want non-nil")
			}
		})
	}
}

func TestPRFBasedDeriverKeyManagerDoesSupport(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}
	if !km.DoesSupport(prfBasedDeriverTypeURL) {
		t.Errorf("DoesSupport(%q) = false, want true", prfBasedDeriverTypeURL)
	}
	if unsupported := "unsupported.key.type"; km.DoesSupport(unsupported) {
		t.Errorf("DoesSupport(%q) = true, want false", unsupported)
	}
}

func TestPRFBasedDeriverKeyManagerTypeURL(t *testing.T) {
	km, err := registry.GetKeyManager(prfBasedDeriverTypeURL)
	if err != nil {
		t.Fatalf("GetKeyManager(%q) err = %v, want nil", prfBasedDeriverTypeURL, err)
	}
	if km.TypeURL() != prfBasedDeriverTypeURL {
		t.Errorf("TypeURL() = %q, want %q", km.TypeURL(), prfBasedDeriverTypeURL)
	}
}
