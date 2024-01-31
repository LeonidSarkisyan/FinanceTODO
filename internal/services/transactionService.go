package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
	"errors"
)

var InCorrectType = errors.New("неверный тип транзакции, тип должен быть INC или EXP")

type Transaction struct {
	repository *repositories.Repository
}

func NewTransactionService(repository *repositories.Repository) *Transaction {
	return &Transaction{
		repository: repository,
	}
}

func (s *Transaction) Register(input models.TransactionInput, balanceId, subCategoryId, userId int) (int, error) {
	transaction := models.NewTransaction(input, balanceId, subCategoryId, userId)
	transactionID, err := s.repository.Transaction.Create(*transaction)
	if err != nil {
		return 0, err
	}
	var valueToUpdate float64
	switch input.Type {
	case "INC":
		valueToUpdate = input.Value
	case "EXP":
		valueToUpdate = -input.Value
	default:
		return 0, InCorrectType
	}
	err = s.repository.Balance.UpdateValue(valueToUpdate, balanceId, userId)
	if err != nil {
		return 0, err
	}
	return transactionID, nil
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
