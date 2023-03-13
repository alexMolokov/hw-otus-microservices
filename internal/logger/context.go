package logger

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

// Context контекст логирования.
type Context map[string]interface{}
