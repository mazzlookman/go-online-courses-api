package repository

import "go-pzn-restful-api/model/domain"

type LessonTitleRepository interface {
	Save(title domain.LessonTitle) domain.LessonTitle
	FindByCourseId(courseId int) ([]domain.LessonTitle, error)
	FindById(ltId int) (domain.LessonTitle, error)
	Update(title domain.LessonTitle) domain.LessonTitle
}
