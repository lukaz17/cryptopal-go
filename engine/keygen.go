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

// Struct KeygenModule handles user requests related HDAccounts key file
// generation and modification.
type KeygenModule struct{}

// Create a new key file with random mnemonic and derive first account using
// specified derivationPath, then save it keyFilePath.
func (m *KeygenModule) RandomKey(outputPath, derivationPath string) error {
	mnemonic, entropy, err := keymngr.NewMnemonic()
	if err != nil {
		return err
	}
	multiAcc := &storage.HDAccounts{
		Mnemonic:         mnemonic,
		Entropy:          entropy,
		EthereumAccounts: make(map[string]*storage.HDAccount),
	}

	hdAccount, err := m.updateHDAccount(multiAcc, derivationPath)
	if err != nil {
		return err
	}

	keyFilePath := storage.FilePath(outputPath, hdAccount.AddressStr+".json")
	err = m.writeHDAccounts(multiAcc, keyFilePath)
	if err != nil {
		return err
	}

	return nil
}

// Derive a HDAccount using specified derivationPath and save it to the HDAccounts.
func (m *KeygenModule) updateHDAccount(multiAcc *storage.HDAccounts, derivationPath string) (*storage.HDAccount, error) {
	newAccount, err := keymngr.DeriveEthereumAccountFromMnemonic(multiAcc.Mnemonic, "", derivationPath)
	if err != nil {
		return nil, err
	}
	hdAccount := NewHDAccountFromEthereumAccount(newAccount)
	multiAcc.EthereumAccounts[derivationPath] = NewHDAccountFromEthereumAccount(newAccount)
	return hdAccount, nil
}

// Serialize HDAccounts into JSON and save it to keyFilePath.
func (m *KeygenModule) writeHDAccounts(multiAcc *storage.HDAccounts, keyFilePath string) error {
	fileBuffer, err := storage.JsonMarshal(multiAcc)
	if err != nil {
		return err
	}
	return storage.WriteFile(keyFilePath, fileBuffer)
}
