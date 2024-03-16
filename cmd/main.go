package main

import (
	"go-quickstart/internal/canonical"
	"go-quickstart/internal/repositories"
)

func main() {
	repositories.Create(canonical.User{
		Id:   "12345",
		Name: "Nelson",
	})

	repositories.Get()
}
