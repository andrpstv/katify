package dto

import (
	"net/http"
	"time"
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

// func (a *AuthData) ToDomain() domain.Account {
// 	if a == nil {
// 		return domain.Account{}
// 	}
// 	return domain.Account{
// 		ID:           a.ID,
// 		Email:        a.Email,
// 		Name:         a.Name,
// 		AccessToken:  a.AccessToken,
// 		RefreshToken: a.RefreshToken,
// 		ExpiresAt:    a.ExpiresAt,
// 		Cookies:      a.Cookies, // если domain.Account.Cookies имеет тот же тип
// 	}
// }

type AuthRequest struct {
	Email        string `json:email`
	Username     string `json:"username"`
	Password     string `json:"password"`
	UserName     string `json:username`
	PasswordHash string `json:password_hash`
	FullName     string `json:fullname`
	// CSRFToken string `json:"csrf_token"`
}
