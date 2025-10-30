package ports

import (
	"context"

	domain "katify/internal/domain/user"
)

type AuthRepository interface {
	GetTokensByUserID(context context.Context, userID string) (*domain.UserData, error)
	SaveTokens(ctx context.Context, userID string, tokens *domain.UserData) error
}
