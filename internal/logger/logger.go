package logger

import (
	"encoding/json"
)

// Logger базовый интерфейс логера.
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

// CallerSkiper интерфейс для игнорироавния колстека.
type CallerSkiper interface {
	WithCallerSkip(int) Logger
}

// FatalSkiper интерфейс который позволяет получить логгер с игнорированием завершения программы при фатальной ошибке.
type FatalSkiper interface {
	WithIgnoreFatal() Logger
}

// Releaser интерфейс который позволяет получить логгер с информацией о релизе.
type Releaser interface {
	WithRelease(string) Logger
}

// Requester интерфейс который позволяет получить логгер с информацией о запросе.
type Requester interface {
	WithRequestID(string) Logger
}

// Context контекст логирования.
type Context map[string]interface{}

// Message сообщение с контекстом.
type Message struct {
	Message string  `json:"message"`
	Context Context `json:"context"`
}

// String преобразует структуру в строчное представление.
func (m *Message) String() string {
	if m.Context == nil {
		return m.Message
	}
	data, err := json.Marshal(&m.Context)
	if err != nil {
		return m.Message
	}
	return m.Message + " | " + string(data)
}
