package internalhttp

import (
	_ "github.com/alexMolokov/hw-otus-microservices/docs/api"
	"github.com/alexMolokov/hw-otus-microservices/internal/common"
	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func (s *Server) NewRouter() *router.Router {
	r := router.New()

	r.GET("/health", common.AddUserContext(true, s.LoggingRequest(s.Health)))
	r.GET("/ready", common.AddUserContext(true, s.LoggingRequest(s.Ready)))

	r.GET("/swagger/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	return r
}
