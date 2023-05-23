package services

import (
	"echo_golang/models"
	"echo_golang/repositories"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserIntService interface {
	GetUserService(id uint) (*models.User, error)
	CreateService(user *models.User) (*models.User, error)
	UpdateService(userId *models.User, id uint) (*models.User, error)
	DeleteService(id uint) error
}

type userStrService struct {
	userR repositories.UserRepositories
}

func NewUserServices(userR repositories.UserRepositories) UserIntService {
	return &userStrService{
		userR: userR,
	}
}

func (u *userStrService) GetUserService(id uint) (*models.User, error) {
	user, err := u.userR.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userStrService) CreateService(user *models.User) (*models.User, error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}

func (u *userStrService) UpdateService(userId *models.User, id uint) (*models.User, error) {
	getUserId, err := u.userR.GetUserRepository(id)

	if err != nil {
		return nil, err
	}

	if userId.Name != "" {
		getUserId.Name = userId.Name
	}
	if userId.Email != "" {
		fmt.Println("masuk email")
		getUserId.Email = userId.Email
	}
	if userId.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userId.Password), bcrypt.DefaultCost)
		getUserId.Password = string(hashedPassword)
	}

	user, err := u.userR.UpdateRepository(getUserId, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userStrService) DeleteService(id uint) error {
	err := u.userR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
