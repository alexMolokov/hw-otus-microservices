package zap

import (
	"fmt"
	"os"

	"github.com/alexMolokov/hw-otus-microservices/internal/api/config"
	"github.com/alexMolokov/hw-otus-microservices/internal/logger"
	"github.com/alexMolokov/hw-otus-microservices/internal/logger/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger ZAP логер.
type Logger struct {
	logger *zap.Logger
}

// NewLogger возвращает новый экземпляр логера.
func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Debug implements the logger interface.
func (l *Logger) Debug(v ...interface{}) {
	m := logger.NewMessageFromParams(v)
	fields := toFields(m.Context)
	l.logger.Debug(m.Message, fields...)
}

// Info implements the logger interface.
func (l *Logger) Info(v ...interface{}) {
	m := logger.NewMessageFromParams(v)
	fields := toFields(m.Context)
	l.logger.Info(m.Message, fields...)
}

// Warning implements the logger interface.
func (l *Logger) Warning(v ...interface{}) {
	m := logger.NewMessageFromParams(v)
	fields := toFields(m.Context)
	l.logger.Warn(m.Message, fields...)
}

// Error implements the logger interface.
func (l *Logger) Error(v ...interface{}) {
	m := logger.NewMessageFromParams(v)
	fields := toFields(m.Context)
	l.logger.Error(m.Message, fields...)
}

// Critical implements the logger interface.
func (l *Logger) Critical(v ...interface{}) {
	m := logger.NewMessageFromParams(v)
	fields := toFields(m.Context)
	l.logger.Panic(m.Message, fields...)
}

// Fatal implements the logger interface.
func (l *Logger) Fatal(v ...interface{}) {
	m := logger.NewMessageFromParams(v)
	fields := toFields(m.Context)
	l.logger.Fatal(m.Message, fields...)
}

// WithCallerSkip возвращает новый логер с пропуском вызовов.
func (l *Logger) WithCallerSkip(skip int) logger.Logger {
	return NewLogger(l.logger.WithOptions(zap.AddCallerSkip(skip)))
}

// WithIgnoreFatal игнорироание фатальных ошибок.
func (l *Logger) WithIgnoreFatal() logger.Logger {
	return NewLogger(l.logger.WithOptions(zap.WithFatalHook(zapcore.WriteThenPanic)))
}

// WithRelease логер с информацией о релизе.
func (l *Logger) WithRelease(release string) logger.Logger {
	return NewLogger(l.logger.With(zap.String("release", release)))
}

// WithRequestID логер с информацией о запросе.
func (l *Logger) WithRequestID(requestID string) logger.Logger {
	return NewLogger(l.logger.With(zap.String("request_id", requestID)))
}

// LogLevel возвращет уровень логирования zap.
func LogLevel(l logger.Level) zapcore.Level {
	switch l {
	case logger.DebugLevel:
		return zapcore.DebugLevel
	case logger.InfoLevel, logger.NoticeLevel:
		return zapcore.InfoLevel
	case logger.WarnLevel:
		return zapcore.WarnLevel
	case logger.ErrorLevel:
		return zapcore.ErrorLevel
	case logger.CriticalLevel:
		return zapcore.PanicLevel
	case logger.FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.Level(l)
	}
}

// NewLoggerWithContextWrapper создает логгер с выделением контекстных полей.
func NewLoggerWithContextWrapper(
	name string,
	encoder zapcore.Encoder,
	contextPrefix string,
	ws zapcore.WriteSyncer,
	logLevel, stacktraceLevel zapcore.Level,
) *Logger {
	ctxtEncoder := NewEncoderContextWrapper(encoder, contextPrefix)
	core := zapcore.NewCore(ctxtEncoder, ws, logLevel)
	return NewLogger(
		zap.New(core,
			zap.AddCaller(),
			zap.AddStacktrace(stacktraceLevel),
			zap.AddCallerSkip(1),
		).Named(name),
	)
}

// NewStdoutJSONLogger логирование в stdout.
func NewStdoutJSONLogger(name string, logLevel zapcore.Level) *Logger {
	return NewLoggerWithContextWrapper(
		name,
		zapcore.NewJSONEncoder(NewProductionEncoderConfig()),
		"_ctxt_",
		os.Stdout,
		logLevel,
		zap.ErrorLevel,
	)
}

// toFields convert context to zap.Filed.
func toFields(c logger.Context) []zap.Field {
	fields := make([]zap.Field, 0, len(c))
	for key, value := range c {
		fields = append(fields, zap.Any(key, value))
	}
	return fields
}

func InitLogger(cfg config.LoggerConf) (context.Logger, error) {
	lvl := logger.InfoLevel

	err := lvl.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal log-level: %w", err)
	}

	l := getZapLogger(lvl, cfg)

	return context.NewLoggerWithCallerSkip(l), nil
}

// getZapLogger создает и настраивает логер на основе Zap.
func getZapLogger(lvl logger.Level, cfg config.LoggerConf) *Logger {
	return NewStdoutJSONLogger(cfg.Facility, LogLevel(lvl))
}
