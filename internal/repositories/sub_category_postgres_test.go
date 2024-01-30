package repositories

import (
	"FinanceTODO/internal/models"
	"FinanceTODO/pkg/tests_utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubCategoryPostgres(t *testing.T) {
	db := tests_utils.GetDB()

	authRepository := NewAuthPostgres(db)
	scRepository := NewSubCategoryPostgres(db)

	t.Run("создание подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		subCategory := models.SubCategoryInput{
			Title: "test title",
		}

		subCategoryID, err := scRepository.Create(subCategory, userID)

		require.NoError(t, err)
		require.NotEmpty(t, subCategoryID)

		subCategoryID, err = scRepository.Create(subCategory, userID)

		require.NoError(t, err)
		require.NotEmpty(t, subCategoryID)
	})

	t.Run("получение всех подкатегорий", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		subCategory := models.SubCategoryInput{
			Title: "test title",
		}

		_, err = scRepository.Create(subCategory, userID)
		require.NoError(t, err)

		_, err = scRepository.Create(subCategory, userID)
		require.NoError(t, err)

		_, err = scRepository.Create(subCategory, userID)
		require.NoError(t, err)

		subCategories, err := scRepository.GetAll(userID)

		require.NoError(t, err)
		require.Equal(t, 3, len(subCategories))

		for _, subCategoryFromDB := range subCategories {
			require.NotEmpty(t, subCategoryFromDB.ID)
			require.Equal(t, subCategoryFromDB.Title, subCategory.Title)
			require.Equal(t, subCategoryFromDB.CategoryID, 5)
		}
	})

	t.Run("получение своей подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		subCategory := models.SubCategoryInput{
			Title: "test title",
		}

		subCategoryID, err := scRepository.Create(subCategory, userID)

		require.NoError(t, err)

		subCategoryFromDB, err := scRepository.GetByID(subCategoryID, userID)

		require.NoError(t, err)
		require.Equal(t, subCategoryFromDB.Title, subCategory.Title)
		require.Equal(t, subCategoryFromDB.ID, subCategoryID)
		require.Equal(t, subCategoryFromDB.CategoryID, 5)
	})

	t.Run("получение несуществующей подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		_, err = scRepository.GetByID(123456, userID)

		require.ErrorIs(t, err, NotFoundError)
	})

	t.Run("получение не своей подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user1 := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID1, err := authRepository.Create(&user1)
		require.NoError(t, err)

		user2 := models.UserInput{
			Phone:    "9876543210",
			Password: "password",
		}

		userID2, err := authRepository.Create(&user2)
		require.NoError(t, err)

		newSubCategory := models.SubCategoryInput{
			Title: "test title",
		}

		subCategoryID, err := scRepository.Create(newSubCategory, userID1)
		require.NoError(t, err)

		subCategory, err := scRepository.GetByID(subCategoryID, userID2)

		require.Equal(t, subCategory.Title, "")
		require.Equal(t, subCategory.ID, 0)
		require.Equal(t, subCategory.CategoryID, 0)

		require.ErrorIs(t, err, NotFoundError)
	})

	t.Run("создание категории со слишком длинным названием", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user1 := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID1, err := authRepository.Create(&user1)
		require.NoError(t, err)

		newSubCategory := models.SubCategoryInput{
			Title: "",
		}

		for i := 0; i < 256; i++ {
			newSubCategory.Title += "a"
		}

		_, err = scRepository.Create(newSubCategory, userID1)

		require.ErrorIs(t, err, TooLongValueError{"title", 255})
	})

	t.Run("успешное обновление подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		subCategory := models.SubCategoryInput{
			Title: "test title",
		}

		subCategoryID, err := scRepository.Create(subCategory, userID)
		require.NoError(t, err)

		newSubCategory := models.SubCategoryInput{
			Title: "new title",
		}

		err = scRepository.Update(newSubCategory, subCategoryID, userID)

		require.NoError(t, err)

		subCategoryUpdated, err := scRepository.GetByID(subCategoryID, userID)

		require.NoError(t, err)
		require.Equal(t, subCategoryUpdated.ID, subCategoryID)
		require.Equal(t, subCategoryUpdated.Title, newSubCategory.Title)
	})

	t.Run("неуспешное обновление подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		subCategory := models.SubCategoryInput{
			Title: "test title",
		}

		subCategoryID, err := scRepository.Create(subCategory, userID)
		require.NoError(t, err)

		newSubCategory := models.SubCategoryInput{
			Title: "new title",
		}

		otherSubCategoryID := 9999

		err = scRepository.Update(newSubCategory, otherSubCategoryID, userID)

		require.ErrorIs(t, err, NotFoundError)

		subCategoryUpdated, err := scRepository.GetByID(subCategoryID, userID)
		require.NoError(t, err)

		require.Equal(t, subCategoryUpdated.ID, subCategoryID)
		require.Equal(t, subCategoryUpdated.Title, subCategory.Title)
		require.NotEqual(t, subCategoryUpdated.Title, newSubCategory.Title)
	})

	t.Run("успешное удаление подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		newSubCategory := models.SubCategoryInput{
			Title: "test title",
		}

		subCategoryID, err := scRepository.Create(newSubCategory, userID)

		require.NoError(t, err)

		subCategory, err := scRepository.GetByID(subCategoryID, userID)

		require.NoError(t, err)

		require.Equal(t, subCategory.ID, subCategoryID)
		require.Equal(t, subCategory.Title, newSubCategory.Title)

		err = scRepository.Delete(subCategoryID, userID)

		require.NoError(t, err)

		subCategory, err = scRepository.GetByID(subCategoryID, userID)

		require.ErrorIs(t, err, NotFoundError)
	})

	t.Run("неуспешное удаление подкатегории", func(t *testing.T) {
		tests_utils.ClearTestDatabase(db)
		user := models.UserInput{
			Phone:    "1234567890",
			Password: "password",
		}
		userID, err := authRepository.Create(&user)
		require.NoError(t, err)

		newSubCategory := models.SubCategoryInput{
			Title: "test title",
		}

		subCategoryID, err := scRepository.Create(newSubCategory, userID)

		require.NoError(t, err)

		subCategory, err := scRepository.GetByID(subCategoryID, userID)

		require.NoError(t, err)

		require.Equal(t, subCategory.ID, subCategoryID)
		require.Equal(t, subCategory.Title, newSubCategory.Title)

		otherSubCategoryID := 2000

		err = scRepository.Delete(otherSubCategoryID, userID)

		require.ErrorIs(t, err, NotFoundError)
	})
}
