package repositories

import (
	"echo_golang/models"
	"errors"

	"gorm.io/gorm"
)

type CategoryInterfaceR interface {
	GetCategoriesRepository() ([]models.Category, error)
	GetCategoryRepository(id string) (*models.Category, error)
	CreateCategoryRepository(Category *models.Category) (*models.Category, error)
	DeleteCategoryRepository(id string) error
	UpdateCategoryRepository(CategoryId *models.Category, id string) (*models.Category, error)
}

type CategoryStructR struct {
	DB *gorm.DB
}

func NewCategoryRepositories(db *gorm.DB) CategoryInterfaceR {
	return &CategoryStructR{
		DB: db,
	}
}

func (cr *CategoryStructR) GetCategoriesRepository() ([]models.Category, error) {
	var categories []models.Category
	check := cr.DB.Find(&categories).Error

	if check != nil {
		return nil, check
	}

	return categories, check
}

func (cr *CategoryStructR) GetCategoryRepository(id string) (*models.Category, error) {
	var category models.Category
	check := cr.DB.Preload("Products").First(&category, id).Error
	if check != nil {
		return nil, check
	}
	return &category, check
}

func (cr *CategoryStructR) DeleteCategoryRepository(id string) error {
	if err := cr.DB.Where("id = ?", id).Take(&models.Category{}).Error; err != nil {
		return errors.New("not_found")
	}
	check := cr.DB.Delete(&models.Category{}, &id).Error
	return check
}

func (cr *CategoryStructR) CreateCategoryRepository(category *models.Category) (*models.Category, error) {
	check := cr.DB.Save(category).Error
	if check != nil {
		return nil, check
	}
	return category, check
}

func (cr *CategoryStructR) UpdateCategoryRepository(categoryId *models.Category, id string) (*models.Category, error) {
	check := cr.DB.Save(categoryId).Error
	if check != nil {
		return nil, check
	}
	return categoryId, check
}
