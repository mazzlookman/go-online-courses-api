package repository

import "go-pzn-restful-api/model/domain"

type AuthorRepository interface {
	Save(author domain.Author) domain.Author
	Update(author domain.Author) domain.Author
	FindByID(authorID int) (domain.Author, error)
	FindByEmail(email string) (domain.Author, error)
	Delete(email string) error
}
