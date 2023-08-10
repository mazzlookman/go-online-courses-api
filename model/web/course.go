package web

type CourseInputRequest struct {
	AuthorID    int
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Perks       string `json:"perks"`
	Price       int    `json:"price"`
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
