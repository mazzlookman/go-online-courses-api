package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(g *gin.Context)
}
