-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
ORDER BY created_at DESC
LIMIT 1;
