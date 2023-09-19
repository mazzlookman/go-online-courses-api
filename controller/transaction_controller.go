package controller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	EarnPaidCourse(ctx *gin.Context)
	MidtransNotification(ctx *gin.Context)
}
