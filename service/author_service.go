package service

import "go-pzn-restful-api/model/web"

type AuthorService interface {
	Register(input web.AuthorRegisterInput) web.AuthorResponse
	Login(input web.AuthorLoginInput) web.AuthorResponse
	UploadAvatar(authorId int, filePath string) web.AuthorResponse
	FindById(authorId int) web.AuthorResponse
	Logout(authorId int) web.AuthorResponse
}
