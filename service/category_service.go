package service

import "go-pzn-restful-api/model/web"

type CategoryService interface {
	Create(input web.CategoryCreateInput) web.CategoryResponse
}
