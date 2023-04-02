package internalhttp

import (
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/common"
	"github.com/alexMolokov/hw-otus-microservices/internal/logger"
	"github.com/alexMolokov/hw-otus-microservices/internal/utils"
	"github.com/valyala/fasthttp"
)

func (s *Server) LoggingRequest(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func(begin time.Time) {
			s.Logger.Info("Request", common.GetLoggerContext(ctx,
				logger.Context{
					"response_http_code": ctx.Response.StatusCode(),
					"response_body":      utils.SafeJSON(string(ctx.Response.Body())),
					"response_time":      time.Since(begin),
				}))
		}(time.Now())

		next(ctx)
	}
}
