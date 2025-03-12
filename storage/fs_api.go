// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package storage

import (
	"io"
	"os"

	"github.com/tforce-io/tf-golib/stdx"
)

// Read content of specified file into memory.
func ReadFile(path string) (stdx.Bytes, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return stdx.Bytes(data), nil
}

// Create new file or overwrite file content from memory.
func WriteFile(path string, data stdx.Bytes) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
