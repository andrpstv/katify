package adapters

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"

	dto "katify/internal/dto/auth"
)

type AmoAuthParserServiceImpl struct{}

func (p *AmoAuthParserServiceImpl) ParseCSRF(html *http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing HTML response from %s: %w", html.Body, err)
	}

	token, exists := doc.Find(`input[name="csrf_token"]`).Attr("value")
	if !exists {
		return "", fmt.Errorf(
			"csrf_token input not found in HTML response from %s",
			html.Body,
		)
	}

	return token, nil
}

func (p *AmoAuthParserServiceImpl) DecodeAuthData(
	ctx context.Context,
	resp *http.Response,
) (*dto.AuthData, error) {
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	authResp := &dto.AuthData{
		Cookies: resp.Cookies(),
	}

	for _, cookie := range authResp.Cookies {
		switch cookie.Name {
		case "access_token":
			authResp.AccessToken = cookie.Value
		case "refresh_token":
			authResp.RefreshToken = cookie.Value
		case "access_token_expires_at":
			ts, err := strconv.ParseInt(cookie.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid access_token_expires_at value: %v", err)
			}
			authResp.ExpiresAt = time.Unix(ts, 0).UTC()
		case "amo_user_id":
			authResp.ID = cookie.Value
		case "amo_user_full_name":
			if v, err := url.QueryUnescape(cookie.Value); err == nil {
				authResp.Name = v
			} else {
				authResp.Name = cookie.Value
			}
		case "amo_user_email":
			authResp.Email = cookie.Value
		}
	}

	if authResp.AccessToken == "" {
		return nil, fmt.Errorf("AuthUseCase failed: no access_token found in cookies")
	}
	if authResp.ID == "" {
		return nil, fmt.Errorf("AuthUseCase failed: amo_user_id not found in cookies")
	}
	if authResp.Email == "" {
		return nil, fmt.Errorf("AuthUseCase warning: amo_user_email not found in cookies")
	}

	return authResp, nil
}
