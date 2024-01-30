package repositories

import (
	"FinanceTODO/internal/models"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) GetDefaultCategories() ([]models.Category, error) {
	query := `
		SELECT c.id, c.title, sc.id, sc.title FROM categories c 
		LEFT JOIN sub_categories sc ON c.id = sc.category_id AND c.id != 5
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[int]*models.Category)

	for rows.Next() {
		var rowCategoryID, rowSubCategoryID sql.NullInt64
		var rowCategoryTitle, rowSubCategoryTitle sql.NullString

		err = rows.Scan(&rowCategoryID, &rowCategoryTitle, &rowSubCategoryID, &rowSubCategoryTitle)
		if err != nil {
			return nil, err
		}

		categoryID := int(rowCategoryID.Int64)

		category, exists := categoryMap[categoryID]
		if !exists {
			category = &models.Category{
				ID:    categoryID,
				Title: rowCategoryTitle.String,
			}
			categoryMap[categoryID] = category
		}

		if rowSubCategoryID.Valid {
			subCategory := models.Subcategory{
				ID:    int(rowSubCategoryID.Int64),
				Title: rowSubCategoryTitle.String,
			}
			category.Subcategories = append(category.Subcategories, subCategory)
		}
	}

	var categories []models.Category
	for _, category := range categoryMap {
		categories = append(categories, *category)
	}

	return categories, nil
}
