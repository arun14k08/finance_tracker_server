-- +goose Up
CREATE TABLE merchants (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category TEXT,
    website TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount NUMERIC(15,2) NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('credit','debit')),
    description TEXT,
    category TEXT,
    status TEXT DEFAULT 'completed' CHECK (status IN ('pending','completed','failed')),
    reference_id TEXT,
    metadata JSONB,
    initiated_by_user BOOLEAN DEFAULT TRUE,
    related_account_id BIGINT,
    is_self_transfer BOOLEAN DEFAULT FALSE,
    merchant_id BIGINT REFERENCES merchants(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE user_merchants (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    merchant_id BIGINT NOT NULL REFERENCES merchants(id) ON DELETE CASCADE,
    last_used TIMESTAMPTZ DEFAULT NOW(),
    usage_count INT DEFAULT 1,
    UNIQUE(user_id, merchant_id)
);


-- +goose Down
DROP TABLE user_merchants;
DROP TABLE merchants;
DROP TABLE transactions;