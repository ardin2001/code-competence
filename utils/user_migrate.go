package utils

import (
	"echo_golang/configs"
	"echo_golang/models"
	"fmt"
)

func UserMigrate() {
	DB, err := configs.InitDB()
	if err != nil {
		fmt.Println("Failed connect to database : ", err.Error())
		return
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
