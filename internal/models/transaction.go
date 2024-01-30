package models

import "time"

type Transaction struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Type            string    `json:"type"`
	Value           float64   `json:"value"`
	CreatedDateTime time.Time `json:"created_date_time" db:"created_date_time"`
	UpdatedDateTime time.Time `json:"updated_date_time" db:"updated_date_time"`
	BalanceId       int       `json:"balance_id" db:"balance_id"`
	SubCategoryId   int       `json:"sub_category_id" db:"sub_category_id"`
	UserId          int       `json:"user_id" db:"user_id"`
}

type TransactionInput struct {
	Title         string  `json:"title" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Type          string  `json:"type" binding:"required"`
	Value         float64 `json:"value" binding:"required"`
	SubCategoryID int     `json:"sub_category_id" binding:"required"`
}

type TransactionUpdate struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Type        *string  `json:"type"`
	Value       *float64 `json:"value"`
}

func NewTransaction(input TransactionInput, balanceId int, subCategoryId int, userId int) *Transaction {
	return &Transaction{
		Title:         input.Title,
		Description:   input.Description,
		Type:          input.Type,
		Value:         input.Value,
		BalanceId:     balanceId,
		SubCategoryId: subCategoryId,
		UserId:        userId,
	}
}
