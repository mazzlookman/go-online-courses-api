package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
)

type UserControllerImpl struct {
	service.UserService
}

func (c *UserControllerImpl) Register(g *gin.Context) {
	input := web.UserRegisterInput{}
	err := g.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	register := c.UserService.Register(input)
	g.JSON(200, helper.APIResponse(200, "Register is successfully", register))
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}
