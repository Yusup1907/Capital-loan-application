package manager

import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetCategoryLoanUsecase() usecase.CategoryLoanUsecase
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
}

type usecaseManager struct {
	repoManager                 RepoManager
	categoryLoanUsecase         usecase.CategoryLoanUsecase
	userUsecase                 usecase.UserUsecase
	loginUsecase                usecase.LoginUsecase
	onceLoadCategoryLoanUsecase sync.Once
	onceLoadUserUsecase         sync.Once
	onceLoadLoginUsecase        sync.Once
}

func (um *usecaseManager) GetCategoryLoanUsecase() usecase.CategoryLoanUsecase {
	um.onceLoadCategoryLoanUsecase.Do(func() {
		um.categoryLoanUsecase = usecase.NewCategoryLoanUsecase(um.repoManager.GetCategoryLoanRepo())
	})
	return um.categoryLoanUsecase
}
func (um *usecaseManager) GetUserUsecase() usecase.UserUsecase {
	um.onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.userUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUsecase {
	um.onceLoadLoginUsecase.Do(func() {
		um.loginUsecase = usecase.NewLoginUsecase(um.repoManager.GetLoginRepo())
	})
	return um.loginUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
