package web

type LessonContentCreateInput struct {
	AuthorId      int
	CourseId      int
	LessonTitleId int
	Content       string
	InOrder       int `form:"in_order"`
}

type LessonContentUpdateInput struct {
	AuthorId int
	CourseId int
	Content  string
	InOrder  int `form:"in_order"`
}

type LessonContentResponse struct {
	Id            int    `json:"id"`
	LessonTitleId int    `json:"lesson_title_id"`
	Content       string `json:"content"`
	InOrder       int    `json:"in_order"`
	Duration      string `json:"duration"`
}
