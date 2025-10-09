package adapters

import (
	"database/sql"

	domain "report/internal/domain/auth"
	sqlc "report/sqlc/repository/auth"
)

func fromSQLUser(u sqlc.User) domain.AccountData {
	var name string
	if u.Name.Valid {
		name = u.Name.String
	}
	var email string
	if u.Email.Valid {
		email = u.Email.String
	}
	return domain.AccountData{
		ID:           u.ID,
		AmoUserID:    u.AmoUserID,
		Name:         name,
		Email:        email,
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
		ExpiresAt:    u.ExpiresAt,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
