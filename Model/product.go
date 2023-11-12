package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id"`
	Category    Category `json:"category"`
}

type Category struct {
	gorm.Model
	Name string `json:"name"`
}