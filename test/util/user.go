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
	Db             = app.DBConnection()
	UserRepository = repository.NewUserRepository(Db)
	JwtAuth        = auth.NewJwtAuth()
	UserService    = service.NewUserService(UserRepository, JwtAuth)
)

func CreateUserTest() web.UserResponse {
	input := web.UserRegisterInput{
		Name:     "test",
		Email:    "test@test.com",
		Password: "123",
	}
	log.Println("User registered")
	return UserService.Register(input)
}

func DeleteUserTest() {
	err := UserRepository.Delete("test")
	if err != nil {
		helper.PanicIfError(err)
	}

	log.Println("User deleted")
}

func GetTokenAfterLogin() string {
	login := UserService.Login(web.UserLoginInput{
		Email:    "test@test.com",
		Password: "123",
	})

	return login.Token
}
