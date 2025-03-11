// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package privmngr

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/lukaz17/cryptotool-go/hasher"
)

// Returns an EVM checksum address using provided chainID following EIP-1191 specification.
// If chainID is nil, the checksum address will follow EIP-55 specification.
// The output address will match the prefix of input address if present.
func CreateChecksumAddress(address string, chainID *uint32) (string, error) {
	isValid, _ := regexp.MatchString(`^(0x)?[0-9a-fA-F]{40}$`, address)
	if !isValid {
		return "", errors.New("invalid input address")
	}
	hasPrefix := strings.HasPrefix(address, "0x")
	sanitizedAddress := strings.ToLower(address)
	if hasPrefix {
		sanitizedAddress = strings.TrimPrefix(sanitizedAddress, "0x")
	}
	hashInput := sanitizedAddress
	if chainID != nil {
		hashInput = fmt.Sprint(*chainID) + "0x" + sanitizedAddress
	}
	hash := hasher.Keccak256([]byte(hashInput)).HexStr()
	var sb strings.Builder
	if hasPrefix {
		sb.WriteString("0x")
	}
	for i, r := range sanitizedAddress {
		if int(hash[i]) >= 48 && int(hash[i]) <= 55 { // hash[i] has value from 0 to 7
			sb.WriteRune(r) // write rune in lowercase
			continue
		}
		if int(r) >= 48 && int(r) <= 57 { // number from 0 to 9 doesn't need to be uppercase
			sb.WriteRune(r)
			continue
		}
		sb.WriteRune(r - 32) // write rune in uppercase
	}
	return sb.String(), nil
}
