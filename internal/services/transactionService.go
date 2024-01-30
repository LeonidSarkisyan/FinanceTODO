package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
)

type Transaction struct {
	repository *repositories.Repository
}

func NewTransactionService(repository *repositories.Repository) *Transaction {
	return &Transaction{
		repository: repository,
	}
}

func (s *Transaction) Create(input models.TransactionInput, balanceId, subCategoryId, userId int) (int, error) {
	transaction := models.NewTransaction(input, balanceId, subCategoryId, userId)
	return s.repository.Transaction.Create(*transaction)
}

func (s *Transaction) GetAll(balanceID, userID int) ([]models.Transaction, error) {
	return s.repository.Transaction.GetAll(balanceID, userID)
}

func (s *Transaction) GetByID(id int, userId int) (models.Transaction, error) {
	return s.repository.Transaction.GetByID(id, userId)
}

func (s *Transaction) Update(input models.TransactionUpdate, transactionId int, userId int) error {
	return s.repository.Transaction.Update(input, transactionId, userId)
}

func (s *Transaction) Delete(transactionId int, userId int) error {
	return s.repository.Transaction.Delete(transactionId, userId)
}
