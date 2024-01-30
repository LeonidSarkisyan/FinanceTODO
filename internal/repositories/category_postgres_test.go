package repositories

import (
	"FinanceTODO/pkg/tests_utils"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCategoryPostgres(t *testing.T) {
	db := tests_utils.GetDB()

	// _ := NewAuthPostgres(db)
	categoryRepository := NewCategoryPostgres(db)

	t.Run("получение дефолтных категорий", func(t *testing.T) {
		categories, err := categoryRepository.GetDefaultCategories()
		require.NoError(t, err)

		jsonCategories, err := json.Marshal(categories)

		require.NoError(t, err)

		t.Log(string(jsonCategories))
	})
}
