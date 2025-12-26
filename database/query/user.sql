-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (firebase_uid, display_name, email)
VALUES ($1, $2, $3)
RETURNING *;
