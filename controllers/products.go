package controllers

import (
	"echo_golang/helpers"
	"echo_golang/models"
	"echo_golang/services"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductInterfaceC interface {
	GetProductsController(c echo.Context) error
	GetProductController(c echo.Context) error
	CreateProductController(c echo.Context) error
	UpdateProductController(c echo.Context) error
	DeleteProductController(c echo.Context) error
}

type ProductStructC struct {
	productS services.ProductInterfaceS
}

func NewProductControllers(productS services.ProductInterfaceS) ProductInterfaceC {
	return &ProductStructC{
		productS,
	}
}

func (pc *ProductStructC) GetProductsController(c echo.Context) error {
	getName := c.QueryParam("keyword")
	products, check := pc.productS.GetProductsService(getName)
	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    products,
		Message: "Successfull get all products by name " + getName,
		Status:  true,
	})
}

func (pc *ProductStructC) GetProductController(c echo.Context) error {
	id := c.Param("id")
	product, check := pc.productS.GetProductService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    product,
		Message: "Successfull get product",
		Status:  true,
	})
}

func (pc *ProductStructC) CreateProductController(c echo.Context) error {

	product := models.Product{}
	c.Bind(&product)

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	product.ID = uuid

	_, check := pc.productS.CreateProductService(&product)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    product,
		Message: "Successfull create product",
		Status:  true,
	})
}

func (pc *ProductStructC) UpdateProductController(c echo.Context) error {

	id := c.Param("id")
	product := models.Product{}
	c.Bind(&product)

	dataProduct, check := pc.productS.UpdateProductService(&product, id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    dataProduct,
		Message: "Successfull update product",
		Status:  true,
	})
}

func (pc *ProductStructC) DeleteProductController(c echo.Context) error {

	id := c.Param("id")
	check := pc.productS.DeleteProductService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    id,
		Message: "Successfull delete product",
		Status:  true,
	})
}
