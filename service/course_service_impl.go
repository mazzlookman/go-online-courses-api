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

func (s *CourseServiceImpl) UserEnrolled(userID int, courseID int) domain.UserCourse {
	userCourse := domain.UserCourse{
		CourseID: courseID,
		UserID:   userID,
	}

	usersEnrolled, err := s.CourseRepository.UsersEnrolled(userCourse)
	helper.PanicIfError(err)

	return usersEnrolled
}

func (s *CourseServiceImpl) FindAll() []web.CourseResponse {
	courses, err := s.CourseRepository.FindAll()
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Courses is not found").Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.ID)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindByAuthorID(authorID int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByAuthorID(authorID)
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Courses is not found").Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.ID)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindBySlug(slug string) web.CourseBySlugResponse {
	findBySlug, err := s.CourseRepository.FindBySlug(slug)
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Course is not found").Error()))
	}

	countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(findBySlug.ID)
	return helper.ToCourseBySlugResponse(findBySlug, countUsersEnrolled)
}

func (s *CourseServiceImpl) FindByID(courseID int) web.CourseResponse {
	findByID, err := s.CourseRepository.FindByID(courseID)
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Course is not found").Error()))
	}
	countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(findByID.ID)
	return helper.ToCourseResponse(findByID, countUsersEnrolled)
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
	return helper.ToCourseResponse(save, 0)
}

func NewCourseService(courseRepository repository.CourseRepository) CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository}
}
