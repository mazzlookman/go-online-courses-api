package service

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

type CourseService interface {
	Create(request web.CourseInputRequest) web.CourseResponse
	UploadBanner(courseID int, pathFile string) bool
	FindByID(courseID int) web.CourseResponse
	FindBySlug(slug string) web.CourseBySlugResponse
	FindByAuthorID(authorID int) []web.CourseResponse
	FindByUserID(userID int) []web.CourseResponse
	FindAll() []web.CourseResponse
	UserEnrolled(userID int, courseID int) domain.UserCourse
}
