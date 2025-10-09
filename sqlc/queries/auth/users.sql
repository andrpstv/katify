-- name: CreateUser :one
INSERT INTO users (amo_user_id, name, email, access_token, refresh_token, expires_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByAmoID :one
SELECT * FROM users WHERE amo_user_id = $1;

-- name: UpdateUserTokens :exec
UPDATE users
SET access_token = $2, refresh_token = $3, expires_at = $4, updated_at = NOW()
WHERE id = $1;
