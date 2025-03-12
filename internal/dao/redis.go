package dao

import (
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func NewRedis() (client *redis.Client, err error) {
	var cfg redis.Options
	// 指针类型本身就是地址
	if err = viper.UnmarshalKey("redis", &cfg); err != nil {
		return nil, err
	}
	log.Info("NewRedis cfg:", cfg.Addr)
	client = redis.NewClient(&cfg)
	return
}

// Ping 测试redis是否连通
func (d *Dao) Ping() (bool, error) {
	ping := d.redis.Ping()
	err := ping.Err()
	if err != nil {
		log.Error("redis Ping err:", err)
		return false, err
	}
	log.Info("redis ping success")
	return true, nil
}

// Set 设置key duration<0 永久存储
func (d *Dao) Set(key string, val string, duration time.Duration) {
	d.redis.Set(key, val, duration)
}

func (d *Dao) Get(key string) string {
	return d.redis.Get(key).Val()
}
