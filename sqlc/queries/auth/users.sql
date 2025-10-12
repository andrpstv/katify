-- Создание аккаунта пользователя в внешнем сервисе
-- name: CreateAccount :one
INSERT INTO accounts (user_id, provider_id, provider_user_id, display_name, email)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- Создание токенов для этого аккаунта
-- name: CreateAccountCredentials :exec
INSERT INTO accounts_credentials (account_id, access_token, refresh_token, expires_at)
VALUES ($1, $2, $3, $4);
-- name: GetAccountsByUser :many
SELECT a.id, a.provider_id, a.provider_user_id, a.display_name, a.email,
       ac.access_token, ac.refresh_token, ac.expires_at
FROM accounts a
LEFT JOIN accounts_credentials ac ON a.id = ac.account_id
WHERE a.user_id = $1;
-- name: GetAccountByProviderUserID :one
SELECT a.id, a.user_id, a.provider_id, a.provider_user_id, a.display_name, a.email,
       ac.access_token, ac.refresh_token, ac.expires_at
FROM accounts a
LEFT JOIN accounts_credentials ac ON a.id = ac.account_id
WHERE a.provider_id = (SELECT id FROM external_providers WHERE service = $1)
  AND a.provider_user_id = $2;
-- name: UpdateAccountTokens :exec
UPDATE accounts_credentials
SET access_token = $2,
    refresh_token = $3,
    expires_at = $4,
    updated_at = NOW()
WHERE account_id = $1;
