// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package keymngr

import (
	"github.com/lukaz17/cryptotool-go/hasher"
	"github.com/tforce-io/tf-golib/stdx"
)

// An EthereumAccount derive key information and address based on Secp256k1Keypair
// following the specification defined in EIPs.
type EthereumAccount struct {
	keypair *Secp256k1Keypair
}

// Returns an EthereumAccount from a Secp256k1Keypair.
func NewEthereumAccount(keypair *Secp256k1Keypair) *EthereumAccount {
	return &EthereumAccount{
		keypair: keypair,
	}
}

// Returns corresponding address bytes of the keypair.
func (a *EthereumAccount) Address() stdx.Bytes {
	uPubkey := a.keypair.UncompressPublicKey()
	hash := hasher.Keccak256(uPubkey[1:])
	address := stdx.Bytes(hash[12:])
	return address
}

// Returns corresponding address string of the keypair following EIP-55 specification.
func (a *EthereumAccount) AddressStr() string {
	uPubkey := a.keypair.UncompressPublicKey()
	hash := hasher.Keccak256(uPubkey[1:])
	address := stdx.Bytes(hash[12:])
	hexStr := stdx.NewHex(address, true)
	checksumAddress, _ := CreateChecksumAddress(hexStr.Value(), nil)
	return checksumAddress
}

// Returns corresponding address string of the keypair following EIP-1191 specification.
func (a *EthereumAccount) AddressWithChecksum(chainID uint32) string {
	uPubkey := a.keypair.UncompressPublicKey()
	hash := hasher.Keccak256(uPubkey[1:])
	address := stdx.Bytes(hash[12:])
	hexStr := stdx.NewHex(address, true)
	checksumAddress, _ := CreateChecksumAddress(hexStr.Value(), &chainID)
	return checksumAddress
}

// Returns the derivation path linked to underlying keypair.
func (a *EthereumAccount) DerivationPath() string {
	return a.keypair.derivationPath
}

// Returns the mnemonic linked to underlying keypair.
func (a *EthereumAccount) Mnemonic() string {
	return a.keypair.mnemonic
}

// Returns the private key of underlying keypair.
func (a *EthereumAccount) PrivateKey() stdx.Bytes {
	return a.keypair.PrivateKey()
}

// Returns the private key of underlying keypair in 0x hex string.
func (a *EthereumAccount) PrivateKeyStr() string {
	hexStr := stdx.NewHex(a.keypair.PrivateKey(), true)
	return hexStr.Value()
}

// Returns the compressed public key of underlying keypair.
func (a *EthereumAccount) PublicKey() stdx.Bytes {
	return a.keypair.PublicKey()
}

// Returns the compressed public key of underlying keypair in 0x hex string.
func (a *EthereumAccount) PublicKeyStr() string {
	hexStr := stdx.NewHex(a.keypair.PublicKey(), true)
	return hexStr.Value()
}

// Returns the uncompressed public key of underlying keypair.
func (a *EthereumAccount) UncompressPublicKey() stdx.Bytes {
	return a.keypair.UncompressPublicKey()
}

// Returns the uncompressed public key of underlying keypair in 0x hex string.
func (a *EthereumAccount) UncompressPublicKeyStr() string {
	hexStr := stdx.NewHex(a.keypair.UncompressPublicKey(), true)
	return hexStr.Value()
}
