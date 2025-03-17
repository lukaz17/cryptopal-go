// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package cmd

import (
	"fmt"
	"os"

	"github.com/lukaz17/cryptotool-go/engine"
	"github.com/spf13/cobra"
)

// Execute the program.
func Execute() {
	rootCmd := &cobra.Command{
		Use:   "cryptotool",
		Short: "Command line utility for key management.",
	}
	rootCmd.AddCommand(engine.KeygenCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
