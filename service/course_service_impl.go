package service

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type CourseServiceImpl struct {
	repository.CourseRepository
}

func (s *CourseServiceImpl) FindBySlug(slug string) web.CourseResponse {
	findBySlug, err := s.CourseRepository.FindBySlug(slug)
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Course is not found").Error()))
	}

	return helper.ToCourseResponse(findBySlug)
}

func (s *CourseServiceImpl) FindByID(courseID int) web.CourseResponse {
	findByID, err := s.CourseRepository.FindByID(courseID)
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Course is not found").Error()))
	}

	return helper.ToCourseResponse(findByID)
}

func (s *CourseServiceImpl) Create(request web.CourseInputRequest) web.CourseResponse {
	course := domain.Course{
		AuthorID:    request.AuthorID,
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Perks:       request.Perks,
		Price:       request.Price,
	}

	if course.AuthorID == 0 {
		panic(helper.NewNotFoundError("You're not an author"))
	}

	save := s.CourseRepository.Save(course)

	return helper.ToCourseResponse(save)
}

func NewCourseService(courseRepository repository.CourseRepository) CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository}
}
