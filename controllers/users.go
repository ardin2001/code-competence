package controllers

import (
	"echo_golang/configs"
	middleware "echo_golang/middlewares"
	"echo_golang/models"
	"echo_golang/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserIntController interface {
	LoginUserController(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateUserController(c echo.Context) error
	DeleteUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
}

type UserStrController struct {
	userR services.UserIntService
}

func NewUserController(uc services.UserIntService) UserIntController {
	return &UserStrController{
		userR: uc,
	}
}
func (us *UserStrController) GetUserController(c echo.Context) error {
	getDataUser := middleware.GetJwtToken(c)
	user, check := us.userR.GetUserService(getDataUser.ID)
	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
			"user":    user,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data user",
		"users":   user,
	})
}

func (us *UserStrController) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	_, check := us.userR.CreateService(&user)

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})

}

func (us *UserStrController) DeleteUserController(c echo.Context) error {
	getDataUser := middleware.GetJwtToken(c)
	check := us.userR.DeleteService(getDataUser.ID)

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    "data user berhasil dihapus",
	})

}

func (us *UserStrController) UpdateUserController(c echo.Context) error {
	getDataUser := middleware.GetJwtToken(c)
	user := models.User{}
	c.Bind(&user)

	dataUser, check := us.userR.UpdateService(&user, getDataUser.ID)

	if check != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": check.Error(),
			"data":    dataUser,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    dataUser,
	})
}

func (us *UserStrController) LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	DB, _ := configs.InitDB()
	err := DB.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "login failed username or password",
			"error":   err.Error(),
		})
	}

	token, _ := middleware.CreateToken(user.ID, user.Name)
	userresponse := models.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"users":   userresponse,
	})
}
