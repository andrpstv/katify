package config

type EnvVariable string

const (
	AMOCRM_URL         EnvVariable = "AMOCRM_URL"
	AMOCRM_BASEURL     EnvVariable = "AMOCRM_BASEURL"
	AMOCRM_ACCOUNTS    EnvVariable = "AMOCRM_ACCOUNTS"
	AMOCRM_CALLSTATS   EnvVariable = "AMOCRM_CALLSTATS"
	AMO_LOGIN          EnvVariable = "AMO_LOGIN"
	AMO_PASSWORD       EnvVariable = "AMO_PASSWORD"
	TELEGRAM_API_TOKEN EnvVariable = "TELEGRAM_API_TOKEN"
	DB_USER            EnvVariable = "DB_USER"
	DB_PASSWORD        EnvVariable = "DB_PASSWORD"
	DB_NAME            EnvVariable = "DB_NAME"
	DB_HOST            EnvVariable = "DB_HOST"
	DB_PORT            EnvVariable = "DB_PORT"
	HTTPPort           EnvVariable = "HTTP_PORT"
	APIVer             EnvVariable = "API_VER"
	LogLevel           EnvVariable = "LOG_LEVEL"
	APP_ENV            EnvVariable = "APP_ENV"
)
