package manager

import "pinjam-modal-app/usecase"

type UsecaseManager interface {
	GetCustomerUsecase() usecase.CustomerUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	cstUsecase usecase.CustomerUsecase
}

func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUsecase {
	um.cstUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo())

	return um.cstUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
