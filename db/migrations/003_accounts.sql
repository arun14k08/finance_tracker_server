-- +goose Up
-- Create Accounts Table
CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    balance NUMERIC(15,2) DEFAULT 0.0 NOT NULL,
    account_type TEXT DEFAULT 'bank', -- bank, card, wallet, crypto
    currency TEXT DEFAULT 'USD',
    bank_name TEXT,
    last_four CHAR(4),
    is_active BOOLEAN DEFAULT TRUE,
    nickname TEXT,
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Index for faster lookup
CREATE INDEX idx_accounts_user_id ON accounts(user_id);

-- +goose Down
DROP TABLE accounts;
