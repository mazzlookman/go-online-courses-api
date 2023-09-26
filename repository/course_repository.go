package repository

import (
	"go-pzn-restful-api/model/domain"
)

type CourseRepository interface {
	Save(course domain.Course) domain.Course
	SaveToCategoryCourse(categoryName string, courseId int) bool
	Update(course domain.Course) domain.Course
	FindById(courseId int) (domain.Course, error)
	FindBySlug(slug string) (domain.Course, error)
	FindByAuthorId(authorId int) ([]domain.Course, error)
	FindByUserId(userId int) ([]domain.Course, error)
	FindByCategory(categoryName string) ([]domain.Course, error)
	FindAll() ([]domain.Course, error)
	UsersEnrolled(userCourse domain.UserCourse) domain.UserCourse
	CountUsersEnrolled(courseId int) int
	FindAllCourseIdByUserId(userId int) []string
}
