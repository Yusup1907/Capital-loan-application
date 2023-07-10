package usecase

import (
	"fmt"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type CategoryLoanUsecase interface {
	// GetCategoryLoanById(int) (*model.CategoryLoanModel, error)
	// GetCategoryLoanByName(string) (*model.CategoryLoanModel, error)
	// GetAllCategoryLoan() ([]*model.CategoryLoanModel, error)
	InsertCategoryLoan(*model.CategoryLoanModel) error
	// UpdateCategoryLoan(*model.CategoryLoanModel) error
	// DeleteCategoryLoan(*model.CategoryLoanModel) error
}

type CategoryLoanUsecaseImpl struct {
	ctrRepo repository.CategoryLoanRepo
}

func NewCategoryLoanUsecase(ctrRepo repository.CategoryLoanRepo) CategoryLoanUsecase {
	return &CategoryLoanUsecaseImpl{
		ctrRepo: ctrRepo,
	}
}

// func (uc *CategoryLoanUsecaseImpl) GetCategoryLoanById(id int) (*model.CategoryLoanModel, error) {
// 	return uc.ctrRepo.GetCategoryLoanById(id)
// }

// func (uc *CategoryLoanUsecaseImpl) GetCategoryLoanByName(name string) (*model.CategoryLoanModel, error) {
// 	return uc.ctrRepo.GetCategoryLoanByName(name)
// }

// func (uc *CategoryLoanUsecaseImpl) GetAllCategoryLoan() ([]*model.CategoryLoanModel, error) {
// 	return uc.ctrRepo.GetAllCategoryLoan()
// }

func (uc *CategoryLoanUsecaseImpl) InsertCategoryLoan(ctr *model.CategoryLoanModel) error {
	categoryLoanByName, err := uc.ctrRepo.GetCategoryLoanByName(ctr.CategoryLoanName)
	if err != nil {
		return fmt.Errorf("CategoryLoanUsecaseImpl.InsertCategoryLoan(): %w", err)
	}

	if categoryLoanByName != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("Data CategoryLoan with name %v already exists", ctr.CategoryLoanName),
		}
	}

	return uc.ctrRepo.InsertCategoryLoan(ctr)
}

// func (uc *CategoryLoanUsecaseImpl) UpdateCategoryLoan(ctr *model.CategoryLoanModel) error {
// 	return uc.ctrRepo.UpdateCategoryLoan(ctr)
// }

// func (uc *CategoryLoanUsecaseImpl) DeleteCategoryLoan(ctr *model.CategoryLoanModel) error {
// 	return uc.ctrRepo.DeleteCategoryLoan(ctr)
// }
