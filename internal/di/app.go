package di

import (
	"myproject/internal/controller"
	"myproject/internal/dao"
	"myproject/internal/service"
)

type App struct {
	Controller *controller.Controller
	Service    *service.Service
	Dao        *dao.Dao
}

func NewApp(controller *controller.Controller, service *service.Service, dao *dao.Dao) *App {
	return &App{
		Controller: controller,
		Service:    service,
		Dao:        dao,
	}
}
