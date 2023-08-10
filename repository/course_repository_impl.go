package repository

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
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
