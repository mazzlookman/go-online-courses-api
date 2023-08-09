package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
)

type AuthorControllerImpl struct {
	service.AuthorService
}

func (c *AuthorControllerImpl) Register(ctx *gin.Context) {
	author := web.AuthorRegisterInput{}
	err := ctx.ShouldBind(&author)
	//if ctx.GetHeader("role") != "admin" {
	//	ctx.AbortWithStatusJSON(401,
	//		helper.APIResponse(401, "Unauthorized", "You're not an admin"),
	//	)
	//	return
	//}
	fileHeader, err := ctx.FormFile("avatar")
	helper.PanicIfError(err)
	filePath := fmt.Sprintf("assets/images/avatars/%s-%s", author.Email, fileHeader.Filename)

	author.Avatar = filePath
	authorResponse := c.AuthorService.Register(author)

	err = ctx.SaveUploadedFile(fileHeader, filePath)
	helper.PanicIfError(err)

	ctx.JSON(200,
		helper.APIResponse(200, "Author has created", authorResponse),
	)
}

func (c *AuthorControllerImpl) Login(ctx *gin.Context) {
	input := web.AuthorLoginInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	authorResponse := c.AuthorService.Login(input)
	ctx.Set("author_id", authorResponse.ID)
	ctx.JSON(200,
		helper.APIResponse(200, "Author has been logged in", authorResponse),
	)
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &AuthorControllerImpl{AuthorService: authorService}
}