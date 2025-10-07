package dto

import (
	"net/http"

	domain "report/internal/domain/accounts"
)

type AuthData struct {
	Token   string
	Cookies []*http.Cookie
}

func (a *AuthData) toModel(data *AuthData) domain.Account {
	accountInfo := domain.Account{}
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// CSRFToken string `json:"csrf_token"`
}
