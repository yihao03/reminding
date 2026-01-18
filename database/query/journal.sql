-- name: CreateJournal :one
INSERT INTO journals (
    user_uid,
    title,
    journal_content
) VALUES ($1, $2, $3)
RETURNING *;

-- name: ListJournals :many
SELECT
    id,
    created_at,
    updated_at,
    title
FROM journals
WHERE user_uid = $1
ORDER BY updated_at DESC;

-- name: GetJournal :one
SELECT * FROM journals
WHERE id = $1 AND user_uid = $2;
