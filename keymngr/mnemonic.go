// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package privmngr

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// Returns a random mnemonic along with the entropy used to derive it
// following BIP-39 specification.
func NewMnemonic() (string, []byte, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", nil, err
	}
	return mnemonic, entropy, nil
}

// Returns the private key derived from mnemonic following BIP-32 specification.
// To get the master key, use empty string "" or "m" as derivationPath.
func DeriveKeyFromMnemonic(mnemonic, password, derivationPath string) (*bip32.Key, error) {
	seed := bip39.NewSeed(mnemonic, password)
	masterKey, _ := bip32.NewMasterKey(seed)
	path, err := ParseDerivationPath(derivationPath)
	if err != nil {
		return nil, err
	}
	key := masterKey
	for _, part := range path {
		index := part.Index
		if part.IsHarden {
			index += bip32.FirstHardenedChild
		}
		key, err = key.NewChildKey(index)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}
