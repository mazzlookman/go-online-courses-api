package util

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"go-pzn-restful-api/service"
	"log"
)

var (
	authorRepo = repository.NewAuthorRepository(Db)
	authorServ = service.NewAuthorService(authorRepo, JwtAuth)
)

func CreateAuthorTest() web.AuthorResponse {
	input := web.AuthorRegisterInput{
		Name:     "test",
		Email:    "test@test.com",
		Password: "123",
		Profile:  "Profile",
		Avatar:   "assets/images/avatars/author.jpg",
	}

	log.Println("Author has been created")
	return authorServ.Register(input)
}

func DeleteAuthorTest() {
	err := authorRepo.Delete("test@test.com")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Author has been deleted")
}

func GetAuthorByID(id int) domain.Author {
	author, err := authorRepo.FindByID(id)
	if err != nil {
		log.Fatalln(err)
	}

	return author
}
