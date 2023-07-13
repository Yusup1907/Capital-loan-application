package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetCategoryProductRepo() repository.CategoryProductRepo
	GetGoodsRepo() repository.GoodsRepo
}

type repoManager struct {
	infraManager InfraManager

	cpRepo repository.CategoryProductRepo
	goodsRepo repository.GoodsRepo
}

var onceLoadCategoryProductRepo sync.Once
var onceLoadGoodsRepo sync.Once

func (rm *repoManager) GetCategoryProductRepo() repository.CategoryProductRepo{
	onceLoadCategoryProductRepo.Do( func() {
		rm.cpRepo = repository.NewCategoryProductRepo(rm.infraManager.GetDB())
	})
	return rm.cpRepo
}

func (rm *repoManager) GetGoodsRepo() repository.GoodsRepo{
	onceLoadGoodsRepo.Do(func() {
		rm.goodsRepo = repository.NewGoodsRepo(rm.infraManager.GetDB())
	})
	return rm.goodsRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}