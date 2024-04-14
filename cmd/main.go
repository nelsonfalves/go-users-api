package main

import "go-quickstart/internal/channels/rest"

func main() {
	server := rest.New()
	server.Start()
}
