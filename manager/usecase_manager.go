package manager

import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetProductUsecase() usecase.ProductUsecase
	GetCategoryLoanUsecase() usecase.CategoryLoanUsecase
}

type usecaseManager struct {
	repoManager         RepoManager
	productUsecase      usecase.ProductUsecase
	loanApp             usecase.LoanApplicationUsecase
	categoryLoanUsecase usecase.CategoryLoanUsecase

	onceLoadUsecase        sync.Once
	onceLoadProductUsecase sync.Once
	onceLoadLoanAppUsecase sync.Once
}

func (um *usecaseManager) GetCategoryLoanUsecase() usecase.CategoryLoanUsecase {
	um.onceLoadUsecase.Do(func() {
		um.categoryLoanUsecase = usecase.NewCategoryLoanUsecase(um.repoManager.GetCategoryLoanRepo())
	})
	return um.categoryLoanUsecase
}
func (um *usecaseManager) GetProductUsecase() usecase.ProductUsecase {
	um.onceLoadProductUsecase.Do(func() {
		um.productUsecase = usecase.NewProductUseCase(um.repoManager.GetProductRepo())
	})
	return um.productUsecase
}

func (um *usecaseManager) GetLoanAppUsecase() usecase.LoanApplicationUsecase {
	um.onceLoadLoanAppUsecase.Do(func() {
		um.loanApp = usecase.NewLoanApplicationUseCase(um.repoManager.GetLoanApplicationRepo())
	})
	return um.loanApp
}
func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
