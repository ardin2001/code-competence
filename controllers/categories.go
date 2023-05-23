package controllers

import (
	"echo_golang/helpers"
	"echo_golang/models"
	"echo_golang/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryInterfaceC interface {
	GetCategoriesController(c echo.Context) error
	GetCategoryController(c echo.Context) error
	CreateCategoryController(c echo.Context) error
	UpdateCategoryController(c echo.Context) error
	DeleteCategoryController(c echo.Context) error
}

type CategoryStructC struct {
	categoryS services.CategoryInterfaceS
}

func NewCategoryControllers(categoryS services.CategoryInterfaceS) CategoryInterfaceC {
	return &CategoryStructC{
		categoryS,
	}
}

func (cc *CategoryStructC) GetCategoriesController(c echo.Context) error {
	categorys, check := cc.categoryS.GetCategoriesService()
	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    categorys,
		Message: "Successfull get all categories",
		Status:  true,
	})
}

func (cc *CategoryStructC) GetCategoryController(c echo.Context) error {
	id := c.Param("category_id")
	category, check := cc.categoryS.GetCategoryService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    category,
		Message: "Successfull get Category",
		Status:  true,
	})
}

func (cc *CategoryStructC) CreateCategoryController(c echo.Context) error {

	category := models.Category{}
	c.Bind(&category)

	_, check := cc.categoryS.CreateCategoryService(&category)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    category,
		Message: "Successfull create Category",
		Status:  true,
	})
}

func (cc *CategoryStructC) UpdateCategoryController(c echo.Context) error {

	id := c.Param("id")
	category := models.Category{}
	c.Bind(&category)

	dataCategory, check := cc.categoryS.UpdateCategoryService(&category, id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    dataCategory,
		Message: "Successfull update Category",
		Status:  true,
	})
}

func (cc *CategoryStructC) DeleteCategoryController(c echo.Context) error {

	id := c.Param("id")
	check := cc.categoryS.DeleteCategoryService(id)

	if check != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: check.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    id,
		Message: "Successfull delete Category",
		Status:  true,
	})
}
