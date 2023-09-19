package web

type CategoryCreateInput struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
