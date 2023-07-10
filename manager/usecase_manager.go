package manager

import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetCategoryLoanUsecase() usecase.CategoryLoanUsecase
}

type usecaseManager struct {
	repoManager         RepoManager
	categoryLoanUsecase usecase.CategoryLoanUsecase
	onceLoadUsecase     sync.Once
}

func (um *usecaseManager) GetCategoryLoanUsecase() usecase.CategoryLoanUsecase {
	um.onceLoadUsecase.Do(func() {
		um.categoryLoanUsecase = usecase.NewCategoryLoanUsecase(um.repoManager.GetCategoryLoanRepo())
	})
	return um.categoryLoanUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
