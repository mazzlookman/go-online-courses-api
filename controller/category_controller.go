package controller

import "github.com/gin-gonic/gin"

type CategoryController interface {
	Create(ctx *gin.Context)
}
