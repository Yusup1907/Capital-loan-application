package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetProductRepo() repository.ProductRepo
	GetLoanApplicationRepo() repository.LoanApplicationRepo
	GetCategoryLoanRepo() repository.CategoryLoanRepo
}

type repoManager struct {
	infraManager     InfraManager
	productRepo      repository.ProductRepo
	loan             repository.LoanApplicationRepo
	categoryLoanRepo repository.CategoryLoanRepo

	onceLoadProductRepo sync.Once
	onceLoadLoanAppRepo sync.Once
	onceLoadRepo        sync.Once
}

func (rm *repoManager) GetCategoryLoanRepo() repository.CategoryLoanRepo {
	rm.onceLoadRepo.Do(func() {
		rm.categoryLoanRepo = repository.NewCategoryLoanRepo(rm.infraManager.GetDB())
	})
	return rm.categoryLoanRepo
}
func (rm *repoManager) GetProductRepo() repository.ProductRepo {
	rm.onceLoadProductRepo.Do(func() {
		rm.productRepo = repository.NewProductRepo(rm.infraManager.GetDB())
	})
	return rm.productRepo
}

func (rm *repoManager) GetLoanApplicationRepo() repository.LoanApplicationRepo {
	rm.onceLoadLoanAppRepo.Do(func() {
		rm.loan = repository.NewLoanApplicationRepository(rm.infraManager.GetDB())
	})
	return rm.loan
}
func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
