package repository

import "go-pzn-restful-api/model/domain"

type TransactionRepository interface {
	Save(transaction domain.Transaction) (domain.Transaction, error)
	Update(transaction domain.Transaction) (domain.Transaction, error)
	FindById(transactionId int) (domain.Transaction, error)
}
