package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var ErrConnectDB = errors.New("can't connect to database")

type Connection struct {
	DB *sqlx.DB
}

func (c *Connection) GetDB() *sqlx.DB {
	return c.DB
}

func (c *Connection) Close() error {
	return c.DB.Close()
}

func NewConnection(cfg config.DBConf) (*Connection, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s binary_parameters=%s",
		cfg.User, cfg.Password, cfg.Name, cfg.Host, cfg.Port, cfg.SslMode, cfg.BinaryParams)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s %w", ErrConnectDB, err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Minute)

	return &Connection{
		DB: db,
	}, nil
}
