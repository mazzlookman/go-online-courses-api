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

func (s *LessonContentServiceImpl) FindByLessonTitleID(ltID int) []web.LessonContentResponse {
	lessonContents, err := s.LessonContentRepository.FindByLessonTitleID(ltID)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToLessonContentsResponse(lessonContents)
}

func (s *LessonContentServiceImpl) Update(lcID int, input web.LessonContentCreateInput) web.LessonContentResponse {
	authorID := s.CourseService.FindByID(input.CourseID).AuthorID
	if authorID != input.AuthorID {
		panic(helper.NewUnauthorizedError("You're not an author for this courses"))
	}

	findByID, err := s.LessonContentRepository.FindByID(lcID)
	oldContent := findByID.Content
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if input.InOrder != 0 {
		findByID.InOrder = input.InOrder
	}

	if input.Content != "" {
		findByID.Content = input.Content
		findByID.Duration = helper.GetLessonContentVideoDuration(input.Content)
		lessonContent, err := s.LessonContentRepository.Update(findByID)
		helper.PanicIfError(err)
		if oldContent != lessonContent.Content {
			os.Remove(oldContent)
		}
		return helper.ToLessonContentResponse(lessonContent)
	}

	lessonContent, err := s.LessonContentRepository.Update(findByID)
	helper.PanicIfError(err)
	return helper.ToLessonContentResponse(lessonContent)
}

func (s *LessonContentServiceImpl) Create(input web.LessonContentCreateInput) web.LessonContentResponse {
	authorID := s.CourseService.FindByID(input.CourseID).AuthorID
	if authorID != input.AuthorID {
		panic(helper.NewUnauthorizedError("You're not an author for this courses"))
	}

	lessonContent := domain.LessonContent{}
	lessonContent.LessonTitleID = input.LessonTitleID
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
