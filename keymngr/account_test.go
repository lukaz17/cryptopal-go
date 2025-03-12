// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package keymngr

import (
	"testing"
)

func TestEthereumAccount_AddressStr(t *testing.T) {
	// Test cases are generated from https://iancoleman.io/bip39
	tests := []struct {
		name           string
		mnemonic       string
		derivationPath string
		address        string
	}{
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/0",
			"0x114A781017506df34B3Ed4C0E6B438889a6Eb3F7"},
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/1",
			"0x3d2F2242a7B705E7865c38a68989A7cde6b6f8Ad"},
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/2",
			"0x4031B9cd5728d9cc26C0819747C00ECaA06d1cbB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, _ := DeriveKeyFromMnemonic(tt.mnemonic, "", tt.derivationPath)
			keypair := NewSecp256k1Keypair(key.Key)
			account := NewEthereumAccount(keypair)
			if account.AddressStr() != tt.address {
				t.Errorf("invalid address. expected %s actual %s", tt.address, account.AddressStr())
			}
		})
	}
}

func TestEthereumAccount_PrivateKeyStr(t *testing.T) {
	// Test cases are generated from https://iancoleman.io/bip39
	tests := []struct {
		name           string
		mnemonic       string
		derivationPath string
		privateKey     string
	}{
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/0",
			"0x6f210f99b79bd5d2d4d93c061aae0351aa00b2b9f1e5f43ffac58ac4a983d355"},
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/1",
			"0x799647d70bb0cbd6fff050defe6299c27771b27362679ded72ec022cdcc0e359"},
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/2",
			"0xceb005d73b46cdc964993d981e9980378770319a8bdd391f9722d22cf162a111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, _ := DeriveKeyFromMnemonic(tt.mnemonic, "", tt.derivationPath)
			keypair := NewSecp256k1Keypair(key.Key)
			account := NewEthereumAccount(keypair)
			if account.PrivateKeyStr() != tt.privateKey {
				t.Errorf("invalid private key. expected %s actual %s", tt.privateKey, account.PrivateKeyStr())
			}
		})
	}
}

func TestEthereumAccount_PublicKeyStr(t *testing.T) {
	// Test cases are generated from https://iancoleman.io/bip39
	tests := []struct {
		name           string
		mnemonic       string
		derivationPath string
		publicKey      string
	}{
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/0",
			"0x036b8387ad386664bb70e326a07a87ae179fbb32705a3c46635bbdd618fa11984f"},
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/1",
			"0x02b6a746c1eeb764e90dec1952a8ea46c24a9101cf1565c663a128aabe5295e512"},
		{"ethereum_account", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/2",
			"0x03b408ca61d93f997156951851aecb997d57869579263d1e9c3c089ecba0c8ca4a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, _ := DeriveKeyFromMnemonic(tt.mnemonic, "", tt.derivationPath)
			keypair := NewSecp256k1Keypair(key.Key)
			account := NewEthereumAccount(keypair)
			if account.PublicKeyStr() != tt.publicKey {
				t.Errorf("invalid public key. expected %s actual %s", tt.publicKey, account.PublicKeyStr())
			}
		})
	}
}
