package api

import "time"

type TelegramConfig struct {
	Timeout time.Duration
	API     string
}
