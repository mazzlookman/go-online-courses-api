package controller

import "github.com/gin-gonic/gin"

type AuthorController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}
