package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
)

type User struct {
	repository *repositories.Repository
}

func NewUserService(repository *repositories.Repository) *User {
	return &User{repository: repository}
}

func (u *User) GetById(userID int) (models.UserPublic, error) {
	user, err := u.repository.User.GetByID(userID)
	if err != nil {
		return models.UserPublic{}, err
	}
	return user, nil
}
