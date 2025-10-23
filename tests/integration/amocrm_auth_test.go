//go:build integration
// +build integration

package integration_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"report/internal/adapters/amocrm/auth"
	"report/internal/adapters/amocrm/client"
	"report/internal/adapters/amocrm/services/accounts"
)

func TestAmocrmLoginAndGetAccounts(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("failed to load .env file: %v", err)
	}

	baseURL := os.Getenv("AMOCRM_BASEURL")
	loginURL := os.Getenv("AMOCRM_URL")
	username := os.Getenv("AMO_LOGIN")
	password := os.Getenv("AMO_PASSWORD")
	accountsURL := os.Getenv("AMOCRM_ACCOUNTS")

	if baseURL == "" || loginURL == "" || username == "" || password == "" || accountsURL == "" {
		t.Skip("AMO credentials or URLs not set in environment, skipping integration test")
	}

	authCfg := &auth.AuthConfig{
		BaseURL:  baseURL,
		LoginURL: loginURL,
		Timeout:  60 * time.Second,
	}

	httpClient := &client.HTTPClientAdapter{
		Client: &http.Client{
			Timeout: authCfg.Timeout,
		},
	}

	authClient := auth.NewAmocrmClient(authCfg, httpClient)

	authService := auth.NewAuthService(authClient, authCfg)

	ctx := context.Background()

	if err := authService.Init(ctx, username, password); err != nil {
		t.Fatalf("failed to init AuthUseCase service: %v", err)
	}
	t.Logf("Login successful, token: %s...", authService.GetToken()[:20]) // только начало токена

	accountsCfg := &accounts.AccountsConfig{
		AccountsURL: accountsURL,
	}
	accountsClient := accounts.NewAccountsClient(httpClient, accountsCfg)

	accountsService := accounts.NewAccountsService(authService, accountsClient)

	accountInfo, err := accountsService.GetAccountsList(ctx)
	if err != nil {
		t.Fatalf("failed to get accounts list: %v", err)
	}

	t.Logf("User has access to %d accounts:", len(accountInfo.Titles))
	for _, title := range accountInfo.Titles {
		t.Logf(" - %s", title)
	}
}
