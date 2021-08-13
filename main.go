//go:generate go run main.go gen

package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"wakago/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Err(err).Msg("execute failed")
		os.Exit(1)
	}
}

//
func bad_function() {
var bad_id string
_ = bad_id
}
