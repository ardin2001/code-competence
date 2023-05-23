package controllers

import (
	"echo_golang/configs"
	"echo_golang/helpers"
	middleware "echo_golang/middlewares"
	"echo_golang/models"
	"echo_golang/services"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

func NewUserControllers(uc services.UserIntService) UserIntController {
	return &UserStrController{
		userR: uc,
	}
}
func (uc *UserStrController) GetUserController(c echo.Context) error {
	getDataUser := middleware.GetJwtToken(c)
	user, check := uc.userR.GetUserService(getDataUser.ID)
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

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
	userRequest := models.User{}
	user := models.User{}
	c.Bind(&userRequest)
	fmt.Println(userRequest)

	DB, _ := configs.InitDB()
	err := DB.Where("name = ?", userRequest.Name).First(&user).Error
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "wrong username",
			Status:  false,
		})
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if errPassword != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "wrong password",
			Status:  false,
		})
	}

	token, _ := middleware.CreateToken(user.ID, user.Name)
	userresponse := models.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"users":   userresponse,
	})
}
