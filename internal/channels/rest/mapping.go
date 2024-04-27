package rest

import "go-quickstart/internal/canonical"

func toCanonical(user user) canonical.User {
	return canonical.User{
		Name: user.Name,
	}

}
