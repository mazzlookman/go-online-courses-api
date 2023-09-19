package service

import "go-pzn-restful-api/model/web"

type UserService interface {
	Login(input web.UserLoginInput) web.UserResponse
	Register(input web.UserRegisterInput) web.UserResponse
	FindById(userId int) web.UserResponse
	UploadAvatar(userId int, filePath string) web.UserResponse
	Logout(userId int) web.UserResponse
}
