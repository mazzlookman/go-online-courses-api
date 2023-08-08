package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
)

type AuthorControllerImpl struct {
	service.AuthorService
}

func (c *AuthorControllerImpl) Create(ctx *gin.Context) {
	author := web.AuthorInputRequest{}
	err := ctx.ShouldBindJSON(&author)
	if ctx.GetHeader("role") != "admin" {
		ctx.JSON(401,
			helper.APIResponse(401, "Unauthorized", "You're not an admin"),
		)
	}
	helper.PanicIfError(err)

	authorResponse := c.AuthorService.Create(author)
	ctx.JSON(200,
		helper.APIResponse(200, "Author has created", authorResponse),
	)
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &AuthorControllerImpl{AuthorService: authorService}
}
