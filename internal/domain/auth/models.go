package domain

import "net/http"

type AccountData struct {
	Token   string
	Cookies []*http.Cookie
}
