package controller

import "github.com/gin-gonic/gin"

type AuthorController interface {
	Create(ctx *gin.Context)
}
