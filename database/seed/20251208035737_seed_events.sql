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
    'An introductory session covering the basics of dementia care and communication strategies. Topics include: identifying early signs, effective verbal and non-verbal communication, and creating a dementia-friendly home environment. Notable guest: Dr. Sarah Chen, Geriatric Specialist.',
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
    'A safe space for caregivers to share experiences and find emotional support online. We will discuss coping mechanisms for caregiver burnout and strategies for self-care. Facilitated by licensed therapist Mark Johnson.',
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
    'Practical techniques for handling agitation, confusion, and other behavioral changes. This workshop includes role-playing scenarios and Q&A sessions. Guest speaker: Nurse Practitioner Emily White, specializing in behavioral health.',
    'Managing Challenging Behaviors Workshop'
),
(
    NOW(),
    NOW(),
    'Legal Aid for Seniors',
    false,
    'Downtown Law Office Conference Room',
    NOW() + INTERVAL '10 days',
    NOW() + INTERVAL '10 days 2 hours',
    'A comprehensive seminar on legal planning for long-term care. Topics covered: Power of Attorney, Living Wills, Guardianship, and Estate Planning basics. Featuring guest attorney Robert Davis from Davis & Associates.',
    'Legal Planning for the Future'
),
(
    NOW(),
    NOW(),
    'Mindful Living Institute',
    true,
    'Zoom Link',
    NOW() + INTERVAL '12 days',
    NOW() + INTERVAL '12 days 1 hour',
    'Learn mindfulness and stress reduction techniques specifically tailored for caregivers. The session will cover breathing exercises, guided meditation, and finding moments of peace in a busy day. Led by mindfulness coach Lisa Wong.',
    'Mindfulness for Caregivers'
),
(
    NOW(),
    NOW(),
    'Nutrition & Wellness Board',
    false,
    'General Hospital Auditorium',
    NOW() + INTERVAL '2 weeks',
    NOW() + INTERVAL '2 weeks 2 hours',
    'Exploring the impact of nutrition on cognitive health. We will discuss brain-boosting foods, meal planning for seniors with swallowing difficulties, and hydration importance. Guest speaker: Clinical Dietitian Michael Brown.',
    'Nutrition and Brain Health'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
TRUNCATE TABLE events RESTART IDENTITY CASCADE;
-- +goose StatementEnd
