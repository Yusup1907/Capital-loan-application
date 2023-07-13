package repository

import "database/sql"

type LoanApplicationRepo interface {
}

type loanApplicationRepo struct {
	db *sql.DB
}

func NewLoanApplicationRepository(db *sql.DB) LoanApplicationRepo {
	return &loanApplicationRepo{
		db: db,
	}
}
