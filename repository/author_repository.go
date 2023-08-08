package repository

import "go-pzn-restful-api/model/domain"

type AuthorRepository interface {
	Save(author domain.Author) domain.Author
}
