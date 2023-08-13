package web

type LessonContentCreateInput struct {
	AuthorID      int
	CourseID      int
	LessonTitleID int
	Content       string
	InOrder       int `form:"in_order"`
}

type LessonContentResponse struct {
	ID            int    `json:"id"`
	LessonTitleID int    `json:"lesson_title_id"`
	Content       string `json:"content"`
	InOrder       int    `json:"in_order"`
	Duration      string `json:"duration"`
}
