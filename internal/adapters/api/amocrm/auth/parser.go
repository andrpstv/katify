package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	dto "report/internal/dto/auth"
)

type AuthParserServiceImpl struct{}

func (p *AuthParserServiceImpl) ParseCSRF(html *http.Response) (string, error) {
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

func (p *AuthParserServiceImpl) DecodeAuthData(
	ctx context.Context,
	resp *http.Response,
) (*dto.AuthData, error) {
	authResp := &dto.AuthData{}
	if err := json.NewDecoder(resp.Body).Decode(authResp); err != nil {
		return nil, fmt.Errorf("error decoding login response: %w", err)
	}
	authResp.Cookies = resp.Cookies()
	return authResp, nil
}
