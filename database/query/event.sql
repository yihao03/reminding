-- name: ListEvents :many
SELECT
    id,
    organiser,
    is_online,
    location_name,
    start_time,
    end_time,
    event_name
FROM events
ORDER BY start_time DESC;

-- name: ListEventsWithRegistrationStatus :many
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
    ON e.id = er.event_id AND er.user_uid = $1
ORDER BY e.start_time DESC;

-- name: GetEventByIdAndUid :one
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

-- Admin Queries
-- name: GetEventById :one
SELECT
    *
FROM events
WHERE id = $1;

-- name: GetEventRegisteredUsers :many
SELECT u.*
FROM event_registrations AS er
INNER JOIN users AS u
    ON er.user_uid = u.firebase_uid
WHERE er.event_id = $1;

-- name: CreateEvent :one
INSERT INTO events (
    event_name,
    organiser,
    is_online,
    location_name,
    start_time,
    end_time,
    details
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: ListEventsAdmin :many
SELECT
    e.id,
    e.organiser,
    e.is_online,
    e.location_name,
    e.start_time,
    e.end_time,
    e.details,
    e.event_name,
    COUNT(er.user_uid) AS user_count
FROM events AS e
LEFT JOIN event_registrations AS er
    ON e.id = er.event_id
GROUP BY e.id
ORDER BY e.start_time DESC;
