package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          string `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CategoryId  uint           `json:"category_id" form:"category_id"`
	Name        string         `json:"name" form:"name"`
	Description string         `json:"description" form:"description"`
	Stock       int            `json:"stock" form:"stock"`
	Price       int            `json:"price" form:"price"`
}
