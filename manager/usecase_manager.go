package manager

import "project-ojt/usecase"

type UsecaseManager interface {
	// GetDataInternetUsecase() usecase.GetInternetDataUsecase
	GetDataInternetSiteUsecase() usecase.GetInternetDataSiteUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

// func (u *usecaseManager) GetDataInternetUsecase() usecase.GetInternetDataUsecase {
// 	return usecase.NewGetInternetDataUsecase(u.repoManager.GetJsonRepo())
// }

func (u *usecaseManager) GetDataInternetSiteUsecase() usecase.GetInternetDataSiteUsecase {
	return usecase.NewGetInternetDataSiteUsecase(u.repoManager.GetJsonRepo(), u.repoManager.GetDataSensorRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
