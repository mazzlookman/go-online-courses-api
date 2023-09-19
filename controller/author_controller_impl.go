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

func (c *AuthorControllerImpl) UploadAvatar(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("avatar")
	helper.PanicIfError(err)
	author := ctx.MustGet("current_author").(web.AuthorResponse)
	filePath := fmt.Sprintf("assets/images/avatars/%s-%s", author.Email, fileHeader.Filename)

	uploadAvatar := c.AuthorService.UploadAvatar(author.Id, filePath)

	err = ctx.SaveUploadedFile(fileHeader, filePath)
	helper.PanicIfError(err)

	ctx.JSON(
		200,
		helper.APIResponse(200, "Your avatar has been uploaded",
			gin.H{"author": uploadAvatar.Name, "avatar": uploadAvatar.Avatar, "is_uploaded": true}),
	)
}

func (c *AuthorControllerImpl) GetById(ctx *gin.Context) {
	// terakhir sampe sini
	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	authorResponse := c.AuthorService.FindById(authorId)
	ctx.JSON(200,
		helper.APIResponse(200, "Detail of author", authorResponse),
	)
}

func (c *AuthorControllerImpl) Logout(ctx *gin.Context) {
	authorId := ctx.MustGet("current_author").(web.AuthorResponse).Id
	logout := c.AuthorService.Logout(authorId)

	if logout.Token == "" {
		ctx.JSON(200,
			helper.APIResponse(200, "You're successfully logged out",
				gin.H{"author": logout.Name, "token": logout.Token}),
		)
	}
}

func (c *AuthorControllerImpl) Register(ctx *gin.Context) {
	author := web.AuthorRegisterInput{}
	err := ctx.ShouldBindJSON(&author)
	helper.PanicIfError(err)

	authorResponse := c.AuthorService.Register(author)

	ctx.JSON(200,
		helper.APIResponse(200, "Author has been registered", authorResponse),
	)
}

func (c *AuthorControllerImpl) Login(ctx *gin.Context) {
	input := web.AuthorLoginInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	authorResponse := c.AuthorService.Login(input)

	ctx.JSON(200,
		helper.APIResponse(200, "Author has been logged in", authorResponse),
	)
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &AuthorControllerImpl{AuthorService: authorService}
}
