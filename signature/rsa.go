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

package signature

import (
	"fmt"
	"math/big"

	internal "github.com/tink-crypto/tink-go/v2/internal/signature"
	commonpb "github.com/tink-crypto/tink-go/v2/proto/common_go_proto"
)

func bytesToBigInt(val []byte) *big.Int {
	return new(big.Int).SetBytes(val)
}

func validateRSAPubKeyParams(h commonpb.HashType, modSizeBits int, pubExponent []byte) error {
	if err := internal.HashSafeForSignature(hashName(h)); err != nil {
		return err
	}
	if err := internal.RSAValidModulusSizeInBits(modSizeBits); err != nil {
		return err
	}
	e := bytesToBigInt(pubExponent)
	if !e.IsInt64() {
		return fmt.Errorf("public exponent can't fit in a 64 bit integer")
	}
	return internal.RSAValidPublicExponent(int(e.Int64()))
}

func hashName(h commonpb.HashType) string {
	return commonpb.HashType_name[int32(h)]
}
