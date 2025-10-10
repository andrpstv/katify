package adapters

import (
	"database/sql"
	"time"

	domain "report/internal/domain/auth"
	sqlc "report/sqlc/repository/auth"
)

func fromSQLUser(u sqlc.User) domain.AccountData {
	var name, email string
	if u.Name.Valid {
		name = u.Name.String
	}
	if u.Email.Valid {
		email = u.Email.String
	}

	return domain.AccountData{
		AmoUserID:    u.AmoUserID,
		Name:         name,
		Email:        email,
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
		ExpiresAt:    u.ExpiresAt,
		CreatedAt:    nullTimeToTime(u.CreatedAt),
		UpdatedAt:    nullTimeToTime(u.UpdatedAt),
	}
}

func nullTimeToTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{} // Valid=false -> будет NULL в INSERT
	}
	return sql.NullString{String: s, Valid: true}
}
