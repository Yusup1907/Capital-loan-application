package manager

<<<<<<< HEAD
import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetCustomerUsecase() usecase.CustomerUsecase
	GetProductUsecase() usecase.ProductUsecase
	GetCategoryLoanUsecase() usecase.CategoryLoanUsecase
	GetCategoryProductUsecase() usecase.CategoryProductUsecase
	GetGoodsUsecase() usecase.GoodsUsecase
}

type usecaseManager struct {
	repoManager         RepoManager
	cstUsecase usecase.CustomerUsecase
	productUsecase      usecase.ProductUsecase
	loanApp             usecase.LoanApplicationUsecase
	categoryLoanUsecase usecase.CategoryLoanUsecase

	onceLoadUsecase        sync.Once
	onceLoadCustomerUsecase        sync.Once
	onceLoadGetCategoryProductUsecase sync.Once
	onceLoadGetGoodsUsecase sync.Once
	onceLoadProductUsecase sync.Once
	onceLoadLoanAppUsecase sync.Once
}

func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUsecase {
	um.onceLoadCustomerUsecase.Do(func()  {
		um.cstUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo())
	})

	return um.cstUsecase
}

func (um *usecaseManager) GetCategoryLoanUsecase() usecase.CategoryLoanUsecase {
	um.onceLoadUsecase.Do(func() {
		um.categoryLoanUsecase = usecase.NewCategoryLoanUsecase(um.repoManager.GetCategoryLoanRepo())
	})
	return um.categoryLoanUsecase
}
func (um *usecaseManager) GetProductUsecase() usecase.ProductUsecase {
	um.onceLoadProductUsecase.Do(func() {
		um.productUsecase = usecase.NewProductUseCase(um.repoManager.GetProductRepo())
	})
	return um.productUsecase
}

func (um *usecaseManager) GetLoanAppUsecase() usecase.LoanApplicationUsecase {
	um.onceLoadLoanAppUsecase.Do(func() {
		um.loanApp = usecase.NewLoanApplicationUseCase(um.repoManager.GetLoanApplicationRepo())
	})
	return um.loanApp
}
func (um *usecaseManager) GetCategoryProductUsecase() usecase.CategoryProductUsecase {
	um.onceLoadGetCategoryProductUsecase.Do(func() {
		um.cpUsecase = usecase.NewCategoryProductUsecase(um.repoManager.GetCategoryProductRepo())
	})
	return um.cpUsecase
}

func (um *usecaseManager) GetGoodsUsecase() usecase.GoodsUsecase{
	um.onceLoadGetGoodsUsecase.Do(func() {
		um.gUsecase = usecase.NewGoodsUsecase(um.repoManager.GetGoodsRepo())
	})
	return um.gUsecase

}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
