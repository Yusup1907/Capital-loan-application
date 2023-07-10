package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetCategoryLoanRepo() repository.CategoryLoanRepo
}

type repoManager struct {
	infraManager     InfraManager
	categoryLoanRepo repository.CategoryLoanRepo
	onceLoadRepo     sync.Once
}

func (rm *repoManager) GetCategoryLoanRepo() repository.CategoryLoanRepo {
	rm.onceLoadRepo.Do(func() {
		rm.categoryLoanRepo = repository.NewCategoryLoanRepo(rm.infraManager.GetDB())
	})
	return rm.categoryLoanRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
