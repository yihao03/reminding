-- name: GetUserByUid :one
SELECT * FROM users
WHERE firebase_uid = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: CreateUserIfAbsent :one
INSERT INTO users (firebase_uid, display_name, email)
VALUES ($1, $2, $3)
ON CONFLICT (firebase_uid) DO NOTHING
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (firebase_uid, display_name, email)
VALUES ($1, $2, $3)
RETURNING *;
