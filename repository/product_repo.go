package repository

import (
	"database/sql"
	"pinjam-modal-app/model"
)

type ProductRepo interface {
	CreateProduct(newProduct *model.ProductModel) error
}

type productRepo struct {
	db *sql.DB
}

func (p *productRepo) CreateProduct(newProduct *model.ProductModel) error {
	insertStatement := "INSERT INTO mst_product (product_name, description, price, stok, category_product_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := p.db.QueryRow(insertStatement, newProduct.ProductName, newProduct.Description, newProduct.Price, newProduct.Stok, newProduct.CategoryProductId, newProduct.Status, newProduct.CreatedAt, newProduct.UpdatedAt).Scan(&newProduct.Id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}
