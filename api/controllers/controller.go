package controllers

import (
  "github.com/Salz-Team/salz/api/config"
)

type Controller struct {
  cfg *config.Config
}

func NewController(cfg *config.Config) *Controller {
  return &Controller{
    cfg: cfg,
  }
}
