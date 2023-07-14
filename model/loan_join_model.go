package model

import "time"

type LoanApplicationJoinModel struct {
	Id               int       `json:"id"`
	CustomerId       int       `json:"customer_id"`
	LoanDate         time.Time `json:"loan_date"`
	DueDate          time.Time `json:"due_date"`
	CategoryLoanID   int       `json:"category_loan_id"`
	Amount           int       `json:"amount"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	FullName         string    `json:"full_name"`
	Address          string    `json:"address"`
	NIK              string    `json:"nik"`
	PhoneNumber      string    `json:"phone_number"`
	EmergencyName    string    `json:"emergencyname"`
	EmergencyContact string    `json:"emergencycontact"`
	NoKK             string    `json:"nokk"`
	LastSalary       float64   `json:"last_salary"`
}
