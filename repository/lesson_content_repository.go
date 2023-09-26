package repository

import "go-pzn-restful-api/model/domain"

type LessonContentRepository interface {
	Save(content domain.LessonContent) domain.LessonContent
	Update(content domain.LessonContent) domain.LessonContent
	FindById(lcId int) (domain.LessonContent, error)
	FindByLessonTitleId(ltId int) ([]domain.LessonContent, error)
}
