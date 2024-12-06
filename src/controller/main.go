package controller

import "api/src/config"

type Controller struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Controller {
	return &Controller{cfg: cfg}
}
