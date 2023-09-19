package controller

import "github.com/gin-gonic/gin"

type AuthorController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	UploadAvatar(ctx *gin.Context)
	Logout(ctx *gin.Context)
	GetById(ctx *gin.Context)
}
