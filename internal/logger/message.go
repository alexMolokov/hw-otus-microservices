package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// NewMessage возвращает экземпляр сообщения.
func NewMessage(message string, context Context) *Message {
	return &Message{
		Message: message,
		Context: context,
	}
}

// NewMessageFromParams возвращает экземпляр сообщения на основе параметров.
func NewMessageFromParams(v []interface{}) *Message {
	size := len(v)
	if size == 0 {
		return NewMessage("", nil)
	}
	if size == 1 {
		return NewMessageFromParam(v[0])
	}
	return NewMessage(strings.TrimRight(fmt.Sprintln(v...), "\n"), nil)
}

// NewMessageFromParam возвращает экземпляр сообщения на основе параметра.
func NewMessageFromParam(v interface{}) *Message {
	switch t := v.(type) {
	case Message:
		return &t
	case *Message:
		if t == nil {
			return NewMessage("", nil)
		}
		return NewMessage(t.Message, t.Context)
	}
	return NewMessage(fmt.Sprint(v), nil)
}

// ErrorContext возвращет контекст с информацией о ошибке.
func ErrorContext(err error) Context {
	if err == nil {
		return SimpleContext("error", nil)
	}
	return SimpleContext("error", err.Error())
}

// SimpleContext возвращает простой контекст.
func SimpleContext(key string, val interface{}) Context {
	return Context{
		key: val,
	}
}

// GetCallerPath Получение файла и номера строки логирования.
func GetCallerPath(callerDepth int) (string, int) {
	_, file, line, ok := runtime.Caller(callerDepth)
	if !ok {
		return "", 0
	}
	return filepath.Base(file), line
}
