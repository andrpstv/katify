-- +goose Up
CREATE TABLE IF NOT EXISTS project_invites (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    invited_email TEXT NOT NULL,
    invited_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role TEXT NOT NULL CHECK (role IN ('manager', 'bot', 'readonly')),
    status TEXT DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'declined')),
    token UUID DEFAULT gen_random_uuid(),
    expires_at TIMESTAMP DEFAULT NOW() + INTERVAL '7 days',
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (project_id, invited_email)
);

-- +goose Down
DROP TABLE IF EXISTS project_invites;
