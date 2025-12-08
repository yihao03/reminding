-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO events (
    created_at,
    updated_at,
    organiser,
    is_online,
    location_name,
    start_time,
    end_time,
    details,
    event_name
) VALUES
(
    NOW(),
    NOW(),
    'Dementia Support Group',
    false,
    'Community Center Hall A',
    NOW() + INTERVAL '2 days',
    NOW() + INTERVAL '2 days 2 hours',
    'An introductory session covering the basics of dementia care and communication strategies.',
    'Understanding Dementia: A Caregiver''s Guide'
),
(
    NOW(),
    NOW(),
    'Caregiver Alliance',
    true,
    'Zoom Link',
    NOW() + INTERVAL '5 days',
    NOW() + INTERVAL '5 days 1 hour 30 minutes',
    'A safe space for caregivers to share experiences and find emotional support online.',
    'Virtual Support Circle'
),
(
    NOW(),
    NOW(),
    'Senior Health Services',
    false,
    'City Library Meeting Room',
    NOW() + INTERVAL '1 week',
    NOW() + INTERVAL '1 week 3 hours',
    'Practical techniques for handling agitation, confusion, and other behavioral changes.',
    'Managing Challenging Behaviors Workshop'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
TRUNCATE TABLE events RESTART IDENTITY CASCADE;
-- +goose StatementEnd
