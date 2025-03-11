package controller

import (
	"myproject/internal/service"
)

type Controller struct {
	svc *service.Service
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc: svc,
	}
}
