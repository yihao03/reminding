-- name: AddMood :one
INSERT INTO mood_tracker (
    user_uid,
    mood
) VALUES ($1, $2)
RETURNING *;

-- name: GetMonthlyMoodCountByUserUid :many
SELECT
    mood,
    COUNT(*) AS occurrence_count
FROM mood_tracker
WHERE user_uid = $1 AND entry_date >= NOW() - INTERVAL '30 days'
GROUP BY mood
ORDER BY occurrence_count DESC;
