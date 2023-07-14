package model

import "time"

type LoanApplicationModel struct {
	Id             int        `json:"id"`
	CustomerId     int        `json:"customer_id"`
	LoanDate       time.Time  `json:"loan_date"`
	DueDate        time.Time  `json:"due_date"`
	CategoryLoanId int        `json:"category_loan_id"`
	Amount         int64      `json:"amount"`
	Description    string     `json:"description"`
	Status         LoanStatus `json:"status"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type LoanJoinRequest struct {
	Id             int        `json:"id"`
	CustomerId     int        `json:"customer_id"`
	LoanDate       time.Time  `json:"loan_date"`
	DueDate        time.Time  `json:"due_date"`
	CategoryLoanId int        `json:"category_loan_id"`
	Amount         int64      `json:"amount"`
	Description    string     `json:"description"`
	Status         LoanStatus `json:"status"`
}

type LoanStatus string

const (
	LoanStatusApprove LoanStatus = "Approve"
	LoanStatusPending LoanStatus = "Pending"
	LoanStatusDenied  LoanStatus = "Denied"
)

type LoanRepaymentModel struct {
	PaymentDate time.Time `json:"payment_date"`
	Payment     float64
}
