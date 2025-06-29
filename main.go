package main

import (
	"context"
	"fmt"
	"log"
	"os"
	auth "report/internal/authAmoCrm"
	"time"
)

func main() {
	baseURL := os.Getenv("AMOCRM")
	loginURL := os.Getenv("AMOCRMURL")

	username := os.Getenv("AMO_LOGIN")
	password := os.Getenv("AMO_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("Не заданы AMO_LOGIN или AMO_PASSWORD")
	}

	cfg := &auth.AmocrmConfig{
		Timeout:  15 * time.Second,
		BaseURL:  baseURL,
		LoginUrl: loginURL,
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

	fmt.Println("CSRF Token:", csrfToken)
	fmt.Println("Cookies:", cookies)

	token, newCookies, err := client.Login(ctx, username, password, csrfToken, cookies)
	if err != nil {
		log.Fatalf("ошибка Login: %v", err)
	}

	fmt.Println("Access Token:", token)
	fmt.Println("New Cookies:", newCookies)
}
