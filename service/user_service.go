package service

import "go-pzn-restful-api/model/web"

type UserService interface {
	Login()
	Register(input web.UserRegisterInput) web.UserResponse
	FindByID()
	Logout()
}
