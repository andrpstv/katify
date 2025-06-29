package models

import "net/http"

type AuthRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CSRFToken string `json:"csrf_token"`
}

type AuthResponse struct {
	AccessToken string         `json:"access_token"`
	Cookies     []*http.Cookie `json:"cookies"`
}

type AccountInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}
