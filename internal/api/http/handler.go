package internalhttp

import (
	"github.com/alexMolokov/hw-otus-microservices/internal/model"
	"github.com/valyala/fasthttp"
)

// Health ...
// @Summary Проверка здоровья сервиса
// @IDs health
// @Produce json
// @Success 200 {object} model.StatusResponse "Сервис работает корректно"
// @Router /health [get]
// @tags system
// .
func (s *Server) Health(ctx *fasthttp.RequestCtx) {
	response := newOk(ctx)
	response.Data(model.StatusResponse{
		Status: "OK",
	})
}

// Ready ...
// @Summary Проверка сервиса принимать трафик
// @IDs ready
// @Produce text/plain
// @Success 200 {string} string "Сервис может принимать трафик"
// @Failure 503 {string} string "Сервис не может принимать трафик"
// @Router /ready [get]
// @tags system
// .
func (s *Server) Ready(ctx *fasthttp.RequestCtx) {
	response := newOk(ctx)
	response.Text("OK")
}
