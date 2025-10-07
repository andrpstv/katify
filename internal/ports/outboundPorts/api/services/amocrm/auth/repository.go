package ports

import (
	"context"

	domain "report/internal/domain/auth"
)

type AmoAuthRepository interface {
	Save(ctx context.Context, account *domain.AccountData) error
	GetByID(ctx context.Context, id int) (*domain.AccountData, error)
}
