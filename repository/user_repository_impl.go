package repository

import (
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*gorm.DB
}

func (r *UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImpl) FindByID(userID int) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImpl) Delete(userID int) error {
	//TODO implement me
	panic("implement me")
}
