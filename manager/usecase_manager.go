package manager

import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetProductUsecase() usecase.ProductUsecase
}

type usecaseManager struct {
	repoManager    RepoManager
	productUsecase usecase.ProductUsecase
	loanApp        usecase.LoanApplicationUsecase

	onceLoadProductUsecase sync.Once
	onceLoadLoanAppUsecase sync.Once
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
