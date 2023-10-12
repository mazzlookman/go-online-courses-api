package repository

import "go-pzn-restful-api/model/domain"

type AuthorRepository interface {
	Save(author domain.Author) domain.Author
	Update(author domain.Author) domain.Author
	FindById(authorId int) (domain.Author, error)
	FindByEmail(email string) (domain.Author, error)
	Delete(email string)
}
