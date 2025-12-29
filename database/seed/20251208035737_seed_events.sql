-- +goose Up
-- +goose StatementBegin

-- 1. Clean up existing data to avoid duplicates/conflicts
TRUNCATE TABLE public.event_registrations, public.events, public.users RESTART IDENTITY CASCADE;

-- 2. Insert Users
INSERT INTO public.users (firebase_uid, display_name, email, is_admin, created_at, updated_at)
VALUES 
    ('admin_uid_001', 'Sarah Jenkins (Admin)', 'sarah.j@caregiversupport.app', true, NOW(), NOW()),
    ('admin_uid_002', 'Dr. Alistair Wu', 'alistair.wu@caregiversupport.app', true, NOW(), NOW()),
    ('user_uid_101', 'Martha Stewart', 'martha.s@example.com', false, NOW(), NOW()),
    ('user_uid_102', 'David Miller', 'dave.miller88@example.com', false, NOW(), NOW()),
    ('user_uid_103', 'Elena Rodriguez', 'elena.r@example.com', false, NOW(), NOW()),
    ('user_uid_104', 'Sam O''Connor', 'soconnor@example.com', false, NOW(), NOW()),
    ('user_uid_105', 'Priya Patel', 'priya.p@example.com', false, NOW(), NOW()),
    ('user_uid_106', 'Robert Chen', 'r.chen@example.com', false, NOW(), NOW()),
    ('user_uid_107', 'Emily Blunt', 'emily.b@example.com', false, NOW(), NOW()),
    ('user_uid_108', 'Marcus Johnson', 'marcus.j@example.com', false, NOW(), NOW()),
    ('user_uid_109', 'Linda K.', 'linda.k@example.com', false, NOW(), NOW()),
    ('user_uid_110', 'Tom Hiddleston', 'tom.h@example.com', false, NOW(), NOW());

-- 3. Insert Events with Long Descriptions
INSERT INTO public.events (event_name, organiser, is_online, location_name, details, start_time, end_time, created_at, updated_at)
VALUES 
    (
        'Virtual Support Group: Late Stage Care', 
        'Sarah Jenkins', 
        true, 
        'Zoom (Link sent upon registration)', 
        'Caring for a loved one in the late stages of dementia brings a unique set of challenges, often involving limited mobility, difficulty swallowing, and non-verbal communication. It is a time that can feel incredibly isolating for the caregiver. This group provides a confidential, safe harbor to share your feelings of grief, exhaustion, and "ambiguous loss" with others who truly understand the weight you are carrying.

In this session, we will open with a brief check-in for every member, followed by a guided discussion on palliative care measures and comfort strategies. We will also discuss how to maintain connection when verbal communication is no longer possible, using touch, music, and presence.

Please note: This group is specifically for those caring for someone in Stage 6 or 7. If you are new to this journey, please look for our Early Stage group to ensure the discussions are relevant to your current needs. Cameras are optional if you are feeling overwhelmed, but we encourage active listening.', 
        NOW() + interval '2 days' + interval '18 hours',
        NOW() + interval '2 days' + interval '19 hours 30 minutes',
        NOW(), NOW()
    ),
    (
        'Understanding Sundowning: Expert Q&A', 
        'Dr. Alistair Wu', 
        true, 
        'Google Meet', 
        'Does your loved one become increasingly confused, agitated, or anxious as the sun begins to set? This phenomenon, known as "Sundowning," is a common symptom of Alzheimer''s disease and other forms of dementia. It can be physically and emotionally draining for caregivers who are already tired from the day''s demands.

Join Dr. Alistair Wu, a geriatric psychiatrist specializing in dementia behaviors, for this deep-dive webinar. Dr. Wu will explain the biological circadian rhythm disruptions that trigger these episodes and offer practical, non-pharmaceutical interventions. We will cover lighting strategies, dietary adjustments, and calming evening routines that can help reduce the severity of symptoms.

The last 30 minutes of the session will be dedicated entirely to your questions. Please feel free to submit specific scenarios you are facing in the chat during the presentation. A recording will be made available to all registered attendees.', 
        NOW() + interval '5 days' + interval '14 hours', 
        NOW() + interval '5 days' + interval '15 hours 30 minutes',
        NOW(), NOW()
    ),
    (
        'Caregiver Wellness: 15-Minute Reset', 
        'Mindful Care Team', 
        true, 
        'Zoom', 
        'As a caregiver, you spend your entire day focused on the needs of someone else. It is easy to forget that you need to breathe, too. This 15-minute micro-session is designed to fit into even the most chaotic schedule, offering you a brief moment of respite to regulate your nervous system.

We will focus on "Box Breathing" and progressive muscle relaxation techniques that you can use anywhere—even in the middle of a difficult caregiving moment—to lower your cortisol levels. No yoga mats or special clothing required; just find a chair and log in.

This is a camera-off event to ensure maximum privacy and relaxation. You do not need to speak or interact; simply listen to the guidance and allow yourself a quarter of an hour to just "be" rather than "do."', 
        NOW() + interval '1 day' + interval '8 hours', 
        NOW() + interval '1 day' + interval '8 hours 15 minutes',
        NOW(), NOW()
    ),
    (
        'Saturday Morning Memory Café', 
        'Community Center', 
        false, 
        'Highland Park Community Center, Room B', 
        'The Memory Café is a stigma-free social gathering for people with mild to moderate dementia and their care partners. We know that going out in public can sometimes feel daunting due to unpredictable behaviors; here, no one judges, and everyone understands. It is a place to leave the disease at the door and just enjoy being a person again.

This month, we will have a local jazz guitarist playing gentle background music, and we will be serving coffee, tea, and pastries. There will be volunteers on hand to engage with your loved ones in simple games or crafts, giving you—the caregiver—a chance to sit back, sip your coffee, and chat with other families.

The venue is fully wheelchair accessible. Accessible restrooms are located directly across the hall. Please RSVP so we can get an accurate headcount for food catering.', 
        NOW() + interval '3 days' + interval '10 hours', 
        NOW() + interval '3 days' + interval '12 hours',
        NOW(), NOW()
    ),
    (
        'Legal Workshop: Power of Attorney & Wills', 
        'Legal Aid Society', 
        false, 
        'City Library - Conference Room 1', 
        'Navigating the legal landscape of dementia care is critical, yet often overwhelmed by complex jargon. This workshop aims to demystify the essential documents you need to protect your loved one''s assets and ensure their medical wishes are honored when they can no longer speak for themselves.

We will cover the difference between Medical Power of Attorney and Financial Power of Attorney, the importance of Living Wills (Advance Directives), and basic estate planning considerations. We will also discuss what happens if these documents are not in place and the process of obtaining Guardianship/Conservatorship.

Disclaimer: This session provides legal information, not legal advice for specific cases. However, we will provide a list of local elder law attorneys and pro-bono resources for low-income families at the end of the session. Please bring a notepad and pen.', 
        NOW() + interval '10 days' + interval '17 hours', 
        NOW() + interval '10 days' + interval '19 hours',
        NOW(), NOW()
    ),
    (
        'Respite Care Information Session', 
        'Senior Services Dept', 
        false, 
        'St. Mary Hospital Auditorium', 
        'Caregiver burnout is not a sign of weakness; it is a reality of the job. "Respite" simply means taking a break, but finding trustworthy care for your loved one while you rest can be difficult. This session explores the various respite options available in our county, from in-home companions to adult day centers and short-term residential stays.

Representatives from the State Department on Aging will be present to explain government vouchers and grant programs that can help cover the cost of respite care. We will break down the application process step-by-step so you can access the funds you are entitled to.

We will also discuss the emotional hurdle of "letting go" and trusting others with your loved one''s care. Hearing from other caregivers who have successfully used respite services can help alleviate the guilt often associated with taking time for oneself.', 
        NOW() + interval '14 days' + interval '14 hours', 
        NOW() + interval '14 days' + interval '15 hours',
        NOW(), NOW()
    );

-- 4. Insert Event Registrations
INSERT INTO public.event_registrations (event_id, user_uid, registration_date)
VALUES 
    -- Event 1: Late Stage Support
    (1, 'user_uid_101', NOW() - interval '2 days'),
    (1, 'user_uid_103', NOW() - interval '1 day'),
    (1, 'user_uid_105', NOW() - interval '5 hours'),

    -- Event 2: Sundowning Q&A
    (2, 'user_uid_101', NOW() - interval '3 days'),
    (2, 'user_uid_102', NOW() - interval '2 days'),
    (2, 'user_uid_104', NOW() - interval '1 day'),
    (2, 'user_uid_106', NOW() - interval '12 hours'),
    (2, 'user_uid_108', NOW() - interval '2 hours'),
    (2, 'user_uid_110', NOW() - interval '1 hour'),

    -- Event 3: Meditation
    (3, 'user_uid_109', NOW() - interval '1 day'),
    (3, 'user_uid_107', NOW() - interval '4 hours'),

    -- Event 4: Memory Café
    (4, 'user_uid_102', NOW() - interval '5 days'),
    (4, 'user_uid_103', NOW() - interval '4 days'),
    (4, 'user_uid_105', NOW() - interval '2 days'),

    -- Event 5: Legal Workshop
    (5, 'user_uid_104', NOW() - interval '6 days'),
    (5, 'user_uid_106', NOW() - interval '3 days'),
    (5, 'user_uid_108', NOW() - interval '1 day'),
    (5, 'user_uid_101', NOW() - interval '12 hours');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE public.event_registrations, public.events, public.users RESTART IDENTITY CASCADE;
-- +goose StatementEnd
