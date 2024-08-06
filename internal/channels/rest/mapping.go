package rest

import (
	"go-quickstart/internal/canonical"
	"time"
)

func toCanonical(user userRequest) canonical.User {
	return canonical.User{
		Name: user.Name,
	}
}

func toResponse(user canonical.User) userResponse {
	return userResponse{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}
