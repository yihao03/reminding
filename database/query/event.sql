-- name: ListEvents :many
SELECT
    id,
    organiser,
    is_online,
    location_name,
    start_time,
    end_time,
    event_name
FROM events;

-- name: GetEventById :one
SELECT * FROM events
WHERE id = $1;
