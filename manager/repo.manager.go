package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetCustomerRepo() repository.CustomerRepo
	GetProductRepo() repository.ProductRepo
	GetLoanApplicationRepo() repository.LoanApplicationRepo
	GetCategoryLoanRepo() repository.CategoryLoanRepo
	GetCategoryProductRepo() repository.CategoryProductRepo
	GetGoodsRepo() repository.GoodsRepo
}

type repoManager struct {
	infraManager     InfraManager
	cstRepo          repository.CustomerRepo
	productRepo      repository.ProductRepo
	loan             repository.LoanApplicationRepo
	categoryLoanRepo repository.CategoryLoanRepo
	CategoryProduct  repository.CategoryProductRepo
	TrxGoods         repository.GoodsRepo

	onceLoadCustomerRepo        sync.Once
	onceLoadCategoryProductRepo sync.Once
	onceLoadGoodsRepo           sync.Once
	onceLoadProductRepo         sync.Once
	onceLoadLoanAppRepo         sync.Once
	onceLoadRepo                sync.Once
}

func (rm *repoManager) GetCustomerRepo() repository.CustomerRepo {
	rm.onceLoadCustomerRepo.Do(func() {
		rm.cstRepo = repository.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.cstRepo
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
func (rm *repoManager) GetCategoryProductRepo() repository.CategoryProductRepo {
	rm.onceLoadCategoryProductRepo.Do(func() {
		rm.CategoryProduct = repository.NewCategoryProductRepo(rm.infraManager.GetDB())
	})
	return rm.CategoryProduct
}

func (rm *repoManager) GetGoodsRepo() repository.GoodsRepo {
	rm.onceLoadGoodsRepo.Do(func() {
		rm.TrxGoods = repository.NewGoodsRepo(rm.infraManager.GetDB())
	})
	return rm.TrxGoods
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
