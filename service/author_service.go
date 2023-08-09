package service

import "go-pzn-restful-api/model/web"

type AuthorService interface {
	Register(input web.AuthorRegisterInput) web.AuthorResponse
	Login(input web.AuthorLoginInput) web.AuthorResponse
	FindByID(authorID int) web.AuthorResponse
}
