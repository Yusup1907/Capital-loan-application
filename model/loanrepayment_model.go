package model

import "time"

type LoanRepaymentModel struct {
	PaymentDate     time.Time        `json:"payment_date"`
	Payment         float64          `json:"payment"`
	RepaymentStatus RepaymentsStatus `json:"repayment_status"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type RepaymentsStatus string

const (
	RepaymentsStatusLunas      RepaymentsStatus = "lunas"
	RepaymentsStatusBelumLunas RepaymentsStatus = "belum lunas"
)
