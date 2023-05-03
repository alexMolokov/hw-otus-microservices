package app

import (
	"context"
	"errors"

	"github.com/alexMolokov/hw-otus-microservices/internal/api/config"
	"github.com/alexMolokov/hw-otus-microservices/internal/db"
	loggerContext "github.com/alexMolokov/hw-otus-microservices/internal/logger/context"
	"github.com/alexMolokov/hw-otus-microservices/internal/model"
)

var (
	ErrUserNotExists  = errors.New("user not exists")
	ErrUserNameExists = errors.New("user name exists")
)

type IStorage interface {
	UserCreate(ctx context.Context, request model.UserCreateRequest) (int64, error)
	UserUpdate(ctx context.Context, request model.UserUpdateRequest) error
	UserGet(ctx context.Context, id int64) (*model.User, error)
	UserDelete(ctx context.Context, id int64) error
}

type App struct {
	Cfg     *config.AppConfig
	Logger  loggerContext.Logger
	Storage IStorage
}

func (a *App) UserCreate(ctx context.Context, createUserRequest model.UserCreateRequest) (int64, error) {
	id, err := a.Storage.UserCreate(ctx, createUserRequest)
	if errors.Is(err, db.ErrUserNameExists) {
		return 0, ErrUserNameExists
	}
	return id, err
}

func (a *App) UserGet(ctx context.Context, id int64) (*model.User, error) {
	user, err := a.Storage.UserGet(ctx, id)
	if err != nil {
		return nil, ErrUserNotExists
	}
	return user, nil
}

func (a *App) UserUpdate(ctx context.Context, request model.UserUpdateRequest) error {
	_, err := a.UserGet(ctx, request.UserID)
	if err != nil {
		return err
	}
	return a.Storage.UserUpdate(ctx, request)
}

func (a *App) UserDelete(ctx context.Context, id int64) error {
	if _, err := a.UserGet(ctx, id); err != nil {
		return err
	}

	return a.Storage.UserDelete(ctx, id)
}

func NewApp(cfg *config.AppConfig, logger loggerContext.Logger, storage IStorage) *App {
	return &App{
		Cfg:     cfg,
		Logger:  logger,
		Storage: storage,
	}
}
