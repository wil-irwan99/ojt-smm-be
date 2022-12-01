package delivery

import (
	"project-ojt/config"
	"project-ojt/delivery/controller"
	"project-ojt/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	managerRepo    manager.RepositoryManager
	infra          manager.Infra
	managerUsecase manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	managerRepo := manager.NewRepositoryManager(infra)
	managerUsecase := manager.NewUseCaseManager(managerRepo)
	host := appConfig.ApiConfig.Url

	return &appServer{
		managerRepo:    managerRepo,
		infra:          infra,
		managerUsecase: managerUsecase,
		engine:         r,
		host:           host,
	}
}

func (a *appServer) initControllers() {
	controller.NewGetDataSiteController(a.engine, a.managerUsecase.GetDataInternetSiteUsecase(), a.managerUsecase.GetDataCPUSiteUsecase(), a.infra.ConfigData())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
