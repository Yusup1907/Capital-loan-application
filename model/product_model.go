package model

import "time"

type ProductModel struct {
	Id                  int       `json:"id"`
	ProductName         string    `json:"product_name" validate:"required,min=3" `
	Description         string    `json:"description"`
	Price               float64   `json:"price" validate:"required,gte=0"`
	Stok                int       `json:"stok" validate:"required"`
	CategoryProductId   int       `json:"category_product_id" validate:"required"`
	CategorProductyName string    `json:"category_product_name"`
	Status              bool      `json:"status" validate:"required"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}