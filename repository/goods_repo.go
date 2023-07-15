package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"time"
)

type GoodsRepo interface {
	InsertGoods(*model.GoodsModel) error
	GetGoodsById(int) (*model.LoanGoodsModel, error)
	GetAllTrxGoods(page, limit int) ([]*model.LoanGoodsModel, error)
	GetCustomerById(int) (*model.ValidasiCustomerModel, error)
	GoodsRepayment(int, *model.LoanRepaymentModel) error
}

type goodsRepoImpl struct {
	db *sql.DB
}


func (goodsRepo *goodsRepoImpl) InsertGoods(goods *model.GoodsModel) error {
	insertQuery := "INSERT INTO trx_goods (customer_id, loan_date, payment_date, due_date, category_loan_id, product_id, quantity, price, amount, created_at, status, repayment_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"

	goods.CreateAt = time.Now()
	goods.LoadDate = time.Now()

	product := &model.ProductModel{}
	err := goodsRepo.db.QueryRow("SELECT price FROM mst_product WHERE id = $1", goods.ProductId).Scan(&product.Price)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan harga produk: %w", err)
	}

	goods.Amount = float64(goods.Quantity) * product.Price


	_, err = goodsRepo.db.Exec(insertQuery, goods.CustomerId, goods.LoadDate, goods.PaymentDate, goods.DueDate, goods.CategoryIdLoan, goods.ProductId, goods.Quantity, product.Price, goods.Amount, goods.CreateAt, goods.Status, goods.RepaymentStatus)
	if err != nil {
		return fmt.Errorf("gagal memasukkan data goods: %w", err)
	}

	return nil
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

func (goodsRepo *goodsRepoImpl) GetGoodsById(id int) (*model.LoanGoodsModel, error) {
	qry := `SELECT g.id, g.customer_id, g.loan_date, g.due_date, g.category_loan_id, g.product_id, p.product_name, g.quantity , p.price , g.amount, p.deskripsi, g.status, g.repayment_status, g.created_at, g.updated_at, c.full_name, c.address, c.phone_number, c.nik, c.nokk, c.emergencyname, c.emergencyphone, c.last_salary FROM trx_goods g JOIN mst_customer c ON g.customer_id = c.id JOIN mst_product p ON g.product_id = p.id WHERE g.id = $1`
	
	loangoods := &model.LoanGoodsModel{}
	err := goodsRepo.db.QueryRow(qry, id).Scan(
					&loangoods.Id, &loangoods.CustomerId,&loangoods.LoanDate,
					&loangoods.DueDate, &loangoods.CategoryLoanID, &loangoods.ProductId, 
					&loangoods.ProductName,&loangoods.Quantity,&loangoods.Price, 
					&loangoods.Amount, &loangoods.Description,&loangoods.Status, 
					&loangoods.RepaymentStatus, &loangoods.CreatedAt,&loangoods.UpdatedAt,
					&loangoods.FullName, &loangoods.Address,&loangoods.PhoneNumber,
					&loangoods.NIK, &loangoods.NoKK,&loangoods.EmergencyName,
					&loangoods.EmergencyContact, &loangoods.LastSalary,
				)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("goods not found")
		}
		return nil, fmt.Errorf("error on goodsRepoImpl.GetGoodsById: %w", err)
	}

	return loangoods, nil
}

func (goodsRepo *goodsRepoImpl) GetAllTrxGoods(page, limit int) ([]*model.LoanGoodsModel, error) {
		offset := (page - 1) * limit
		query := 
		`SELECT g.id, g.customer_id, g.loan_date, g.due_date, g.category_loan_id, g.product_id, p.product_name, g.quantity , p.price , g.amount, p.deskripsi, g.status, g.repyement_status, g.created_at, g.updated_at, c.full_name, c.address, c.phone_number, c.nik, c.nokk, c.emergencyname, c.emergencyphone, c.last_salary FROM trx_goods g JOIN mst_customer c ON g.customer_id = c.id JOIN mst_product p ON g.product_id = p.id ORDER BY g.id ASC OFFSET $1 LIMIT $2`
		rows, err := goodsRepo.db.Query(query, offset, limit)
		if err != nil {
			return nil, fmt.Errorf("failed to get loan applications: %w", err)
		}
		defer rows.Close()
					
		loangoods := []*model.LoanGoodsModel{}
		for rows.Next(){
			loangood := &model.LoanGoodsModel{}
			err := rows.Scan( &loangood.Id, &loangood.CustomerId,&loangood.LoanDate, &loangood.DueDate, &loangood.CategoryLoanID, &loangood.ProductId, &loangood.ProductName,&loangood.Quantity,&loangood.Price, &loangood.Amount, &loangood.Description,&loangood.Status, &loangood.RepaymentStatus, &loangood.CreatedAt,&loangood.UpdatedAt, &loangood.FullName, &loangood.Address,&loangood.PhoneNumber, &loangood.NIK, &loangood.NoKK,&loangood.EmergencyName, &loangood.EmergencyContact, &loangood.LastSalary)
			if err != nil {
				return nil, fmt.Errorf("failed to scan loan application: %w", err)
			}
			loangoods = append(loangoods, loangood)
			}
			if err := rows.Err(); err != nil {
				return nil, fmt.Errorf("failed to get loan applications: %w", err)
			}
			return loangoods, nil
}

func (goodsRepo *goodsRepoImpl) GoodsRepayment(id int, repayment *model.LoanRepaymentModel) error {
	updateStatment := "UPDATE trx_goods SET payment_date = $1, payment = $2, repayment_status = $3, updated_at = $4 WHERE id = $5"
	_, err := goodsRepo.db.Exec(updateStatment, repayment.PaymentDate, repayment.Payment, model.StatusEnum(repayment.RepaymentStatus), repayment.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("error on loanApplicationRepo.LoanRepayment() : %w", err)
	}
	return nil
}


func NewGoodsRepo(db *sql.DB) GoodsRepo {
	return &goodsRepoImpl{
		db: db,
	}
}

