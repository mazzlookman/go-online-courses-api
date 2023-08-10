package web

type CategoryCreateInput struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
