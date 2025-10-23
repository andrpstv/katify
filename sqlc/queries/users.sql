-- name: CreateUser :one
INSERT INTO users (id, username, email, password_hash, full_name, mfa_enabled, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
-- name: GetUserByUserID :one
SELECT * FROM users WHERE id = $1;
-- name: UpdateUser :exec
UPDATE users
SET
    username = $2,
    email = $3,
    password_hash = $4,
    full_name = $5,
    mfa_enabled = $6,
    updated_at = $7
WHERE id = $1;
-- name: GetTokensByUserId :one
SELECT * FROM users_credentials WHERE user_id = $1;
-- name: CreateTokensByUserId :one
INSERT INTO users_credentials  (user_id, access_token, refresh_token, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING access_token, refresh_token, expires_at;
-- name: UpdateTokensByUserId :exec
UPDATE users_credentials
SET
    access_token = $2,
    refresh_token = $3,
    expires_at = $4
WHERE user_id = $1;