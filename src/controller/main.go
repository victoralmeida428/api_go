package controller

import (
	"api/src/config"
	"api/src/model"
	"api/src/repository"
)

type Controller struct {
	cfg        *config.Config
	model      *model.Model
	repository *repository.Repository
}

func New(cfg *config.Config, model *model.Model, repository *repository.Repository) *Controller {
	return &Controller{cfg: cfg, model: model, repository: repository}
}
