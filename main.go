//go:generate go run main.go gen

package main

import (
	"github.com/rs/zerolog/log"
	"os"
	"wakago/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Err(err).Msg("execute failed")
		os.Exit(1)
	}
}
