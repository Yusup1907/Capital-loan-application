package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetProductRepo() repository.ProductRepo
	GetLoanApplicationRepo() repository.LoanApplicationRepo
}

type repoManager struct {
	infraManager InfraManager
	productRepo  repository.ProductRepo
	loan         repository.LoanApplicationRepo

	onceLoadProductRepo sync.Once
	onceLoadLoanAppRepo sync.Once
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
