package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetCategoryLoanRepo() repository.CategoryLoanRepo
	GetUserRepo() repository.UserRepo
	GetLoginRepo() repository.LoginRepo
}

type repoManager struct {
	infraManager         InfraManager
	categoryLoanRepo     repository.CategoryLoanRepo
	userRepo             repository.UserRepo
	loginRepo            repository.LoginRepo
	onceLoadCategoryLoan sync.Once
	onceLoadUserRepo     sync.Once
	onceLoadLoginRepo    sync.Once
}

func (rm *repoManager) GetCategoryLoanRepo() repository.CategoryLoanRepo {
	rm.onceLoadCategoryLoan.Do(func() {
		rm.categoryLoanRepo = repository.NewCategoryLoanRepo(rm.infraManager.GetDB())
	})
	return rm.categoryLoanRepo
}
func (rm *repoManager) GetUserRepo() repository.UserRepo {
	rm.onceLoadUserRepo.Do(func() {
		rm.userRepo = repository.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.userRepo
}
func (rm *repoManager) GetLoginRepo() repository.LoginRepo {
	rm.onceLoadLoginRepo.Do(func() {
		rm.loginRepo = repository.NewLoginRepo(rm.infraManager.GetDB())
	})
	return rm.loginRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
