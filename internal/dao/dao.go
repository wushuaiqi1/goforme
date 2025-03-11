package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	forgeDb *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		forgeDb: db,
	}
}
