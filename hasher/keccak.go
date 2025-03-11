// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package hasher

import (
	"github.com/tforce-io/tf-golib/stdx"
	"golang.org/x/crypto/sha3"
)

func Keccak256(data stdx.Bytes) stdx.Bytes {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return stdx.Bytes(hash)
}
