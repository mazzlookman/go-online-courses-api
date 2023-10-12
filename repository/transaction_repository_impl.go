package repository

import (
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func (r *TransactionRepositoryImpl) Save(transaction domain.Transaction) domain.Transaction {
	err := r.db.Create(&transaction).Error
	helper.PanicIfError(err)

	return transaction
}

func (r *TransactionRepositoryImpl) Update(transaction domain.Transaction) domain.Transaction {
	err := r.db.Save(&transaction).Error
	helper.PanicIfError(err)

	return transaction
}

func (r *TransactionRepositoryImpl) FindById(transactionId int) (domain.Transaction, error) {
	trx := domain.Transaction{}
	err := r.db.Find(&trx, "id=?", transactionId).Error
	if err != nil || trx.Id == 0 {
		return trx, errors.New("Transaction not found")
	}

	return trx, nil
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}
