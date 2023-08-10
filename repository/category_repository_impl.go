package repository

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (r *CategoryRepositoryImpl) Save(category domain.Category) domain.Category {
	err := r.db.Create(&category).Error
	helper.PanicIfError(err)

	return category
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}
