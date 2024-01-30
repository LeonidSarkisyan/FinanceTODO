package repositories

import (
	"FinanceTODO/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetByID(userID int) (models.UserPublic, error) {
	var user models.UserPublic
	err := r.db.Get(
		&user, "SELECT id, phone, username, email, created_datetime, updated_datetime FROM users WHERE id = $1", userID)
	return user, err
}
