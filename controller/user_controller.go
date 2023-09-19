package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetById(ctx *gin.Context)
	UploadAvatar(ctx *gin.Context)
	Logout(ctx *gin.Context)
}
