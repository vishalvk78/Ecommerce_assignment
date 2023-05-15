package model

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID          uint    `json:"id" gorm:"primary_key"`
	Category    string  `json:"category"`
	ProductName string  `json:"productname"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Attributes  string  `json:"attributes"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type CreateProduct struct {
	Category    string  `json:"category"`
	ProductName string  `json:"productname"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Attributes  string  `json:"attributes"`
}

type UpdateProducts struct {
	Category    string  `json:"category"`
	ProductName string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
