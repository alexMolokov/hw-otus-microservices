package logger

import (
	"github.com/alexMolokov/hw-otus-microservices/internal/api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warning(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type ZLogger struct {
	zp *zap.SugaredLogger
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func New(c config.LoggerConf) (*ZLogger, error) {
	cfg := zap.NewProductionConfig()

	cfg.Level = zap.NewAtomicLevelAt(getLevel(c.Level))
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05 -0700")

	loggerZap, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	logger := &ZLogger{
		zp: loggerZap.Sugar(),
	}

	return logger, nil
}

func (l *ZLogger) Debug(msg string, args ...interface{}) {
	l.zp.Debugf(msg, args...)
}

func (l *ZLogger) Info(msg string, args ...interface{}) {
	l.zp.Infof(msg, args...)
}

func (l *ZLogger) Warning(msg string, args ...interface{}) {
	l.zp.Warnf(msg, args...)
}

func (l *ZLogger) Error(msg string, args ...interface{}) {
	l.zp.Errorf(msg, args...)
}
