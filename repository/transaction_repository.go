package repository

import "go-pzn-restful-api/model/domain"

type TransactionRepository interface {
	Save(transaction domain.Transaction) domain.Transaction
	Update(transaction domain.Transaction) domain.Transaction
	FindById(transactionId int) (domain.Transaction, error)
}
