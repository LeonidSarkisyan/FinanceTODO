package repositories

import (
	"FinanceTODO/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"strings"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) Create(input models.Transaction) (int, error) {
	var id int

	stmt := "INSERT INTO transactions (title, description, type, value, balance_id, sub_category_id, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	err := r.db.QueryRow(stmt, input.Title, input.Description, input.Type, input.Value, input.BalanceId, input.SubCategoryId, input.UserId).Scan(&id)

	return id, err
}

func (r *TransactionPostgres) GetAll(balanceID, userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	log.Info().Int("balanceID", balanceID).Int("userID", userID).Msg("GetAll")
	log.Info().Int("userID", userID).Msg("GetAll")

	query := "SELECT id, title, description, type, value, created_date_time, updated_date_time, balance_id, sub_category_id, user_id FROM transactions WHERE balance_id = $1 AND user_id = $2"

	err := r.db.Select(&transactions, query, balanceID, userID)

	return transactions, err
}

func (r *TransactionPostgres) GetByID(id int, userId int) (models.Transaction, error) {
	var transactions models.Transaction

	query := "SELECT id, title, description, type, value, created_date_time, updated_date_time, balance_id, sub_category_id, user_id FROM transactions WHERE id = $1 AND user_id = $2"

	err := r.db.Get(&transactions, query, id, userId)

	return transactions, err
}

func (r *TransactionPostgres) Update(input models.TransactionUpdate, transactionId int, userId int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, "title=$1")
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, "description=$2")
		args = append(args, *input.Description)
		argID++
	}

	if input.Type != nil {
		setValues = append(setValues, "type=$3")
		args = append(args, *input.Type)
		argID++
	}

	if input.Value != nil {
		setValues = append(setValues, "value=$4")
		args = append(args, *input.Value)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	stmt := fmt.Sprintf("UPDATE transactions SET %s WHERE id = $%d AND user_id = $%d", setQuery, argID, argID+1)

	args = append(args, transactionId, userId)

	result, err := r.db.Exec(stmt, args...)

	row, _ := result.RowsAffected()

	if row == 0 {
		return NotFoundError
	}

	return err
}

func (r *TransactionPostgres) Delete(transactionId int, userId int) error {
	query := "DELETE FROM transactions WHERE id = $1 AND user_id = $2"

	result, err := r.db.Exec(query, transactionId, userId)

	row, _ := result.RowsAffected()

	if row == 0 {
		return NotFoundError
	}

	return err
}
