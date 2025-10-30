package user

import (
	domain "katify/internal/domain/user"
)

type UserService interface {
	GenerateTokens(userID string) (*domain.UserCredentials, error)
}
