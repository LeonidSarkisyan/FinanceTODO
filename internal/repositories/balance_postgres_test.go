package repositories

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/pkg/tests_utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBalancePostgres(t *testing.T) {
	db := tests_utils.GetDB()

	authRepo := NewAuthPostgres(db)
	balanceRepo := NewBalancePostgres(db)

	user := models.UserInput{
		Phone:    "1234",
		Password: "1234",
	}

	t.Run("создание баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Type:     "test",
			Value:    200,
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)
		require.NotEmpty(t, balanceID)
	})

	t.Run("создание баланса со слишком большим названием", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		for i := 0; i < 16; i++ {
			newBalance.Title += "test"
		}

		_, err = balanceRepo.Create(newBalance, userID)

		require.ErrorIs(t, err, TooLongValueError{Field: "title", MaxValue: 16})
	})

	t.Run("получение всех балансов", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		_, err = balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		_, err = balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		_, err = balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		_, err = balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		balances, err := balanceRepo.GetAll(userID)

		require.NoError(t, err)
		require.Equal(t, 4, len(balances))

		for _, balance := range balances {
			require.NotEmpty(t, balance.ID)
			require.Equal(t, balance.Title, newBalance.Title)
			require.Equal(t, balance.Value, newBalance.Value)
			require.Equal(t, balance.Type, newBalance.Type)
			require.Equal(t, balance.Currency, newBalance.Currency)
		}
	})

	t.Run("получение своего баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		balance, err := balanceRepo.GetByID(balanceID, userID)

		require.NoError(t, err)
		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Value, newBalance.Value)
		require.Equal(t, balance.Type, newBalance.Type)
		require.Equal(t, balance.Currency, newBalance.Currency)
	})

	t.Run("получение баланса, которого нет", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		_, err = balanceRepo.GetByID(1000, userID)

		require.ErrorIs(t, err, NotFoundError)
	})

	t.Run("успешное обновление баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		balance, err := balanceRepo.GetByID(balanceID, userID)
		require.NoError(t, err)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Value, newBalance.Value)
		require.Equal(t, balance.Type, newBalance.Type)
		require.Equal(t, balance.Currency, newBalance.Currency)

		balanceUpdate := models.BalanceUpdate{
			Title:    tests_utils.GetPointer("test (UPDATE) 2"),
			Currency: tests_utils.GetPointer("test (UPDATE) 2"),
		}

		err = balanceRepo.Update(balanceUpdate, balanceID, userID)
		require.NoError(t, err)

		balance, err = balanceRepo.GetByID(balanceID, userID)

		require.Equal(t, balance.Title, *balanceUpdate.Title)
		require.Equal(t, balance.Currency, *balanceUpdate.Currency)
	})

	t.Run("неуспешное обновление баланса (другой user_id)", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		balance, err := balanceRepo.GetByID(balanceID, userID)
		require.NoError(t, err)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Value, newBalance.Value)
		require.Equal(t, balance.Type, newBalance.Type)
		require.Equal(t, balance.Currency, newBalance.Currency)

		balanceUpdate := models.BalanceUpdate{
			Title:    tests_utils.GetPointer("test (UPDATE) 2"),
			Currency: tests_utils.GetPointer("test (UPDATE) 2"),
		}

		otherUserID := 20000

		err = balanceRepo.Update(balanceUpdate, balanceID, otherUserID)
		require.ErrorIs(t, err, NotFoundError)

		balance, err = balanceRepo.GetByID(balanceID, userID)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Currency, newBalance.Currency)
	})

	t.Run("неуспешное обновление баланса (другой balance_id)", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)
		require.NoError(t, err)

		balance, err := balanceRepo.GetByID(balanceID, userID)
		require.NoError(t, err)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Value, newBalance.Value)
		require.Equal(t, balance.Type, newBalance.Type)
		require.Equal(t, balance.Currency, newBalance.Currency)

		balanceUpdate := models.BalanceUpdate{
			Title:    tests_utils.GetPointer("test (UPDATE) 2"),
			Currency: tests_utils.GetPointer("test (UPDATE) 2"),
		}

		otherBalanceID := 20000

		err = balanceRepo.Update(balanceUpdate, otherBalanceID, userID)
		require.ErrorIs(t, err, NotFoundError)

		balance, err = balanceRepo.GetByID(balanceID, userID)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Currency, newBalance.Currency)
	})

	t.Run("успешное удаление баланса", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)

		require.NoError(t, err)

		balance, err := balanceRepo.GetByID(balanceID, userID)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Value, newBalance.Value)
		require.Equal(t, balance.Type, newBalance.Type)
		require.Equal(t, balance.Currency, newBalance.Currency)

		err = balanceRepo.Delete(balanceID, userID)

		require.NoError(t, err)
	})

	t.Run("неуспешное удаление баланса (другой balance_id)", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)

		require.NoError(t, err)

		balance, err := balanceRepo.GetByID(balanceID, userID)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Value, newBalance.Value)
		require.Equal(t, balance.Type, newBalance.Type)
		require.Equal(t, balance.Currency, newBalance.Currency)

		otherBalanceID := 20000

		err = balanceRepo.Delete(otherBalanceID, userID)

		require.ErrorIs(t, err, NotFoundError)
	})

	t.Run("неуспешное удаление баланса (другой user_id)", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := authRepo.Create(&user)
		require.NoError(t, err)

		newBalance := models.BalanceInput{
			Title:    "test",
			Value:    255,
			Type:     "test",
			Currency: "test",
		}

		balanceID, err := balanceRepo.Create(newBalance, userID)

		require.NoError(t, err)

		balance, err := balanceRepo.GetByID(balanceID, userID)

		require.Equal(t, balance.Title, newBalance.Title)
		require.Equal(t, balance.Value, newBalance.Value)
		require.Equal(t, balance.Type, newBalance.Type)
		require.Equal(t, balance.Currency, newBalance.Currency)

		otherUserID := 20000

		err = balanceRepo.Delete(balanceID, otherUserID)

		require.ErrorIs(t, err, NotFoundError)
	})
}
