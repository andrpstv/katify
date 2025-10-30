package ports

import (
	domain "katify/internal/domain/user"
)

type AuthService interface {
	GenerateToken(userID string) (*domain.UserData, error)
	ValidateToken(accessToken string) (userID string, err error)
	ValidatePassword(password string) (bool, error)
}
