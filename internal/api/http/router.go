package internalhttp

import (
	_ "github.com/alexMolokov/hw-otus-microservices/docs/api"
	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (s *Server) NewRouter() *router.Router {
	r := router.New()

	r.GET("/health", s.Health)
	r.GET("/ready", s.Ready)
	r.POST("/api/v1/user", s.UserCreate)
	r.GET("/api/v1/user/{id}", s.UserGet)
	r.DELETE("/api/v1/user/{id}", s.UserDelete)
	r.PUT("/api/v1/user/{id}", s.UserUpdate)

	r.GET("/swagger/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	r.GET("/sometimes/error", s.SometimesError)
	return r
}
