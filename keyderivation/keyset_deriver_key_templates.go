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

package keyderivation

import (
	"google.golang.org/protobuf/proto"
	"github.com/tink-crypto/tink-go/keyset"
	prfderpb "github.com/tink-crypto/tink-go/proto/prf_based_deriver_go_proto"
	tinkpb "github.com/tink-crypto/tink-go/proto/tink_go_proto"
)

// CreatePRFBasedKeyTemplate creates a PRF-Based Deriver key template with the
// specified PRF and derived key templates. If either the PRF or derived key
// templates are not supported by the registry, an error is returned.
func CreatePRFBasedKeyTemplate(prfKeyTemplate, derivedKeyTemplate *tinkpb.KeyTemplate) (*tinkpb.KeyTemplate, error) {
	keyFormat := &prfderpb.PrfBasedDeriverKeyFormat{
		PrfKeyTemplate: prfKeyTemplate,
		Params: &prfderpb.PrfBasedDeriverParams{
			DerivedKeyTemplate: derivedKeyTemplate,
		},
	}
	serializedFormat, _ := proto.Marshal(keyFormat)
	template := &tinkpb.KeyTemplate{
		TypeUrl:          prfBasedDeriverTypeURL,
		OutputPrefixType: derivedKeyTemplate.GetOutputPrefixType(),
		Value:            serializedFormat,
	}
	handle, err := keyset.NewHandle(template)
	if err != nil {
		return nil, err
	}
	if _, err = New(handle); err != nil {
		return nil, err
	}
	return template, nil
}
