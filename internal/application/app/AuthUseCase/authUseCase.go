package authUseCase

import (
	"context"
	"fmt"
	"report/internal/ports/outboundPorts/app/user"

	domain "report/internal/domain/user"
	dto "report/internal/dto/auth"
)

type AuthUseCaseImpl struct {
	userRepo    user.UserRepository
	userService user.UserService
}

func (a *AuthUseCaseImpl) Login(
	ctx context.Context,
	data *dto.AuthRequest,
) (*domain.UserCredentials, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found, please register")
	}
	_, err = user.VerifyPassword(data.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	tokens, err := a.userRepo.GetTokensByUserId(ctx, user.ID)
	if err != nil || tokens.IsExpired() {
		newUserCreds, err := a.userService.GenerateTokens(user.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to generate new tokens")
		}

		_, err = a.userRepo.CreateTokensByUserId(ctx, user.ID, newUserCreds)
		if err != nil {
			return nil, fmt.Errorf("failed to save new tokens")
		}

		return newUserCreds, nil
	}

	return tokens, nil
}

func (a *AuthUseCaseImpl) Register(
	ctx context.Context,
	data *dto.AuthRequest,
) (*domain.User, error) {
	existingUser, err := a.userRepo.GetUserByEmail(ctx, data.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPassword, err := existingUser.HashPassword(data.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := domain.User{
		Email:    data.Email,
		Password: hashedPassword,
	}

	_, err = a.userRepo.CreateUser(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	userCreds, err := a.userService.GenerateTokens(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	_, err = a.userRepo.CreateTokensByUserId(ctx, user.ID, userCreds)
	if err != nil {
		return nil, fmt.Errorf("failed to save tokens: %w", err)
	}

	userData := &domain.User{
		ID:    user.ID,
		Email: user.Email,
	}
	return userData, nil
}
