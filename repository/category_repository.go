package repository

import "go-pzn-restful-api/model/domain"

type CategoryRepository interface {
	Save(category domain.Category) domain.Category
}
