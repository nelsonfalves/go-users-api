package rest

import "go-quickstart/internal/canonical"

func toCanonical(user userRequest) canonical.User {
	return canonical.User{
		Name: user.Name,
	}
}
func toUserResponse(user canonical.User) userResponse {
	return userResponse{
		Id:   user.Id,
		Name: user.Name,
	}
}
