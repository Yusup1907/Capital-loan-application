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
	GetCategoryProductByName(string) (*model.CategoryProductModel, error)
	GetAllCategoryProduct() ([]model.CategoryProductModel, error)
	UpdateCategoryProduct(int, *model.CategoryProductModel) error
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

func (cpRepo *categoryProductRepoImpl) GetCategoryProductByName(name string) (*model.CategoryProductModel, error) {
	qry := "SELECT id, category_product_name, created_at, updated_at FROM category_product WHERE category_product_name = $1"

	cp := &model.CategoryProductModel{}
	err := cpRepo.db.QueryRow(qry, name).Scan(&cp.Id, &cp.CategoryProductName, &cp.CreateAt, &cp.UpdateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on categoryProductRepoImpl.categoryProductRepoImpl : %w", err)
	}
	return cp, nil
}

func (cpRepo *categoryProductRepoImpl) GetAllCategoryProduct() ([]model.CategoryProductModel, error){
	qry := "SELECT  id, category_product_name, created_at, updated_at FROM category_product ORDER BY id"
	rows, err := cpRepo.db.Query(qry)
	if  err != nil{
		return nil, fmt.Errorf("error oncategoryProductRepoImpl.GetAllCategoryProduct : %w", err)
	}
	defer rows.Close()
	var arrCp []model.CategoryProductModel
	for rows.Next(){
		cp := &model.CategoryProductModel{}
		rows.Scan(&cp.Id, &cp.CategoryProductName, &cp.CreateAt, cp.UpdateAt)
		arrCp = append(arrCp, *cp)
	}
	return arrCp, nil
}

func (cpRepo *categoryProductRepoImpl) UpdateCategoryProduct(id int , cp *model.CategoryProductModel) error{
	qryId := "SELECT id FROM category_product WHERE id = $1"
	err := cpRepo.db.QueryRow(qryId, cp.Id).Scan(&cp.Id)
	if err != nil{
		return fmt.Errorf("data category product not found")
	}
	cp.UpdateAt = time.Now()
	 qry := "UPDATE category_product SET category_product_name = $1, updated_at = $2 WHERE id = $3"
	 _, err = cpRepo.db.Exec(qry, &cp.CategoryProductName, &cp.UpdateAt, &cp.Id)
	 if err != nil {
		return fmt.Errorf("err on categoryProductRepoImpl.UpdateCategoryProduct : %w ", err)
	}
	return nil
}

func NewCategoryProductRepo(db *sql.DB) CategoryProductRepo {
	return &categoryProductRepoImpl{
		db: db,
	}
}