package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetCategoryProductRepo() repository.CategoryProductRepo
}

type repoManager struct {
	infraManager InfraManager

	cpRepo repository.CategoryProductRepo
}

var onceLoadRepoManager sync.Once

func (rm *repoManager) GetCategoryProductRepo() repository.CategoryProductRepo{
	onceLoadRepoManager.Do( func() {
		rm.cpRepo = repository.NewCategoryProductRepo(rm.infraManager.GetDB())
	})
	return rm.cpRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}