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
	DEV_DB_USER        EnvVariable = "DEV_DB_USER"
	DEV_DB_PASSWORD    EnvVariable = "DEV_DB_PASSWORD"
	DEV_DB_NAME        EnvVariable = "DEV_DB_NAME"
	DEV_DB_HOST        EnvVariable = "DEV_DB_HOST"
	DEV_DB_PORT        EnvVariable = "DEV_DB_PORT"
	PROD_DB_USER       EnvVariable = "PROD_DB_USER"
	PROD_DB_PASSWORD   EnvVariable = "PROD_DB_PASSWORD"
	PROD_DB_NAME       EnvVariable = "PROD_DB_NAME"
	PROD_DB_HOST       EnvVariable = "PROD_DB_HOST"
	PROD_DB_PORT       EnvVariable = "PROD_DB_PORT"
	HTTPPort           EnvVariable = "HTTP_PORT"
	APIVer             EnvVariable = "API_VER"
	LogLevel           EnvVariable = "LOG_LEVEL"
	APP_ENV            EnvVariable = "APP_ENV"
)
