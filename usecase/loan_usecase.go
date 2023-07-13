package usecase

import (
	"pinjam-modal-app/repository"
)

type LoanApplicationUsecase interface {
}

type loanApplicationUsecase struct {
	repo repository.LoanApplicationRepo
}

func NewLoanApplicationUseCase(repo repository.LoanApplicationRepo) LoanApplicationUsecase {
	return &loanApplicationUsecase{
		repo: repo,
	}
}
