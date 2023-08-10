package repository

import "go-pzn-restful-api/model/domain"

type CourseRepository interface {
	Save(course domain.Course) domain.Course
	FindByID(courseID int) (domain.Course, error)
	FindBySlug(slug string) (domain.Course, error)
}
