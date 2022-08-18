package utils

type dataBaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type runConfig struct {
	Host     string
	Port     string
	LogLevel string
}

type config struct {
	RunConfig      runConfig
	DataBaseConfig dataBaseConfig
}
