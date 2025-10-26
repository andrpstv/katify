package user

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func toNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{Valid: false}
}

func toNullBool(b *bool) sql.NullBool {
	if b != nil {
		return sql.NullBool{Bool: *b, Valid: true}
	}
	return sql.NullBool{Valid: false}
}

func toNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{Time: *t, Valid: true}
	}
	return sql.NullTime{Valid: false}
}
func toUUID(s string) uuid.UUID {
	u, err := uuid.Parse(s)
	if err != nil {
		panic("invalid UUID string: " + s) // или вернуть ошибку по-другому
	}
	return u
}
func fromNullString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}
func fromNullTime(s sql.NullTime) *time.Time {
	if s.Valid {
		return &s.Time
	}
	return nil
}
func fromNullBool(s sql.NullBool) *bool {
	if s.Valid {
		return &s.Bool
	}
	return nil
}
