-- name: CreateMerchant :one
INSERT INTO merchants (name, category, website)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetMerchantByID :one
SELECT * FROM merchants
WHERE id = $1;

-- name: GetMerchantByName :one
SELECT * FROM merchants
WHERE name = $1;

-- name: ListMerchants :many
SELECT * FROM merchants
ORDER BY name;

-- name: AddOrUpdateUserMerchant :one
INSERT INTO user_merchants (user_id, merchant_id, last_used, usage_count)
VALUES ($1, $2, NOW(), 1)
ON CONFLICT (user_id, merchant_id)
DO UPDATE SET 
    last_used = NOW(),
    usage_count = user_merchants.usage_count + 1
RETURNING *;

-- name: GetFrequentMerchantsByUser :many
SELECT m.*
FROM merchants m
JOIN user_merchants um ON m.id = um.merchant_id
WHERE um.user_id = $1
ORDER BY um.usage_count DESC
LIMIT $2;
