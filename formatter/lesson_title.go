package formatter

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

func ToLessonTitleResponse(title domain.LessonTitle) web.LessonTitleResponse {
	return web.LessonTitleResponse{
		ID:       title.ID,
		CourseID: title.CourseID,
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
