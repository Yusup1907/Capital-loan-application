package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"time"
)

type CategoryProductRepo interface {
	InsertCategoryProduct(*model.CategoryProductModel) error
	GetCategoryProductById(int) (*model.CategoryProductModel, error)
}

type categoryProductRepoImpl struct {
	db *sql.DB
}

func (cpRepo *categoryProductRepoImpl) InsertCategoryProduct(cp *model.CategoryProductModel) error {
	qry := "INSERT INTO category_product(category_product_name, created_at, updated_at) VALUES ($1 ,$2, DEFAULT) RETURNING id"

	cp.CreateAt = time.Now()
	err := cpRepo.db.QueryRow(qry, cp.CategoryProductName, cp.CreateAt).Scan(&cp.Id)
	if err != nil {
		return fmt.Errorf("error on categoryProductRepoImpl.InsertCategoryProduct: %w", err)
	}
	return nil
}

func (cpRepo *categoryProductRepoImpl) GetCategoryProductById(id int) (*model.CategoryProductModel, error) {
	qry := "SELECT id, category_product_name, created_at, updated_at FROM category_product WHERE id = $1"

	cp := &model.CategoryProductModel{}
	err := cpRepo.db.QueryRow(qry, id).Scan(&cp.Id, &cp.CategoryProductName, &cp.CreateAt, &cp.UpdateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category product not found")
		}
		return nil, fmt.Errorf("error on categoryProductRepoImpl.GetCategoryProductById: %w", err)
	}
	return cp, nil
}


func NewCategoryProductRepo(db *sql.DB) CategoryProductRepo {
	return &categoryProductRepoImpl{
		db: db,
	}
}