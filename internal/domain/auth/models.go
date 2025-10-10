package domain

import "time"

type AccountData struct {
	AmoUserID    string
	Email        string
	Name         string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
