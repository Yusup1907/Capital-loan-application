package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetProductRepo() repository.ProductRepo
}

type repoManager struct {
	infraManager InfraManager
	productRepo  repository.ProductRepo

	onceLoadProductRepo sync.Once
}

func (rm *repoManager) GetProductRepo() repository.ProductRepo {
	rm.onceLoadProductRepo.Do(func() {
		rm.productRepo = repository.NewProductRepo(rm.infraManager.GetDB())
	})
	return rm.productRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
