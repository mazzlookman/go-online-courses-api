package controller

import "github.com/gin-gonic/gin"

type CourseController interface {
	Create(ctx *gin.Context)
	GetBySlug(ctx *gin.Context)
	GetByAuthorID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	UserEnrolled(ctx *gin.Context)
	UploadBanner(ctx *gin.Context)
}
