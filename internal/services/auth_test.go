package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
	"FinanceTODO/pkg/tests_utils"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
	"testing"
)

func clearTestDatabase(db *sqlx.DB) {
	_, err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY")
	if err != nil {
		log.Fatal().Err(err).Msg("ошибка при удалении данных в конце теста")
	}
}

func TestAuthService(t *testing.T) {

	db := tests_utils.GetDB()
	repository := repositories.NewRepository(db)
	service := NewService(repository)

	tests_utils.ClearTestDatabase(db)

	t.Run("корректное регистрация пользователя", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		newUser := models.UserInput{
			Phone:    "89774108103900",
			Username: "Leonid",
			Email:    "Leonid",
			Password: "1234567890",
		}

		id, err := service.Auth.Register(newUser)

		require.NoError(t, err)
		require.NotEmpty(t, id)
		tests_utils.ClearTestDatabase(db)
	})

	t.Run("выдача дружелюбной ошибки при регистрации с того же номера", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		newUser := models.UserInput{
			Phone:    "89774108103333",
			Username: "Leonid",
			Email:    "Leonid",
			Password: "1234567890",
		}
		_, err := service.Auth.Register(newUser)
		require.NoError(t, err)
		_, err = service.Auth.Register(newUser)
		require.ErrorAs(t, err, &UserAlreadyExists)
		tests_utils.ClearTestDatabase(db)
	})

	t.Run("выдача ошибки при неправильной аутентифицкации", func(t *testing.T) {
		t.Run("неправильный пароль", func(t *testing.T) {
			tests_utils.ClearTestDatabase(db)
			newUser := models.UserInput{
				Phone:    "1234",
				Username: "Leonid",
				Email:    "Leonid",
				Password: "1234",
			}

			_, err := service.Auth.Register(newUser)
			require.NoError(t, err)

			newUser.Password = "12345"

			_, err = service.Auth.Login(newUser)

			require.ErrorAs(t, err, &BadLoginOrPassword)
			tests_utils.ClearTestDatabase(db)
		})

		t.Run("неправильный телефон", func(t *testing.T) {
			tests_utils.ClearTestDatabase(db)

			newUser := models.UserInput{
				Phone:    "5432",
				Username: "Leonid",
				Email:    "Leonid",
				Password: "1234",
			}

			_, err := service.Auth.Register(newUser)
			require.NoError(t, err)

			newUser.Phone = "1234999999"

			_, err = service.Auth.Login(newUser)

			require.ErrorAs(t, err, &BadLoginOrPassword)
		})
	})
}
