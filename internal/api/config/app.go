package config

import (
	"context"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
)

type AppConfig struct {
	Logger LoggerConf
	HTTP   HTTPConf
}

func NewAppConfig(fileName string) (*AppConfig, error) {
	loader := confita.NewLoader(
		file.NewBackend(fileName),
		env.NewBackend(),
	)

	cfg := &AppConfig{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := loader.Load(ctx, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
