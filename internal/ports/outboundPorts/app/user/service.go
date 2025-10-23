package user

import (
	domain "report/internal/domain/user"
)

type UserService interface {
	GenerateTokens(userID string) (*domain.UserCredentials, error)
}
