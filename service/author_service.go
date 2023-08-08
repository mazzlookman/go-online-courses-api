package service

import "go-pzn-restful-api/model/web"

type AuthorService interface {
	Create(request web.AuthorInputRequest) web.AuthorResponse
}
