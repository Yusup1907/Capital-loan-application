package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
)

type LoanApplicationRepo interface {
	CreateLoanApplication(application *model.LoanApplicationModel) error
	GetCustomerById(int) (*model.ValidasiCustomerModel, error)
}

type loanApplicationRepo struct {
	db *sql.DB
}

func (r *loanApplicationRepo) CreateLoanApplication(application *model.LoanApplicationModel) error {
	insertStatement := `
		INSERT INTO trx_loan (customer_id, loan_date, due_date, category_loan_id, amount, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
	`

	err := r.db.QueryRow(insertStatement, application.CustomerId, application.LoanDate, application.DueDate, application.CategoryLoanId, application.Amount, application.Description, application.Status, application.CreatedAt, application.UpdatedAt).Scan(&application.Id)
	if err != nil {
		return fmt.Errorf("error on loanApplicationRepo.CreateLoanApplication: %w", err)
	}

	return nil
}

func (r *loanApplicationRepo) GetCustomerById(id int) (*model.ValidasiCustomerModel, error) {
	qry := "SELECT id, nik, nokk, emergencyname, emergencycontact, last_salary FROM mst_customer WHERE id = $1"

	customer := &model.ValidasiCustomerModel{}
	err := r.db.QueryRow(qry, id).Scan(
		&customer.Id, &customer.NIK, &customer.NoKK, &customer.EmergencyName, &customer.EmergencyContact, &customer.LastSalary)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, fmt.Errorf("error on GetCustomerById: %w", err)
	}

	return customer, nil
}

func NewLoanApplicationRepository(db *sql.DB) LoanApplicationRepo {
	return &loanApplicationRepo{
		db: db,
	}
}
