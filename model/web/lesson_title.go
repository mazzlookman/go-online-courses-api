package web

type LessonTitleCreateInput struct {
	CourseID int
	Title    string `json:"title" binding:"required"`
	InOrder  int    `json:"in_order" binding:"required"`
	AuthorID int
}

type LessonTitleResponse struct {
	ID       int    `json:"id"`
	CourseID int    `json:"course_id"`
	Title    string `json:"title"`
	InOrder  int    `json:"in_order"`
}
