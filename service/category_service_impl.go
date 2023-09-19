package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
)

type CategoryServiceImpl struct {
	repository.CategoryRepository
}

func (s *CategoryServiceImpl) Create(input web.CategoryCreateInput) web.CategoryResponse {
	ctg := domain.Category{
		Name: input.Name,
	}
	category := s.CategoryRepository.Save(ctg)

	return helper.ToCategoryResponse(category)
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository}
}
