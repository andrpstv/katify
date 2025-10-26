package config

import (
	"log"
	postgres "report/internal/adapters/db"
	"report/pkg/logger"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   *ServerConfig
	Logger   *logger.LoggerConfig
	Postgres *postgres.PostgresConfig
}

func ReadConfig() (Config, error) {
	preloadEnv()

	cfg := Config{}

	cfg.Server = readServerConfig()
	cfg.Logger = readLoggerConfig()
	cfg.Postgres = readPostgresConfig()

	return cfg, nil
}

func preloadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}
}

type ServerConfig struct {
	HTTPPort     string
	APIVer       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func readServerConfig() *ServerConfig {
	return &ServerConfig{
		HTTPPort: GetOrDefaultEnv(HTTPPort, ":8080"),
		APIVer:   GetOrDefaultEnv(APIVer, "/api/v1"),
	}
}

func readLoggerConfig() *logger.LoggerConfig {
	loggerConfig := &logger.LoggerConfig{}
	loggerConfig.LogLevel = GetOrDefaultEnv(LogLevel, "info")
	return loggerConfig
}

func readPostgresConfig() *postgres.PostgresConfig {
	appEnv := GetOrDefaultEnv(APP_ENV, "dev")

	if appEnv == "prod" {
		return &postgres.PostgresConfig{
			Host:     GetOrDefaultEnv(PROD_DB_HOST, "localhost"),
			Port:     GetOrDefaultEnv(PROD_DB_PORT, "5432"),
			User:     GetOrDefaultEnv(PROD_DB_USER, "admin"),
			Password: GetOrDefaultEnv(PROD_DB_PASSWORD, "pwd"),
			Database: GetOrDefaultEnv(PROD_DB_NAME, "db"),
		}
	}

	return &postgres.PostgresConfig{
		Host:     GetOrDefaultEnv(DEV_DB_HOST, "localhost"),
		Port:     GetOrDefaultEnv(DEV_DB_PORT, "5432"),
		User:     GetOrDefaultEnv(DEV_DB_USER, "admin"),
		Password: GetOrDefaultEnv(DEV_DB_PASSWORD, "pwd"),
		Database: GetOrDefaultEnv(DEV_DB_NAME, "db"),
	}

}
