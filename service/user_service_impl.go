package service

import (
	"errors"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type UserServiceImpl struct {
	repository.UserRepository
	auth.JwtAuth
}

func (s *UserServiceImpl) Logout(userID int) web.UserResponse {
	findByID, err := s.UserRepository.FindByID(userID)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	findByID.Token = ""
	update := s.UserRepository.Update(findByID)

	return helper.ToUserResponse(update)
}

func (s *UserServiceImpl) UploadAvatar(userID int, filePath string) web.UserResponse {
	findByID, err := s.UserRepository.FindByID(userID)
	oldAvatar := findByID.Avatar
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if oldAvatar != filePath {
		if findByID.Avatar == "" {
			return updateWhenUploadAvatar(findByID, filePath, s.UserRepository)
		}
		err := os.Remove(oldAvatar)
		helper.PanicIfError(err)
		return updateWhenUploadAvatar(findByID, filePath, s.UserRepository)
	}

	return updateWhenUploadAvatar(findByID, filePath, s.UserRepository)
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
		panic(helper.NewBadRequestError(errors.New("Email or password is wrong").Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.Password), []byte(input.Password))
	if err != nil {
		panic(helper.NewBadRequestError(errors.New("Email or password is wrong").Error()))
	}

	token, _ := s.JwtAuth.GenerateJwtToken("user", findByEmail.ID)
	findByEmail.Token = token

	update := s.UserRepository.Update(findByEmail)

	return helper.ToUserResponse(update)
}

func (s *UserServiceImpl) Register(input web.UserRegisterInput) web.UserResponse {
	findByEmail, err := s.UserRepository.FindByEmail(input.Email)
	if findByEmail.ID != 0 {
		panic(helper.NewNotFoundError(errors.New("Email has been registered").Error()))
	}

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

func NewUserService(userRepository repository.UserRepository, jwtAuth auth.JwtAuth) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		JwtAuth:        jwtAuth,
	}
}

func updateWhenUploadAvatar(user domain.User, filePath string, userRepository repository.UserRepository) web.UserResponse {
	user.Avatar = filePath
	update := userRepository.Update(user)
	response := helper.ToUserResponse(update)
	return response
}
