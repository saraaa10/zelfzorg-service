package usecases

import (
	"service-api/src/main/entities"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	FindAll() ([]entities.Transaction, error)
	FindById(id uint) (entities.Transaction, error)
	FindByUserId(userId uint) ([]entities.Transaction, error)
	Save(transaction entities.Transaction) (entities.Transaction, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *transactionRepo {
	return &transactionRepo{
		db: db,
	}
}
