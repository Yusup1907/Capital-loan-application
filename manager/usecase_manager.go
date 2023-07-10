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

	onceLoadProductUsecase sync.Once
}

func (um *usecaseManager) GetProductUsecase() usecase.ProductUsecase {
	um.onceLoadProductUsecase.Do(func() {
		um.productUsecase = usecase.NewProductUseCase(um.repoManager.GetProductRepo())
	})
	return um.productUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
