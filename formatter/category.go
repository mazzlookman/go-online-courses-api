package formatter

import (
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
