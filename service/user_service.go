package service

import "go-pzn-restful-api/model/web"

type UserService interface {
	Login(input web.UserLoginInput) web.UserResponse
	Register(input web.UserRegisterInput) web.UserResponse
	FindByID(userID int) web.UserResponse
	UploadAvatar(userID int, filePath string) web.UserResponse
	Logout(userID int) web.UserResponse
}
