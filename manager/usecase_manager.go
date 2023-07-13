package manager

import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetCategoryProductUsecase() usecase.CategoryProductUsecase
	GetGoodsUsecase() usecase.GoodsUsecase
}

type usecaseManager struct {
	repoManager RepoManager
	cpUsecase   usecase.CategoryProductUsecase
	gUsecase usecase.GoodsUsecase
}

var onceLoadGetCategoryProductUsecase sync.Once
var onceLoadGetGoodsUsecase sync.Once

func (um *usecaseManager) GetCategoryProductUsecase() usecase.CategoryProductUsecase {
	onceLoadGetCategoryProductUsecase.Do(func() {
		um.cpUsecase = usecase.NewCategoryProductUsecase(um.repoManager.GetCategoryProductRepo())
	})
	return um.cpUsecase
}

func (um *usecaseManager) GetGoodsUsecase() usecase.GoodsUsecase{
	onceLoadGetGoodsUsecase.Do(func() {
		um.gUsecase = usecase.NewGoodsUsecase(um.repoManager.GetGoodsRepo())
	})
	return um.gUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}