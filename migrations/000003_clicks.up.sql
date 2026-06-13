CREATE TABLE IF NOT EXISTS clicks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    link_id UUID REFERENCES links ON DELETE CASCADE,
    click_time TIMESTAMP DEFAULT NOW(),
    user_agent TEXT,
    user_ip VARCHAR(16)
);