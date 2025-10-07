package logger

type LoggerConfig struct {
	LogLevel string
}

type Logger interface {
	Infof(string)
	Debugf(string)
	Errorf(string, error)
	Fatalf(string, error)
}
