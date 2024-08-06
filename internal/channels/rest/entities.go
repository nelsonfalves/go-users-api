package rest

type userRequest struct {
	Name string `json:"name"`
}

type userResponse struct {
	Id        string `json:"_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
