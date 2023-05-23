package services

import (
	"echo_golang/models"
	"echo_golang/repositories"
)

type ProductInterfaceS interface {
	GetProductsService(name string) ([]models.Product, error)
	GetProductService(id string) (*models.Product, error)
	CreateProductService(Product *models.Product) (*models.Product, error)
	UpdateProductService(ProductId *models.Product, id string) (*models.Product, error)
	DeleteProductService(id string) error
}

type ProductStructS struct {
	productR repositories.ProductInterfaceR
}

func NewProductServices(productR repositories.ProductInterfaceR) ProductInterfaceS {
	return &ProductStructS{
		productR: productR,
	}
}

func (ps *ProductStructS) GetProductsService(name string) ([]models.Product, error) {
	products, err := ps.productR.GetProductsRepository(name)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductStructS) GetProductService(id string) (*models.Product, error) {
	product, err := ps.productR.GetProductRepository(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductStructS) CreateProductService(product *models.Product) (*models.Product, error) {

	productR, err := ps.productR.CreateProductRepository(product)
	if err != nil {
		return nil, err
	}

	return productR, nil
}

func (ps *ProductStructS) UpdateProductService(productId *models.Product, id string) (*models.Product, error) {
	getProductId, err := ps.productR.GetProductRepository(id)

	if err != nil {
		return nil, err
	}

	if productId.Name != "" {
		getProductId.Name = productId.Name
	}
	if productId.Description != "" {
		getProductId.Description = productId.Description
	}
	if productId.Stock != 0 {
		getProductId.Stock = productId.Stock
	}
	if productId.Price != 0 {
		getProductId.Price = productId.Price
	}

	user, err := ps.productR.UpdateProductRepository(getProductId, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *ProductStructS) DeleteProductService(id string) error {
	err := u.productR.DeleteProductRepository(id)
	if err != nil {
		return err
	}

	return nil
}
