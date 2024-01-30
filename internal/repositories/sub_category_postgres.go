package repositories

import (
	"FinanceTODO/internal/models"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type SubCategoryPostgres struct {
	db *sqlx.DB
}

func NewSubCategoryPostgres(db *sqlx.DB) *SubCategoryPostgres {
	return &SubCategoryPostgres{db: db}
}

func (r *SubCategoryPostgres) Create(subCategory models.SubCategoryInput, userID int) (int, error) {
	var id int

	stmt := `INSERT INTO sub_categories (title, category_id, user_id) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRow(stmt, subCategory.Title, 5, userID).Scan(&id)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "22001" {
				return id, TooLongValueError{
					Field:    "title",
					MaxValue: 255,
				}
			}
		}
	}

	return id, err
}

func (r *SubCategoryPostgres) GetAll(userID int) ([]models.SubCategoryOutput, error) {
	var categories []models.SubCategoryOutput

	query := `SELECT id, title, category_id FROM sub_categories WHERE user_id = $1`

	err := r.db.Select(&categories, query, userID)

	return categories, err
}

func (r *SubCategoryPostgres) GetByID(subCategoryID int, userID int) (models.SubCategoryOutput, error) {
	var subCategory models.SubCategoryOutput

	query := `SELECT id, title, category_id FROM sub_categories WHERE id = $1 AND user_id = $2`

	err := r.db.Get(&subCategory, query, subCategoryID, userID)

	if errors.Is(err, sql.ErrNoRows) {
		return subCategory, NotFoundError
	}

	return subCategory, nil
}

func (r *SubCategoryPostgres) Update(subCategory models.SubCategoryInput, subCategoryID int, userID int) error {
	stmt := `UPDATE sub_categories SET title = $1 WHERE id = $2 AND user_id = $3`

	result, err := r.db.Exec(stmt, subCategory.Title, subCategoryID, userID)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return NotFoundError
	}

	return nil
}

func (r *SubCategoryPostgres) Delete(subCategoryID, userID int) error {
	stmt := `DELETE FROM sub_categories WHERE id = $1 AND user_id = $2`

	result, err := r.db.Exec(stmt, subCategoryID, userID)

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

	return nil
}
