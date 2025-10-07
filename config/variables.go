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
)
