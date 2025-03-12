// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package storage

import (
	"encoding/json"

	"github.com/tforce-io/tf-golib/stdx"
)

// Serialize any data object into JSON.
func JsonMarshal[T any](data T) (stdx.Bytes, error) {
	jsonBuffer, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return stdx.Bytes(jsonBuffer), nil
}

// Deserialize JSON into object of type T.
func JsonUnmarshal[T any](jsonBuffer stdx.Bytes) (*T, error) {
	var data T
	err := json.Unmarshal(jsonBuffer, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
