package ports

import (
	"context"

	domain "report/internal/domain/auth"
)

type AmoAccountClient interface {
	FetchProjects(
		ctx context.Context,
		token string,
	) (*domain.AccountInfo, error)
}
