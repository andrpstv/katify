package config

import (
	"report/internal/api"
	"report/internal/auth"
	"report/internal/server"
)

type Config struct {
	Server   *server.ServerConfig
	Telegram *api.TelegramConfig
	AmoCrm *auth.AmoCrmConfig
}
