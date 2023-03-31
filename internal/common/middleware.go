package common

import (
	"os"

	"github.com/alexMolokov/hw-otus-microservices/internal/model"
	"github.com/alexMolokov/hw-otus-microservices/internal/utils"
	"github.com/valyala/fasthttp"
)

var hostName string

func init() {
	hostName, _ = os.Hostname()
}

func AddUserContext(addBody bool, next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		req := &ctx.Request
		userContext := model.UserContext{
			"url":         string(req.RequestURI()),
			"ip":          ctx.RemoteIP().String(),
			"http_method": string(ctx.Method()),
			"server":      hostName,
			"referrer":    string(ctx.Referer()),
			"user_agent":  string(ctx.UserAgent()),
			"request_id":  string(req.Header.Peek("X-Request-Id")),
		}
		if addBody {
			userContext["body"] = utils.SafeJSON(string(ctx.PostBody()))
		}
		ctx.SetUserValue(model.UserCtxKey, userContext)

		next(ctx)
	}
}
