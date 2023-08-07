package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (s *UserServiceImpl) Login() {

}

func (s *UserServiceImpl) Register(input web.UserRegisterInput) web.UserResponse {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	domainUser := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
	}
	save := s.UserRepository.Save(domainUser)
	helper.PanicIfError(err)

	return helper.ToResponseUser(save)
}

func (s *UserServiceImpl) FindByID() {
	//TODO implement me
	panic("implement me")
}

func (s *UserServiceImpl) Logout() {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}
