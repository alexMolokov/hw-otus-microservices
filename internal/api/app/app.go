package app

import (
	"github.com/alexMolokov/hw-otus-microservices/internal/api/config"
	"github.com/alexMolokov/hw-otus-microservices/internal/logger"
)

type App struct {
	Cfg    *config.AppConfig
	Logger logger.Logger
}

func NewApp(cfg *config.AppConfig, logger logger.Logger) *App {
	return &App{
		Cfg:    cfg,
		Logger: logger,
	}
}
