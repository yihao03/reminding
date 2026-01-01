-- name: GetUserByUid :one
SELECT * FROM users
WHERE firebase_uid = $1
ORDER BY created_at
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (firebase_uid, display_name, email, state, age)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (firebase_uid) DO NOTHING
RETURNING *;
