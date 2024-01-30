package repositories

import (
	"FinanceTODO/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(user *models.UserInput) (int, error) {

	stmt := "INSERT INTO users (phone, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id"

	row := r.db.QueryRow(stmt, user.Phone, user.Username, user.Email, user.Password)

	var id int

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetByPhone(phone string) (models.UserOutput, error) {
	var user models.UserOutput

	err := r.db.Get(&user, "SELECT * FROM users WHERE phone = $1;", phone)

	if err != nil {
		return models.UserOutput{}, err
	}

	if err != nil {
		return models.UserOutput{}, err
	}

	return user, nil
}
