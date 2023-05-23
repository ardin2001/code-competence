package repositories

import (
	"echo_golang/configs"
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

	DB, _ := configs.InitDB()
	check := DB.First(&user, id).Error
	if check != nil {
		return nil, check
	}
	return &user, check
}

func (us *UserRepository) DeleteRepository(id uint) error {
	DB, _ := configs.InitDB()
	check := DB.Delete(&models.User{}, &id).Error

	return check
}

func (us *UserRepository) CreateRepository(user *models.User) (*models.User, error) {
	DB, _ := configs.InitDB()
	check := DB.Save(user).Error
	if check != nil {
		return nil, check
	}
	return user, check
}

func (us *UserRepository) UpdateRepository(userId *models.User, id uint) (*models.User, error) {
	DB, _ := configs.InitDB()
	var user models.User

	err := DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	if userId.Name != "" {
		user.Name = userId.Name
	}
	if userId.Email != "" {
		user.Email = userId.Email
	}
	if userId.Password != "" {
		user.Password = userId.Password
	}

	check := DB.Save(user).Error
	if check != nil {
		return nil, check
	}
	return &user, check
}
