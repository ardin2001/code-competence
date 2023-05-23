package repositories

import (
	"echo_golang/models"

	"gorm.io/gorm"
)

type UserRepositories interface {
	GetUserRepository(id uint) (*models.User, error)
	CreateRepository(User *models.User) (*models.User, error)
	DeleteRepository(id uint) error
	UpdateRepository(userId *models.User, id uint) (*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepositories(db *gorm.DB) UserRepositories {
	return &UserRepository{
		DB: db,
	}
}

func (us *UserRepository) GetUserRepository(id uint) (*models.User, error) {
	var user models.User

	check := us.DB.First(&user, id).Error
	if check != nil {
		return nil, check
	}
	return &user, check
}

func (us *UserRepository) DeleteRepository(id uint) error {

	check := us.DB.Delete(&models.User{}, &id).Error

	return check
}

func (us *UserRepository) CreateRepository(user *models.User) (*models.User, error) {

	check := us.DB.Save(user).Error
	if check != nil {
		return nil, check
	}
	return user, check
}

func (us *UserRepository) UpdateRepository(userId *models.User, id uint) (*models.User, error) {
	check := us.DB.Save(&userId).Error
	if check != nil {
		return nil, check
	}
	return userId, check
}
