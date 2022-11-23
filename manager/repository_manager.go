package manager

import "project-ojt/repository"

type RepositoryManager interface {
	GetJsonRepo() repository.GetDataRepository
	GetDataSensorRepo() repository.GetDataSensorRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) GetJsonRepo() repository.GetDataRepository {
	return repository.NewGetDataRepository(r.infra.Dummy())
}

func (r *repositoryManager) GetDataSensorRepo() repository.GetDataSensorRepository {
	return repository.NewGetDataSensorRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
