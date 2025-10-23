package user

import (
	"context"

	"github.com/google/uuid"

	domain "report/internal/domain/user"
	sqlc "report/sqlc/repository/users"
)

type UserRepositoryImpl struct {
	Queries sqlc.Queries
}

func (a *UserRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	arg := sqlc.CreateUserParams{
		ID:           toUUID(user.ID),
		Username:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		FullName:     toNullString(user.FullName),
		CreatedAt:    toNullTime(user.CreatedAt),
		UpdatedAt:    toNullTime(user.UpdatedAt),
	}

	id, err := a.Queries.CreateUser(ctx, arg)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
func (a *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	dbUser, err := a.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:           dbUser.ID.String(),
		UserName:     dbUser.Username,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		FullName:     fromNullString(dbUser.FullName),
		MfaEnabled:   fromNullBool(dbUser.MfaEnabled),
		CreatedAt:    fromNullTime(dbUser.CreatedAt),
		UpdatedAt:    fromNullTime(dbUser.UpdatedAt),
	}

	return user, nil
}
func (a *UserRepositoryImpl) GetUserByUserID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	dbUser, err := a.Queries.GetUserByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:           dbUser.ID.String(),
		UserName:     dbUser.Username,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		FullName:     fromNullString(dbUser.FullName),
		MfaEnabled:   fromNullBool(dbUser.MfaEnabled),
		CreatedAt:    fromNullTime(dbUser.CreatedAt),
		UpdatedAt:    fromNullTime(dbUser.UpdatedAt),
	}

	return user, nil
}

func (a *UserRepositoryImpl) UpdateUser(ctx context.Context, user *domain.User) error {
	arg := sqlc.UpdateUserParams{
		ID:           toUUID(user.ID),
		Username:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		FullName:     toNullString(user.FullName),
	}

	err := a.Queries.UpdateUser(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (a *UserRepositoryImpl) GetTokensByUserId(ctx context.Context, userId uuid.UUID) (domain.UserCredentials, error) {
	dbUserCreds, err := a.Queries.GetTokensByUserId(ctx, userId)
	if err != nil {
		return domain.UserCredentials{}, err
	}
	user := &domain.UserCredentials{
		AccessToken:  dbUserCreds.AccessToken,
		RefreshToken: dbUserCreds.RefreshToken,
		ExpiresAt:    dbUserCreds.ExpiresAt,
	}
	return *user, nil
}

func (a *UserRepositoryImpl) CreateTokensByUserId(ctx context.Context, userId uuid.UUID, userCreds *domain.UserCredentials) (string, error) {
	arg := sqlc.CreateTokensByUserIdParams{
		UserID:       userId,
		AccessToken:  userCreds.AccessToken,
		RefreshToken: userCreds.RefreshToken,
		ExpiresAt:    userCreds.ExpiresAt,
	}
	dbUserTokens, err := a.Queries.CreateTokensByUserId(ctx, arg)
	if err != nil {
		return "", err
	}
	return dbUserTokens.AccessToken, nil
}

func (a *UserRepositoryImpl) UpdateTokensByUserId(ctx context.Context, userId uuid.UUID, userCreds *domain.UserCredentials) error {
	arg := sqlc.UpdateTokensByUserIdParams{
		UserID:       userId,
		AccessToken:  userCreds.AccessToken,
		RefreshToken: userCreds.RefreshToken,
		ExpiresAt:    userCreds.ExpiresAt,
	}
	err := a.Queries.UpdateTokensByUserId(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}
