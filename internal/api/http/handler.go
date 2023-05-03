package internalhttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/api/app"
	"github.com/alexMolokov/hw-otus-microservices/internal/common"
	"github.com/alexMolokov/hw-otus-microservices/internal/model"
	valid "github.com/asaskevich/govalidator"
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

// UserCreate ...
// @Summary Создание пользователя.
// @Accept  json
// @Produce json
// @Param _ body model.UserCreateRequest true "Запрос на создание"
// @Success 202 {object} model.UserCreateResponse "OK"
// @Failure 400 {object} ResponseErrors "Bad request"
// @Failure 500 {object} ResponseError "Some error"
// @Router /api/v1/user [post]
// @tags v1
// .
func (s *Server) UserCreate(ctx *fasthttp.RequestCtx) {
	var createUserRequest model.UserCreateRequest
	responseErrors := NewResponseErrors()

	body := ctx.PostBody()
	err := json.Unmarshal(body, &createUserRequest)
	if err != nil {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "can't parse json", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	responseErrors = getValidateErrorResponse(createUserRequest)
	if responseErrors != nil {
		response := newBadRequest(ctx)
		response.Data(responseErrors)
		return
	}

	id, err := s.App.UserCreate(ctx, createUserRequest)
	if err == nil {
		response := newOk(ctx)
		response.Data(model.NewUserCreateResponse(id))
		return
	}
	if errors.Is(err, app.ErrUserNameExists) {
		response := newOk(ctx)
		response.Data(NewResponseError("User name exists"))
		return
	}

	s.Logger.Error("Some troubles with db",
		common.GetErrorLoggerContext(ctx, err, map[string]interface{}{
			"request": string(body),
			"action":  "UserCreate",
		}))
	response := newInternalError(ctx)
	response.Data(NewResponseError("Can't create user"))
}

// UserUpdate ...
// @Summary Изменение пользователя.
// @Accept  json
// @Produce json
// @Param id path int64 true "ID пользователя" example(1)
// @Param _ body model.UserUpdateRequest true "Запрос на изменение данных"
// @Success 202 {object} ResponseOk "OK"
// @Failure 400 {object} ResponseErrors "Bad request"
// @Failure 500 {object} ResponseError "Some error"
// @Router /api/v1/user/{id} [put]
// @tags v1
// .
func (s *Server) UserUpdate(ctx *fasthttp.RequestCtx) {
	responseErrors := NewResponseErrors()
	val, ok := ctx.UserValue("id").(string)
	if !ok {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "can't get id from query", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "not valid id. id must be int", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	var userUpdateRequest model.UserUpdateRequest
	body := ctx.PostBody()
	err = json.Unmarshal(body, &userUpdateRequest)
	if err != nil {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "can't parse json", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	responseErrors = getValidateErrorResponse(userUpdateRequest)
	if responseErrors != nil {
		response := newBadRequest(ctx)
		response.Data(responseErrors)
		return
	}

	userUpdateRequest.UserID = id
	err = s.App.UserUpdate(ctx, userUpdateRequest)
	if err != nil {
		if userNotFound(ctx, id, err) {
			return
		}
		s.Logger.Error("Some troubles with db",
			common.GetErrorLoggerContext(ctx, err, map[string]interface{}{
				"request": string(body),
				"action":  "UserUpdate",
			}))
		response := newInternalError(ctx)
		response.Data(NewResponseError(fmt.Sprintf("Can't update user by id  %d", id)))
		return
	}

	response := newOk(ctx)
	response.Data(NewResponseOk())
}

// UserGet ...
// @Summary Получение пользователя по ID
// @Produce json
// @Param id path int64 true "ID пользователя" example(1)
// @Success 200 {object} model.User "Пользователь"
// @Failure 400 {object} ResponseError "Bad request"
// @Failure 404 {object} ResponseError "Not found"
// @Failure 500 {object} ResponseError "Some error"
// @Router  /api/v1/user/{id} [get]
// @tags v1
// .
func (s *Server) UserGet(ctx *fasthttp.RequestCtx) {
	responseErrors := NewResponseErrors()
	val, ok := ctx.UserValue("id").(string)
	if !ok {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "can't get id from query", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "not valid id. id must be int", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	user, err := s.App.UserGet(ctx, id)
	if err == nil {
		response := newOk(ctx)
		response.Data(user)
		return
	}

	if userNotFound(ctx, id, err) {
		return
	}

	s.Logger.Error("Some troubles with db",
		common.GetErrorLoggerContext(ctx, err, map[string]interface{}{
			"request": id,
			"action":  "UserGet",
		}))
	response := newInternalError(ctx)
	response.Data(NewResponseError(fmt.Sprintf("Can't get user by id = %d", id)))
}

// UserDelete ...
// @Summary Удаление пользователя по ID
// @Produce json
// @Param id path int64 true "ID пользователя" example(1)
// @Success 200 {object} ResponseOk "Ok"
// @Failure 400 {object} ResponseError "Bad request"
// @Failure 404 {object} ResponseError "Not found"
// @Failure 500 {object} ResponseError "Some error"
// @Router  /api/v1/user/{id} [delete]
// @tags v1
// .
func (s *Server) UserDelete(ctx *fasthttp.RequestCtx) {
	responseErrors := NewResponseErrors()
	val, ok := ctx.UserValue("id").(string)
	if !ok {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "can't get id from query", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		response := newBadRequest(ctx)
		responseErrors.Add(CommonError{Message: "not valid id. id must be int", Code: CodeErrorParse})
		response.Data(responseErrors)
		return
	}

	err = s.App.UserDelete(ctx, id)
	if err == nil {
		response := newOk(ctx)
		response.Data(NewResponseOk())
		return
	}

	if userNotFound(ctx, id, err) {
		return
	}

	s.Logger.Error("Some troubles with db",
		common.GetErrorLoggerContext(ctx, err, map[string]interface{}{
			"request": id,
			"action":  "UserDelete",
		}))
	response := newInternalError(ctx)
	response.Data(NewResponseError(fmt.Sprintf("Can't delete user by id = %d", id)))
}

func (s *Server) SometimesError(ctx *fasthttp.RequestCtx) {
	rand.Seed(time.Now().UnixNano())
	tm := rand.Intn(100) // nolint
	switch tm % 5 {
	case 0:
		response := newInternalError(ctx)
		response.Data(NewResponseError("Generated 500"))
	case 4:
		response := newBadRequest(ctx)
		response.Data(NewResponseError("Generated 400"))
	default:
		response := newOk(ctx)
		response.Text("Ok")
	}
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

func getValidateErrorResponse(i interface{}) *ResponseErrors {
	_, err := valid.ValidateStruct(i)
	if err == nil {
		return nil
	}

	responseErrors := NewResponseErrors()
	makeValidationErrors(responseErrors, err, 5)
	return responseErrors
}

func makeValidationErrors(responseErrors *ResponseErrors, err error, level int) {
	if level < 0 {
		return
	}

	switch value := err.(type) { //nolint
	case valid.Errors:
		for _, e := range value {
			makeValidationErrors(responseErrors, e, level-1)
		}
	case valid.Error:
		responseErrors.Add(FieldError{
			Field: value.Name,
			CommonError: CommonError{
				Message: value.Error(),
				Code:    value.Validator,
			},
		})
	default:
		return
	}
}

func userNotFound(ctx *fasthttp.RequestCtx, id int64, err error) bool {
	if errors.Is(err, app.ErrUserNotExists) {
		response := newPageNotFoundError(ctx)
		response.Data(NewResponseError(fmt.Sprintf("user by id = %d not found", id)))
		return true
	}
	return false
}
