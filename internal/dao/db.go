package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type DbConfig struct {
	DSN          string
	Active       int           // pool
	Idle         int           // pool
	IdleTimeout  time.Duration // connect max life time.
	QueryTimeout time.Duration // query sql timeout
	ExecTimeout  time.Duration // execute sql timeout
	TranTimeout  time.Duration // transaction sql timeout
}

func NewDb() (db *gorm.DB, err error) {
	var config DbConfig

	if err := viper.UnmarshalKey("db", &config); err != nil {
		log.Error(err)
		return nil, err
	}
	db, err = gorm.Open("mysql", config.DSN)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	db.DB().SetMaxIdleConns(config.Idle)
	db.DB().SetMaxOpenConns(config.Active)
	db.DB().SetConnMaxLifetime(config.IdleTimeout)
	return
}
