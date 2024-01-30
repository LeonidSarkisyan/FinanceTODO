package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
)

type SubCategory struct {
	repository *repositories.Repository
}

func NewSubCategoryService(repository *repositories.Repository) *SubCategory {
	return &SubCategory{
		repository: repository,
	}
}

func (s *SubCategory) Create(subCategory models.SubCategoryInput, userID int) (int, error) {
	return s.repository.SubCategory.Create(subCategory, userID)
}

func (s *SubCategory) GetAll(userID int) ([]models.SubCategoryOutput, error) {
	return s.repository.SubCategory.GetAll(userID)
}

func (s *SubCategory) GetByID(subCategoryID, userID int) (models.SubCategoryOutput, error) {
	return s.repository.SubCategory.GetByID(subCategoryID, userID)
}

func (s *SubCategory) Update(subCategory models.SubCategoryInput, subCategoryID, userID int) error {
	return s.repository.SubCategory.Update(subCategory, subCategoryID, userID)
}

func (s *SubCategory) Delete(subCategoryID, userID int) error {
	return s.repository.SubCategory.Delete(subCategoryID, userID)
}
