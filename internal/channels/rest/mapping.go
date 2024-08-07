package rest

import (
	"time"

	"github.com/nelsonalves117/go-users-api/internal/canonical"
)

func toCanonical(user userRequest) canonical.User {
	return canonical.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func toResponse(user canonical.User) userResponse {
	return userResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}
