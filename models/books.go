package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"name" form:"name"`
	Author string `json:"email" form:"email"`
}
