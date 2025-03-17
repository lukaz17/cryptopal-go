// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package engine

import (
	"errors"

	"github.com/lukaz17/cryptotool-go/keymngr"
	"github.com/lukaz17/cryptotool-go/storage"
	"github.com/tforce-io/tf-golib/stdx/stringxt"
	"github.com/tyler-smith/go-bip39"
)

// Struct KeygenModule handles user requests related HDAccounts key file
// generation and modification.
type KeygenModule struct{}

// Read a key file and derive new account from mnemonic in the file and specified
// derivationPath, then save it.
func (m *KeygenModule) DeriveKey(keyFilePath, derivationPath string) error {
	multiAcc, err := m.readHDAccounts(keyFilePath)
	if err != nil {
		return err
	}

	if stringxt.IsEmptyOrWhitespace(multiAcc.Mnemonic) {
		return errors.New("mnemonic is not available")
	}
	if !bip39.IsMnemonicValid(multiAcc.Mnemonic) {
		return errors.New("invalid mnemonic")
	}

	_, err = m.updateHDAccount(multiAcc, derivationPath)
	if err != nil {
		return err
	}

	err = m.writeHDAccounts(multiAcc, keyFilePath)
	if err != nil {
		return err
	}

	return nil
}

// Grind for a new key file with vanity address defined by predicate using derivationPath,
// then save it keyFilePath.
func (m *KeygenModule) GenerateKey(outputPath, derivationPath string, predicate stringxt.Predicate) error {
	var multiAcc *storage.HDAccounts
	var hdAccount *storage.HDAccount
	found := false
	for !found {
		mnemonic, entropy, err := keymngr.NewMnemonic()
		if err != nil {
			return err
		}
		multiAcc = &storage.HDAccounts{
			Mnemonic:         mnemonic,
			Entropy:          entropy,
			EthereumAccounts: make(map[string]*storage.HDAccount),
		}
		hdAccount, err = m.updateHDAccount(multiAcc, derivationPath)
		if err != nil {
			return err
		}
		found, err = predicate.Match(hdAccount.AddressStr)
		if err != nil {
			return err
		}
	}

	keyFilePath := storage.FilePath(outputPath, hdAccount.AddressStr+".json")
	err := m.writeHDAccounts(multiAcc, keyFilePath)
	if err != nil {
		return err
	}

	return nil
}

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

// Read a key file and regenerate all accounts inside.
func (m *KeygenModule) RefreshKeys(keyFilePath string) error {
	multiAcc, err := m.readHDAccounts(keyFilePath)
	if err != nil {
		return err
	}

	if stringxt.IsEmptyOrWhitespace(multiAcc.Mnemonic) {
		return errors.New("mnemonic is not available")
	}
	if !bip39.IsMnemonicValid(multiAcc.Mnemonic) {
		return errors.New("invalid mnemonic")
	}

	if multiAcc.EthereumAccounts != nil {
		for derivationPath := range multiAcc.EthereumAccounts {
			_, err = m.updateHDAccount(multiAcc, derivationPath)
			if err != nil {
				return err
			}
		}
	}

	err = m.writeHDAccounts(multiAcc, keyFilePath)
	if err != nil {
		return err
	}

	return nil
}

// Read HDAccounts from keyFilePath and deserialize it from JSON.
func (m *KeygenModule) readHDAccounts(keyFilePath string) (*storage.HDAccounts, error) {
	fileContent, err := storage.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}
	return storage.ParseHDAccountsJson([]byte(fileContent))
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
