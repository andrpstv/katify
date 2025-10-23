package domain

import (
	"errors"
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

func (u *User) VerifyPassword(dtoPassword string) (bool, error) {
	if dtoPassword == "" || dtoPassword != u.PasswordHash {
		return false, errors.New("invalid password, try again")
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
