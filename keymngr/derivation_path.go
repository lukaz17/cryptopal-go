// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package privmngr

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/tforce-io/tf-golib/stdx/stringxt"
)

// struct DerivationPart contains information of each part in a DerivationPath.
type DerivationPart struct {
	Index    uint32
	IsHarden bool
}

// Parse a derivationPath into DerivationPart structs.
// Empty string "" and "m" are considered empty path.
func ParseDerivationPath(path string) ([]DerivationPart, error) {
	if stringxt.IsEmptyOrWhitespace(path) {
		return []DerivationPart{}, nil
	}
	isValid, _ := regexp.MatchString(`^m(\/\d{1,10}(')?)*$`, path)
	if !isValid {
		return []DerivationPart{}, errors.New("invalid derivation path")
	}
	parts := strings.Split(path, "/")
	results := make([]DerivationPart, len(parts)-1)
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		isHarden := strings.HasSuffix(part, "'")
		var index uint32
		if isHarden {
			trimmedPart := strings.TrimSuffix(part, "'")
			u64, _ := strconv.ParseUint(trimmedPart, 10, 64)
			index = uint32(u64)
		} else {
			u64, _ := strconv.ParseUint(part, 10, 64)
			index = uint32(u64)
		}
		results[i-1] = DerivationPart{
			Index:    index,
			IsHarden: isHarden,
		}
	}
	return results, nil
}
