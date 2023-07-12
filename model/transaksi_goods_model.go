package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type GoodsTransactionModel struct {
	Id             int       `json:"id"`
	CustomerId     int       `json:"customer_id"`
	LoanDate       time.Time `json:"loan_date"`
	PaymentDate    time.Time `json:"payment_date"`
	DueDate        time.Time `json:"due_date" validate:"required"`
	CategoryLoanID int       `json:"category_loan_id"`
	ProductID      int       `json:"product_id"`
	Quantity       int       `json:"quantity" validate:"required"`
	Amount         int64     `json:"amount"`
	Price          int64     `json:"price" validate:"required,gte=0"`
	Status         bool      `json:"status" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateGoodsTransactionRequest struct {
	Id             int       `json:"id"`
	CustomerId     int       `json:"customer_id"`
	LoanDate       time.Time `json:"loan_date"`
	PaymentDate    time.Time `json:"payment_date"`
	DueDate        time.Time `json:"due_date"`
	CategoryLoanID int       `json:"category_loan_id"`
	ProductID      int       `json:"product_id"`
	Quantity       int       `json:"quantity"`
	Amount         int64     `json:"amount"`
	Price          int64     `json:"price"`
	Status         bool      `json:"status"`
}

func (p *GoodsTransactionModel) ValidateUpdate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		// Mengembalikan error dengan pesan validasi yang lebih spesifik
		errs := err.(validator.ValidationErrors)
		errMsg := ""
		for _, e := range errs {
			errMsg += fmt.Sprintf("Field %s: validation failed on tag '%s'\n", e.Field(), e.Tag())
		}
		return errors.New(errMsg)
	}

	return nil
}
