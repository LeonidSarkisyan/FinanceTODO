package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
)

type Balance struct {
	repository *repositories.Repository
}

func NewBalanceService(repository *repositories.Repository) *Balance {
	return &Balance{
		repository: repository,
	}
}

func (b *Balance) Create(balance models.BalanceInput, userID int) (int, error) {
	return b.repository.Balance.Create(balance, userID)
}

func (b *Balance) GetAll(userID int) ([]models.Balance, error) {
	return b.repository.Balance.GetAll(userID)
}

func (b *Balance) GetByID(balanceID, userID int) (models.Balance, error) {
	return b.repository.Balance.GetByID(balanceID, userID)
}

func (b *Balance) Update(balance models.BalanceUpdate, balanceID, userID int) error {
	return b.repository.Balance.Update(balance, balanceID, userID)
}

func (b *Balance) Delete(balanceID, userID int) error {
	return b.repository.Balance.Delete(balanceID, userID)
}
