CREATE TABLE IF NOT EXISTS links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users ON DELETE CASCADE,
    short_code VARCHAR(6) UNIQUE,
    original_url VARCHAR(512) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    is_delete BOOL DEFAULT FALSE
);