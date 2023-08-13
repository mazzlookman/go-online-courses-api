package service

import (
	"go-pzn-restful-api/formatter"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type LessonTitleServiceImpl struct {
	repository.LessonTitleRepository
	CourseService
}

func (s *LessonTitleServiceImpl) Update(ltID int, input web.LessonTitleCreateInput) web.LessonTitleResponse {
	findByID, err := s.LessonTitleRepository.FindByID(ltID)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	if input.Title != "" {
		findByID.Title = input.Title
	}
	if input.InOrder != 0 {
		findByID.InOrder = input.InOrder
	}

	return formatter.ToLessonTitleResponse(s.LessonTitleRepository.Update(findByID))
}

func (s *LessonTitleServiceImpl) FindByCourseID(courseID int) []web.LessonTitleResponse {
	lessonTitles, err := s.LessonTitleRepository.FindByCourseID(courseID)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return formatter.ToLessonTitlesResponse(lessonTitles)
}

func (s *LessonTitleServiceImpl) Create(input web.LessonTitleCreateInput) web.LessonTitleResponse {

	lt := domain.LessonTitle{}
	lt.CourseID = input.CourseID
	lt.Title = input.Title
	lt.InOrder = input.InOrder

	authorID := s.CourseService.FindByID(input.CourseID).AuthorID
	if authorID != input.AuthorID {
		panic(helper.NewUnauthorizedError("You're not an author for this courses"))
	}

	lessonTitle := s.LessonTitleRepository.Save(lt)
	return formatter.ToLessonTitleResponse(lessonTitle)
}

func NewLessonTitleService(lessonTitleRepository repository.LessonTitleRepository, courseService CourseService) LessonTitleService {
	return &LessonTitleServiceImpl{
		LessonTitleRepository: lessonTitleRepository,
		CourseService:         courseService,
	}
}
