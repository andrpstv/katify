package domain

import "time"

type Account struct {
	ID       string
	Name     string
	Email    string
	Projects []Project
	Data     AccountData
}

type Project struct {
	ID         int
	UUID       string
	Name       string
	Subdomain  string
	ShardType  int
	IsAdmin    bool
	Version    int
	IsKommo    bool
	Domain     string
	IsTrial    bool
	TrialEnded bool
	IsPayed    bool
	PayedEnded bool
	MFAEnabled bool
	ge         string
}

type AccountData struct {
	ProviderUserID string
	Email          string
	Name           string
	Login          string
	Password       string
	AccessToken    string
	RefreshToken   string
	ExpiresAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
