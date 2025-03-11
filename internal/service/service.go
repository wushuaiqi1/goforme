package service

import "myproject/internal/dao"

type Service struct {
	dao *dao.Dao
}

func NewService(dao *dao.Dao) *Service {
	return &Service{
		dao: dao,
	}
}
