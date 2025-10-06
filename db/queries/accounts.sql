-- name: CreateAccount :one
INSERT INTO accounts (user_id, name, account_type, currency, bank_name, last_four, nickname, notes, balance)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountsByUserId :many
SELECT * FROM accounts
WHERE user_id = $1;


-- name: GetAccountByBankAndLastFour :one
SELECT * FROM accounts
WHERE bank_name = $1 AND last_four = $2 AND user_id = $3;

-- name: UpdateAccount :one
UPDATE accounts
SET name = COALESCE($2, name),
    account_type = COALESCE($3, account_type),
    currency = COALESCE($4, currency),
    bank_name = COALESCE($5, bank_name),
    last_four = COALESCE($6, last_four),
    nickname = COALESCE($7, nickname),
    notes = COALESCE($8, notes),
    is_active = COALESCE($9, is_active)
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;