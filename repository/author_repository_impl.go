package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func (r *AuthorRepositoryImpl) Update(author domain.Author) domain.Author {
	err := r.db.Save(&author).Error
	helper.PanicIfError(err)

	return author
}

func (r *AuthorRepositoryImpl) FindByEmail(email string) (domain.Author, error) {
	author := domain.Author{}
	err := r.db.Where("email=?", email).Find(&author).Error
	if err != nil {
		return author, errors.New("Author is not found")
	}

	return author, nil
}

func (r *AuthorRepositoryImpl) FindByID(authorID int) (domain.Author, error) {
	author := domain.Author{}
	err := r.db.Where("id=?", authorID).Find(&author).Error
	if err != nil {
		return author, errors.New("Author is not found")
	}

	return author, nil
}

func (r *AuthorRepositoryImpl) Save(author domain.Author) domain.Author {
	err := r.db.Create(&author).Error
	helper.PanicIfError(err)

	return author
}

func (r *AuthorRepositoryImpl) Delete(email string) error {
	user := domain.Author{}
	err := r.db.Where("email=?", email).Delete(&user).Error
	helper.PanicIfError(err)

	return nil
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{db: db}
}
