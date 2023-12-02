package controller

import (
	"github.com/fnazarov97/json_crud/config"
	"github.com/fnazarov97/json_crud/storage"
)

type Controller struct {
	cfg   *config.Config
	store storage.StorageI
}

// Controller Constructor
func NewController(cfg *config.Config, store storage.StorageI) *Controller {
	return &Controller{
		cfg:   cfg,
		store: store,
	}
}
