package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Telegram *api.TelegramConfig
	AmoCrm   *auth.AmocrmConfig
}

func Configload() *Config {
	preloadEnv()

	cfg := &Config{}
	cfg.Telegram = readTelegramConfig()
	cfg.AmoCrm = readAmoCrmConfig()

	return cfg
}

func preloadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)

func readTelegramConfig() *api.TelegramConfig {
	telegram := &api.TelegramConfig{
		TelegramAPItoken: GetEnv(TELEGRAM_API_TOKEN),
	}
	return telegram
}

func readAmoCrmConfig() *auth.AmocrmConfig {
	amocrm := &auth.AmocrmConfig{
		Timeout:     20,
		BaseURL:     GetEnv(AMOCRM_BASEURL),
		LoginURL:    GetEnv(AMO_LOGIN),
		AccountsURL: GetEnv(AMOCRM_ACCOUNTS),
	}
	return amocrm
}
