package service

import "go-pzn-restful-api/model/web"

type LessonTitleService interface {
	Create(title web.LessonTitleCreateInput) web.LessonTitleResponse
	FindByCourseId(courseId int) []web.LessonTitleResponse
	Update(ltId int, input web.LessonTitleCreateInput) web.LessonTitleResponse
}
