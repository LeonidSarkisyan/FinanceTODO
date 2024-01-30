package models

import "time"

type Balance struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Type            string    `json:"type"`
	Value           float64   `json:"value"`
	Currency        string    `json:"currency"`
	CreatedDateTime time.Time `json:"created_date_time" db:"created_date_time"`
	UpdatedDateTime time.Time `json:"updated_date_time" db:"updated_date_time"`
	UserID          int       `json:"user_id" db:"user_id"`
}

type BalanceInput struct {
	Title    string  `json:"title" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	Value    float64 `json:"value" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type BalanceUpdate struct {
	Title    *string `json:"title"`
	Type     *string `json:"type"`
	Currency *string `json:"currency"`
}

func (b *BalanceUpdate) IsValid() bool {
	if b.Title == nil && b.Type == nil && b.Currency == nil {
		return false
	}
	return true
}
