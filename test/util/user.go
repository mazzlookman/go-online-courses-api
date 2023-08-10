package util

import (
	"go-pzn-restful-api/app"
	"go-pzn-restful-api/auth"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
	"log"
)

var (
	R              = app.NewRouter()
	Db             = app.DBConnection()
	userRepository = repository.NewUserRepository(Db)
	JwtAuth        = auth.NewJwtAuth()
	userService    = service.NewUserService(userRepository, JwtAuth)
)

func CreateUserTest() web.UserResponse {
	input := web.UserRegisterInput{
		Name:     "test",
		Email:    "test@test.com",
		Password: "123",
	}
	log.Println("User registered")
	return userService.Register(input)
}

func DeleteUserTest() {
	err := userRepository.Delete("test")
	if err != nil {
		helper.PanicIfError(err)
	}

	log.Println("User deleted")
}

func GetTokenAfterLogin() string {
	login := userService.Login(web.UserLoginInput{
		Email:    "test@test.com",
		Password: "123",
	})

	return login.Token
}
