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
WHERE user_uid = $1 AND created_at >= $2
GROUP BY mood
ORDER BY occurrence_count DESC;

-- name: CheckUserLoggedMoodToday :one
SELECT EXISTS(
    SELECT 1
    FROM mood_tracker
    WHERE user_uid = $1 AND created_at::DATE = CURRENT_DATE
) AS has_logged_mood_today;
