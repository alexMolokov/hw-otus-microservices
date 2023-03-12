package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexMolokov/hw-otus-microservices/internal/api/app"
	"github.com/alexMolokov/hw-otus-microservices/internal/api/config"
	internalhttp "github.com/alexMolokov/hw-otus-microservices/internal/api/http"
	logger2 "github.com/alexMolokov/hw-otus-microservices/internal/logger"
)

var configFile string

func init() {
	flag.StringVar(
		&configFile,
		"config",
		"./configs/api.json",
		"Path to configuration file",
	)
}

// @title Microservices Otus Service (Курс микросервисы Otus)
// @version 1.0
// @description Описание API методов
// @BasePath /
// .
func main() {
	flag.Parse()

	cfg, err := config.NewAppConfig(configFile)
	if err != nil {
		log.Fatalf("Can't load config: %#v", err)
	}

	logger, err := logger2.New(cfg.Logger)
	if err != nil {
		log.Fatalf("Can't init logger: %#v", err)
	}

	application := app.NewApp(cfg, logger)
	server := internalhttp.NewServer(logger, application, fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port))

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		logger.Info("otus-microservices api is running...")
		if err := server.Start(); err != nil {
			logger.Error("failed to start http server", logger2.ErrorContext(err))
			cancel()
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	if err := server.Stop(); err != nil {
		logger.Error("failed to stop http server error correctly", logger2.ErrorContext(err))
	} else {
		logger.Info("otus-microservices is stopped", logger2.Context{})
	}
}
