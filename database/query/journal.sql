-- name: CreateJournal :one
INSERT INTO journals (
    user_uid,
    title,
    journal_content
) VALUES ($1, $2, $3)
RETURNING *;

-- name: ListJournals :many
SELECT
    title,
    journal_content
FROM journals
WHERE user_uid = $1
ORDER BY updated_at;

-- name: GetJournal :one
SELECT * FROM journals
WHERE id = $1;
