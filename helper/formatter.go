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
	userResponse.Courses = ToCoursesResponse(user.Courses)

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

func ToCourseResponse(course domain.Course) web.CourseResponse {
	courseResponse := web.CourseResponse{}
	courseResponse.ID = course.ID
	courseResponse.AuthorID = course.AuthorID
	courseResponse.Title = course.Title
	courseResponse.Slug = course.Slug
	courseResponse.Description = course.Description
	courseResponse.Perks = course.Perks
	courseResponse.Price = course.Price
	courseResponse.Banner = course.Banner
	courseResponse.UsersEnrolled = 0
	courseResponse.Author = ToAuthorResponse(course.Author)

	return courseResponse
}

func ToCoursesResponse(courses []domain.Course) []web.CourseResponse {
	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		courseResponse := ToCourseResponse(course)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}
