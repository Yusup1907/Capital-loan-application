package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"time"
)

type GoodsRepo interface {
	InsertGoods(*model.GoodsModel) error
	GetGoodsById(int) (*model.GoodsModel, error)
	GetCustomerById(int) (*model.ValidasiCustomerModel, error)
}

type goodsRepoImpl struct {
	db *sql.DB
}

func (goodsRepo *goodsRepoImpl) InsertGoods(goods *model.GoodsModel) error {
	tx, err := goodsRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("trxGoodsRepo.InsertTrxGoods: %w", err)
	} 

	qry := "INSERT INTO trx_goods(customer_id, loan_date, payment_date, due_date, category_loan_id, product_id, quantity, price, amount, created_at, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)RETURNING id"

	goods.CreateAt = time.Now()
	goods.LoadDate = time.Now()
	
	product := &model.ProductModel{}
	goods.Amount = float64(goods.Quantity) * product.Price

	err = tx.QueryRow(qry, goods.CustomerId, goods.LoadDate, goods.PaymentDate, goods.DueDate, goods.CategoryIdLoan, goods.ProductId, goods.Quantity, product.Price, goods.Amount, goods.CreateAt, goods.Status).Scan(&goods.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on trxGoodsRepo.InsertTrxGoods: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on GoodsRepo.InsertGoods (commit): %w", err)
	}
	return nil
}

func (goodsRepo *goodsRepoImpl) GetGoodsById(id int)(*model.GoodsModel, error) {
	qry := "SELECT tg.id, mc.full_name, mc.phone_number, mp.product_name, tg.quantity, mp.price, tg.amount, tg.loan_date, tg.payment_date, tg.due_date FROM trx_goods AS tg JOIN mst_customer AS mc ON tg.customer_id = mc.id JOIN category_loan AS cl ON tg.category_loan_id = cl.id JOIN mst_product AS mp ON tg.product_id = mp.id WHERE tg.id = $1"


	goods := &model.GoodsModel{}
	customer := &model.CustomerModel{}
	product := &model.ProductModel{}
	err := goodsRepo.db.QueryRow(qry, id).Scan(
							&goods.Id, 
							&customer.FullName,
							&customer.Phone,
							&product.ProductName,
							&goods.Quantity,
							&product.Price,
							&goods.Amount,
							&goods.LoadDate,
							&goods.PaymentDate,
							&goods.DueDate,
							&goods.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaksi not found")
		}
		return nil, fmt.Errorf("error on categoryProductRepoImpl.GetCategoryProductById: %w", err)
	}
	return goods, nil
}


func (goodsRepo *goodsRepoImpl) GetCustomerById(id int) (*model.ValidasiCustomerModel, error) {
	qry := "SELECT id, nik, nokk, emergencyname, emergencyphone, last_salary FROM mst_customer WHERE id = $1"

	customer := &model.ValidasiCustomerModel{}
	err := goodsRepo.db.QueryRow(qry, id).Scan(
		&customer.Id, &customer.NIK, &customer.NoKK, &customer.EmergencyName, &customer.EmergencyContact, &customer.LastSalary)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, fmt.Errorf("error on GetCustomerById: %w", err)
	}

	return customer, nil
}

func NewGoodsRepo(db *sql.DB) GoodsRepo {
	return &goodsRepoImpl{
		db: db,
	}
}

