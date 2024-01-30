package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
	"FinanceTODO/pkg/tests_utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBalanceService(t *testing.T) {
	db := tests_utils.GetDB()
	repository := repositories.NewRepository(db)
	service := NewService(repository)

	newUser := models.UserInput{
		Phone:    "12347777777777777777",
		Password: "12347777777777777777",
	}

	newBalance := models.BalanceInput{
		Title:    "test",
		Type:     "test",
		Currency: "test",
		Value:    500,
	}

	t.Run("создание баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		balanceID, err := service.Balance.Create(newBalance, userID)
		require.NoError(t, err)

		require.NotEmpty(t, balanceID)
	})

	t.Run("получение всех балансов", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		_, err = service.Balance.Create(newBalance, userID)
		require.NoError(t, err)

		_, err = service.Balance.Create(newBalance, userID)
		require.NoError(t, err)

		_, err = service.Balance.Create(newBalance, userID)
		require.NoError(t, err)

		balances, err := service.Balance.GetAll(userID)

		require.NoError(t, err)
		require.Equal(t, 3, len(balances))
	})

	t.Run("получение баланса по ID", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		balanceID, err := service.Balance.Create(newBalance, userID)

		require.NoError(t, err)

		balance, err := service.Balance.GetByID(balanceID, userID)

		require.NoError(t, err)
		require.Equal(t, balanceID, balance.ID)
		require.Equal(t, newBalance.Title, balance.Title)
		require.Equal(t, newBalance.Type, balance.Type)
		require.Equal(t, newBalance.Currency, balance.Currency)
		require.Equal(t, newBalance.Value, balance.Value)
	})

	t.Run("получение несуществующего баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		balanceID, err := service.Balance.Create(newBalance, userID)

		require.NoError(t, err)

		otherBalanceID := balanceID + 1

		_, err = service.Balance.GetByID(otherBalanceID, userID)

		require.ErrorIs(t, err, repositories.NotFoundError)
	})

	t.Run("успешное обновление баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		balanceID, err := service.Balance.Create(newBalance, userID)

		require.NoError(t, err)

		balanceUpdate := models.BalanceUpdate{
			Title: tests_utils.GetPointer("test (UPDATE)"),
			Type:  tests_utils.GetPointer("test (UPDATE)"),
		}

		err = service.Balance.Update(balanceUpdate, balanceID, userID)

		require.NoError(t, err)

		balance, err := service.Balance.GetByID(balanceID, userID)

		require.NoError(t, err)

		require.Equal(t, *balanceUpdate.Title, balance.Title)
		require.Equal(t, *balanceUpdate.Title, balance.Title)
	})

	t.Run("неуспешное обновление баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		balanceID, err := service.Balance.Create(newBalance, userID)

		require.NoError(t, err)

		balanceUpdate := models.BalanceUpdate{
			Title: tests_utils.GetPointer("test (UPDATE)"),
			Type:  tests_utils.GetPointer("test (UPDATE)"),
		}

		otherBalanceID := balanceID + 1

		err = service.Balance.Update(balanceUpdate, otherBalanceID, userID)

		require.ErrorIs(t, err, repositories.NotFoundError)

		balance, err := service.Balance.GetByID(balanceID, userID)

		require.NoError(t, err)

		require.NotEqual(t, *balanceUpdate.Title, balance.Title)
		require.NotEqual(t, *balanceUpdate.Title, balance.Title)
	})

	t.Run("успешное удаление баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		balanceID, err := service.Balance.Create(newBalance, userID)

		require.NoError(t, err)

		err = service.Balance.Delete(balanceID, userID)

		require.NoError(t, err)
	})

	t.Run("неуспешное удаление баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(newUser)

		require.NoError(t, err)

		balanceID, err := service.Balance.Create(newBalance, userID)

		require.NoError(t, err)

		otherBalanceID := balanceID + 1

		err = service.Balance.Delete(otherBalanceID, userID)

		require.ErrorIs(t, err, repositories.NotFoundError)

		balance, err := service.Balance.GetByID(balanceID, userID)

		require.NoError(t, err)
		require.Equal(t, balanceID, balance.ID)
		require.Equal(t, newBalance.Title, balance.Title)
		require.Equal(t, newBalance.Type, balance.Type)
		require.Equal(t, newBalance.Currency, balance.Currency)
		require.Equal(t, newBalance.Value, balance.Value)
	})
}
