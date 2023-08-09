package helper

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
		Token:  user.Token,
	}
}

func ToAuthorResponse(author domain.Author) web.AuthorResponse {
	return web.AuthorResponse{
		ID:      author.ID,
		Name:    author.Name,
		Email:   author.Email,
		Profile: author.Profile,
		Avatar:  author.Avatar,
	}
}
