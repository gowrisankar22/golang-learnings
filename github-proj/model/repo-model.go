package model

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRepoResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
