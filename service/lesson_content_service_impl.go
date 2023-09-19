package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"os"
)

type LessonContentServiceImpl struct {
	repository.LessonContentRepository
	CourseService
}

func (s *LessonContentServiceImpl) FindById(lcId int) web.LessonContentResponse {
	lessonContent, err := s.LessonContentRepository.FindById(lcId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToLessonContentResponse(lessonContent)
}

func (s *LessonContentServiceImpl) FindByLessonTitleId(ltId int) []web.LessonContentResponse {
	lessonContents, err := s.LessonContentRepository.FindByLessonTitleId(ltId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToLessonContentsResponse(lessonContents)
}

func (s *LessonContentServiceImpl) Update(lcId int, input web.LessonContentUpdateInput) web.LessonContentResponse {
	course := s.CourseService.FindById(input.CourseId)
	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this courses"))
	}

	findById, err := s.LessonContentRepository.FindById(lcId)
	oldContent := findById.Content
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if input.InOrder != 0 {
		findById.InOrder = input.InOrder
	}

	if input.Content != "" {
		findById.Content = input.Content
		findById.Duration = helper.GetLessonContentVideoDuration(input.Content)
		lessonContent, err := s.LessonContentRepository.Update(findById)
		helper.PanicIfError(err)
		if oldContent != lessonContent.Content {
			os.Remove(oldContent)
		}
		return helper.ToLessonContentResponse(lessonContent)
	}

	lessonContent, err := s.LessonContentRepository.Update(findById)
	helper.PanicIfError(err)
	return helper.ToLessonContentResponse(lessonContent)
}

func (s *LessonContentServiceImpl) Create(input web.LessonContentCreateInput) web.LessonContentResponse {
	course := s.CourseService.FindById(input.CourseId)
	if course.AuthorId != input.AuthorId {
		panic(helper.NewUnauthorizedError("You're not an author of this courses"))
	}

	lessonContent := domain.LessonContent{}
	lessonContent.LessonTitleId = input.LessonTitleId
	lessonContent.Content = input.Content
	lessonContent.InOrder = input.InOrder
	lessonContent.Duration = helper.GetLessonContentVideoDuration(input.Content)

	content, err := s.LessonContentRepository.Save(lessonContent)
	if err != nil {
		os.Remove(input.Content)
		helper.PanicIfError(err)
	}

	return helper.ToLessonContentResponse(content)
}

func NewLessonContentService(lessonContentRepository repository.LessonContentRepository, courseService CourseService) LessonContentService {
	return &LessonContentServiceImpl{
		LessonContentRepository: lessonContentRepository,
		CourseService:           courseService,
	}
}
