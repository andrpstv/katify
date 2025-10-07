package loggerPort

type Logger interface {
	Infof(string)
	Debugf(string)
	Errorf(string, error)
	Fatalf(string, error)
}
