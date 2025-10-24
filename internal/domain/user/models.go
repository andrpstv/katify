package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string
	UserName     string
	Email        string
	PasswordHash string
	FullName     *string
	MfaEnabled   *bool
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (u *User) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) VerifyPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

type UserCredentials struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

func (t *UserCredentials) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}
