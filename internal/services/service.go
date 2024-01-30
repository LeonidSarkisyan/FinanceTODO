package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
)

type AuthService interface {
	Register(userInput models.UserInput) (int, error)
	Login(userInput models.UserInput) (string, error)
}

type UserService interface {
	GetById(userID int) (models.UserPublic, error)
}

type BalanceService interface {
	Create(balance models.BalanceInput, userID int) (int, error)
	GetAll(userID int) ([]models.Balance, error)
	GetByID(balanceID, userID int) (models.Balance, error)
	Update(balance models.BalanceUpdate, balanceID, userID int) error
	Delete(balanceID, userID int) error
}

type SubCategoryService interface {
	Create(subCategory models.SubCategoryInput, userID int) (int, error)
	GetAll(userID int) ([]models.SubCategoryOutput, error)
	GetByID(subCategoryID, userID int) (models.SubCategoryOutput, error)
	Update(subCategory models.SubCategoryInput, subCategoryID, userID int) error
	Delete(subCategoryID, userID int) error
}

type TransactionService interface {
	Create(input models.TransactionInput, balanceId, subCategoryId, userId int) (int, error)
	GetAll(balanceID, userID int) ([]models.Transaction, error)
	GetByID(id int, userId int) (models.Transaction, error)
	Update(input models.TransactionUpdate, transactionId int, userId int) error
	Delete(transactionId int, userId int) error
}

type Service struct {
	Auth        AuthService
	User        UserService
	Balance     BalanceService
	SubCategory SubCategoryService
	Transaction TransactionService
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		Auth:        NewAuthService(repository),
		User:        NewUserService(repository),
		Balance:     NewBalanceService(repository),
		SubCategory: NewSubCategoryService(repository),
		Transaction: NewTransactionService(repository),
	}
}
