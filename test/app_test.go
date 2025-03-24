package test

import (
	"flag"
	"github.com/magiconair/properties/assert"
	log "github.com/sirupsen/logrus"
	"myproject/internal/di"
	"testing"
)

// 注意测试包下 寻找配置是从当前目录寻找的

var b string

func init() {
	// 测试函数执行前 go-test会提前帮我们调用flag.Parse() 不需要显示调用
	// 定义在init函数才可以
	flag.StringVar(&b, "b", "", "测试参数读取")
}

func Test_db(t *testing.T) {
	app := di.BuildApp()
	u := app.Dao.GetUser(2)
	log.Info("查询用户名称：", u.Username)
}

func Test_redisPing(t *testing.T) {
	app := di.BuildApp()
	app.Dao.Ping()
}

func Test_mongoPing(t *testing.T) {
	app := di.BuildApp()
	res, _ := app.Dao.PingMongo()
	assert.Equal(t, true, res)
}

func Test_redisInsertAndGet(t *testing.T) {
	app := di.BuildApp()
	name := "name"
	app.Dao.Set(name, "张三", -1)
	val := app.Dao.Get(name)
	log.Info("Test_redisInsertAndGet val:", val)
}
