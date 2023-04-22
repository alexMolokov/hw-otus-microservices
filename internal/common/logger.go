package common

import (
	"github.com/alexMolokov/hw-otus-microservices/internal/logger"
	"github.com/alexMolokov/hw-otus-microservices/internal/model"
	"github.com/valyala/fasthttp"
)

func GetErrorLoggerContext(ctx *fasthttp.RequestCtx, err error, m map[string]interface{}) logger.Context {
	contextError := GetLoggerContext(ctx, m)
	if err == nil {
		contextError["error"] = nil
		return contextError
	}
	contextError["error"] = err.Error()
	return contextError
}

func GetLoggerContext(ctx *fasthttp.RequestCtx, m map[string]interface{}) logger.Context {
	uv, ok := ctx.Value(model.UserCtxKey).(model.UserContext)
	if ok {
		for k, v := range uv {
			m[k] = v
		}
	}

	return m
}
