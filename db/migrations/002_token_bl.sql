-- +goose Up
CREATE TABLE token_blacklist (
    jti TEXT PRIMARY KEY,
    expires_at TIMESTAMPTZ NOT NULL
);

-- +goose Down
DROP TABLE token_blacklist;
