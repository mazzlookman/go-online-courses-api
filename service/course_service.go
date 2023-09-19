package service

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

type CourseService interface {
	Create(request web.CourseCreateInput) web.CourseResponse
	UploadBanner(courseId int, pathFile string) bool
	FindById(courseId int) web.CourseResponse
	FindBySlug(slug string) web.CourseBySlugResponse
	FindByAuthorId(authorId int) []web.CourseResponse
	FindByUserId(userId int) []web.CourseResponse
	FindByCategory(categoryName string) []web.CourseResponse
	FindAll() []web.CourseResponse
	UserEnrolled(userId int, courseId int) domain.UserCourse
	FindAllCourseIdByUserId(userId int) []string
}
