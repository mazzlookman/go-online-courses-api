package service

import "go-pzn-restful-api/model/web"

type TransactionService interface {
	Create(input web.CreateTransactionInput) web.MidtransTransactionResponse
	PaymentProcess(midtransNotif web.TransactionNotificationFromMidtrans)
	//Update(trxID int, input web.CreateTransactionInput) web.MidtransTransactionResponse
}
