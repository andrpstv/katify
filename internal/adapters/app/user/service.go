package user

import (
	domain "report/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

var jwtSecret = []byte("super_secret_key") // обычно берётся из .env

func (a *UserServiceImpl) GenerateTokens(userID string) (*domain.UserCredentials, error) {
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refresh.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	userCred := &domain.UserCredentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(15 * time.Minute),
	}

	return userCred, nil
}
