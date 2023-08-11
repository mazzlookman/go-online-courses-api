package repository

import "go-pzn-restful-api/model/domain"

type CourseRepository interface {
	Save(course domain.Course) domain.Course
	SaveToCategoryCourse(categoryName string, courseID int) bool
	Update(course domain.Course) domain.Course
	FindByID(courseID int) (domain.Course, error)
	FindBySlug(slug string) (domain.Course, error)
	FindByAuthorID(authorID int) ([]domain.Course, error)
	FindByUserID(userID int) ([]domain.Course, error)
	FindByCategory(categoryName string) ([]domain.Course, error)
	FindAll() ([]domain.Course, error)
	UsersEnrolled(userCourse domain.UserCourse) (domain.UserCourse, error)
	CountUsersEnrolled(courseID int) int
}
