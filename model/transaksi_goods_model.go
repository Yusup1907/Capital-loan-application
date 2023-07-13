package model

import "time"

type LoanApplicationModel struct {
	Id             int        `json:"id"`
	CustomerId     int        `json:"customer_id"`
	LoanDate       time.Time  `json:"loan_date"`
	PaymentDate    time.Time  `json:"payment_date"`
	DueDate        time.Time  `json:"due_date"`
	CategoryLoanId int        `json:"category_loan_id"`
	Amount         int64      `json:"amount"`
	Description    string     `json:"description"`
	Status         LoanStatus `json:"status" validate:"required"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type LoanJoinRequest struct {
	Id               int        `json:"id"`
	CustomerId       int        `json:"customer_id"`
	FullName         string     `json:"full_name"`
	LoanDate         time.Time  `json:"loan_date"`
	PaymentDate      time.Time  `json:"payment_date"`
	DueDate          time.Time  `json:"due_date"`
	CategoryLoanId   int        `json:"category_loan_id"`
	CategoryLoanName string     `json:"category_loan_name"`
	Amount           int64      `json:"amount"`
	Description      string     `json:"description"`
	Status           LoanStatus `json:"status" validate:"required"`
}

type LoanStatus string

const (
	LoanStatusApprove LoanStatus = "approve"
	LoanStatusPending LoanStatus = "pending"
	LoanStatusDenied  LoanStatus = "denied"
)
