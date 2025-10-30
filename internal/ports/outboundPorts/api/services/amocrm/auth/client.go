package ports

import (
	"context"
	"net/http"

	dto "katify/internal/dto/auth"
)

type AuthClient interface {
	GetCSRFtoken(
		ctx context.Context,
	) (string, error)
	Login(
		ctx context.Context,
		data *dto.AuthRequest,
	) (*http.Response, error)
}
