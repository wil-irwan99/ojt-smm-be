package manager

import "project-ojt/usecase"

type UsecaseManager interface {
	GetDataInternetSiteUsecase() usecase.GetInternetDataSiteUsecase
	GetDataCPUSiteUsecase() usecase.GetCPUDataSiteUsecase
	AddNewSensorUsecase() usecase.AddNewSensorUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func (u *usecaseManager) GetDataInternetSiteUsecase() usecase.GetInternetDataSiteUsecase {
	return usecase.NewGetInternetDataSiteUsecase(u.repoManager.GetJsonRepo(), u.repoManager.GetDataSensorRepo())
}

func (u *usecaseManager) GetDataCPUSiteUsecase() usecase.GetCPUDataSiteUsecase {
	return usecase.NewGetCPUDataSiteUsecase(u.repoManager.GetJsonRepo(), u.repoManager.GetDataSensorRepo())
}

func (u *usecaseManager) AddNewSensorUsecase() usecase.AddNewSensorUsecase {
	return usecase.NewAddNewSensorUsecase(u.repoManager.GetDataSensorRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
