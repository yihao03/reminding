-- +goose Up
-- +goose StatementBegin

-- 1. Clean up existing data to ensure a fresh seed
TRUNCATE TABLE event_registrations, events, users RESTART IDENTITY CASCADE;

-- 2. Insert Users (Admins, Medical Professionals, and Family Caregivers)
INSERT INTO users (firebase_uid, display_name, email, is_admin, state, age)
VALUES 
    -- Admins / Professionals
    ('uid_admin_001', 'Dr. Aminah Razak', 'aminah.razak@carebridge.my', true, 'Selangor', 45),
    ('uid_admin_002', 'Sarah Lee (Social Worker)', 'sarah.lee@carebridge.my', true, 'Penang', 32),
    
    -- Family Caregivers
    ('uid_user_101', 'Ravi Muthusamy', 'ravi.muthu88@gmail.com', false, 'Johor', 38),
    ('uid_user_102', 'Lim Mei Ling', 'mei.ling.lim@yahoo.com', false, 'Selangor', 52),
    ('uid_user_103', 'Ahmad Zulkifli', 'ahmad.zul@hotmail.com', false, 'Kelantan', 29),
    ('uid_user_104', 'Grace O''Connor', 'grace.oc@gmail.com', false, 'Sabah', 61),
    ('uid_user_105', 'Farid bin Harun', 'farid.h@gmail.com', false, 'Terengganu', 44);

-- 3. Insert Events (Mix of Online and In-Person with Long-Form Content)
INSERT INTO events (organiser, is_online, location_name, start_time, end_time, event_name, state, registration_link, details)
VALUES
    (
        'Alzheimer’s Support Network',
        true,
        'Zoom Meeting',
        NOW() + interval '2 days 10 hours', -- Upcoming in 2 days
        NOW() + interval '2 days 12 hours',
        'Understanding Sundowning: Strategies for Evening Anxiety',
        NULL, -- Online, no specific state state
        'https://zoom.us/j/123456789',
        E'Sundowning can be one of the most challenging behaviors for caregivers to manage. As the sun begins to set, you may notice your loved one becoming increasingly confused, anxious, or agitated. This webinar is designed to help you understand the biological and environmental triggers behind this phenomenon. We will discuss how lighting, routine adjustments, and dietary changes can significantly reduce the severity of symptoms.\n\nIn this session, Dr. Aminah Razak will share medical insights into the circadian rhythms of dementia patients. We will also open the floor to three experienced caregivers who will share their personal "evening routine" checklists that have brought peace back to their homes. You are not alone in this struggle, and there are practical steps we can take together.\n\nPlease note: This is a safe space for sharing. We will begin with a 45-minute presentation followed by a supportive Q&A session. A recording will be made available to all registered participants if you are unable to attend the full duration due to caregiving duties.'
    ),
    (
        'Legal Aid Bureau',
        true,
        'Google Meet',
        NOW() + interval '5 days 18 hours',
        NOW() + interval '5 days 20 hours',
        'Navigating Legalities: Wills & Power of Attorney',
        NULL,
        'https://meet.google.com/abc-defg-hij',
        E'Thinking about legal matters while managing the emotional toll of a dementia diagnosis can feel overwhelming. However, establishing a Power of Attorney (POA) and updating wills is crucial to ensuring your loved one’s wishes are respected when they can no longer advocate for themselves. This session aims to demystify the legal jargon and provide a clear, step-by-step roadmap for families in Malaysia.\n\nWe will cover the specific differences between a standard POA and a Lasting Power of Attorney (LPA), and why the timing of these documents is critical regarding mental capacity assessments. Our guest lawyer will also touch upon handling bank accounts, insurance claims, and managing assets to fund long-term care without running into bureaucratic roadblocks.\n\nThis session is free of charge. We encourage you to submit your questions anonymously via the registration link beforehand so we can address specific concerns regarding inheritance and guardianship without compromising your privacy. Handouts with legal templates will be emailed after the session.'
    ),
    (
        'Penang Carers Alliance',
        false,
        'Georgetown Community Centre, Hall B',
        NOW() + interval '1 week 09 hours',
        NOW() + interval '1 week 13 hours',
        'Caregiver Burnout: A Morning of Respite',
        'Penang',
        NULL, -- Walk-in / RSVP
        E'Caregiving is a marathon, not a sprint, and too often the caregiver’s health is neglected. We invite you to a physical gathering at the Georgetown Community Centre specifically focused on YOU—the caregiver. This is a morning dedicated to mental health, relaxation, and connecting with others who truly understand the weight you carry every day.\n\nThe morning will begin with a guided breathing and mindfulness session led by a clinical psychologist, focusing on techniques you can use in high-stress moments at home. Following this, we will have a "Respite Roundtable" where we share resources for temporary care options in Penang, giving you the practical tools to take necessary breaks. Light refreshments and brunch will be served.\n\nVolunteers from the local nursing college will be present in the adjacent room to provide supervised activities for your loved ones with dementia, should you need to bring them along. This allows you to fully engage in the workshop with peace of mind. Please RSVP so we can arrange adequate staffing for the supervised care room.'
    ),
    (
        'Selangor Memory Café',
        false,
        'The Bee Cafe, Jaya One',
        NOW() + interval '2 weeks 14 hours',
        NOW() + interval '2 weeks 16 hours',
        'Monthly Memory Café: Music & Connection',
        'Selangor',
        'https://carebridge.my/events/memory-cafe-sept',
        E'Music has the power to unlock memories and emotions that words sometimes cannot. Join us for our monthly Memory Café, a comfortable, judgement-free social gathering for people with dementia and their care partners. This month, we are featuring a live acoustic performance of classic hits from the 60s and 70s, designed to encourage sing-alongs and toe-tapping.\n\nThe Memory Café is about focusing on what our loved ones can still do, rather than what they cannot. It is an opportunity to socialize without the fear of stigma or embarrassment if a drink is spilled or a sentence is repeated. The environment is relaxed, the lighting is gentle, and the staff are trained to be dementia-friendly.\n\nCoffee, tea, and cakes are on the house, sponsored by the Rotary Club. Whether you want to dance, sing, or just sit back and listen, you are welcome here. It is a wonderful opportunity for caregivers to chat with one another while their loved ones are engaged and happy. We look forward to seeing you there!'
    );

-- 4. Insert Registrations (Linking Users to Events via Firebase UID and Event IDs)
-- Note: We assume IDs 1, 2, 3, 4 based on the TRUNCATE RESTART IDENTITY above.

INSERT INTO event_registrations (event_id, user_uid)
VALUES
    -- Event 1: Sundowning (Online) - Popular
    (1, 'uid_user_101'), -- Ravi
    (1, 'uid_user_102'), -- Mei Ling
    (1, 'uid_user_105'), -- Farid
    
    -- Event 2: Legalities (Online)
    (2, 'uid_user_103'), -- Ahmad
    (2, 'uid_user_101'), -- Ravi (active user)
    
    -- Event 3: Caregiver Burnout (In-Person Penang)
    -- Only the user in Penang and maybe a admin registered
    (3, 'uid_admin_002'), -- Sarah (Admin attending)
    
    -- Event 4: Memory Cafe (In-Person Selangor)
    (4, 'uid_user_102'), -- Mei Ling (Lives in Selangor)
    (4, 'uid_admin_001'); -- Dr Aminah (Lives in Selangor)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE event_registrations, events, users RESTART IDENTITY CASCADE;
-- +goose StatementEnd
