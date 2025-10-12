package ports

import (
	"context"

	domain "report/internal/domain/accounts"
)

type AmoAccountClient interface {
	FetchAccounts(
		ctx context.Context,
		token string,
	) (*domain.Account, error)
}
