package manager

import "project-ojt/config"

type Infra interface {
	ConfigData() config.Config
}

type infra struct {
	cfg config.Config
}

func (i *infra) ConfigData() config.Config {
	return i.cfg
}

func NewInfra(config config.Config) Infra {
	infra := infra{
		cfg: config,
	}
	return &infra
}
