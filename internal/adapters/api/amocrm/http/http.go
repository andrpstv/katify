package adapters

import (
	"bytes"
	"context"
	"net/http"
)

type HTTPClient struct {
	Client *http.Client
}

func (h *HTTPClient) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return h.Client.Do(req)
}

func (h *HTTPClient) Post(
	ctx context.Context,
	url string,
	body []byte,
	headers map[string]string,
) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return h.Client.Do(req)
}
