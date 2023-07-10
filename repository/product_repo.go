package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
)

type ProductRepo interface {
	CreateProduct(newProduct *model.ProductModel) error
	GetProductByName(nameProduct string) (*model.ProductModel, error)
	GetAllProduct() ([]*model.ProductModel, error)
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

func (p *productRepo) GetAllProduct() ([]*model.ProductModel, error) {
	selectStatement := "SELECT id, product_name, description, price, stok, category_product_id, status, created_at, updated_at FROM mst_product ORDER BY id ASC"

	rows, err := p.db.Query(selectStatement)
	if err != nil {
		return nil, fmt.Errorf("GetAllProduct() : %w", err)

	}
	defer rows.Close()

	var products []*model.ProductModel
	for rows.Next() {
		product := &model.ProductModel{}
		err := rows.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Stok, &product.CategoryProductId, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("GetAllProduct() : %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (p *productRepo) GetProductByName(nameProduct string) (*model.ProductModel, error) {
	selectStatement := "SELECT id, product_name, description, price, stok, category_product_id, status, created_at, updated_at FROM mst_product WHERE product_name = $1"

	row := p.db.QueryRow(selectStatement, nameProduct)

	product := &model.ProductModel{}
	err := row.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Stok, &product.CategoryProductId, &product.Status, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetProductByName() : %w", err)
	}

	return product, nil
}



func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}
