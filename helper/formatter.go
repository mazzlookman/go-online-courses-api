package helper

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	userResponse := web.UserResponse{}
	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Avatar = user.Avatar
	userResponse.Token = user.Token

	return userResponse
}

func ToAuthorResponse(author domain.Author) web.AuthorResponse {
	authorResponse := web.AuthorResponse{}
	authorResponse.Id = author.Id
	authorResponse.Name = author.Name
	authorResponse.Email = author.Email
	authorResponse.Profile = author.Profile
	authorResponse.Avatar = author.Avatar
	authorResponse.Token = author.Token

	return authorResponse
}

func ToCourseResponse(course domain.Course, countUserEnrolled int) web.CourseResponse {
	courseResponse := web.CourseResponse{}
	courseResponse.Id = course.Id
	courseResponse.AuthorId = course.AuthorId
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
	courseResponse.Id = course.Id
	courseResponse.AuthorId = course.AuthorId
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

func ToLessonContentResponse(content domain.LessonContent) web.LessonContentResponse {
	return web.LessonContentResponse{
		Id:            content.Id,
		LessonTitleId: content.LessonTitleId,
		Content:       content.Content,
		InOrder:       content.InOrder,
		Duration:      content.Duration,
	}
}

func ToLessonContentsResponse(contents []domain.LessonContent) []web.LessonContentResponse {
	lessonContents := []web.LessonContentResponse{}
	for _, content := range contents {
		lessonContents = append(lessonContents, ToLessonContentResponse(content))
	}

	return lessonContents
}

func ToMidtransTransactionResponse(transaction domain.Transaction, trxID string) web.MidtransTransactionResponse {
	return web.MidtransTransactionResponse{
		Id:         trxID,
		UserId:     transaction.UserId,
		CourseId:   transaction.CourseId,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		PaymentUrl: transaction.PaymentUrl,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToLessonTitleResponse(title domain.LessonTitle) web.LessonTitleResponse {
	return web.LessonTitleResponse{
		Id:       title.Id,
		CourseId: title.CourseId,
		Title:    title.Title,
		InOrder:  title.InOrder,
	}
}

func ToLessonTitlesResponse(titles []domain.LessonTitle) []web.LessonTitleResponse {
	lessonTitlesResponse := []web.LessonTitleResponse{}
	for _, lessonTitle := range titles {
		lessonTitleResponse := ToLessonTitleResponse(lessonTitle)
		lessonTitlesResponse = append(lessonTitlesResponse, lessonTitleResponse)
	}

	return lessonTitlesResponse
}
