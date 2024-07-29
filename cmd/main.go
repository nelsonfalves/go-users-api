package main

import (
	"fmt"
	"go-quickstart/internal/channels/rest"
	"go-quickstart/internal/config"

	"github.com/rs/zerolog/log"
)

func main() {
	config.Parse()

	fmt.Println(config.Get())

	server := rest.New()

	err := server.Start()
	if err != nil {
		log.Panic().Err(err).Msg("An error occurred while trying to start the server")
	}
}
