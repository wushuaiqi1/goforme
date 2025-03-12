package di

import (
	log "github.com/sirupsen/logrus"
	"myproject/configs/mconfig"
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

func BuildApp() (app *App) {
	// 初始化全局viper配置
	if err := mconfig.Init(); err != nil {
		log.Error(err)
		return
	}
	// 初始化实例
	app, f, err := initApp()
	if err != nil {
		// 调用处理函数
		f()
		log.Error(err)
		return
	}
	return
}
