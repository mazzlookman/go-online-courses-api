package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type LessonTitleRepositoryImpl struct {
	db *gorm.DB
}

func (r *LessonTitleRepositoryImpl) FindById(ltId int) (domain.LessonTitle, error) {
	lessonTitle := domain.LessonTitle{}
	err := r.db.Find(&lessonTitle, "id=?", ltId).Error
	if lessonTitle.Id == 0 || err != nil {
		return lessonTitle, errors.New("Lesson title not found")
	}

	return lessonTitle, nil
}

func (r *LessonTitleRepositoryImpl) Update(title domain.LessonTitle) domain.LessonTitle {
	err := r.db.Save(&title).Error
	helper.PanicIfError(err)

	return title
}

func (r *LessonTitleRepositoryImpl) FindByCourseId(courseId int) ([]domain.LessonTitle, error) {
	lessonTitles := []domain.LessonTitle{}
	err := r.db.Order("in_order asc").Find(&lessonTitles, "course_id=?", courseId).Error
	if len(lessonTitles) == 0 || err != nil {
		return nil, errors.New("Lesson titles not found")
	}

	return lessonTitles, nil
}

func (r *LessonTitleRepositoryImpl) Save(title domain.LessonTitle) domain.LessonTitle {
	err := r.db.Create(&title).Error
	helper.PanicIfError(err)

	return title
}

func NewLessonTitleRepository(db *gorm.DB) LessonTitleRepository {
	return &LessonTitleRepositoryImpl{db: db}
}
