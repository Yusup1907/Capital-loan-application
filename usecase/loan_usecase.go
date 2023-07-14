package usecase

import (
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
	"time"
)

type LoanApplicationUsecase interface {
	CreateLoanApplication(application *model.LoanApplicationModel) error
}

type loanApplicationUsecase struct {
	repo repository.LoanApplicationRepo
}

func (uc *loanApplicationUsecase) CreateLoanApplication(application *model.LoanApplicationModel) error {

	customerDB, err := uc.repo.GetCustomerById(application.CustomerId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	if customerDB.NIK.Valid && customerDB.NIK.String != "" &&
		customerDB.NoKK.Valid && customerDB.NoKK.String != "" &&
		customerDB.EmergencyName.Valid && customerDB.EmergencyName.String != "" &&
		customerDB.EmergencyContact.Valid && customerDB.EmergencyContact.String != "" &&
		customerDB.LastSalary.Valid && customerDB.LastSalary.Float64 != 0 {
		application.Status = model.LoanStatusApprove
		application.DueDate = time.Now().AddDate(0, 2, 0)
	} else {
		application.Status = model.LoanStatusDenied
		fmt.Println("Silakan lengkapi data customer")
	}

	err = uc.repo.CreateLoanApplication(application)
	if err != nil {
		return fmt.Errorf("failed to insert loan: %v", err)
	}

	return nil
}

func NewLoanApplicationUseCase(repo repository.LoanApplicationRepo) LoanApplicationUsecase {
	return &loanApplicationUsecase{
		repo: repo,
	}
}
