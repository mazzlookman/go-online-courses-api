package web

type AuthorInputRequest struct {
	Name         string `json:"name" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
}

type AuthorResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}
