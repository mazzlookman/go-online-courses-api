package service

import (
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"strconv"
	"strings"
)

type TransactionServiceImpl struct {
	repository.TransactionRepository
	CourseService
}

func (s *TransactionServiceImpl) PaymentProcess(midtransNotif web.TransactionNotificationFromMidtrans) {
	split := strings.Split(midtransNotif.OrderId, "-")
	transactionId, _ := strconv.Atoi(split[0])
	findById, err := s.TransactionRepository.FindById(transactionId)
	if err != nil || findById.Id == 0 {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if midtransNotif.PaymentType == "credit_card" && midtransNotif.TransactionStatus == "capture" && midtransNotif.FraudStatus == "accept" {
		findById.Status = "paid"
	} else if midtransNotif.TransactionStatus == "settlement" {
		findById.Status = "paid"
	} else if midtransNotif.TransactionStatus == "deny" || midtransNotif.TransactionStatus == "expire" || midtransNotif.TransactionStatus == "cancel" {
		findById.Status = "canceled"
	}
	findById.PaymentUrl = ""

	update := s.TransactionRepository.Update(findById)

	s.CourseService.UserEnrolled(update.UserId, update.CourseId)
}

func (s *TransactionServiceImpl) Create(input web.CreateTransactionInput) web.MidtransTransactionResponse {
	// Find the course
	coursePrice := s.CourseService.FindById(input.CourseId).Price

	transaction := domain.Transaction{}
	transaction.CourseId = input.CourseId
	transaction.UserId = input.User.Id
	transaction.Amount = coursePrice
	transaction.Status = "pending"

	// Save the transaction (no payment URL)
	save := s.TransactionRepository.Save(transaction)

	// Get payment URL
	paymentResponse := helper.GetPaymentUrl(save, input.User)
	save.PaymentUrl = paymentResponse[0]

	// Update transaction with payment URL
	update := s.TransactionRepository.Update(save)

	return helper.ToMidtransTransactionResponse(update, paymentResponse[1])
}

func NewTransactionService(transactionRepository repository.TransactionRepository, courseService CourseService) TransactionService {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository, CourseService: courseService}
}
