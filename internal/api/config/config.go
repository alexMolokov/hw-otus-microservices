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

type DBConf struct {
	Host         string `config:"db-host"`
	Port         int    `config:"db-port"`
	Name         string `config:"db-name"`
	User         string `config:"db-user"`
	Password     string `config:"db-password"`
	SslMode      string `config:"db-ssl"`
	BinaryParams string `config:"db-binary-params"`
	MaxOpenConn  int    `config:"db-max-open-conn"`
	MaxIdleConn  int    `config:"db-max-idle-conn"`
	MaxLifetime  int    `config:"db-max-life-time"`
}
