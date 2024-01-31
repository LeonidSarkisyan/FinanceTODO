package repositories

import (
	"FinanceTODO/internal/models"
	"errors"
	"github.com/jmoiron/sqlx"
)

var (
	NotFoundError = errors.New("ничего не найдено")
)

type AuthRepository interface {
	Create(user *models.UserInput) (int, error)
	GetByPhone(phone string) (models.UserOutput, error)
}

type UserRepository interface {
	GetByID(userID int) (models.UserPublic, error)
}

type CategoryRepository interface {
	GetDefaultCategories() ([]models.Category, error)
}

type SubCategoryRepository interface {
	Create(subCategory models.SubCategoryInput, userID int) (int, error)
	GetAll(userID int) ([]models.SubCategoryOutput, error)
	GetByID(subCategoryID, userID int) (models.SubCategoryOutput, error)
	Update(subCategory models.SubCategoryInput, subCategoryID, userID int) error
	Delete(subCategoryID, userID int) error
}

type BalanceRepository interface {
	Create(balance models.BalanceInput, userID int) (int, error)
	GetAll(userID int) ([]models.Balance, error)
	GetByID(balanceID, userID int) (models.Balance, error)
	Update(balance models.BalanceUpdate, balanceID, userID int) error
	UpdateValue(amount float64, balanceID, userID int) error
	Delete(balanceID, userID int) error
}

type TransactionRepository interface {
	Create(input models.Transaction) (int, error)
	GetAll(balanceID, userID int) ([]models.Transaction, error)
	GetByID(id int, userId int) (models.Transaction, error)
	Update(input models.TransactionUpdate, transactionId int, userId int) error
	Delete(transactionId int, userId int) error
}

type Repository struct {
	Auth        AuthRepository
	User        UserRepository
	Category    CategoryRepository
	SubCategory SubCategoryRepository
	Balance     BalanceRepository
	Transaction TransactionRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:        NewAuthPostgres(db),
		User:        NewUserPostgres(db),
		Category:    NewCategoryPostgres(db),
		SubCategory: NewSubCategoryPostgres(db),
		Balance:     NewBalancePostgres(db),
		Transaction: NewTransactionPostgres(db),
	}
}
