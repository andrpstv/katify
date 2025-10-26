package authUseCase

import (
	"context"
	"fmt"
	postgres "report/internal/adapters/db"
	domain "report/internal/domain/user"
	dto "report/internal/dto/auth"
	"report/internal/ports/outboundPorts/app/user"
)

type AuthUseCaseImpl struct {
	userRepo    user.UserRepository
	userService user.UserService
	db          postgres.PgTxManager
}

func NewAuthUseCaseImpl(userRepo user.UserRepository, userService user.UserService, db postgres.PgTxManager) *AuthUseCaseImpl {
	return &AuthUseCaseImpl{
		userRepo:    userRepo,
		userService: userService,
		db:          db,
	}
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
	if data.UserName == "" {
		return nil, fmt.Errorf("username is required")
	}
	existingUserByEmail, err := a.userRepo.GetUserByEmail(ctx, data.Email)
	if err == nil && existingUserByEmail != nil {
		return nil, fmt.Errorf("user with this email already exists")
	}

	var userData *domain.User

	err = a.db.WithTx(ctx, func(stores postgres.TxStores) error {
		var tempUser domain.User
		hashedPassword, err := tempUser.HashPassword(data.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		newUser := domain.User{
			Email:        data.Email,
			UserName:     data.UserName,
			PasswordHash: hashedPassword,
			FullName:     &data.FullName,
			MfaEnabled:   nil,
		}

		id, err := stores.User().CreateUser(ctx, &newUser)
		if err != nil {
			return fmt.Errorf("failed to create newUser: %w", err)
		}

		userCreds, err := a.userService.GenerateTokens(id)
		if err != nil {
			return fmt.Errorf("failed to generate tokens: %w", err)
		}

		_, err = stores.User().CreateTokensByUserId(ctx, id, userCreds)
		if err != nil {
			return fmt.Errorf("failed to save tokens: %w", err)
		}

		userData = &domain.User{
			ID:    newUser.ID,
			Email: newUser.Email,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return userData, nil
}
