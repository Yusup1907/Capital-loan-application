package model

import "time"

type ProductModel struct {
	Id                int       `json:"id"`
	ProductName       string    `json:"product_name"`
	Description       string    `json:"description"`
	Price             float64   `json:"price"`
	Stok              int       `json:"stok"`
	CategoryProductId int       `json:"category_product_id"`
	Status            bool      `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Id                int     `json:"id"`
	ProductName       string  `json:"product_name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Stok              int     `json:"stok"`
	CategoryProductId int     `json:"category_product_id"`
	Status            bool    `json:"status"`
}
