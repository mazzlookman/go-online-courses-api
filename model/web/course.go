package web

type CourseCreateInput struct {
	AuthorID    int
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description" binding:"required"`
	Perks       string `json:"perks" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
}

type CourseResponse struct {
	ID            int    `json:"id"`
	AuthorID      int    `json:"author_id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Description   string `json:"description"`
	Perks         string `json:"perks"`
	Price         int    `json:"price"`
	Banner        string `json:"banner"`
	UsersEnrolled int    `json:"users_enrolled"`
}

type CourseBySlugResponse struct {
	ID            int            `json:"id"`
	AuthorID      int            `json:"author_id"`
	Title         string         `json:"title"`
	Slug          string         `json:"slug"`
	Description   string         `json:"description"`
	Perks         string         `json:"perks"`
	Price         int            `json:"price"`
	Banner        string         `json:"banner"`
	UsersEnrolled int            `json:"users_enrolled"`
	Author        AuthorResponse `json:"author"`
}
