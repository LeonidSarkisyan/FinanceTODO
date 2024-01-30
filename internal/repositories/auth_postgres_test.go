package repositories

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/pkg/tests_utils"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/lib/pq"
)

func TestAuthPostgres(t *testing.T) {

	db := tests_utils.GetDB()

	tests_utils.ClearTestDatabase(db)

	t.Run("корректное добавление пользователя", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		newUser := models.UserInput{
			Phone:    "89774108103777",
			Username: "Leonid",
			Email:    "Leonid",
			Password: "1234567890",
		}

		id, err := NewAuthPostgres(db).Create(&newUser)
		if err != nil {
			t.Fatal(err)
		}
		log.Info().Int("id", id).Msg("Пользователь добавлен")
		require.NotEmpty(t, id)
		tests_utils.ClearTestDatabase(db)
	})

	t.Run("проверка уникальности телефона", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		newUser := models.UserInput{
			Phone:    "89774108103888",
			Username: "Leonid",
			Email:    "Leonid",
			Password: "1234567890",
		}

		_, err := NewAuthPostgres(db).Create(&newUser)

		require.NoError(t, err)

		_, err = NewAuthPostgres(db).Create(&newUser)

		require.Error(t, err, "Пользователь с таким телефоном уже существует")
		tests_utils.ClearTestDatabase(db)
	})

	t.Run("получение пользователя по телефону", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userInput := models.UserInput{
			Phone:    "89774108103456",
			Username: "Leonid",
			Email:    "Leonid",
			Password: "1234567890",
		}

		_, err := NewAuthPostgres(db).Create(&userInput)

		require.NoError(t, err)

		user, err := NewAuthPostgres(db).GetByPhone("89774108103456")

		require.NoError(t, err)
		require.Equal(t, "89774108103456", user.Phone)
		require.Equal(t, "Leonid", user.Username)
		require.Equal(t, "Leonid", user.Email)
		tests_utils.ClearTestDatabase(db)
	})
}
