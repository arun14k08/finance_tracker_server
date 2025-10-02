-- +goose Up
-- Create Transactions Table
CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount NUMERIC(15,2) NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('credit','debit')),
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Indexes for querying transactions efficiently
CREATE INDEX idx_transactions_account_id ON transactions(account_id);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);

-- +goose Down
DROP TABLE transactions;