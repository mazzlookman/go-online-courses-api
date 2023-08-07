package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) Save(user domain.User) domain.User {
	err := r.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) Update(user domain.User) domain.User {
	err := r.db.Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) FindByID(userID int) (domain.User, error) {
	user := domain.User{}
	err := r.db.Where("id=?", userID).Find(&user).Error
	if err != nil {
		return user, errors.New("User is not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(userID int) {
	user := domain.User{}
	err := r.db.Where("id=?", userID).Delete(&user).Error
	helper.PanicIfError(err)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
