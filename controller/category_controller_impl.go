package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
)

type CategoryControllerImpl struct {
	service.CategoryService
}

func (c *CategoryControllerImpl) Create(ctx *gin.Context) {
	input := web.CategoryCreateInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	categoryResponse := c.CategoryService.Create(input)

	ctx.JSON(200,
		helper.APIResponse(200, "Category has created", categoryResponse))
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService}
}
