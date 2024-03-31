package main

import (
	"fmt"
	"go-quickstart/internal/repositories"
)

func main() {
	// repositories.Create(canonical.User{
	// 	Id:   uuid.New().String(),
	// 	Name: "nelson",
	// })

	// repositories.Update("1cac8b79-e25d-440b-8f59-6262a4cdd6ca", canonical.User{
	// 	Id:   "1cac8b79-e25d-440b-8f59-6262a4cdd6ca",
	// 	Name: "José da Silva Pinto Grosso",
	// })

	// result, err := repositories.GetById("1cac8b79-e25d-440b-8f59-6262a4cdd6ca")

	// repositories.Delete("")

	users, err := repositories.Get()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}

}
