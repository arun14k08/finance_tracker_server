-- name: CreateTransaction :one
INSERT INTO transactions 
(account_id, amount, type, description, category, status, reference_id, metadata, initiated_by_user, related_account_id, is_self_transfer, merchant_id)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING *;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE id = $1;

-- name: GetTransactionsByAccount :many
SELECT * FROM transactions
WHERE account_id = $1
ORDER BY created_at DESC
LIMIT $2;

-- name: GetSelfTransfersByUser :many
SELECT * FROM transactions t
JOIN accounts a ON t.account_id = a.id
WHERE a.user_id = $1 AND t.is_self_transfer = TRUE
ORDER BY t.created_at DESC
LIMIT $2;

-- name: GetTransactionsByMerchant :many
SELECT * FROM transactions
WHERE merchant_id = $1
ORDER BY created_at DESC
LIMIT $2;
