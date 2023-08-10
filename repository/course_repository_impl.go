package repository

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func (r *CourseRepositoryImpl) Update(course domain.Course) domain.Course {
	err := r.db.Save(&course).Error
	helper.PanicIfError(err)

	return course
}

func (r *CourseRepositoryImpl) CountUsersEnrolled(courseID int) int {
	var count int64
	userCourse := domain.UserCourse{}
	err := r.db.Find(&userCourse, "course_id=?", courseID).Count(&count).Error
	helper.PanicIfError(err)

	return int(count)
}

func (r *CourseRepositoryImpl) UsersEnrolled(userCourse domain.UserCourse) (domain.UserCourse, error) {
	err := r.db.Create(&userCourse).Error
	helper.PanicIfError(err)
	return userCourse, nil
}

func (r *CourseRepositoryImpl) FindAll() ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Find(&courses).Error
	helper.PanicIfError(err)

	return courses, nil
}

func (r *CourseRepositoryImpl) FindByAuthorID(authorID int) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Find(&courses, "author_id=?", authorID).Error
	helper.PanicIfError(err)

	return courses, nil
}

func (r *CourseRepositoryImpl) FindBySlug(slug string) (domain.Course, error) {
	course := domain.Course{}
	err := r.db.Preload("Users").Preload("Author").Find(&course, "slug=?", slug).Error
	helper.PanicIfError(err)

	return course, nil
}

func (r *CourseRepositoryImpl) FindByID(courseID int) (domain.Course, error) {
	course := domain.Course{}
	err := r.db.Find(&course, "id=?", courseID).Error
	helper.PanicIfError(err)

	return course, nil
}

func (r *CourseRepositoryImpl) Save(course domain.Course) domain.Course {
	err := r.db.Create(&course).Error
	helper.PanicIfError(err)

	return course
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{db: db}
}
