package repository

import (
	"errors"
	"product-service/internal/app/entity"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepositoryInterface interface {
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id uint) (*entity.Category, error)
	CreateCategory(category *entity.Category) error
	DeleteCategory(id uint) error
}

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.DB.Where("deleted_at IS NULL").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(id uint) (*entity.Category, error) {
	var category entity.Category
	if err := r.DB.Where("id = ?", id).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Category not found
		}
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) CreateCategory(category *entity.Category) error {
	if err := r.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) DeleteCategory(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.Category{}).Error; err != nil {
		return err
	}
	return nil
}
