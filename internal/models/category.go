package models

type Category struct {
	ID            int           `db:"id" json:"id"`
	Title         string        `db:"title" json:"title"`
	Subcategories []Subcategory `json:"sub_categories" db:"sub_categories"`
}

type Subcategory struct {
	ID    int    `db:"sub_category_id" json:"id"`
	Title string `db:"sub_category_title" json:"title"`
}

type SubCategoryInput struct {
	Title string `json:"title" binding:"required"`
}

type SubCategoryOutput struct {
	ID         int    `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	CategoryID int    `json:"category_id" db:"category_id"`
}
