package service

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (s *UserServiceImpl) FindByID(userID int) web.UserResponse {
	findByID, err := s.UserRepository.FindByID(userID)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(findByID)
}

func (s *UserServiceImpl) Login(input web.UserLoginInput) web.UserResponse {
	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Email or password is wrong").Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.Password), []byte(input.Password))
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Email or password is wrong").Error()))
	}

	findByEmail.Token = "token"
	update := s.UserRepository.Update(findByEmail)

	return helper.ToUserResponse(update)
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

	return helper.ToUserResponse(save)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}
