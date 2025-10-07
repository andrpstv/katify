package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type LoggerConfig struct {
	LogLevel string
}

type Logger interface {
	Infof(string, ...any)
	Debugf(string, ...any)
	Errorf(string, error, ...any)
	Fatalf(string, error, ...any)
}

type logger struct {
	log zerolog.Logger
}

func NewLogger(cfg LoggerConfig) Logger {
	time.Local = time.UTC
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	lvl := loggerLevel(cfg.LogLevel)
	zerolog.SetGlobalLevel(lvl)

	log.Info().Msgf("Logger configured is successful on level: %s", lvl.String())

	return &logger{log: log}
}

func loggerLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

func (l *logger) Infof(msg string, args ...any) {
	l.log.Info().Msgf(msg, args...)
}

func (l *logger) Debugf(msg string, args ...any) {
	l.log.Debug().Msgf(msg, args...)
}

func (l *logger) Errorf(msg string, err error, args ...any) {
	l.log.Error().Err(err).Msgf(msg, args...)
}

func (l *logger) Fatalf(msg string, err error, args ...any) {
	l.log.Fatal().Err(err).Msgf(msg, args...)
}
