package test

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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

func Test_ListRubricRecord(t *testing.T) {
	app := di.BuildApp()
	classId := 1908816
	lessonIndex := 1
	receive, err := app.Dao.ListRubricRecord(context.Background(), classId, lessonIndex, "503560")
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, 0, len(receive))
}

func Test_EsQuery(t *testing.T) {
	app := di.BuildApp()
	var EsAnswerInteractionList = []string{
		"INTERACTION_CHOICE_OPEN",
		"CHOICE_LIGHT_OPEN",
		// "INTERACTION_CHOICE_SCENE_OPEN", // 老师现场发起选择题
	}
	total, _ := app.Dao.QueryInteractionTotal("1908781", 1, EsAnswerInteractionList)
	assert.Equal(t, 5, total)
}
