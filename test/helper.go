package test

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"log"
)

// User
func CreateUserTest() web.UserResponse {
	input := web.UserRegisterInput{
		Name:     "user",
		Email:    "user@user.com",
		Password: "123",
	}
	log.Println("User registered")
	return UserService.Register(input)
}

func DeleteUserTest() {
	err := UserRepository.Delete("user")
	if err != nil {
		helper.PanicIfError(err)
	}

	log.Println("User deleted")
}

func GetTokenAfterLogin() string {
	login := UserService.Login(web.UserLoginInput{
		Email:    "user@user.com",
		Password: "123",
	})

	return login.Token
}

// Author
func CreateAuthorTest() web.AuthorResponse {
	input := web.AuthorRegisterInput{
		Name:     "author",
		Email:    "author@author.com",
		Password: "123",
		Profile:  "Profile",
		Avatar:   "assets/images/avatars/author.jpg",
	}

	log.Println("Author has been created")
	return AuthorService.Register(input)
}

func DeleteAuthorTest() {
	err := AuthorRepository.Delete("author@author.com")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Author has been deleted")
}

func GetAuthorToken() string {
	login := AuthorService.Login(web.AuthorLoginInput{
		Email:    "author@author.com",
		Password: "123",
	})

	return login.Token
}
