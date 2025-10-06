-- +goose Up
-- Alter accounts table: set currency default to INR and add trigger for cash accounts

-- 1️⃣ Change currency default to 'INR'
ALTER TABLE accounts
ALTER COLUMN currency SET DEFAULT 'INR';

-- +goose Down
-- Drop the trigger and function, revert defaults
ALTER TABLE accounts
ALTER COLUMN currency SET DEFAULT 'USD';`