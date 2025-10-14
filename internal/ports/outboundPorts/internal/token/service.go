package ports

import (
	domain "report/internal/domain/user"
)

type TokenService interface {
	Generate(userID string) (*domain.TokenPair, error)
	Validate(accessToken string) (userID string, err error)
}
