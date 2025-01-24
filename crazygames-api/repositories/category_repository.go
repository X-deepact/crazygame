package repositories

import (
	"time"

	"crazygames.io/entities"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

type CategoryRepositoryInterface interface {
	Create(category *entities.Category) error
	GetAll() ([]entities.Category, error)
	GetMenu() ([]entities.Category, error)
	GetByID(id uint) (*entities.Category, error)
	Update(category *entities.Category) (*entities.Category, error)
	Delete(id uint) error
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) GetAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := r.db.Find(&categories).Error
	return categories, err
}
func (r *CategoryRepository) GetMenu() ([]entities.Category, error) {
	var categories []entities.Category
	err := r.db.Where("is_menu = ?", true).Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetByID(id uint) (*entities.Category, error) {
	var category entities.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(category *entities.Category) (*entities.Category, error) {
	err := r.db.Model(&entities.Category{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"category_name": category.CategoryName,
		"description":   category.Description,
		"icon":          category.Icon,
		"path":          category.Path,
		"updated_at":    time.Now(),
	}).Error
	if err != nil {
		return nil, err
	}

	// Fetch updated record
	var updatedCategory entities.Category
	err = r.db.First(&updatedCategory, category.ID).Error
	if err != nil {
		return nil, err
	}

	return &updatedCategory, nil
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Category{}, id).Error
}
