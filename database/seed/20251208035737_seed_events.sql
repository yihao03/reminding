-- +goose Up
-- +goose StatementBegin

-- 1. Clean up existing data to ensure a fresh seed
TRUNCATE TABLE event_registrations, events, users, journals, mood_tracker RESTART IDENTITY CASCADE;

-- 2. Insert Users (Caregivers, Medical Professionals, and Admins)
-- Note: firebase_uid serves as the foreign key reference for other tables
INSERT INTO public.users (firebase_uid, display_name, email, is_admin, state, date_of_birth)
VALUES
    ('uid_admin_001', 'Dr. Sarah Lim', 'sarah.lim@careconnect.my', true, 'Selangor', '1980-05-15'),
    ('uid_user_001', 'Ahmad Razak', 'ahmad.razak@email.com', false, 'Johor', '1975-03-22'),
    ('uid_user_002', 'Mei Ling Tan', 'mei.ling@email.com', false, 'Penang', '1968-11-02'),
    ('uid_user_003', 'Rajiv Menon', 'rajiv.m@email.com', false, 'Selangor', '1990-07-19'),
    ('uid_user_004', 'Grace Woong', 'grace.w@email.com', false, 'Sarawak', '1985-01-30');

-- 3. Insert Events (Mix of Online and Physical events in Malaysia)
INSERT INTO public.events (
    event_name, 
    organiser, 
    is_online, 
    location_name, 
    state, 
    start_time, 
    end_time, 
    registration_link, 
    details
)
VALUES
    (
        'Understanding Sundowning: Strategies for Late-Day Confusion',
        'Alzheimer''s Disease Foundation Malaysia',
        true,
        'Zoom Webinar',
        NULL,
        NOW() + interval '2 days 18 hours', -- Upcoming in 2 days at 6 PM
        NOW() + interval '2 days 19 hours 30 minutes',
        'https://zoom.us/j/sundowning-webinar',
        'Join us for a compassionate and informative session dedicated to understanding "Sundowning"—a symptom of Alzheimer''s disease and other forms of dementia. Many caregivers observe increased confusion, anxiety, and aggression in their loved ones as the sun begins to set. This phenomenon can be one of the most exhausting challenges for family caregivers to manage alone.

In this webinar, Dr. Sarah Lim will explain the physiological triggers behind sundowning and offer practical, non-medical interventions. We will discuss the importance of lighting, routine adjustments, and dietary impacts. You are not alone in this struggle; we will open the floor for a Q&A session to share experiences and coping mechanisms that have worked for others in our community.

We invite you to bring a cup of tea and your notebook. This session is designed to be a safe space where you can learn how to create a calmer environment for your loved one and, crucially, how to maintain your own patience and well-being during these difficult hours.'
    ),
    (
        'Legal & Financial Planning Workshop for Caregivers',
        'Messrs. Lee & Partners',
        false,
        'Subang Jaya Community Centre, Hall A',
        'Selangor',
        NOW() + interval '5 days 10 hours', -- Upcoming in 5 days at 10 AM
        NOW() + interval '5 days 13 hours',
        'https://events.my/legal-planning-dementia',
        'Navigating the legal landscape while caring for a loved one with cognitive decline is daunting, but taking proactive steps now can prevent significant stress in the future. This physical workshop focuses on the essential legal instruments every caregiver in Malaysia should know about, specifically the Power of Attorney (POA) and the nuances of mental capacity laws.

We will break down complex legal jargon into plain English. Topics covered include: How to set up a Trust, managing joint bank accounts, and the specific requirements for making medical decisions on behalf of a parent or spouse. We will also touch upon the importance of having difficult conversations about wills and inheritance while your loved one can still participate in the decision-making process.

Light refreshments and lunch will be provided. Please bring any existing legal documents you have questions about (copies only). This is an opportunity to get professional advice in a relaxed, supportive setting where we understand that your priority is the care and dignity of your family member.'
    ),
    (
        'George Town Memory Café: Music & Connection',
        'Penang Dementia Care Network',
        false,
        'Hin Bus Depot Event Space',
        'Penang',
        NOW() + interval '10 days 14 hours',
        NOW() + interval '10 days 16 hours',
        'https://pdcn.org.my/memory-cafe',
        'The George Town Memory Café is a monthly gathering designed exclusively for people living with dementia and their care partners. It is a judgment-free zone where the focus is on social connection, joy, and breaking the isolation often felt by families navigating this journey. There is no agenda here other than to be present and enjoy the company of others who "get it."

This month, we are delighted to host a nostalgic music therapy session featuring hits from the 60s and 70s. Research shows that musical memory is often retained even when other cognitive functions decline. We encourage singing along, clapping, or simply listening. Volunteers will be on hand to assist, allowing caregivers a moment to relax, chat with other caregivers, and enjoy a slice of cake and local kopi.

Accessibility note: The venue is wheelchair friendly with ample parking nearby. Please register in advance so we can prepare enough refreshments. Let’s create beautiful moments together, proving that life and joy continue despite a diagnosis.'
    ),
    (
        'Caregiver Burnout: Putting Your Oxygen Mask First',
        'Mental Health Association',
        true,
        'Google Meet',
        NULL,
        NOW() + interval '1 day 20 hours',
        NOW() + interval '1 day 21 hours 30 minutes',
        'https://meet.google.com/burnout-support',
        'It is a cliché because it is true: you cannot pour from an empty cup. Yet, as dementia caregivers, guilt often drives us to sacrifice our sleep, health, and social lives for our loved ones until we reach a breaking point. This online support group meeting focuses specifically on "Compassion Fatigue" and the signs of burnout that often go unnoticed until it is too late.

Led by a clinical psychologist, we will explore the emotional rollercoaster of caregiving—from grief and anger to love and acceptance. We will practice short, actionable mindfulness techniques that can be done in 3 minutes or less during a stressful day. We will also discuss how to ask for help from family members who may be in denial about the severity of the situation.

This is a camera-optional event. If you are in your pyjamas or need to step away to tend to your loved one, that is perfectly fine. We are here to support you, validate your feelings, and remind you that prioritizing your mental health is actually the best thing you can do for the person you care for.'
    );

-- 4. Insert Event Registrations (Linking Users to Events)
-- We use subqueries to dynamically find IDs to ensure the seed is robust
INSERT INTO public.event_registrations (event_id, user_uid)
VALUES
    -- User Ahmad registers for Sundowning (Online)
    (
        (SELECT id FROM public.events WHERE event_name LIKE 'Understanding Sundowning%' LIMIT 1),
        'uid_user_001'
    ),
    -- User Mei Ling registers for Memory Cafe (Physical, in her state of Penang)
    (
        (SELECT id FROM public.events WHERE event_name LIKE 'George Town Memory Café%' LIMIT 1),
        'uid_user_002'
    ),
    -- User Rajiv registers for Legal Workshop (Physical, in his state of Selangor)
    (
        (SELECT id FROM public.events WHERE event_name LIKE 'Legal & Financial Planning%' LIMIT 1),
        'uid_user_003'
    ),
    -- User Rajiv also registers for Burnout (Online)
    (
        (SELECT id FROM public.events WHERE event_name LIKE 'Caregiver Burnout%' LIMIT 1),
        'uid_user_003'
    ),
    -- Dr Sarah (Admin) registers for Sundowning to moderate
    (
        (SELECT id FROM public.events WHERE event_name LIKE 'Understanding Sundowning%' LIMIT 1),
        'uid_admin_001'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE event_registrations, events, users, journals, mood_tracker RESTART IDENTITY CASCADE;
-- +goose StatementEnd
