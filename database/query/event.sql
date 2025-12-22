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

-- name: ListEventsWithUserRegistration :many
SELECT
    e.id,
    e.organiser,
    e.is_online,
    e.location_name,
    e.start_time,
    e.end_time,
    e.event_name,
    (er.user_uid IS NOT NULL)::boolean AS is_registered
FROM events AS e
LEFT JOIN event_registrations AS er
    ON e.id = er.event_id AND er.user_uid = $1;

-- name: GetEventById :one
SELECT
    e.*,
    (er.user_uid IS NOT NULL)::boolean AS is_registered
FROM events AS e
LEFT JOIN event_registrations AS er
    ON e.id = er.event_id AND er.user_uid = $2
WHERE e.id = $1;

-- name: RegisterEvent :one
INSERT INTO event_registrations (user_uid, event_id)
VALUES ($1, $2)
RETURNING *;
