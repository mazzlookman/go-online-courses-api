package service

import "go-pzn-restful-api/model/web"

type LessonTitleService interface {
	Create(title web.LessonTitleCreateInput) web.LessonTitleResponse
	FindByCourseID(courseID int) []web.LessonTitleResponse
	Update(ltID int, input web.LessonTitleCreateInput) web.LessonTitleResponse
}
