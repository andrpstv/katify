package user

import (
	"context"

	domain "katify/internal/domain/user"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByUserID(ctx context.Context, userID string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (string, error)
	UpdateUser(ctx context.Context, user *domain.User) error

	GetTokensByUserId(ctx context.Context, userId string) (*domain.UserCredentials, error)
	CreateTokensByUserId(ctx context.Context, userId string, userCreds *domain.UserCredentials) (string, error)
	UpdateTokensByUserId(ctx context.Context, userId string, userCreds *domain.UserCredentials) error
}
