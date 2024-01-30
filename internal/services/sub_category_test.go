package services

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/internal/repositories"
	"FinanceTODO/pkg/tests_utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubCategoryService(t *testing.T) {
	db := tests_utils.GetDB()
	repository := repositories.NewRepository(db)
	service := NewService(repository)

	user := models.UserInput{
		Phone:    "12345",
		Password: "password",
	}

	newSubCategory := models.SubCategoryInput{
		Title: "test title",
	}

	t.Run("создание подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(user)

		require.NoError(t, err)

		subCategoryID, err := service.SubCategory.Create(newSubCategory, userID)

		require.NoError(t, err)

		require.NotEmpty(t, subCategoryID)
	})

	t.Run("получение всех подкатегорий", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(user)

		require.NoError(t, err)

		_, err = service.SubCategory.Create(newSubCategory, userID)

		require.NoError(t, err)

		_, err = service.SubCategory.Create(newSubCategory, userID)

		require.NoError(t, err)

		_, err = service.SubCategory.Create(newSubCategory, userID)

		require.NoError(t, err)

		subCategories, err := service.SubCategory.GetAll(userID)

		require.NoError(t, err)
		require.NotEmpty(t, subCategories)
		require.Equal(t, 3, len(subCategories))
	})

	t.Run("получение подкатегории по ID", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(user)

		require.NoError(t, err)

		subCategoryID, err := service.SubCategory.Create(newSubCategory, userID)

		require.NoError(t, err)

		subCategory, err := service.SubCategory.GetByID(subCategoryID, userID)

		require.NoError(t, err)
		require.NotEmpty(t, subCategory.ID)
		require.Equal(t, subCategory.Title, newSubCategory.Title)
		require.Equal(t, subCategory.CategoryID, 5)
	})

	t.Run("получение несуществующей категории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(user)

		require.NoError(t, err)

		_, err = service.SubCategory.GetByID(867543, userID)

		require.ErrorIs(t, err, repositories.NotFoundError)
	})

	t.Run("обновление подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(user)

		require.NoError(t, err)

		subCategoryID, err := service.SubCategory.Create(newSubCategory, userID)

		require.NoError(t, err)

		subCategory := models.SubCategoryInput{
			Title: "new title",
		}

		err = service.SubCategory.Update(subCategory, subCategoryID, userID)

		require.NoError(t, err)

		subCategoryFromDB, err := service.SubCategory.GetByID(subCategoryID, userID)

		require.NoError(t, err)
		require.Equal(t, subCategoryFromDB.Title, subCategory.Title)
		require.Equal(t, subCategoryFromDB.CategoryID, 5)
	})

	t.Run("удаление подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		userID, err := service.Auth.Register(user)

		require.NoError(t, err)

		subCategoryID, err := service.SubCategory.Create(newSubCategory, userID)

		require.NoError(t, err)

		err = service.SubCategory.Delete(subCategoryID, userID)

		require.NoError(t, err)

		subCategory, err := service.SubCategory.GetByID(subCategoryID, userID)

		require.ErrorIs(t, err, repositories.NotFoundError)
		require.Equal(t, subCategory.ID, 0)
	})
}
