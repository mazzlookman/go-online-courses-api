package web

type LessonTitleCreateInput struct {
	CourseId int
	Title    string `json:"title" binding:"required"`
	InOrder  int    `json:"in_order" binding:"required"`
	AuthorId int
}

type LessonTitleResponse struct {
	Id       int    `json:"id"`
	CourseId int    `json:"course_id"`
	Title    string `json:"title"`
	InOrder  int    `json:"in_order"`
}
