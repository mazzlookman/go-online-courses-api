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
	if err != nil || user.ID == 0 {
		return user, errors.New("User is not found")
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

func (r *UserRepositoryImpl) FindByID(userID int) (domain.User, error) {
	user := domain.User{}
	err := r.db.Where("id=?", userID).Find(&user).Error
	if err != nil || user.ID == 0 {
		return user, errors.New("User is not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(userName string) error {
	user := domain.User{}
	err := r.db.Where("name=?", userName).Delete(&user).Error
	helper.PanicIfError(err)

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
