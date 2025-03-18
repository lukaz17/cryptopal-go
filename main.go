// Copyright (C) 2025 Nguyen Nhat Tung
//
// CryptoTool is licensed under the MIT license.
// You should receive a copy of MIT along with this software.
// If not, see <https://opensource.org/license/mit>

package main

import (
	"os"
	"time"

	"github.com/lukaz17/cryptotool-go/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Entrypoint for the application
func main() {
	consoleWriter := &zerolog.FilteredLevelWriter{
		Writer: zerolog.LevelWriterAdapter{Writer: zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.DateTime}},
		Level:  zerolog.InfoLevel,
	}
	multiWriter := zerolog.MultiLevelWriter(consoleWriter)
	log.Logger = zerolog.New(multiWriter).With().Timestamp().Logger()
	cmd.Execute()
}
