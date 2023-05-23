package controllers

import (
	"echo_golang/configs"
	"echo_golang/helpers"
	middleware "echo_golang/middlewares"
	"echo_golang/models"
	"echo_golang/services"
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
	user, err := uc.userR.GetUserService(getDataUser.ID)
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    user,
		Message: "Successfully Get Data",
		Status:  true,
	})
}

func (us *UserStrController) CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err := us.userR.CreateService(&user)

	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    user,
		Message: "Successfully Created Data",
		Status:  true,
	})

}

func (us *UserStrController) DeleteUserController(c echo.Context) error {
	getDataUser := middleware.GetJwtToken(c)
	err := us.userR.DeleteService(getDataUser.ID)

	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    getDataUser,
		Message: "Successfully deleted Data",
		Status:  true,
	})

}

func (us *UserStrController) UpdateUserController(c echo.Context) error {
	getDataUser := middleware.GetJwtToken(c)
	user := models.User{}
	c.Bind(&user)

	dataUser, err := us.userR.UpdateService(&user, getDataUser.ID)

	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    dataUser,
		Message: "Successfully Updated Data",
		Status:  true,
	})

}

func (us *UserStrController) LoginUserController(c echo.Context) error {
	userRequest := models.User{}
	user := models.User{}
	c.Bind(&userRequest)

	DB, _ := configs.InitDB()
	err := DB.Where("name = ?", userRequest.Name).First(&user).Error
	if err != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "Wrong username",
			Status:  false,
		})
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if errPassword != nil {
		return helpers.Response(c, http.StatusBadRequest, helpers.ResponseModel{
			Data:    nil,
			Message: "Wrong password",
			Status:  false,
		})
	}

	token, _ := middleware.CreateToken(user.ID, user.Name)
	userResponse := models.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}
	return helpers.Response(c, http.StatusOK, helpers.ResponseModel{
		Data:    userResponse,
		Message: "Login Success",
		Status:  true,
	})
}
