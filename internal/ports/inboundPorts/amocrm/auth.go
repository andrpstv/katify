package ports

import (
	"context"

	domain "report/internal/domain/AuthUseCase"
	dto "report/internal/dto/auth"
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
