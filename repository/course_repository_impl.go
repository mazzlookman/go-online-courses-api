package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func (r *CourseRepositoryImpl) FindByCategory(categoryName string) ([]domain.Course, error) {
	category := domain.Category{}
	err := r.db.Find(&category, "name=?", categoryName).Error

	courses := []domain.Course{}
	err = r.db.
		Joins("JOIN category_courses ON category_courses.course_id=courses.id").
		Joins("JOIN categories ON category_courses.category_id=categories.id").
		Where("categories.id=?", category.ID).
		Find(&courses).Error
	if len(courses) == 0 || err != nil {
		return nil, errors.New("Courses is not found")
	}

	return courses, nil
}

func (r *CourseRepositoryImpl) SaveToCategoryCourse(categoryName string, courseID int) bool {
	category := domain.Category{}
	err := r.db.Find(&category, "name=?", categoryName).Error

	categoryCourse := domain.CategoryCourse{}
	categoryCourse.CategoryID = category.ID
	categoryCourse.CourseID = courseID

	err = r.db.Create(&categoryCourse).Error
	helper.PanicIfError(err)

	return true
}

func (r *CourseRepositoryImpl) FindByUserID(userID int) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.
		Joins("JOIN user_courses ON user_courses.course_id=courses.id").
		Joins("JOIN users ON users.id=user_courses.user_id").
		Where("users.id=?", userID).
		Find(&courses).Error
	if len(courses) == 0 || err != nil {
		return nil, errors.New("Courses is not found")
	}
	return courses, nil
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
	if len(courses) == 0 || err != nil {
		return nil, errors.New("Courses is not found")
	}

	return courses, nil
}

func (r *CourseRepositoryImpl) FindByAuthorID(authorID int) ([]domain.Course, error) {
	courses := []domain.Course{}
	err := r.db.Find(&courses, "author_id=?", authorID).Error
	if err != nil || len(courses) == 0 {
		return nil, errors.New("Courses is not found")
	}
	return courses, nil
}

func (r *CourseRepositoryImpl) FindBySlug(slug string) (domain.Course, error) {
	course := domain.Course{}
	err := r.db.Preload("Author").Find(&course, "slug=?", slug).Error
	if course.ID == 0 || err != nil {
		return course, errors.New("Course is not found")
	}

	return course, nil
}

func (r *CourseRepositoryImpl) FindByID(courseID int) (domain.Course, error) {
	course := domain.Course{}
	err := r.db.Find(&course, "id=?", courseID).Error
	if err != nil || course.ID == 0 {
		return course, errors.New("Course is not found")
	}

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
