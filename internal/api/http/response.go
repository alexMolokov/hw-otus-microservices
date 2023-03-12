package internalhttp

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
)

type ResponseError struct {
	Error   bool   `json:"error" example:"true"`
	Message string `json:"message" example:"Some error message"`
}

type responseWriter struct {
	ctx *fasthttp.RequestCtx
}

func (rw *responseWriter) ErrorMessage(message string) {
	resp := ResponseError{
		Error:   true,
		Message: message,
	}
	jsonResp, _ := json.Marshal(resp)
	rw.ctx.SetBody(jsonResp)
}

func (rw *responseWriter) Data(v interface{}) {
	jsonResp, _ := json.Marshal(v)
	rw.ctx.SetBody(jsonResp)
}

func (rw *responseWriter) Text(message string) {
	rw.ctx.SetBody([]byte(message))
}

// Возвращает статус code = 200 OK.
func newOk(ctx *fasthttp.RequestCtx) *responseWriter {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	return &responseWriter{ctx: ctx}
}
