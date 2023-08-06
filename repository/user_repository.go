package repository

import "go-pzn-restful-api/model/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	FindByID(userID int) (domain.User, error)
	Delete(userID int) error
}
