package repository

import "go-pzn-restful-api/model/domain"

type UserRepository interface {
	Save(user domain.User) domain.User
	Update(user domain.User) domain.User
	FindByID(userID int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Delete(userID int)
}