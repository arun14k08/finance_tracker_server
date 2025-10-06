-- +goose Up
-- Update unique constraint on accounts table to include user_id

-- Drop the existing unique constraint
ALTER TABLE accounts
DROP CONSTRAINT IF EXISTS accounts_bank_name_last_four_key;

-- Add the new unique constraint including user_id
ALTER TABLE accounts
ADD CONSTRAINT accounts_user_bank_last_four_key UNIQUE (user_id, bank_name, last_four);

-- +goose Down
-- Revert back to old unique constraint (without user_id)

ALTER TABLE accounts
DROP CONSTRAINT IF EXISTS accounts_user_bank_last_four_key;

ALTER TABLE accounts
ADD CONSTRAINT accounts_bank_name_last_four_key UNIQUE (bank_name, last_four);
