package config

// В параметре config указывается имя переменной окружения
// Например если указано  config:"api-logger-level, то имя переменной окружения будет API_LOGGER_LEVEL
// таким образом возможно переопределить значение взятое из файла ../configs/api.json

type LoggerConf struct {
	Level    string `config:"api-logger-level"`
	Facility string `config:"api-logger-facility"`
}

type HTTPConf struct {
	Host string `config:"api-http-host"`
	Port int    `config:"api-http-port"`
}
