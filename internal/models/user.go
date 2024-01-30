package models

import "time"

type UserInput struct {
	Phone    string `json:"phone" binding:"required"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type UserPublic struct {
	ID              int       `json:"id" db:"id"`
	Phone           string    `json:"phone" db:"phone"`
	Username        string    `json:"username" db:"username"`
	Email           string    `json:"email" db:"email"`
	CreatedDatetime time.Time `db:"created_datetime"`
	UpdatedDatetime time.Time `db:"updated_datetime"`
}

type UserOutput struct {
	ID              int       `json:"id" db:"id"`
	Phone           string    `json:"phone" db:"phone"`
	Username        string    `json:"username" db:"username"`
	Email           string    `json:"email" db:"email"`
	Password        string    `json:"password" db:"password"`
	CreatedDatetime time.Time `db:"created_datetime"`
	UpdatedDatetime time.Time `db:"updated_datetime"`
}
