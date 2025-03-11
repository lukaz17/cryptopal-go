// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package privmngr

import "testing"

func TestCreateChecksumAddress(t *testing.T) {
	tests := []struct {
		name       string
		address    string
		hasChainID bool
		chainID    uint32
		expected   string
	}{
		// Test cases are referenced from https://eips.ethereum.org/EIPS/eip-1191
		{"eth_mainnet", "dbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB", false, 1, "dbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"rsk_mainnet", "DBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB", true, 30, "DBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB"},
		{"rsk_testnet", "dbF03B407C01E7cd3cbEa99509D93f8dDDc8C6fB", true, 31, "dbF03B407C01E7cd3cbEa99509D93f8dDDc8C6fB"},
		{"eth_mainnet", "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB", false, 1, "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"rsk_mainnet", "0xDBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB", true, 30, "0xDBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB"},
		{"rsk_testnet", "0xdbF03B407C01E7cd3cbEa99509D93f8dDDc8C6fB", true, 31, "0xdbF03B407C01E7cd3cbEa99509D93f8dDDc8C6fB"},
		{"all_uppercase", "DBF03B407C01E7CD3CBEA99509D93F8DDDC8C6FB", false, 1, "dbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"all_uppercase", "0xDBF03B407C01E7CD3CBEA99509D93F8DDDC8C6FB", false, 1, "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"all_uppercase", "DBF03B407C01E7CD3CBEA99509D93F8DDDC8C6FB", true, 30, "DBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB"},
		{"all_uppercase", "0xDBF03B407C01E7CD3CBEA99509D93F8DDDC8C6FB", true, 30, "0xDBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB"},
		{"all_lowercase", "dbf03b407c01e7cd3cbea99509d93f8dddc8c6fb", false, 1, "dbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"all_lowercase", "0xdbf03b407c01e7cd3cbea99509d93f8dddc8c6fb", false, 1, "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"all_uppercase", "dbf03b407c01e7cd3cbea99509d93f8dddc8c6fb", true, 30, "DBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB"},
		{"all_uppercase", "0xdbf03b407c01e7cd3cbea99509d93f8dddc8c6fb", true, 30, "0xDBF03B407c01E7CD3cBea99509D93F8Dddc8C6FB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chainID := &tt.chainID
			if !tt.hasChainID {
				chainID = nil
			}
			checksumAddress, _ := CreateChecksumAddress(tt.address, chainID)
			if checksumAddress != tt.expected {
				t.Errorf("invalid checksum address. expected %s actual %s", tt.expected, checksumAddress)
			}
		})
	}
}
