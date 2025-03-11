package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"myproject/configs/mconfig"
	"myproject/internal/di"
)

func init() {
	log.Info("main init")
	flag.Parse()
}

var app *di.App

func main() {
	log.Info("main start")

	// 初始化全局viper配置
	if err := mconfig.Init(); err != nil {
		log.Error(err)
		return
	}

	// 初始化bean实例
	app, _, _ := di.InitApp()
	u := app.Dao.GetUser(1)
	marshal, _ := json.Marshal(u)
	log.Info("user:1:", string(marshal))
	log.Info("main end")
}
