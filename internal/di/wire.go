//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"myproject/internal/controller"
	"myproject/internal/dao"
	"myproject/internal/service"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(controller.NewController, service.NewService, dao.NewDao, dao.NewDb, NewApp))
}
