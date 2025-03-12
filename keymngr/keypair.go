// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package keymngr

import (
	"bytes"
	"math/big"

	btcutil "github.com/FactomProject/btcutilecc"
	"github.com/tforce-io/tf-golib/stdx"
)

const (
	Secp256k1PointLength = 32
)

// Secp256k1Keypair struct implements key management based on Secp256k1 elliptic curve.
type Secp256k1Keypair struct {
	privateKey     stdx.Bytes
	mnemonic       string
	derivationPath string
}

// Returns a new Secp256k1Keypair from a private key.
// If you want to log the mnemonic and derivationPath associated with this private key,
// please use NewSecp256k1KeypairWithMetadata instead.
func NewSecp256k1Keypair(privateKey stdx.Bytes) *Secp256k1Keypair {
	return &Secp256k1Keypair{
		privateKey: privateKey,
	}
}

// Returns a new Secp256k1Keypair from a private key along with the mnemonic and derivationPath.
// This function does not validate if the mnemonic, derivationPath and privateKey having relationship
// with each others.
func NewSecp256k1KeypairWithMetadata(privateKey stdx.Bytes, mnemonic, derivationPath string) *Secp256k1Keypair {
	return &Secp256k1Keypair{
		privateKey:     privateKey,
		mnemonic:       mnemonic,
		derivationPath: derivationPath,
	}
}

// Returns the derivation path linked to this keypair.
// This struct does not validate if the mnemonic, derivationPath and privateKey having relationship
// with each others.
func (p *Secp256k1Keypair) DerivationPath() string {
	return p.derivationPath
}

// Returns the mnemonic linked to this keypair.
// This struct does not validate if the mnemonic, derivationPath and privateKey having relationship
// with each others.
func (p *Secp256k1Keypair) Mnemonic() string {
	return p.mnemonic
}

// Returns the private key in Bitcoin format.
func (p *Secp256k1Keypair) PrivateKey() stdx.Bytes {
	return p.privateKey
}

// Returns the compressed public key in Bitcoin format.
func (p *Secp256k1Keypair) PublicKey() stdx.Bytes {
	curve := btcutil.Secp256k1()
	pubkey := compressPublicKey(curve.ScalarBaseMult(p.privateKey))
	return stdx.Bytes(pubkey)
}

// Returns the uncompressed public key in Bitcoin format.
func (p *Secp256k1Keypair) UncompressPublicKey() stdx.Bytes {
	curve := btcutil.Secp256k1()
	pubkey := uncompressPublicKey(curve.ScalarBaseMult(p.privateKey))
	return stdx.Bytes(pubkey)
}

// Turns a point with coordinate (x, y) into compressed format: header + x value.
// Header value will be 0x2 for even y value, 0x3 for odd y value.
func compressPublicKey(x *big.Int, y *big.Int) []byte {
	var key bytes.Buffer

	// Write header; 0x2 for even y value; 0x3 for odd
	key.WriteByte(byte(0x2) + byte(y.Bit(0)))

	// Write X coord; Pad the key so x is aligned with the LSB. Pad size point length - xBytes size
	xBytes := x.Bytes()
	for i := 0; i < Secp256k1PointLength-len(xBytes); i++ {
		key.WriteByte(0x0)
	}
	key.Write(xBytes)

	return key.Bytes()
}

// Turns a point with coordinate (x, y) into uncompressed format: header + x value + y value.
// Header value will always be 0x4.
func uncompressPublicKey(x *big.Int, y *big.Int) []byte {
	var key bytes.Buffer

	// Write header; 0x4 for uncompress
	key.WriteByte(byte(0x4))

	// Write X coord; Pad the key so x is aligned with the LSB. Pad size point length - xBytes size
	xBytes := x.Bytes()
	for i := 0; i < Secp256k1PointLength-len(xBytes); i++ {
		key.WriteByte(0x0)
	}
	key.Write(xBytes)

	// Write Y coord; Pad the key so y is aligned with the LSB. Pad size point length - yBytes size
	yBytes := y.Bytes()
	for i := 0; i < Secp256k1PointLength-len(yBytes); i++ {
		key.WriteByte(0x0)
	}
	key.Write(yBytes)

	return key.Bytes()
}
