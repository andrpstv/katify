package ports

import (
	"context"

	domain "report/internal/domain/user"
	dto "report/internal/dto/auth"
)

type AuthUseCase interface {
	Login(ctx context.Context, data *dto.AuthRequest) (*domain.TokenPair, error)
	Register(ctx context.Context, data *dto.AuthRequest) (*domain.TokenPair, error)
}
