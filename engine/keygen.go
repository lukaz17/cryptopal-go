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
	"github.com/spf13/cobra"
	"github.com/tforce-io/tf-golib/stdx/stringxt"
	"github.com/tyler-smith/go-bip39"
)

// Struct KeygenModule handles user requests related HDAccounts key file
// generation and modification.
type KeygenModule struct{}

// Read a key file and derive new account from mnemonic in the file and specified
// derivationPath, then save it.
func (m *KeygenModule) Add(keyFilePath, derivationPath string) error {
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
func (m *KeygenModule) Grind(outputPath, derivationPath string, predicate *stringxt.Predicate) error {
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
func (m *KeygenModule) New(outputPath, derivationPath string) error {
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
func (m *KeygenModule) Refresh(keyFilePath string) error {
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

// Define Cobra Command for Keygen module.
func KeygenCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "keygen",
		Short: "HD Accounts generation and modification",
	}

	addCmd := &cobra.Command{
		Use:  "add <key file path>",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			flags := ParseKeygenFlags(cmd)
			m := &KeygenModule{}
			m.Add(args[0], flags.DerivationPath)
		},
	}
	addCmd.Flags().StringP("ckd", "p", "", "Child key perivation path. Must start with 'm'.")
	rootCmd.AddCommand(addCmd)

	grindCmd := &cobra.Command{
		Use:  "grind <output path>",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			flags := ParseKeygenFlags(cmd)
			m := &KeygenModule{}
			predicate := &stringxt.Predicate{
				Prefix: flags.AccountPrefix,
				Suffix: flags.AccountSuffix,
				Regexp: flags.AccountRegexp,
			}
			m.Grind(args[0], flags.DerivationPath, predicate)
		},
	}
	grindCmd.Flags().StringP("ckd", "p", "", "Child key perivation path. Must start with 'm'.")
	grindCmd.Flags().Uint16P("count", "n", 1, "Number of accounts to search for.")
	grindCmd.Flags().String("prefix", "", "Prefix of the output address. Case sensitive.")
	grindCmd.Flags().String("suffix", "", "Suffix of the output address. Case sensitive.")
	grindCmd.Flags().String("regexp", "", "Regular expression to match the output address. Prefix and suffix flags will be ignored.")
	rootCmd.AddCommand(grindCmd)

	newCmd := &cobra.Command{
		Use:  "new <output path>",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			flags := ParseKeygenFlags(cmd)
			m := &KeygenModule{}
			m.New(args[0], flags.DerivationPath)
		},
	}
	newCmd.Flags().StringP("ckd", "p", "", "Child key perivation path. Must start with 'm'.")
	rootCmd.AddCommand(newCmd)

	refreshCmd := &cobra.Command{
		Use:  "refresh <key file path>",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			m := &KeygenModule{}
			m.Refresh(args[0])
		},
	}
	rootCmd.AddCommand(refreshCmd)

	return rootCmd
}

// Struct KeygenFlags contains all flags used by Keygen module.
type KeygenFlags struct {
	AccountPrefix  string
	AccountRegexp  string
	AccountSuffix  string
	DerivationPath string
	KeyCount       uint16
}

// Extract all flags from a Cobra Command.
func ParseKeygenFlags(cmd *cobra.Command) *KeygenFlags {
	cdk, _ := cmd.Flags().GetString("cdk")
	count, _ := cmd.Flags().GetUint16("count")
	prefix, _ := cmd.Flags().GetString("prefix")
	regexp, _ := cmd.Flags().GetString("regexp")
	suffix, _ := cmd.Flags().GetString("suffix")

	return &KeygenFlags{
		AccountPrefix:  prefix,
		AccountRegexp:  regexp,
		AccountSuffix:  suffix,
		DerivationPath: cdk,
		KeyCount:       count,
	}
}
