package dao

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
)

// NewSubjectMongoClient 创建mongo客户端
func NewSubjectMongoClient() (qmgoClient *qmgo.QmgoClient, err error) {
	ctx := context.Background()
	cfg := &qmgo.Config{
		Uri:      viper.GetString("mgo.host"),
		Database: viper.GetString("mgo.db"),
		Coll:     viper.GetString("mgo.collection"),
	}
	qmgoClient, err = qmgo.Open(ctx, cfg)
	if err != nil {
		return
	}
	defer func() {
		if err = qmgoClient.Close(ctx); err != nil {
			panic(err)
		}
	}()
	return
}

func (d *Dao) PingMongo() (bool, error) {
	err := d.qmgoClient.Ping(1)
	if err != nil {
		return false, err
	}
	return true, nil
}
