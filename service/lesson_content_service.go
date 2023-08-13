package service

import "go-pzn-restful-api/model/web"

type LessonContentService interface {
	Create(input web.LessonContentCreateInput) web.LessonContentResponse
	Update(lcID int, input web.LessonContentCreateInput) web.LessonContentResponse
}
