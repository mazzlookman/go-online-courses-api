package service

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"os"
	"strings"
	"time"
)

type CourseServiceImpl struct {
	repository.CourseRepository
	TransactionService
}

func (s *CourseServiceImpl) FindAllCourseIdByUserId(userId int) []string {
	return s.CourseRepository.FindAllCourseIdByUserId(userId)
}

func (s *CourseServiceImpl) FindByCategory(categoryName string) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByCategory(categoryName)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindByUserId(userId int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByUserId(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) UploadBanner(courseId int, pathFile string) bool {
	findById, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if findById.Banner != pathFile {
		if findById.Banner == "" {
			return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
		}
		os.Remove(findById.Banner)
		return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
	}

	return updateWhenUploadBanner(findById, pathFile, s.CourseRepository)
}

func (s *CourseServiceImpl) UserEnrolled(userId int, courseId int) domain.UserCourse {
	_, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	userCourse := domain.UserCourse{
		CourseId: courseId,
		UserId:   userId,
	}

	usersEnrolled := s.CourseRepository.UsersEnrolled(userCourse)

	return usersEnrolled
}

func (s *CourseServiceImpl) FindAll() []web.CourseResponse {
	courses, err := s.CourseRepository.FindAll()
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindByAuthorId(authorId int) []web.CourseResponse {
	courses, err := s.CourseRepository.FindByAuthorId(authorId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	coursesResponse := []web.CourseResponse{}
	for _, course := range courses {
		countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(course.Id)
		courseResponse := helper.ToCourseResponse(course, countUsersEnrolled)
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return coursesResponse
}

func (s *CourseServiceImpl) FindBySlug(slug string) web.CourseBySlugResponse {
	findBySlug, err := s.CourseRepository.FindBySlug(slug)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(findBySlug.Id)
	return helper.ToCourseBySlugResponse(findBySlug, countUsersEnrolled)
}

func (s *CourseServiceImpl) FindById(courseId int) web.CourseResponse {
	findById, err := s.CourseRepository.FindById(courseId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	countUsersEnrolled := s.CourseRepository.CountUsersEnrolled(findById.Id)
	return helper.ToCourseResponse(findById, countUsersEnrolled)
}

func (s *CourseServiceImpl) Create(request web.CourseCreateInput) web.CourseResponse {
	course := domain.Course{
		AuthorId:    request.AuthorId,
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Perks:       request.Perks,
		Price:       request.Price,
	}

	if course.AuthorId == 0 {
		panic(helper.NewUnauthorizedError("You're not an author"))
	}

	save := s.CourseRepository.Save(course)
	categoryCourse := s.CourseRepository.SaveToCategoryCourse(strings.ToLower(request.Category), save.Id)
	if !categoryCourse {
		panic(errors.New("Failed to create category for this course"))
	}

	return helper.ToCourseResponse(save, 0)
}

func updateWhenUploadBanner(course domain.Course, pathFile string, courseRepository repository.CourseRepository) bool {
	course.Banner = pathFile
	course.UpdatedAt = time.Now()
	courseRepository.Update(course)
	return true
}

func NewCourseService(courseRepository repository.CourseRepository) CourseService {
	return &CourseServiceImpl{CourseRepository: courseRepository}
}
