package main

import (
	"fmt"
	"go-quickstart/internal/canonical"
	"go-quickstart/internal/service"

	"github.com/google/uuid"
)

func main() {
	service := service.New()
	service.Create(canonical.User{
		Id:   uuid.New().String(),
		Name: "teste service",
	})

	// service.Update("1cac8b79-e25d-440b-8f59-6262a4cdd6ca", canonical.User{
	// 	Id:   "1cac8b79-e25d-440b-8f59-6262a4cdd6ca",
	// 	Name: "José da Silva Pinto Grosso",
	// })

	// result, err := service.GetById("1cac8b79-e25d-440b-8f59-6262a4cdd6ca")

	// service.Delete("")

	users, err := service.Get()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}

}
