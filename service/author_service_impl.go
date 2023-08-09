package service

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthorServiceImpl struct {
	repository.AuthorRepository
}

func (s *AuthorServiceImpl) Register(input web.AuthorRegisterInput) web.AuthorResponse {
	author := domain.Author{}
	findByID, _ := s.AuthorRepository.FindByEmail(input.Email)
	if findByID.ID != 0 {
		panic(errors.New("Email has been registered").Error())
	}

	author.Name = input.Name
	author.Email = input.Email
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	author.Password = string(password)
	author.Profile = input.Profile
	author.Avatar = input.Avatar

	save := s.AuthorRepository.Save(author)

	return helper.ToAuthorResponse(save)
}

func (s *AuthorServiceImpl) Login(input web.AuthorLoginInput) web.AuthorResponse {
	findByID, err := s.AuthorRepository.FindByEmail(input.Email)
	if err != nil || findByID.ID == 0 {
		panic(helper.NewNotFoundError(errors.New("Email or password is wrong").Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByID.Password), []byte(input.Password))
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Email or password is wrong").Error()))
	}

	return helper.ToAuthorResponse(findByID)
}

func (s *AuthorServiceImpl) FindByID(authorID int) web.AuthorResponse {
	findByID, err := s.AuthorRepository.FindByID(authorID)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToAuthorResponse(findByID)
}

func NewAuthorService(authorRepository repository.AuthorRepository) AuthorService {
	return &AuthorServiceImpl{AuthorRepository: authorRepository}
}
