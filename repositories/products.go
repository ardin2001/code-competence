package repositories

import (
	"echo_golang/models"
	"errors"

	"gorm.io/gorm"
)

type ProductInterfaceR interface {
	GetProductsRepository(name string) ([]models.Product, error)
	GetProductRepository(id string) (*models.Product, error)
	CreateProductRepository(Product *models.Product) (*models.Product, error)
	DeleteProductRepository(id string) error
	UpdateProductRepository(ProductId *models.Product, id string) (*models.Product, error)
}

type ProductStructR struct {
	DB *gorm.DB
}

func NewProductRepositories(db *gorm.DB) ProductInterfaceR {
	return &ProductStructR{
		DB: db,
	}
}

func (pr *ProductStructR) GetProductsRepository(name string) ([]models.Product, error) {
	var Products []models.Product
	check := pr.DB.Where("name LIKE ?", "%"+name+"%").Find(&Products).Error

	if check != nil {
		return nil, check
	}

	return Products, check
}

func (pr *ProductStructR) GetProductRepository(id string) (*models.Product, error) {
	var product models.Product
	check := pr.DB.First(&product, id).Error
	if check != nil {
		return nil, check
	}
	return &product, check
}

func (pr *ProductStructR) DeleteProductRepository(id string) error {
	if err := pr.DB.Where("id = ?", id).Take(&models.Product{}).Error; err != nil {
		return errors.New("not_found")
	}
	check := pr.DB.Delete(&models.Product{}, &id).Error
	return check
}

func (pr *ProductStructR) CreateProductRepository(cart *models.Product) (*models.Product, error) {
	check := pr.DB.Save(cart).Error
	if check != nil {
		return nil, check
	}
	return cart, check
}

func (pr *ProductStructR) UpdateProductRepository(cartId *models.Product, id string) (*models.Product, error) {
	check := pr.DB.Save(cartId).Error
	if check != nil {
		return nil, check
	}
	return cartId, check
}
