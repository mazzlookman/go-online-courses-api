package repository

import (
	"go-pzn-restful-api/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func (r *TransactionRepositoryImpl) Save(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (r *TransactionRepositoryImpl) Update(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (r *TransactionRepositoryImpl) FindById(transactionId int) (domain.Transaction, error) {
	trx := domain.Transaction{}
	err := r.db.Find(&trx, "id=?", transactionId).Error
	if err != nil {
		return trx, err
	}

	return trx, nil
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}
