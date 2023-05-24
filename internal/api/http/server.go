package internalhttp

import (
	"context"
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/common"
	fasthttpp "github.com/alexMolokov/hw-otus-microservices/internal/common/prometheus"
	loggerCtx "github.com/alexMolokov/hw-otus-microservices/internal/logger/context"
	"github.com/alexMolokov/hw-otus-microservices/internal/model"
	"github.com/valyala/fasthttp"
)

type Application interface {
	UserCreate(ctx context.Context, createUserRequest model.UserCreateRequest) (int64, error)
	UserGet(ctx context.Context, id int64) (*model.User, error)
	UserUpdate(ctx context.Context, request model.UserUpdateRequest) error
	UserDelete(ctx context.Context, id int64) error
}

type Server struct {
	App        Application
	Logger     loggerCtx.Logger
	httpServer *fasthttp.Server
	addr       string
}

func (s *Server) Start() error {
	err := s.httpServer.ListenAndServe(s.addr)
	return err
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown()
}

func NewServer(logger loggerCtx.Logger, app Application, addr string) *Server {
	s := &Server{
		App:    app,
		Logger: logger,
		addr:   addr,
	}

	router := s.NewRouter()
	prm := fasthttpp.NewWith("", "", "")
	prm.Register(router)

	s.httpServer = &fasthttp.Server{
		Handler:      common.AddUserContext(false, s.LoggingRequest(prm.Middleware(router.Handler))),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}
