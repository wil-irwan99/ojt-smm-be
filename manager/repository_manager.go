package manager

import "project-ojt/repository"

type RepositoryManager interface {
	GetJsonRepo() repository.GetDataRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) GetJsonRepo() repository.GetDataRepository {
	return repository.NewGetDataRepository(r.infra.ConfigData())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
