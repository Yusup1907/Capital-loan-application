package usecase

import (
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type LoanApplicationUsecase interface {
}

type loanApplicationUsecase struct {
	repo repository.LoanApplicationRepo
}

func (uc *loanApplicationUsecase) CreateLoanApplication(application *model.LoanApplicationModel) error {
	if err := application.Validate(); err != nil {
		return err
	}

	percentage := calculateDataPercentage(application)
	if percentage == 100 {
		application.Status = "approve"
	} else if percentage >= 60 {
		application.Status = "pending"
	} else {
		application.Status = "denied"
	}

	if err := uc.loanRepo.CreateLoanApplication(application); err != nil {
		return err
	}

	return nil
}

func calculateDataPercentage(customer *model.CustomerModel) float64 {
	requiredFields := []interface{}{
		customer.NIK,
		customer.NoKK,
		customer.EmergencyName,
		customer.EmergencyContact,
		customer.LastSalary,
		Address.Amount,
	}

	totalFields := len(requiredFields)
	filledFields := 0

	for _, field := range requiredFields {
		if field != nil {
			filledFields++
		}
	}

	return float64(filledFields) / float64(totalFields) * 100
}

func NewLoanApplicationUseCase(repo repository.LoanApplicationRepo) LoanApplicationUsecase {
	return &loanApplicationUsecase{
		repo: repo,
	}
}
