package domain

import (
	"context"
	"fmt"

	amocrm "report/internal/ports/amocrm/auth"
)

type AuthService struct {
	cfg    *AuthConfig
	client amocrm.AuthPort
	token  string
}

func NewAuthService(client amocrm.AuthPort, cfg *AuthConfig) *AuthService {
	return &AuthService{
		client: client,
		cfg:    cfg,
	}
}

func (a *AuthService) Init(ctx context.Context, login, password string) error {
	csrf, cookies, err := a.client.GetCSRFtoken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get CSRF token: %w", err)
	}

	token, err := a.client.Login(ctx, login, password, csrf, cookies)
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}

	a.token = token
	return nil
}

func (a *AuthService) GetToken() string {
	return a.token
}
