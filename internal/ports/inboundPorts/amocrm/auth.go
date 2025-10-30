package ports

import (
	"context"

	domain "katify/internal/domain/AuthUseCase"
	dto "katify/internal/dto/auth"
)

type AuthUseCase interface {
	Login(
		ctx context.Context,
		data *dto.AuthRequest,
	) (*domain.AccountData, error)
}

// type TokenProvider interface {
// 	ValidateToken(ctx, accountID string) (domain.Token, error)
// }
