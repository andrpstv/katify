package ports

import (
	"context"

	domain "report/internal/domain/user"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
}
