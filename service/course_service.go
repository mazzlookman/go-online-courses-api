package service

import "go-pzn-restful-api/model/web"

type CourseService interface {
	Create(request web.CourseInputRequest) web.CourseResponse
	FindByID(courseID int) web.CourseResponse
	FindBySlug(slug string) web.CourseResponse
}
