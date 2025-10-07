package ports

import (
	"context"
	"net/http"
)

// for all services requesting
type HTTPClient interface {
	Get(ctx context.Context, url string, headers map[string]string) (*http.Response, error)
	Post(
		ctx context.Context,
		url string,
		body []byte,
		headers map[string]string,
	) (*http.Response, error)
}
