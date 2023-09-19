package repository

import (
	"errors"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type LessonContentRepositoryImpl struct {
	db *gorm.DB
}

func (r *LessonContentRepositoryImpl) FindByLessonTitleId(ltId int) ([]domain.LessonContent, error) {
	lessonContents := []domain.LessonContent{}
	err := r.db.Order("in_order asc").Find(&lessonContents, "lesson_title_id=?", ltId).Error
	if len(lessonContents) == 0 || err != nil {
		return nil, errors.New("Lesson contents not found")
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

func (r *LessonContentRepositoryImpl) FindById(lcId int) (domain.LessonContent, error) {
	lc := domain.LessonContent{}
	err := r.db.Find(&lc, "Id=?", lcId).Error
	if lc.Id == 0 || err != nil {
		return lc, errors.New("Lesson content not found")
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
