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

type AuthorServiceImpl struct {
	repository.AuthorRepository
	auth.JwtAuth
}

func (s *AuthorServiceImpl) UploadAvatar(authorId int, filePath string) web.AuthorResponse {
	findById, err := s.AuthorRepository.FindById(authorId)
	oldAvatar := findById.Avatar
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if oldAvatar != filePath {
		if findById.Avatar == "" {
			return authorUploadAvatar(findById, filePath, s.AuthorRepository)
		}
		err := os.Remove(oldAvatar)
		helper.PanicIfError(err)
		return authorUploadAvatar(findById, filePath, s.AuthorRepository)
	}

	return authorUploadAvatar(findById, filePath, s.AuthorRepository)
}

func (s *AuthorServiceImpl) Logout(authorId int) web.AuthorResponse {
	findById, err := s.AuthorRepository.FindById(authorId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	findById.Token = ""
	update := s.AuthorRepository.Update(findById)

	return helper.ToAuthorResponse(update)
}

func (s *AuthorServiceImpl) Register(input web.AuthorRegisterInput) web.AuthorResponse {
	author := domain.Author{}
	findById, _ := s.AuthorRepository.FindByEmail(input.Email)
	if findById.Id != 0 {
		panic(errors.New("Email has been registered").Error())
	}

	author.Name = input.Name
	author.Email = input.Email
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	author.Password = string(password)
	author.Profile = input.Profile

	save := s.AuthorRepository.Save(author)

	return helper.ToAuthorResponse(save)
}

func (s *AuthorServiceImpl) Login(input web.AuthorLoginInput) web.AuthorResponse {
	findByEmail, err := s.AuthorRepository.FindByEmail(input.Email)
	if err != nil || findByEmail.Id == 0 {
		panic(helper.NewNotFoundError(errors.New("Email or password is wrong").Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.Password), []byte(input.Password))
	if err != nil {
		panic(helper.NewNotFoundError(errors.New("Email or password is wrong").Error()))
	}

	token, _ := s.JwtAuth.GenerateJwtToken("author", findByEmail.Id)
	findByEmail.Token = token

	update := s.AuthorRepository.Update(findByEmail)

	return helper.ToAuthorResponse(update)
}

func (s *AuthorServiceImpl) FindById(authorId int) web.AuthorResponse {
	findById, err := s.AuthorRepository.FindById(authorId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return helper.ToAuthorResponse(findById)
}

func NewAuthorService(authorRepository repository.AuthorRepository, jwtAuth auth.JwtAuth) AuthorService {
	return &AuthorServiceImpl{
		AuthorRepository: authorRepository,
		JwtAuth:          jwtAuth,
	}
}

func authorUploadAvatar(author domain.Author, filePath string, authorRepository repository.AuthorRepository) web.AuthorResponse {
	author.Avatar = filePath
	update := authorRepository.Update(author)
	response := helper.ToAuthorResponse(update)
	return response
}
