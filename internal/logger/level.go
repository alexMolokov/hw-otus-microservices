package logger

import (
	"bytes"
	"fmt"
)

// Level уровень логирования.
type Level uint8

const (
	// DebugLevel debug.
	DebugLevel Level = iota
	// InfoLevel info.
	InfoLevel
	// NoticeLevel notice.
	NoticeLevel
	// WarnLevel warning.
	WarnLevel
	// ErrorLevel errors.
	ErrorLevel
	// CriticalLevel panic, then calls panic().
	CriticalLevel
	// FatalLevel fatal, then calls os.Exit(1).
	FatalLevel
)

// UnmarshalNilLevelError ошибки декодирования для нулевого указателя.
type UnmarshalNilLevelError struct{}

// Error возвращает текст ошибки.
func (err *UnmarshalNilLevelError) Error() string {
	return "can't unmarshal a nil *Level"
}

// NewUnmarshalNilLevelError констуктор UnmarshalNilLevelError.
func NewUnmarshalNilLevelError() *UnmarshalNilLevelError {
	return &UnmarshalNilLevelError{}
}

// UnmarshalTextLevelError ошибка при декодировании текста.
type UnmarshalTextLevelError struct {
	text string
}

// Error возвращает текст ошибки.
func (err *UnmarshalTextLevelError) Error() string {
	return fmt.Sprintf("unrecognized level: %s", err.text)
}

// Text возвращает текст, который стал причиной ошибки.
func (err *UnmarshalTextLevelError) Text() string {
	return err.text
}

// NewUnmarshalTextLevelError констуктор UnmarshalTextLevelError.
func NewUnmarshalTextLevelError(text string) *UnmarshalTextLevelError {
	return &UnmarshalTextLevelError{text}
}

// String возвращает человекочитаемое представление уровня в ASCII.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case NoticeLevel:
		return "notice"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case CriticalLevel:
		return "critical"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("level_%d", l)
	}
}

// MarshalText преобразование в текст.
func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

// UnmarshalText конвертация из текста.
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return NewUnmarshalNilLevelError()
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return NewUnmarshalTextLevelError(string(text))
	}
	return nil
}

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO", "":
		*l = InfoLevel
	case "notice", "NOTICE":
		*l = NoticeLevel
	case "warn", "WARN", "warning", "WARNING":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "crit", "CRIT", "critical", "CRITICAL":
		*l = CriticalLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}
