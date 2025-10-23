-- +goose Up
CREATE TABLE external_providers (
id SERIAL PRIMARY KEY,
service TEXT UNIQUE NOT NULL
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider_id INT NOT NULL REFERENCES external_providers(id) ON DELETE RESTRICT,
    provider_user_id TEXT NOT NULL,
    display_name TEXT,
    email TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (provider_id, provider_user_id)
);

CREATE INDEX idx_accounts_user_id ON accounts(user_id);

INSERT INTO external_providers (service) VALUES 
('amo'), 
('getcourse'), 
('bitrix');

-- +goose Down
DROP INDEX IF EXISTS idx_accounts_user_id;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS external_providers;

