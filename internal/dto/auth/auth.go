package dto

import (
	"net/http"
	"time"

	domain "report/internal/domain/accounts"
)

type AuthData struct {
	ID           string
	Email        string
	Name         string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	Cookies      []*http.Cookie
}

func (a *AuthData) toModel(data *AuthData) domain.Account {
	accountInfo := domain.Account{}
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// CSRFToken string `json:"csrf_token"`
}
