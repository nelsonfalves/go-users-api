package main

import (
	"github.com/nelsonalves117/go-users-api/internal/channels/rest"
	"github.com/nelsonalves117/go-users-api/internal/config"

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
