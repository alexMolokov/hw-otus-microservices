package zap

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// EncoderContextWrapper обертка для отделения названий полей.
type EncoderContextWrapper struct {
	zapcore.Encoder
	prefix string
}

// NewEncoderContextWrapper конструктор для EncoderContextWrapper.
func NewEncoderContextWrapper(enc zapcore.Encoder, prefix string) *EncoderContextWrapper {
	return &EncoderContextWrapper{
		Encoder: enc,
		prefix:  prefix,
	}
}

// Clone клонирование кодировкщика.
func (wrap *EncoderContextWrapper) Clone() zapcore.Encoder {
	return NewEncoderContextWrapper(wrap.Encoder.Clone(), wrap.prefix)
}

// EncodeEntry кодирование записи.
func (wrap *EncoderContextWrapper) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	total := len(fields)
	if total == 0 {
		return wrap.Encoder.EncodeEntry(entry, fields)
	}
	newFields := make([]zapcore.Field, total)
	for i, v := range fields {
		field := v
		field.Key = wrap.prefix + field.Key
		newFields[i] = field
	}
	return wrap.Encoder.EncodeEntry(entry, newFields)
}

// NewProductionEncoderConfig настройки конфигурации для Production.
func NewProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "facility",
		CallerKey:      "caller",
		FunctionKey:    "function",
		MessageKey:     "short_message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
