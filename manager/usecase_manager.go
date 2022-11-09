package manager

import "project-ojt/usecase"

type UsecaseManager interface {
	GetDataInternetUsecase() usecase.GetInternetDataUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func (u *usecaseManager) GetDataInternetUsecase() usecase.GetInternetDataUsecase {
	return usecase.NewGetInternetDataUsecase(u.repoManager.GetJsonRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
