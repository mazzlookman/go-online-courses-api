package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"
)

type UserControllerImpl struct {
	service.UserService
}

func (c *UserControllerImpl) GetByID(ctx *gin.Context) {
	// user_id from token
	param := ctx.Param("userID")
	userID, _ := strconv.Atoi(param)
	findByID := c.UserService.FindByID(userID)
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
