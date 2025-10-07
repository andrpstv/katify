package domain_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"report/internal/adapters/amocrm/auth"
	"report/internal/adapters/amocrm/client"
)

func TestGetCSRFtoken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := `<html><body><input type="hidden" name="csrf_token" value="fake-token"></body></html>`
		_, _ = w.Write([]byte(html))
	}))
	defer ts.Close()

	cfg := &auth.AuthConfig{
		BaseURL: ts.URL,
		Timeout: time.Second,
	}
	httpClient := &client.HTTPClientAdapter{
		Client: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
	client := auth.NewAmocrmClient(cfg, httpClient)

	token, cookies, err := client.GetCSRFtoken(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token != "fake-token" {
		t.Errorf("expected fake-token, got %s", token)
	}
	if cookies == nil {
		t.Errorf("expected cookies, got nil")
	}
}

func TestLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req auth.AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("decode failed: %v", err)
		}

		if req.Username != "admin" || req.Password != "secret" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{Name: "sid", Value: "123"})

		resp := auth.AuthData{
			AccessToken: "fake-token",
			Cookies:     []*http.Cookie{{Name: "sid", Value: "123"}},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	cfg := &auth.AuthConfig{
		LoginURL: ts.URL,
		Timeout:  time.Second,
	}
	httpClient := &client.HTTPClientAdapter{
		Client: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
	client := auth.NewAmocrmClient(cfg, httpClient)

	token, err := client.Login(
		context.Background(),
		"admin",
		"secret",
		"csrf-token",
		nil,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token != "fake-token" {
		t.Errorf("expected fake-token, got %s", token)
	}
}
