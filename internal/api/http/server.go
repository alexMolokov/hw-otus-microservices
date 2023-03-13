package internalhttp

import (
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/logger"
	"github.com/valyala/fasthttp"
)

type Application interface{}

type Server struct {
	App        Application
	Logger     logger.Logger
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

func NewServer(logger logger.Logger, app Application, addr string) *Server {
	s := &Server{
		App:    app,
		Logger: logger,
		addr:   addr,
	}

	router := s.NewRouter()
	s.httpServer = &fasthttp.Server{
		Handler:      router.Handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}
