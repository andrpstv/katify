package adapters

import "time"

type AmoAuthConfig struct {
	Timeout     time.Duration
	BaseURL     string
	LoginURL    string
	AccountsURL string
}
