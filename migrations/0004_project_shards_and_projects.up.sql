-- +goose Up
CREATE TABLE workspaces (
id BIGSERIAL PRIMARY KEY,
account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
uuid UUID DEFAULT gen_random_uuid(),
name TEXT NOT NULL,
subdomain TEXT,
shard_type INT,
version INT DEFAULT 0,
is_kommo BOOLEAN DEFAULT FALSE,
is_trial BOOLEAN DEFAULT FALSE,
trial_ended BOOLEAN DEFAULT FALSE,
is_payed BOOLEAN DEFAULT FALSE,
payed_ended BOOLEAN DEFAULT FALSE,
mfa_enabled BOOLEAN DEFAULT FALSE,
created_at TIMESTAMP DEFAULT NOW(),
updated_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    amo_workspace_id BIGINT REFERENCES workspaces(id) ON DELETE CASCADE,
    getcourse_workspace_id BIGINT REFERENCES workspaces(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (amo_workspace_id, getcourse_workspace_id)
);

CREATE INDEX idx_projects_account_id ON projects(account_id);

-- +goose Down
DROP INDEX IF EXISTS idx_projects_account_id;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS project_links;
DROP TABLE IF EXISTS project_shards;

