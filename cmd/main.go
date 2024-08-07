package main

import (
	"go-quickstart/internal/channels/rest"
	"go-quickstart/internal/config"

	"github.com/rs/zerolog/log"
)

func main() {
	config.Parse()

	server := rest.New()

	err := server.Start()
	if err != nil {
		log.Panic().Err(err).Msg("an error occurred while trying to start the server")
	}
}
