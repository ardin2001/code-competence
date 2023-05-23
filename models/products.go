package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryId  uint   `json:"category_id" form:"category_id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
}
