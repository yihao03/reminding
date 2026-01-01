-- +goose Up
-- +goose StatementBegin

-- 1. Clean up existing data to ensure a fresh seed
-- We use CASCADE to handle the foreign key constraints automatically.
TRUNCATE TABLE event_registrations, events, users RESTART IDENTITY CASCADE;

-- 2. Insert Users
-- We create a mix of admins (medical professionals) and standard users (family caregivers).
INSERT INTO users (firebase_uid, display_name, email, is_admin, state, age)
VALUES
    -- Admin: A Geriatric Specialist
    ('uid_admin_sarah', 'Dr. Sarah Lim', 'sarah.lim@healthcare.com', true, 'Selangor', 45),
    
    -- User: A son caring for his father
    ('uid_caregiver_ahmad', 'Ahmad Zulkifli', 'ahmad.z@gmail.com', false, 'Johor', 34),
    
    -- User: A wife caring for her husband with early-onset
    ('uid_caregiver_mei', 'Mei Ling', 'mei.ling@yahoo.com', false, 'Penang', 58),
    
    -- User: A professional live-in nurse
    ('uid_nurse_ravi', 'Ravi Chandran', 'ravi.care@nursing.com', false, 'Perak', 29);

-- 3. Insert Events
-- Includes a mix of Online and Physical events with long-form, empathetic descriptions.
INSERT INTO events (
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
        'Understanding Sundowning: Strategies for Evening Anxiety',
        'Alzheimers Disease Foundation Malaysia',
        true, -- Online
        'Zoom Webinar',
        'Selangor', -- Host location
        NOW() + interval '2 days 18 hours', -- 6:00 PM roughly
        NOW() + interval '2 days 20 hours',
        'https://www.google.com',
        'Many caregivers notice a distinct change in their loved one''s behavior as the late afternoon light begins to fade. This phenomenon, known as "Sundowning," can result in increased confusion, anxiety, aggression, or pacing. It is one of the most exhausting aspects of dementia care because it happens when you, the caregiver, are likely most tired. In this compassionate workshop, Dr. Sarah Lim will explain the biological triggers behind these behavioral shifts and why they occur specifically during transition times of the day.

We will move beyond medical definitions to discuss practical, home-based strategies. Topics will include lighting adjustments to reset circadian rhythms, dietary timing to reduce agitation, and calming sensory activities that can ground your loved one. We understand how isolating these evenings can feel, and we want to equip you with a toolkit to make nights more peaceful for the whole family.

The session will conclude with a 30-minute open Q&A where you can share your specific challenges. Please note that this is a safe, judgment-free space. Whether you are dealing with wandering or emotional outbursts, you are not alone in this journey. Recording will be available for those who cannot attend live.'
    ),
    (
        'Penang Memory Café: Coffee, Cake & Connection',
        'George Town Care Network',
        false, -- Physical
        'The Heritage Hall, George Town',
        'Penang',
        NOW() + interval '5 days 10 hours', -- Morning event
        NOW() + interval '5 days 13 hours',
        'https://www.google.com',
        'We invite you and your loved ones living with dementia to our monthly Memory Café, a safe and welcoming gathering designed to bring joy back into social outings. For many families, going out to public restaurants can be stressful due to fear of judgment or behavioral unpredictability. Our Memory Café is different—it is a sanctuary where everyone understands, and everyone belongs.

Come enjoy complimentary local coffee, tea, and traditional kuih in a relaxed atmosphere. This month, we are featuring a gentle music therapy session led by local musicians playing nostalgic hits from the 60s and 70s. Music has a unique way of unlocking memories and sparking joy even when words fail. There is no pressure to participate; sitting back and listening is perfectly encouraged.

Volunteers will be on hand to assist, allowing caregivers a moment to breathe and connect with one another. This is not a clinical setting or a support group meeting, but simply a time to relax and enjoy the company of others who walk a similar path. Wheelchair access is available via the main ramp.'
    ),
    (
        'Legal Essentials: Power of Attorney & Wills',
        'Legal Aid for Seniors',
        true, -- Online
        'Microsoft Teams',
        'Selangor',
        NOW() + interval '10 days 14 hours',
        NOW() + interval '10 days 16 hours',
        'https://www.google.com',
        'Navigating the legal landscape while caring for someone with dementia can be overwhelming, yet it is one of the most critical steps in ensuring their long-term protection. This seminar addresses the difficult but necessary conversations regarding the Mental Capacity Act and the specific legal instruments available in Malaysia. We will discuss the crucial difference between a Power of Attorney (PA) and a Lasting Power of Attorney (LPA), and why timing is everything.

Our guest legal experts will walk you through the process of freezing accounts, managing joint assets, and what happens if a loved one loses the capacity to sign documents before arrangements are made. We will break down complex legal jargon into plain language, ensuring you understand your rights and responsibilities as a next-of-kin or appointed guardian.

We know thinking about the future can be frightening. This session aims to replace fear with preparedness. By handling these administrative burdens now, you can focus more on providing care and love in the future. Checklists and template guides will be provided to all registered participants via email after the session.'
    ),
    (
        'Caregiver Burnout: Putting Your Oxygen Mask First',
        'Mental Health Association',
        false, -- Physical
        'Community Center, Johor Bahru',
        'Johor',
        NOW() + interval '1 day 9 hours',
        NOW() + interval '1 day 11 hours',
        'https://www.google.com',
        'You spend your days (and often nights) caring for someone else, but who is caring for you? Caregiver burnout is a very real, physical, and emotional state of exhaustion that affects nearly all family caregivers at some point. Symptoms include irritability, sleep problems, weight changes, and feelings of hopelessness. In this intimate support circle, we validate these feelings: it is not selfish to need a break.

Facilitated by a licensed counselor, this workshop focuses on "micro-self-care"—small, manageable actions you can take even on your busiest days to lower cortisol levels. We will practice simple breathing techniques and cognitive reframing exercises to help manage the guilt that often accompanies taking time for oneself.

We will also facilitate a sharing circle. Hearing others say, "I feel that way too," is a powerful antidote to the isolation of dementia care. Refreshments will be provided. Please note: Respite care volunteers are available in the adjacent room to look after your loved one while you attend this session, so you can attend with full peace of mind.'
    );

-- 4. Insert Event Registrations
-- We map specific users to the events created above using subqueries to find Event IDs.

INSERT INTO event_registrations (event_id, user_uid)
VALUES
    -- Ahmad (Johor) registers for the Online Sundowning talk and the physical Burnout session in Johor
    (
        (SELECT id FROM events WHERE event_name LIKE 'Understanding Sundowning%' LIMIT 1), 
        'uid_caregiver_ahmad'
    ),
    (
        (SELECT id FROM events WHERE event_name LIKE 'Caregiver Burnout%' LIMIT 1), 
        'uid_caregiver_ahmad'
    ),

    -- Mei (Penang) registers for the Memory Cafe in Penang and the Online Legal talk
    (
        (SELECT id FROM events WHERE event_name LIKE 'Penang Memory Café%' LIMIT 1), 
        'uid_caregiver_mei'
    ),
    (
        (SELECT id FROM events WHERE event_name LIKE 'Legal Essentials%' LIMIT 1), 
        'uid_caregiver_mei'
    ),

    -- Nurse Ravi registers for the Sundowning talk (for professional development)
    (
        (SELECT id FROM events WHERE event_name LIKE 'Understanding Sundowning%' LIMIT 1), 
        'uid_nurse_ravi'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE event_registrations, events, users RESTART IDENTITY CASCADE;
-- +goose StatementEnd
