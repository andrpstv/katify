package adapters

import (
	"context"
	"database/sql"
	"fmt"

	domain "report/internal/domain/auth"
)

type AmoAuthRepositoryImpl struct {
	db *sql.DB
}

func NewAmoAuthRepository(db *sql.DB) *AmoAuthRepositoryImpl {
	return &AmoAuthRepositoryImpl{db: db}
}

func (r *AmoAuthRepositoryImpl) Save(ctx context.Context, account *domain.AccountData) error {
	query := `
		INSERT INTO accounts (id, access_token, refresh_token, expires_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
			SET access_token = EXCLUDED.access_token,
			    refresh_token = EXCLUDED.refresh_token,
			    expires_at = EXCLUDED.expires_at;
	`

	_, err := r.db.ExecContext(ctx, query,
		account.ID,
		account.AccessToken,
		account.RefreshToken,
		account.ExpiresAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save account: %w", err)
	}

	return nil
}

func (r *AmoAuthRepositoryImpl) GetByID(ctx context.Context, id int) (*domain.AccountData, error) {
	query := `
		SELECT id, access_token, refresh_token, expires_at
		FROM accounts
		WHERE id = $1;
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var acc domain.AccountData
	err := row.Scan(
		&acc.ID,
		&acc.AccessToken,
		&acc.RefreshToken,
		&acc.ExpiresAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ничего не нашли — не ошибка
		}
		return nil, fmt.Errorf("failed to get account by id: %w", err)
	}

	return &acc, nil
}
