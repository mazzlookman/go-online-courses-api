package repository

import "go-pzn-restful-api/model/domain"

type LessonTitleRepository interface {
	Save(title domain.LessonTitle) domain.LessonTitle
	FindByCourseID(courseID int) ([]domain.LessonTitle, error)
	FindByID(ltID int) (domain.LessonTitle, error)
	Update(title domain.LessonTitle) domain.LessonTitle
}
