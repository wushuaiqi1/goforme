package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"myproject/configs/mconfig"
	"myproject/internal/di"
)

// 标志定义顺序：先定义标志变量(init函数中)->调用flag.Parse函数解析赋值(main函数或者main函数init函数之中)
// 定义标志变量要在调用flag.Parse函数之前
// flag.Parse函数更推荐在主函数开头使用 这样能保证所有引入的包初始化完成
// 如果在init函数使用要在主函数的init函数使用更合适，如果定义在其他包中，可能导致部分声明的变量标志未能生效 在定义之前
func init() {
	log.Info("main init")
	flag.Parse()
}

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
	log.Info("查询数据库结果:user:1:", string(marshal))
	log.Info("main end")
}
