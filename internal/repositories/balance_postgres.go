package repositories

import (
	"FinanceTODO/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) Create(balance models.BalanceInput, userID int) (int, error) {
	var id int
	stmt := `INSERT INTO balances (title, type, value, currency, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRow(stmt, balance.Title, balance.Type, balance.Value, balance.Currency, userID).Scan(&id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "22001" {
				return id, TooLongValueError{
					Field:    "title",
					MaxValue: 16,
				}
			}
		}
	}

	return id, err
}

func (r *BalancePostgres) GetAll(userID int) ([]models.Balance, error) {
	var balances []models.Balance

	query := `
	SELECT id, title, type, value, currency, created_date_time, updated_date_time, user_id FROM balances WHERE user_id = $1
	`

	err := r.db.Select(&balances, query, userID)

	return balances, err
}

func (r *BalancePostgres) GetByID(balanceID int, userID int) (models.Balance, error) {
	var balance models.Balance

	query := `
	SELECT id, title, type, value, currency, created_date_time, updated_date_time, user_id FROM balances WHERE id = $1 AND user_id = $2
	`

	err := r.db.Get(&balance, query, balanceID, userID)

	if errors.Is(err, sql.ErrNoRows) {
		return balance, NotFoundError
	}

	return balance, err
}

func (r *BalancePostgres) Update(balanceUpdate models.BalanceUpdate, balanceID, userID int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if balanceUpdate.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *balanceUpdate.Title)
		argID++
	}

	if balanceUpdate.Type != nil {
		setValues = append(setValues, fmt.Sprintf("type=$%d", argID))
		args = append(args, *balanceUpdate.Type)
		argID++
	}

	if balanceUpdate.Currency != nil {
		setValues = append(setValues, fmt.Sprintf("currency=$%d", argID))
		args = append(args, *balanceUpdate.Currency)
		argID++
	}

	stmt := fmt.Sprintf("UPDATE balances SET %s WHERE id = $%d AND user_id = $%d", strings.Join(setValues, ", "), argID, argID+1)

	args = append(args, balanceID, userID)

	result, err := r.db.Exec(stmt, args...)

	if err != nil {
		return err
	}

	if result != nil {
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return NotFoundError
		}
	}

	return err
}

func (r *BalancePostgres) UpdateValue(amount float64, balanceID, userID int) error {
	stmt := `UPDATE balances SET value = value + $1 WHERE id = $2 AND user_id = $3`

	_, err := r.db.Exec(stmt, amount, balanceID, userID)

	return err
}

func (r *BalancePostgres) Delete(balanceID, userID int) error {
	stmt := `DELETE FROM balances WHERE id = $1 AND user_id = $2`

	result, err := r.db.Exec(stmt, balanceID, userID)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return NotFoundError
	}

	return nil
}
