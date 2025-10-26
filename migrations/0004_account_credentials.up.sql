-- +goose Up
CREATE TABLE IF NOT EXISTS accounts_credentials (
account_id UUID PRIMARY KEY REFERENCES accounts(id) ON DELETE CASCADE,
access_token TEXT NOT NULL,
refresh_token TEXT NOT NULL,
expires_at TIMESTAMP NOT NULL,
encrypted_at TIMESTAMP DEFAULT NOW(),
created_at TIMESTAMP DEFAULT NOW(),
updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS account_credentials;
