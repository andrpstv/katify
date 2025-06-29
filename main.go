package main

import (
	"context"
	"fmt"
	"log"
	"os"
	auth "report/internal/authAmoCrm"
	amocrm "report/internal/report_service/amocrm_service"
	"time"
)

func main() {
	baseURL := os.Getenv("AMOCRM")
	loginURL := os.Getenv("AMOCRMURL")
	accountsURL := os.Getenv("AMOCRM_ACCOUNTS_URL")

	username := os.Getenv("AMO_LOGIN")
	password := os.Getenv("AMO_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("Не заданы AMO_LOGIN или AMO_PASSWORD")
	}

	cfg := &auth.AmocrmConfig{
		Timeout:     15 * time.Second,
		BaseURL:     baseURL,
		LoginURL:    loginURL,
		AccountsURL: accountsURL,
	}

	client, err := auth.NewAmocrmClient(cfg)
	if err != nil {
		log.Fatalf("ошибка создания клиента: %v", err)
	}

	ctx := context.Background()

	csrfToken, cookies, err := client.GetCSRFtoken(ctx)
	if err != nil {
		log.Fatalf("ошибка GetCSRFtoken: %v", err)
	}

	fmt.Println("✅ CSRF Token:", csrfToken)
	fmt.Println("✅ Cookies:", cookies)

	token, newCookies, err := client.Login(ctx, username, password, csrfToken, cookies)
	if err != nil {
		log.Fatalf("ошибка Login: %v", err)
	}

	client.Cookies = newCookies
	client.CsrfToken = csrfToken

	fmt.Println("✅ Access Token:", token)
	fmt.Println("✅ New Cookies:", newCookies)

	service := amocrm.NewAmocrmService(client)

	accounts, err := service.GetAccountsList(ctx)
	if err != nil {
		log.Fatalf("ошибка GetAccountsList: %v", err)
	}

	fmt.Println("=== ✅ Accounts List ===")
	for _, acc := range accounts {
		fmt.Printf("ID: %d | Name: %s | Domain: %s\n", acc.ID, acc.Name, acc.Domain)
	}

	
}
