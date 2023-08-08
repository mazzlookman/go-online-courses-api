package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type AuthorServiceImpl struct {
	repository.AuthorRepository
}

func (s *AuthorServiceImpl) Create(request web.AuthorInputRequest) web.AuthorResponse {
	author := domain.Author{}
	author.Name = request.Name
	author.Introduction = request.Introduction

	save := s.AuthorRepository.Save(author)

	return helper.ToAuthorResponse(save)
}

func NewAuthorService(authorRepository repository.AuthorRepository) AuthorService {
	return &AuthorServiceImpl{AuthorRepository: authorRepository}
}
