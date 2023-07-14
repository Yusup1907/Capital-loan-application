package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
)

type LoanApplicationRepo interface {
	CreateLoanApplication(application *model.LoanApplicationModel) error
	GetCustomerById(int) (*model.ValidasiCustomerModel, error)
	GetLoanApplications(page, limit int) ([]*model.LoanApplicationJoinModel, error)
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

func (r *loanApplicationRepo) GetLoanApplications(page, limit int) ([]*model.LoanApplicationJoinModel, error) {
	offset := (page - 1) * limit

	selectStatement := `
		SELECT la.id, la.customer_id, la.loan_date, la.due_date, la.category_loan_id, la.amount, la.description, la.status, la.created_at, la.updated_at,
			   mc.full_name, mc.address, mc.nik, mc.phone_number, mc.nokk, mc.emergencyname, mc.emergencycontact, mc.last_salary
		FROM trx_loan la
		INNER JOIN mst_customer mc ON la.customer_id = mc.id
		ORDER BY la.id
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(selectStatement, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}
	defer rows.Close()

	applications := []*model.LoanApplicationJoinModel{}
	for rows.Next() {
		application := &model.LoanApplicationJoinModel{}
		err := rows.Scan(
			&application.Id, &application.CustomerId, &application.LoanDate, &application.DueDate, &application.CategoryLoanID,
			&application.Amount, &application.Description, &application.Status, &application.CreatedAt, &application.UpdatedAt,
			&application.FullName, &application.Address, &application.NIK, &application.PhoneNumber, &application.NoKK, &application.EmergencyName,
			&application.EmergencyContact, &application.LastSalary,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan loan application: %w", err)
		}
		applications = append(applications, application)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}

	return applications, nil
}

func (r *loanApplicationRepo) GetLoanApplicationById(id int) (*model.LoanApplicationModel, error) {
	selectStatement := `
		SELECT id, customer_id, loan_date, due_date, category_loan_id, amount, description, status, created_at, updated_at
		FROM trx_loan 
		WHERE id = $1
		ORDER BY 
		id ASC
	`

	loan := &model.LoanApplicationModel{}
	err := r.db.QueryRow(selectStatement, id).Scan(
		&loan.Id, &loan.CustomerId, &loan.LoanDate, &loan.DueDate, &loan.CategoryLoanId,
		&loan.Amount, &loan.Description, &loan.Status, &loan.CreatedAt, &loan.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("loan application not found")
		}
		return nil, fmt.Errorf("failed to get loan application: %w", err)
	}

	return loan, nil
}

func NewLoanApplicationRepository(db *sql.DB) LoanApplicationRepo {
	return &loanApplicationRepo{
		db: db,
	}
}
