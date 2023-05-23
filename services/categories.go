package services

import (
	"echo_golang/models"
	"echo_golang/repositories"
)

type CategoryInterfaceS interface {
	GetCategoriesService() ([]models.Category, error)
	GetCategoryService(id string) (*models.Category, error)
	CreateCategoryService(category *models.Category) (*models.Category, error)
	UpdateCategoryService(categoryId *models.Category, id string) (*models.Category, error)
	DeleteCategoryService(id string) error
}

type CategoryStructS struct {
	categoryR repositories.CategoryInterfaceR
}

func NewCategoryServices(categoryR repositories.CategoryInterfaceR) CategoryInterfaceS {
	return &CategoryStructS{
		categoryR: categoryR,
	}
}

func (cs *CategoryStructS) GetCategoriesService() ([]models.Category, error) {
	categories, err := cs.categoryR.GetCategoriesRepository()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryStructS) GetCategoryService(id string) (*models.Category, error) {
	category, err := cs.categoryR.GetCategoryRepository(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *CategoryStructS) CreateCategoryService(category *models.Category) (*models.Category, error) {

	categoryR, err := cs.categoryR.CreateCategoryRepository(category)
	if err != nil {
		return nil, err
	}

	return categoryR, nil
}

func (cs *CategoryStructS) UpdateCategoryService(categoryId *models.Category, id string) (*models.Category, error) {
	getCategoryId, err := cs.categoryR.GetCategoryRepository(id)

	if err != nil {
		return nil, err
	}

	if categoryId.Name != "" {
		getCategoryId.Name = categoryId.Name
	}

	category, err := cs.categoryR.UpdateCategoryRepository(getCategoryId, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (u *CategoryStructS) DeleteCategoryService(id string) error {
	err := u.categoryR.DeleteCategoryRepository(id)
	if err != nil {
		return err
	}

	return nil
}
