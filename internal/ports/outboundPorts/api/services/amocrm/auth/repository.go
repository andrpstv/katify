package ports

import (
	"context"
	"time"

	"github.com/google/uuid"

	"report/internal/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, a domain.AccountData) (domain.AccountData, error)
	GetUserByAmoID(ctx context.Context, amoID string) (domain.AccountData, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (domain.AccountData, error)
	UpdateUserTokens(
		ctx context.Context,
		id uuid.UUID,
		accessToken, refreshToken string,
		expiresAt time.Time,
	) error
}
