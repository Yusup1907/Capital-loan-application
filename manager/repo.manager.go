package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetCustomerRepo() repository.CustomerRepo
}

type repoManager struct {
	infraManager InfraManager

	cstRepo repository.CustomerRepo
}

var onceLoadCustomerRepo sync.Once

func (rm *repoManager) GetCustomerRepo() repository.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.cstRepo = repository.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.cstRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
