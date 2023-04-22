package internalhttp

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
)

const (
	CodeErrorParse = "parse"
)

type ResponseError struct {
	Error   bool   `json:"error" example:"true"`
	Message string `json:"message" example:"Some error message"`
}

func NewResponseError(message string) *ResponseError {
	return &ResponseError{
		Error:   true,
		Message: message,
	}
}

type ResponseOk struct {
	Error   bool   `json:"error" example:"false"`
	Message string `json:"message" example:"Ok"`
}

func NewResponseOk() *ResponseOk {
	return &ResponseOk{
		Error:   false,
		Message: "OK",
	}
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

type CommonError struct {
	Message string      `json:"message" example:"error_message"`
	Code    string      `json:"code" example:"required"`
	Data    interface{} `json:"data,omitempty"`
}

type FieldError struct {
	CommonError
	Field string `json:"field" example:"name_field"`
}

type ResponseErrors struct {
	Error  bool          `json:"error" example:"true"`
	Errors []interface{} `json:"errors"`
}

func (re *ResponseErrors) Add(e interface{}) {
	re.Errors = append(re.Errors, e)
}

func NewResponseErrors() *ResponseErrors {
	return &ResponseErrors{
		Error: true,
	}
}

// Возвращает статус code = 200 OK.
func newOk(ctx *fasthttp.RequestCtx) *responseWriter {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	return &responseWriter{ctx: ctx}
}

// Возвращает статус code = 400 bad request.
func newBadRequest(ctx *fasthttp.RequestCtx) *responseWriter {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusBadRequest)
	return &responseWriter{ctx: ctx}
}

// Возвращает статус code = 404 page not found.
func newPageNotFoundError(ctx *fasthttp.RequestCtx) *responseWriter {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusNotFound)
	return &responseWriter{ctx: ctx}
}

// Возвращает статус code = 500 internal error.
func newInternalError(ctx *fasthttp.RequestCtx) *responseWriter {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusInternalServerError)
	return &responseWriter{ctx: ctx}
}
