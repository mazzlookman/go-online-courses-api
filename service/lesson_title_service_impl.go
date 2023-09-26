package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type LessonTitleServiceImpl struct {
	repository.LessonTitleRepository
	CourseService
}

func (s *LessonTitleServiceImpl) Update(ltId int, input web.LessonTitleCreateInput) web.LessonTitleResponse {
	findById, err := s.LessonTitleRepository.FindById(ltId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	if input.Title != "" {
		findById.Title = input.Title
	}
	if input.InOrder != 0 {
		findById.InOrder = input.InOrder
	}

	return helper.ToLessonTitleResponse(s.LessonTitleRepository.Update(findById))
}

func (s *LessonTitleServiceImpl) FindByCourseId(courseId int) []web.LessonTitleResponse {
	lessonTitles, err := s.LessonTitleRepository.FindByCourseId(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToLessonTitlesResponse(lessonTitles)
}

func (s *LessonTitleServiceImpl) Create(input web.LessonTitleCreateInput) web.LessonTitleResponse {
	lt := domain.LessonTitle{}
	lt.CourseId = input.CourseId
	lt.Title = input.Title
	lt.InOrder = input.InOrder

	course := s.CourseService.FindById(input.CourseId)
	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this course"))
	}

	lessonTitle := s.LessonTitleRepository.Save(lt)
	return helper.ToLessonTitleResponse(lessonTitle)
}

func NewLessonTitleService(lessonTitleRepository repository.LessonTitleRepository, courseService CourseService) LessonTitleService {
	return &LessonTitleServiceImpl{
		LessonTitleRepository: lessonTitleRepository,
		CourseService:         courseService,
	}
}
