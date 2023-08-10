package helper

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	userResponse := web.UserResponse{}
	userResponse.ID = user.ID
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Avatar = user.Avatar
	userResponse.Token = user.Token

	return userResponse
}

func ToAuthorResponse(author domain.Author) web.AuthorResponse {
	authorResponse := web.AuthorResponse{}
	authorResponse.ID = author.ID
	authorResponse.Name = author.Name
	authorResponse.Email = author.Email
	authorResponse.Profile = author.Profile
	authorResponse.Avatar = author.Avatar
	authorResponse.Token = author.Token

	return authorResponse
}

func ToCourseResponse(course domain.Course, countUserEnrolled int) web.CourseResponse {
	courseResponse := web.CourseResponse{}
	courseResponse.ID = course.ID
	courseResponse.AuthorID = course.AuthorID
	courseResponse.Title = course.Title
	courseResponse.Slug = course.Slug
	courseResponse.Description = course.Description
	courseResponse.Perks = course.Perks
	courseResponse.Price = course.Price
	courseResponse.Banner = course.Banner
	courseResponse.UsersEnrolled = countUserEnrolled

	return courseResponse
}

func ToCourseBySlugResponse(course domain.Course, countUserEnrolled int) web.CourseBySlugResponse {
	courseResponse := web.CourseBySlugResponse{}
	courseResponse.ID = course.ID
	courseResponse.AuthorID = course.AuthorID
	courseResponse.Title = course.Title
	courseResponse.Slug = course.Slug
	courseResponse.Description = course.Description
	courseResponse.Perks = course.Perks
	courseResponse.Price = course.Price
	courseResponse.Banner = course.Banner
	courseResponse.UsersEnrolled = countUserEnrolled
	courseResponse.Author = ToAuthorResponse(course.Author)

	return courseResponse
}
