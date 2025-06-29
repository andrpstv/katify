package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"report/internal/models"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type AmocrmClientMeth interface {
	GetCSRFtoken(ctx context.Context) (CSRFtoken string, cookies []*http.Cookie, err error)
	Login(ctx context.Context, login, password, CSRFtoken string, cookies []*http.Cookie) (token string, err error)
}

type AmocrmClient struct {
	Cfg        *AmocrmConfig
	HttpClient *http.Client
	CsrfToken  string
	Cookies    []*http.Cookie
}

type AmocrmConfig struct {
	Timeout     time.Duration
	BaseURL     string
	LoginURL    string
	AccountsURL string
}

func NewAmocrmClient(cfg *AmocrmConfig) (*AmocrmClient, error) {
	return &AmocrmClient{
		Cfg: cfg,
		HttpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
	}, nil
}

func (a *AmocrmClient) GetCSRFtoken(ctx context.Context) (CSRFToken string, cookies []*http.Cookie, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.Cfg.BaseURL, nil)
	if err != nil {
		return "", nil, fmt.Errorf("error creating request to %s: %w", a.Cfg.BaseURL, err)
	}
	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("error sending request to %s: %w", a.Cfg.BaseURL, err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing HTML response from %s: %w", a.Cfg.BaseURL, err)
	}

	val, exists := doc.Find(`input[name="csrf_token"]`).Attr("value")
	if !exists {
		return "", nil, fmt.Errorf("csrf_token input not found in HTML response from %s", a.Cfg.BaseURL)
	}

	a.CsrfToken = val
	a.Cookies = resp.Cookies()

	return val, resp.Cookies(), nil
}

func (a *AmocrmClient) Login(ctx context.Context, username, password, csrfToken string, cookies []*http.Cookie) (string, []*http.Cookie, error) {
	reqBody := models.AuthRequest{
		Username:  username,
		Password:  password,
		CSRFToken: csrfToken,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", nil, fmt.Errorf("error marshaling auth request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.Cfg.LoginURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", nil, fmt.Errorf("error creating request to %s: %w", a.Cfg.LoginURL, err)
	}

	req.Header.Set("Content-Type", "application/json")

	for _, c := range cookies {
		req.AddCookie(c)
	}

	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("error sending login request: %w", err)
	}
	defer resp.Body.Close()

	var authResp models.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", nil, fmt.Errorf("error decoding login response: %w", err)
	}

	fmt.Printf("Логин прошел успешно, твой acces token: %s, твои куки: %s", authResp.AccessToken, authResp.Cookies)
	return authResp.AccessToken, resp.Cookies(), nil
}
