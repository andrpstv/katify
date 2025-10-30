package adapters

import (
	"context"
	"fmt"
	"net/http"

	domain "katify/internal/domain/accounts"
	portsClient "katify/internal/ports/outboundPorts/api/client"
	portsParser "katify/internal/ports/outboundPorts/api/parser"
)

type AmoAccountClientImpl struct {
	client portsClient.HTTPClient
	cfg    *AccountsConfig
	parser portsParser.AccountParserService
}

func NewAmocrmAccountsClient(
	http portsClient.HTTPClient,
	cfg *AccountsConfig,
	parser portsParser.AccountParserService,
) *AmoAccountClientImpl {
	return &AmoAccountClientImpl{
		client: http,
		cfg:    cfg,
		parser: parser,
	}
}

func (c *AmoAccountClientImpl) FetchAccounts(
	ctx context.Context,
	token string,
) (*domain.Account, error) {
	headers := map[string]string{
		"Authorization": "Bearer " + token,
	}

	resp, err := c.client.Get(ctx, c.cfg.AccountsURL, headers)
	if err != nil {
		return &domain.Account{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &domain.Account{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return c.parser.ParseAccounts(resp)
}
