-- +goose Up
CREATE TABLE users_credentials (
user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
access_token TEXT NOT NULL,
refresh_token TEXT NOT NULL,
expires_at TIMESTAMP NOT NULL,
encrypted_at TIMESTAMP DEFAULT NOW(),
created_at TIMESTAMP DEFAULT NOW(),
updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users_credentials;
