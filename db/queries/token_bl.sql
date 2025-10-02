-- name: CreateBlackList :one
INSERT INTO token_blacklist (jti, expires_at) VALUES ($1, $2)
RETURNING jti, expires_at;

-- name: GetBlackListByJti :one
SELECT jti, expires_at 
FROM token_blacklist
WHERE jti = $1;

-- name: DeleteExpiredBlackList :execrows
DELETE FROM token_blacklist
WHERE expires_at < NOW()
RETURNING *;
