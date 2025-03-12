// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package storage

import (
	"github.com/tforce-io/tf-golib/stdx"
)

// Struct HDAccounts (Hierarchical Deterministic Accounts) contains all accounts related to
// a mnemonic following BIP-32 specification.
type HDAccounts struct {
	Mnemonic string     `json:"mnemonic,omitempty"`
	Entropy  stdx.Bytes `json:"entropy,omitempty"`

	EthereumAccounts map[string]HDAccount `json:"ethereumAccounts,omitempty"`
}

// Struct HDAccount (Hierarchical Deterministic Account) contains generic information of
// an account: private key, public key and address.
type HDAccount struct {
	PrivateKey    stdx.Bytes `json:"privateKey,omitempty"`
	PrivateKeyStr string     `json:"privateKeyString,omitempty"`
	PublicKey     stdx.Bytes `json:"publicKey,omitempty"`
	PublicKeyStr  string     `json:"publicKeyString,omitempty"`
	Address       stdx.Bytes `json:"address,omitempty"`
	AddressStr    string     `json:"addressString,omitempty"`
}

// Returns a new HDAccounts from JSON.
func ParseHDAccountsJson(jsonBuffer stdx.Bytes) (*HDAccounts, error) {
	return JsonUnmarshal[HDAccounts](jsonBuffer)
}
