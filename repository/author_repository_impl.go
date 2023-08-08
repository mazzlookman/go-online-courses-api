package repository

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func (r *AuthorRepositoryImpl) Save(author domain.Author) domain.Author {
	err := r.db.Create(&author).Error
	helper.PanicIfError(err)

	return author
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{db: db}
}
