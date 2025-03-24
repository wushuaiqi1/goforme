package dao

import (
	"github.com/go-redis/redis/v7"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/qiniu/qmgo"
)

var Provider = wire.NewSet(NewRedis, NewDb, NewDao, NewSubjectMongoClient)

type Dao struct {
	forgeDb    *gorm.DB         // mysql客户端
	redis      *redis.Client    // redis客户端
	qmgoClient *qmgo.QmgoClient // mongo客户端
}

func NewDao(db *gorm.DB, client *redis.Client, qmgoClient *qmgo.QmgoClient) *Dao {
	return &Dao{
		forgeDb:    db,
		redis:      client,
		qmgoClient: qmgoClient,
	}
}
