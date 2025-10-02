-- name: CreateAccount :one
INSERT INTO accounts (user_id, name, account_type, currency, bank_name, last_four, nickname, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8 )
RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountsByUser :many
SELECT * FROM accounts
WHERE user_id = $1;
