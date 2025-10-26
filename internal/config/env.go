package config

import (
	"os"
)

func GetEnv(e EnvVariable) string {
	return os.Getenv(string(e))
}
func GetOrDefaultEnv(e EnvVariable, defaultValue string) string {
	env := os.Getenv(string(e))
	if env == "" {
		return defaultValue
	}
	return env
}
