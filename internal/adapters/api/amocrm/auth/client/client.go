package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	dto "katify/internal/dto/auth"
	portsClient "katify/internal/ports/outboundPorts/api/client"
	portsMapper "katify/internal/ports/outboundPorts/api/mapper"
	portsParser "katify/internal/ports/outboundPorts/api/parser"
)

type AmoAuthClientImpl struct {
	Cfg    *AmoAuthConfig
	client portsClient.HTTPClient
	parser portsParser.AuthParserService
	mapper portsMapper.AuthMapperService
}

func NewAmocrmAuthClientImpl(
	cfg *AmoAuthConfig,
	httpClient portsClient.HTTPClient,
	parser portsParser.AuthParserService,
	mapper portsMapper.AuthMapperService,
) *AmoAuthClientImpl {
	return &AmoAuthClientImpl{
		Cfg:    cfg,
		client: httpClient,
		parser: parser,
	}
}

func (a *AmoAuthClientImpl) GetCSRFtoken(
	ctx context.Context,
) (string, error) {
	resp, err := a.client.Get(ctx, a.Cfg.BaseURL, nil)
	if err != nil {
		return "", fmt.Errorf("error sending request to %s: %w", a.Cfg.BaseURL, err)
	}
	defer resp.Body.Close()

	return a.parser.ParseCSRF(resp)
}

func (a *AmoAuthClientImpl) Login(
	ctx context.Context,
	data *dto.AuthRequest,
) (*http.Response, error) {
	bodyBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling AuthUseCase request: %w", err)
	}

	resp, err := a.client.Post(
		ctx,
		a.Cfg.LoginURL,
		bodyBytes,
		map[string]string{"Content-Type": "application/json"},
	)
	if err != nil {
		return nil, fmt.Errorf("error sending login request: %w", err)
	}
	defer resp.Body.Close()

	return resp, nil
}
