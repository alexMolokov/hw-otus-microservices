package app

import (
	"github.com/alexMolokov/hw-otus-microservices/internal/api/config"
	"github.com/alexMolokov/hw-otus-microservices/internal/logger/context"
)

type App struct {
	Cfg    *config.AppConfig
	Logger context.Logger
}

func NewApp(cfg *config.AppConfig, logger context.Logger) *App {
	return &App{
		Cfg:    cfg,
		Logger: logger,
	}
}
