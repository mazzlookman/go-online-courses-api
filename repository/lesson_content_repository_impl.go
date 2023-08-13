package repository

import (
	"errors"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type LessonContentRepositoryImpl struct {
	db *gorm.DB
}

func (r *LessonContentRepositoryImpl) FindByLessonTitleID(ltID int) ([]domain.LessonContent, error) {
	lessonContents := []domain.LessonContent{}
	err := r.db.Order("in_order asc").Find(&lessonContents, "lesson_title_id=?", ltID).Error
	if err != nil {
		return nil, errors.New("Lesson contents is not found")
	}

	return lessonContents, nil
}

func (r *LessonContentRepositoryImpl) Update(content domain.LessonContent) (domain.LessonContent, error) {
	err := r.db.Save(&content).Error
	if err != nil {
		return content, errors.New("Failed to update lesson content")
	}
	return content, nil
}

func (r *LessonContentRepositoryImpl) FindByID(lcID int) (domain.LessonContent, error) {
	lc := domain.LessonContent{}
	err := r.db.Find(&lc, "id=?", lcID).Error
	if err != nil {
		return lc, errors.New("Lesson content is not found")
	}
	return lc, nil
}

func (r *LessonContentRepositoryImpl) Save(content domain.LessonContent) (domain.LessonContent, error) {
	err := r.db.Create(&content).Error
	if err != nil {
		return content, errors.New("Error to create a lesson content")
	}

	return content, nil
}

func NewLessonContentRepository(db *gorm.DB) LessonContentRepository {
	return &LessonContentRepositoryImpl{db: db}
}
