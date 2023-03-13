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
	r.GET("/swagger/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	return r
}
