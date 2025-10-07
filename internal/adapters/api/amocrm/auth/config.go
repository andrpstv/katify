package adapters

import "time"

type AuthConfig struct {
	Timeout     time.Duration
	BaseURL     string
	LoginURL    string
	AccountsURL string
}
