package internalhttp

import (
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/logger/context"
	fasthttpprom "github.com/carousell/fasthttp-prometheus-middleware"
	"github.com/valyala/fasthttp"
)

type Application interface{}

type Server struct {
	App        Application
	Logger     context.Logger
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

func NewServer(logger context.Logger, app Application, addr string) *Server {
	s := &Server{
		App:    app,
		Logger: logger,
		addr:   addr,
	}

	router := s.NewRouter()
	p := fasthttpprom.NewPrometheus("otus-microservice")
	p.Use(router)

	s.httpServer = &fasthttp.Server{
		Handler:      router.Handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}
