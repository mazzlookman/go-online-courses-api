package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
)

type UserControllerImpl struct {
	service.UserService
}

func (c *UserControllerImpl) Logout(ctx *gin.Context) {
	userID := ctx.MustGet("current_user").(web.UserResponse).ID
	userResponse := c.UserService.Logout(userID)
	if userResponse.Token == "" {
		ctx.JSON(200,
			helper.APIResponse(200, "You're logged out",
				gin.H{"user": userResponse.Name, "token": userResponse.Token}),
		)
	}
}

func (c *UserControllerImpl) UploadAvatar(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("avatar")
	helper.PanicIfError(err)
	userID := ctx.MustGet("current_user").(web.UserResponse).ID
	filePath := fmt.Sprintf("assets/images/avatars/%d-%s", userID, fileHeader.Filename)

	err = ctx.SaveUploadedFile(fileHeader, filePath)
	helper.PanicIfError(err)

	uploadAvatar := c.UserService.UploadAvatar(userID, filePath)

	ctx.JSON(
		200,
		helper.APIResponse(200, "Your avatar has been uploaded",
			gin.H{"user": uploadAvatar.Name, "avatar": uploadAvatar.Avatar, "is_uploaded": true}),
	)
}

func (c *UserControllerImpl) GetByID(ctx *gin.Context) {
	// user_id from token
	user := ctx.MustGet("current_user").(web.UserResponse)
	findByID := c.UserService.FindByID(user.ID)
	ctx.JSON(
		200,
		helper.APIResponse(200, "Current user: "+findByID.Name, findByID),
	)
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	input := web.UserLoginInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	login := c.UserService.Login(input)

	ctx.JSON(
		200,
		helper.APIResponse(200, "You're logged in now", login),
	)
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	input := web.UserRegisterInput{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	register := c.UserService.Register(input)
	ctx.JSON(200, helper.APIResponse(200, "Register is successfully", register))
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}
