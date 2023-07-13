package model

import "time"

type GoodsModel struct {
	Id				int			`json:"id"`
	CustomerId      int			`json:"customer_id"`
	CategoryIdLoan  int			`json:"category_loan_id"`
	ProductId       int			`json:"product_id"`
	LoadDate		time.Time	`json:"loan_date"`
	PaymentDate		time.Time	`json:"payment_date"`
	DueDate			time.Time	`json:"due_date"`
	Quantity		int			`json:"quantity"`
	Payment			float64		`json:"price"`
	Amount			float64		`json:"amount"`
	Status 			 string		`json:"status"`
	CreateAt		time.Time	`json:"created_at"`
	UpdateAt		time.Time	`json:"updated_at"`
}