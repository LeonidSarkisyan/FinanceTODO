package repositories

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/pkg/tests_utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserPostgres(t *testing.T) {
	db := tests_utils.GetDB()

	authRepository := NewAuthPostgres(db)
	repository := NewUserPostgres(db)

	user := models.UserInput{
		Phone:    "test",
		Username: "test",
		Password: "test",
	}

	t.Run("получение пользователя по id", func(t *testing.T) {
		id, err := authRepository.Create(&user)
		require.NoError(t, err)

		gettingUser, err := repository.GetByID(id)
		require.NoError(t, err)

		require.Equal(t, id, gettingUser.ID)
		require.Equal(t, user.Username, gettingUser.Username)
	})
}
