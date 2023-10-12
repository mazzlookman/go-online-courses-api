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

func (r *UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := r.db.Find(&user, "email=?", email).Error
	if err != nil {
		return user, errors.New("User not found")
	}

	return user, nil
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

func (r *UserRepositoryImpl) FindById(userId int) (domain.User, error) {
	user := domain.User{}
	err := r.db.Preload("Courses").Where("id=?", userId).Find(&user).Error
	if err != nil || user.Id == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(userId int) {
	user := domain.User{}
	err := r.db.Where("id=?", userId).Delete(&user).Error

	helper.PanicIfError(err)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
