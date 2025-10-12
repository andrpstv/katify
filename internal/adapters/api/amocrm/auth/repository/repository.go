package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	domain "report/internal/domain/auth"
	sqlc "report/sqlc/repository/auth"
)

type authRepository struct {
	querier sqlc.Querier
}

func New(querier sqlc.Querier) *authRepository {
	return &authRepository{querier: querier}
}

func (r *authRepository) CreateUser(
	ctx context.Context,
	a domain.AccountData,
) (domain.AccountData, error) {
	params := sqlc.CreateUserParams{
		AmoUserID:    a.AmoUserID,
		Name:         toNullString(a.Name),
		Email:        toNullString(a.Email),
		AccessToken:  a.AccessToken,
		RefreshToken: a.RefreshToken,
		ExpiresAt:    a.ExpiresAt,
	}

	u, err := r.querier.CreateUser(ctx, params)
	if err != nil {
		return domain.AccountData{}, fmt.Errorf("create user: %w", err)
	}
	return fromSQLUser(u), nil
}

func (r *authRepository) GetUserByAmoID(
	ctx context.Context,
	amoID string,
) (domain.AccountData, error) {
	u, err := r.querier.GetUserByAmoID(ctx, amoID)
	if err != nil {
		return domain.AccountData{}, fmt.Errorf("get user by amo id: %w", err)
	}
	return fromSQLUser(u), nil
}

func (r *authRepository) GetUserByID(
	ctx context.Context,
	id uuid.UUID,
) (domain.AccountData, error) {
	u, err := r.querier.GetUserByID(ctx, id)
	if err != nil {
		return domain.AccountData{}, fmt.Errorf("get user by id: %w", err)
	}
	return fromSQLUser(u), nil
}

func (r *authRepository) UpdateUserTokens(
	ctx context.Context,
	id uuid.UUID,
	accessToken, refreshToken string,
	expiresAt time.Time,
) error {
	err := r.querier.UpdateUserTokens(ctx, sqlc.UpdateUserTokensParams{
		ID:           id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	})
	if err != nil {
		return fmt.Errorf("update user tokens: %w", err)
	}
	return nil
}
