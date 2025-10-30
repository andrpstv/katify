package ports

import (
	"context"

	domain "katify/internal/domain/accounts"
)

type AmoAccountClient interface {
	FetchAccounts(
		ctx context.Context,
		token string,
	) (*domain.Account, error)
}
