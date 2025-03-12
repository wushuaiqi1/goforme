package test

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"myproject/configs/mconfig"
	"myproject/internal/di"
	"testing"
)

// 注意测试包下 寻找配置是从当前目录寻找的

var b string

func init() {
	// 测试函数执行前 go-test会提前帮我们调用flag.Parse() 不需要显示调用
	// 定义在init函数才可以
	flag.StringVar(&b, "ss", "", "ss")
}

func Test_app(t *testing.T) {
	if b != "" {
		log.Info(b)
	}
	if err := mconfig.Init(); err != nil {
		log.Error(err)
	}
	app, _, err := di.InitApp()
	if err != nil {
		log.Error(err)
	}
	u := app.Dao.GetUser(2)
	log.Info("查询用户名称：", u.Username)
}
