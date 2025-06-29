package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type AmocrmClient interface {
	GetCSRFtoken(ctx context.Context) (CSRFtoken string, cookies []*http.Cookie, err error)
	Login(ctx context.Context, login, password, CSRFtoken string, cookies []*http.Cookie) (token string, err error)
}

type amocrmClient struct {
	cfg        *AmocrmConfig
	httpClient *http.Client
	csrfToken  string
	cookies    []*http.Cookie
}

type AmocrmConfig struct {
	Timeout  time.Duration
	BaseURL  string
	LoginUrl string
}

func NewAmocrmClient(cfg *AmocrmConfig) (*amocrmClient, error) {
	return &amocrmClient{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}, nil
}

type AuthRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CSRFToken string `json:"csrf_token"`
}

type AuthResponse struct {
	AccessToken string         `json:"access_token"`
	Cookies     []*http.Cookie `json:"cookies"`
}

func (a *amocrmClient) GetCSRFtoken(ctx context.Context) (CSRFToken string, cookies []*http.Cookie, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.cfg.BaseURL, nil)
	if err != nil {
		return "", nil, fmt.Errorf("error creating request to %s: %w", a.cfg.BaseURL, err)
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("error sending request to %s: %w", a.cfg.BaseURL, err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing HTML response from %s: %w", a.cfg.BaseURL, err)
	}

	val, exists := doc.Find(`input[name="csrf_token"]`).Attr("value")
	if !exists {
		return "", nil, fmt.Errorf("csrf_token input not found in HTML response from %s", a.cfg.BaseURL)
	}

	a.csrfToken = val
	a.cookies = resp.Cookies()

	return val, resp.Cookies(), nil
}

func (a *amocrmClient) Login(ctx context.Context, username, password, csrfToken string, cookies []*http.Cookie) (string, []*http.Cookie, error) {
	reqBody := AuthRequest{
		Username:  username,
		Password:  password,
		CSRFToken: csrfToken,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", nil, fmt.Errorf("error marshaling auth request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.cfg.LoginUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", nil, fmt.Errorf("error creating request to %s: %w", a.cfg.LoginUrl, err)
	}

	req.Header.Set("Content-Type", "application/json")

	for _, c := range cookies {
		req.AddCookie(c)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("error sending login request: %w", err)
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", nil, fmt.Errorf("error decoding login response: %w", err)
	}

	fmt.Printf("Логин прошел успешно, твой acces token: %s, твои куки: %s", authResp.AccessToken, authResp.Cookies)
	return authResp.AccessToken, resp.Cookies(), nil
}
