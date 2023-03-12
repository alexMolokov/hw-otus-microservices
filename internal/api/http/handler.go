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
