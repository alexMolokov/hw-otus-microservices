package context

import "github.com/alexMolokov/hw-otus-microservices/internal/logger"

// Logger интерфейс логера с контекстом.
type Logger interface {
	Debug(msg string, c logger.Context)
	Info(msg string, c logger.Context)
	Warning(msg string, c logger.Context)
	Error(msg string, c logger.Context)
	Critical(msg string, c logger.Context)
	Fatal(msg string, c logger.Context)
}

// Releaser интерфейс который позволяет получить логгер с информацией о релизе.
type Releaser interface {
	WithRelease(string) Logger
}

// Requester интерфейс который позволяет получить логгер с информацией о запросе.
type Requester interface {
	WithRequestID(string) Logger
}

// LoggerWrapper обертка для реализации интерфейса ContextLogger.
type LoggerWrapper struct {
	logger logger.Logger
}

// NewLogger возвращает новый интерфейс логирования с контекстом.
func NewLogger(l logger.Logger) Logger {
	return &LoggerWrapper{
		logger: l,
	}
}

// NewLoggerWithCallerSkip новый интерфейс логирования с контекстом с учетом игнорирования уровня CallStack.
func NewLoggerWithCallerSkip(l logger.Logger) Logger {
	tmp := l
	if skiperLogger, ok := tmp.(logger.CallerSkiper); ok {
		// Игнорируем один уровень CallStack из за обертки в ContextLogger
		tmp = skiperLogger.WithCallerSkip(1)
	}
	return NewLogger(tmp)
}

// Debug implements the logger interface.
func (l *LoggerWrapper) Debug(msg string, c logger.Context) {
	l.logger.Debug(logger.NewMessage(msg, c))
}

// Info implements the logger interface.
func (l *LoggerWrapper) Info(msg string, c logger.Context) {
	l.logger.Info(logger.NewMessage(msg, c))
}

// Warning implements the logger interface.
func (l *LoggerWrapper) Warning(msg string, c logger.Context) {
	l.logger.Warning(logger.NewMessage(msg, c))
}

// Error implements the logger interface.
func (l *LoggerWrapper) Error(msg string, c logger.Context) {
	l.logger.Error(logger.NewMessage(msg, c))
}

// Critical implements the logger interface.
func (l *LoggerWrapper) Critical(msg string, c logger.Context) {
	l.logger.Critical(logger.NewMessage(msg, c))
}

// Fatal implements the logger interface.
func (l *LoggerWrapper) Fatal(msg string, c logger.Context) {
	l.logger.Fatal(logger.NewMessage(msg, c))
}

// WithRelease возвращает логер с информацией о релизе.
func (l *LoggerWrapper) WithRelease(release string) Logger {
	if loggerReleaser, ok := l.logger.(logger.Releaser); ok {
		return NewLogger(loggerReleaser.WithRelease(release))
	}
	return l
}

// WithRequestID возвращает логер с информацией о запросе.
func (l *LoggerWrapper) WithRequestID(requestID string) Logger {
	if loggerRequester, ok := l.logger.(logger.Requester); ok {
		return NewLogger(loggerRequester.WithRequestID(requestID))
	}
	return l
}
