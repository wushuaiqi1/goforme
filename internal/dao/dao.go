package dao

import (
	"github.com/go-redis/redis/v7"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var Provider = wire.NewSet(NewRedis, NewDb, NewDao)

type Dao struct {
	forgeDb *gorm.DB
	redis   *redis.Client
}

func NewDao(db *gorm.DB, client *redis.Client) *Dao {
	return &Dao{
		forgeDb: db,
		redis:   client,
	}
}
