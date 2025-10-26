package postgres

import (
	"context"
	"database/sql"
	"report/internal/adapters/app/user"
	userPort "report/internal/ports/outboundPorts/app/user"
	sqlc "report/sqlc/repository/users"
)

type TxManager interface {
	WithTx(ctx context.Context, fn func(stores TxStores) error) error
}

type TxStores interface {
	User() userPort.UserRepository
}

type PgTxManager struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewTxManager(db *sql.DB, queries *sqlc.Queries) *PgTxManager {
	return &PgTxManager{db: db, queries: queries}
}

func (m *PgTxManager) WithTx(ctx context.Context, fn func(stores TxStores) error) error {
	return WithTx(ctx, m.db, func(tx *sql.Tx) error {
		q := m.queries.WithTx(tx)
		stores := &txStores{
			user: user.NewUserRepositoryImpl(*q),
		}
		return fn(stores)
	})
}

type txStores struct {
	user userPort.UserRepository
}

func (t *txStores) User() userPort.UserRepository { return t.user }
