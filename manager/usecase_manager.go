package manager

import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetCategoryProductUsecase() usecase.CategoryProductUsecase
}

type usecaseManager struct {
	repoManager RepoManager
	cpUsecase   usecase.CategoryProductUsecase
}

var onceLoadGetCategoryProductUsecase sync.Once

func (um *usecaseManager) GetCategoryProductUsecase() usecase.CategoryProductUsecase {
	onceLoadGetCategoryProductUsecase.Do(func() {
		um.cpUsecase = usecase.NewCategoryProductUsecase(um.repoManager.GetCategoryProductRepo())
	})
	return um.cpUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}