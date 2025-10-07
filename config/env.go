package config

import (
	"os"
)

func GetEnv(e EnvVariable) string {
	return os.Getenv(string(e))
}
