// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package engine

import (
	"github.com/lukaz17/cryptotool-go/keymngr"
	"github.com/lukaz17/cryptotool-go/storage"
)

// Create a new HDAccount from Ethereum Account.
func NewHDAccountFromEthereumAccount(acc *keymngr.EthereumAccount) *storage.HDAccount {
	return &storage.HDAccount{
		PrivateKey:    acc.PrivateKey(),
		PrivateKeyStr: acc.PrivateKeyStr(),
		PublicKey:     acc.PublicKey(),
		PublicKeyStr:  acc.PublicKeyStr(),
		Address:       acc.Address(),
		AddressStr:    acc.AddressStr(),
	}
}
