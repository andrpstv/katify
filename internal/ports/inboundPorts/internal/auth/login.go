package ports

import (
	"context"

	domain "katify/internal/domain/user"
	dto "katify/internal/dto/auth"
)

type AuthUseCase interface {
	Login(ctx context.Context, data *dto.AuthRequest) (*domain.User, error)
	Register(ctx context.Context, data *dto.AuthRequest) (*domain.User, error)
}
