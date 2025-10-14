package application

import (
	"context"
	"fmt"

	domain "report/internal/domain/user"
	dto "report/internal/dto/auth"
	userRepo "report/internal/ports/outboundPorts/repositories/user"
	tokenService "report/internal/ports/outboundPorts/token"
)

type AuthUseCaseImpl struct {
	userRepo     userRepo.UserRepository
	tokenService tokenService.TokenService
}

func (a *AuthUseCaseImpl) Login(
	ctx context.Context,
	data *dto.AuthRequest,
) (*domain.TokenPair, error) {
	user, err := a.userRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found. please register", err)
	}

	tokenPair, err := a.tokenService.Generate(user.ID)
	if err != nil {
		return nil, fmt.Errorf("can't generate token", err)
	}

	return tokenPair, err
}

func (a *AuthUseCaseImpl) Register(
	ctx context.Context,
	data *dto.AuthRequest,
) (*domain.TokenPair, error) {
	user := domain.User{}
	error := a.userRepo.Create(ctx)
}
