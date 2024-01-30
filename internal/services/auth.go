package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
	"FinanceTODO/pkg/utils"
	"errors"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	UserAlreadyExists  = errors.New("пользователь с таким телефоном уже существует")
	BadLoginOrPassword = errors.New("неправильный логин или пароль")
)

type Auth struct {
	repository *repositories.Repository
}

func NewAuthService(repository *repositories.Repository) *Auth {
	return &Auth{
		repository: repository,
	}
}

func (s *Auth) Register(userInput models.UserInput) (int, error) {
	id, err := s.repository.Auth.Create(&userInput)
	if err != nil {
		var pqErr *pq.Error
		ok := errors.As(err, &pqErr)
		if ok {
			switch pqErr.Code {
			case "23503":
				return 0, UserAlreadyExists
			default:
				return 0, err
			}
		}
	}
	return id, err
}

func (s *Auth) Login(userInput models.UserInput) (string, error) {
	userExists, err := s.repository.Auth.GetByPhone(userInput.Phone)
	if err != nil {
		log.Error().Err(err).Msg("ошибка при поиске пользователя")
		return "", BadLoginOrPassword
	}
	if userExists.ID == 0 {
		log.Error().Err(err).Msg("Пользователь не найден")
		return "", BadLoginOrPassword
	}
	if userExists.Password != userInput.Password {
		log.Error().Err(err).Msg("Неправильный пароль")
		return "", BadLoginOrPassword
	}
	return utils.GenerateJWTToken(userExists.ID)
}
