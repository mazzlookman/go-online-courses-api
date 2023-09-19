package controller

import (
	"github.com/gin-gonic/gin"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/service"
	"strconv"
)

type TransactionControllerImpl struct {
	service.TransactionService
}

func (c *TransactionControllerImpl) MidtransNotification(ctx *gin.Context) {
	input := web.TransactionNotificationFromMidtrans{}
	err := ctx.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	if input.TransactionStatus == "pending" {
		return
	}

	c.TransactionService.PaymentProcess(input)

	ctx.JSON(200, nil)
}

func (c *TransactionControllerImpl) EarnPaidCourse(ctx *gin.Context) {
	cid := ctx.Param("courseId")
	courseID, _ := strconv.Atoi(cid)

	user := ctx.MustGet("current_user").(web.UserResponse)

	input := web.CreateTransactionInput{}
	input.CourseId = courseID
	input.User = user

	transactionResponse := c.TransactionService.Create(input)

	ctx.JSON(200, helper.APIResponse(200, "Happy learning", transactionResponse))
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &TransactionControllerImpl{TransactionService: transactionService}
}
