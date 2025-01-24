package repositories

import (
	"strconv"
	"testing"

	"crazygames.io/entities"
	"github.com/stretchr/testify/assert"
)

func createCategory(categoryName string, isMenu bool) (*entities.Category, error) {
	category := &entities.Category{
		CategoryName: categoryName,
		Description:  "description",
		Icon:         "icon",
		Path:         "path",
		IsMenu:       isMenu,
	}
	err := categoryRepository.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func Test_CreateCategory(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE game_categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	t.Run("create category with isMenu true should succeed", func(t *testing.T) {
		category, err := createCategory("category1", true)
		assert.NoError(t, err, "failed to create category for test")
		assert.Equal(t, category.CategoryName, "category1")
		assert.Equal(t, category.IsMenu, true)
	})

	t.Run("create category with isMenu false should succeed", func(t *testing.T) {
		category, err := createCategory("category2", false)
		assert.NoError(t, err, "failed to create category for test")
		assert.Equal(t, category.CategoryName, "category2")
		assert.Equal(t, category.IsMenu, false)
	})

	t.Run("create category without categoryName should fail", func(t *testing.T) {
		category := &entities.Category{
			Description: "description",
			Icon:        "icon",
			Path:        "path",
		}
		err := categoryRepository.Create(category)
		assert.Error(t, err, "failed to create category for test")
	})

	t.Run("create category with the same categoryName should fail", func(t *testing.T) {
		_, err := createCategory("category3", true)
		assert.NoError(t, err, "failed to create category for test")
		_, err = createCategory("category3", false)
		assert.Error(t, err, "failed to create category for test")
	})
}

func Test_GetAllCategory(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE game_categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	for i := 1; i <= 10; i++ {
		_, err := createCategory("category"+strconv.Itoa(i), false)
		assert.NoError(t, err, "failed to create category for test")
	}

	categories, err := categoryRepository.GetAll()
	assert.NoError(t, err, "failed to get all categories for test")

	assert.Equal(t, len(categories), 10)
}

func Test_GetCategoryMenu(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE game_categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	for i := 1; i <= 5; i++ {
		_, err := createCategory("category"+strconv.Itoa(i), false)
		assert.NoError(t, err, "failed to create category for test")
	}

	for i := 6; i <= 10; i++ {
		_, err := createCategory("category"+strconv.Itoa(i), true)
		assert.NoError(t, err, "failed to create category for test")
	}

	categories, err := categoryRepository.GetMenu()
	assert.NoError(t, err, "failed to get all categories for test")

	assert.Equal(t, len(categories), 5)
}

func Test_GetCategoryById(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE game_categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	category, err := createCategory("category", false)
	assert.NoError(t, err, "failed to get all categories for test")

	t.Run("get category by valid id should succeed", func(t *testing.T) {
		retrievedCategory, err := categoryRepository.GetByID(category.ID)
		assert.NoError(t, err, "failed to fetch category by ID")
		assert.Equal(t, category.CategoryName, retrievedCategory.CategoryName)
	})

	t.Run("get category by invalid id should fail", func(t *testing.T) {
		_, err := categoryRepository.GetByID(uint(100_000))
		assert.Error(t, err, "failed to fetch category by ID")
	})
}

func Test_UpdateCategory(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE game_categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	category, err := createCategory("category", false)
	assert.NoError(t, err, "failed to get all categories for test")

	t.Run("update category by valid id should succeed", func(t *testing.T) {
		category, err := categoryRepository.Update(&entities.Category{
			ID:           category.ID,
			CategoryName: "update category",
		})
		assert.NoError(t, err, "failed to update category for test")
		assert.Equal(t, category.CategoryName, "update category")
	})

	t.Run("update category by invalid id should fail", func(t *testing.T) {
		_, err := categoryRepository.Update(&entities.Category{
			ID:           uint(100_000),
			CategoryName: "update category",
		})
		assert.Error(t, err, "failed to update category for test")
	})
}

func Test_DeleteCategory(t *testing.T) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("TRUNCATE TABLE games;")
	db.Exec("TRUNCATE TABLE game_categories;")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")

	category, err := createCategory("category", false)
	assert.NoError(t, err, "failed to get all categories for test")

	t.Run("delete category by valid id should succeed", func(t *testing.T) {
		err := categoryRepository.Delete(category.ID)
		assert.NoError(t, err, "failed to delete category for test")
	})

	t.Run("delete category by invalid id should succeed", func(t *testing.T) {
		err := categoryRepository.Delete(uint(100_000))
		assert.NoError(t, err, "failed to delete category for test")
	})
}
