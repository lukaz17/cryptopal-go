// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package keymngr

import "testing"

func TestDeriveKeyFromMnemonic(t *testing.T) {
	// Test case is generated from https://iancoleman.io/bip39
	tests := []struct {
		name           string
		mnemonic       string
		derivationPath string
		privKey        string
		pubKey         string
	}{
		{"master_key", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "",
			"xprv9s21ZrQH143K3Z9S6rgvgYDyssDzbxzM5ErxtaGAyocVQ33jFXc2adKJtGn6ij8BxpsCRoCYp926pCJwkXN8CabsZdew4f8z5CA4mkxBin7",
			"xpub661MyMwAqRbcG3DuCtDw3gAiRu4V1RiCSTnZgxfnY99UGqNso4vH8RdnjXty9J9938imV9TcFZTC9Ly7qAuDZzhWjPdEATSuSZRYkPboRui"},
		{"master_key", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m",
			"xprv9s21ZrQH143K3Z9S6rgvgYDyssDzbxzM5ErxtaGAyocVQ33jFXc2adKJtGn6ij8BxpsCRoCYp926pCJwkXN8CabsZdew4f8z5CA4mkxBin7",
			"xpub661MyMwAqRbcG3DuCtDw3gAiRu4V1RiCSTnZgxfnY99UGqNso4vH8RdnjXty9J9938imV9TcFZTC9Ly7qAuDZzhWjPdEATSuSZRYkPboRui"},
		{"only_normal", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44/60/0",
			"xprv9z6gcJ9675Gg2AEhr3hF8DexkMv3zRr5HgEztXyiaeu4wTkfzmj5C13hBh3QHFpXFGhy16mKY9vg8CxLbJMhSR4fcYPNYMjZSAVVHTdeCdH",
			"xpub6D631ofywSpyEeKAx5EFVMbhJPkYPtZveuAbgvPL8zS3pG5pYK3KjoNB2x3JjmxKKGWZcatYYuZFGYdKuSY6eQFXpAPQUfX8WUcGHdEpShf"},
		{"only_harden", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'",
			"xprv9ytYSVKzPnEhjEcUYxFHxvSqiC3XwQbvzzZ55YwkkWp9Pua48zXPMbSZ86vf3Qk3eY4KuVsuEEXmP7MkxYZf5qTFsEJNXTGww7BrVZHES7F",
			"xpub6CstqzrtE9nzwigweynJL4PaGDt2LsKnNDUfswMNJrM8GhuCgXqduPm2yQz1cJJuMVkxcWmcxBxr5jK3fjtxuWTbR6vjAdmAgStZf5ff8Kz"},
		{"mixed_harden", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0",
			"xprvA2DNdUBdfQDn3vxHXdtWqvbPfENhNfLfBAbsCrrtSBs87yGtkdrRQsPYDgodeWHdPxzgoBL47bh6yvmT1cxmvm5XQj1AaPvSjigPCJuCHQG",
			"xpub6FCj2yiXVmn5GR2kdfRXD4Y8DGDBn84WYPXU1FGVzXQ6zmc3JBAfxfi24yvyceDw7uaW2esnFtmWXfy3GdgpMUctCBXu3izT3rjWNbdc2FW"},
		{"mixed_harden", "repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat repeat rescue", "m/44'/60'/0'/0/0'",
			"xprvA45KNSUFisXek7VFHyw4qtganWfHoqWJf2rYdywB7y1aVnL22DzYgMW936aLoUoD7q5SZZG1yPUMtb5EQd28w3spJmGxJgoFJ1DW2NjuqeH",
			"xpub6H4fmx19ZF5wxbZiQ1U5D2dKLYVnDJEA2Fn9SNLngJYZNafAZmJoE9pctPsWkgQB135uxFvqvSPR93w9MJ2yBh1QA4dyXmHsTG825McJ5Sf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, _ := DeriveKeyFromMnemonic(tt.mnemonic, "", tt.derivationPath)
			if key.String() != tt.privKey {
				t.Errorf("invalid private key. expected %s actual %s", tt.privKey, key.String())
			}
			if key.PublicKey().String() != tt.pubKey {
				t.Errorf("invalid public key. expected %s actual %s", tt.pubKey, key.PublicKey().String())
			}
		})
	}
}
