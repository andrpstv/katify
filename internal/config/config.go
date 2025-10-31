package config

import (
	postgres "katify/internal/adapters/db"
	"katify/pkg/logger"
	"log"
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
		HTTPPort:     GetOrDefaultEnv(HTTPPort, ":9000"),
		APIVer:       GetOrDefaultEnv(APIVer, "/api/v1"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func readLoggerConfig() *logger.LoggerConfig {
	loggerConfig := &logger.LoggerConfig{}
	loggerConfig.LogLevel = GetOrDefaultEnv(LogLevel, "debug")
	return loggerConfig
}

func readPostgresConfig() *postgres.PostgresConfig {

	return &postgres.PostgresConfig{
		Host:     GetOrDefaultEnv(DB_HOST, "64.188.97.71"),
		Port:     GetOrDefaultEnv(DB_PORT, "5432"),
		User:     GetOrDefaultEnv(DB_USER, "dev_user"),
		Password: GetOrDefaultEnv(DB_PASSWORD, "dev_pass"),
		Database: GetOrDefaultEnv(DB_NAME, "katify_dev"),
	}
}
