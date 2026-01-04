-- +goose Up
-- +goose StatementBegin

-- 1. Clean up existing data to start fresh
TRUNCATE TABLE event_registrations, journals, events, users RESTART IDENTITY CASCADE;

-- 2. Insert Users (A diverse mix of Admins, Medical Pros, and Caregivers across states)
INSERT INTO users (firebase_uid, display_name, email, is_admin, state, date_of_birth)
VALUES
    -- Admin
    ('uid_admin_001', 'System Admin', 'admin@carebridge.google.com', true, 'Selangor', '1985-01-01'),

    -- Medical Professionals
    ('uid_pro_001', 'Dr. Sarah Lim (Geriatrician)', 'sarah.lim@health.google.com', false, 'Penang', '1978-03-15'),
    ('uid_pro_002', 'Nurse Hafiz (OT)', 'hafiz.ot@health.google.com', false, 'Selangor', '1992-07-20'),

    -- Family Caregivers
    ('uid_cg_01', 'Uncle Robert Tan', 'robert.tan@gmail.com', false, 'Johor', '1955-06-20'), -- Caring for wife
    ('uid_cg_02', 'Aishah Binti Malik', 'aishah.malik@gmail.com', false, 'Selangor', '1988-11-02'), -- Caring for mother
    ('uid_cg_03', 'Rajiv Kumar', 'rajiv.k@gmail.com', false, 'Perak', '1990-09-12'), -- Long distance son
    ('uid_cg_04', 'Mei Ling', 'mei.ling@gmail.com', false, 'Sabah', '1982-04-05'), -- Sandwich generation (kids + parent)
    ('uid_cg_05', 'Puan Zaiton', 'zaiton.abd@gmail.com', false, 'Kelantan', '1960-12-30'); -- Caring for husband

-- 3. Insert Events (Long-form content, mix of Online/Physical)
INSERT INTO events (event_name, organiser, is_online, location_name, state, start_time, end_time, registration_link, details)
VALUES
    -- Event 1: Online Support Group
    (
        'Virtual Caregiver Circle: Finding Calm in the Storm',
        'Dementia Support Alliance',
        true,
        'Google Meet',
        NULL,
        NOW() + interval '2 days' + interval '10 hours', -- Upcoming
        NOW() + interval '2 days' + interval '11 hours 30 minutes',
        'https://meet.google.com/abc-defg-hij',
        'Navigating the journey of dementia caregiving can often feel isolating, but remember, you are never truly alone. Our Virtual Caregiver Circle is a dedicated sanctuary for you to pause, breathe, and connect with others who truly understand the nuances of your daily reality. Whether you are caring for a parent with early-onset Alzheimer’s or a spouse with vascular dementia, this session offers a judgment-free zone to share your triumphs, however small, and your struggles, however heavy.

        During this 90-minute session, we will begin with a guided 10-minute mindfulness grounding exercise designed to lower cortisol levels and help you find a moment of inner peace. Following this, we will open the floor for a structured yet flexible sharing circle. You are welcome to speak, but please know that simply listening is also a valid and honored form of participation. We understand that some days you just need to be in the presence of those who "get it" without having to explain yourself.

        We will conclude the session with a brief presentation on "Micro-Breaks: Restorative techniques that take less than 5 minutes." Please bring a warm cup of tea and a comfortable blanket; this time is strictly for you. Cameras can be on or off depending on your comfort level.'
    ),

    -- Event 2: Physical Workshop (Selangor) - Mobility
    (
        'Safe Mobility & Transfer Techniques Workshop',
        'Active Ageing Centre',
        false,
        'Sunway Medical Centre, Tower B, Room 402',
        'Selangor',
        NOW() + interval '5 days' + interval '14 hours',
        NOW() + interval '5 days' + interval '17 hours',
        'https://docs.google.com/forms/d/e/1FAIpQLSd_mobility_form/viewform',
        'Caring for a loved one with limited mobility can be physically demanding and often leads to strain or injury for the caregiver. We invite you to join us for a practical, hands-on workshop designed to empower you with the skills to move your loved one safely and with dignity. This workshop is specifically tailored for home caregivers who manage transfers from bed to wheelchair, or wheelchair to car, on a daily basis.

        Led by senior physiotherapist Mr. Chong Wei Ming, the agenda focuses on body mechanics and leverage rather than brute strength. We will demonstrate how to utilize assistive devices like gait belts and slide sheets effectively. Participants will have the opportunity to practice these techniques on medical mannequins and with each other under professional supervision to build muscle memory and confidence. We want you to leave feeling capable and protected against back strain.

        Beyond the physical skills, this afternoon is an opportunity to meet fellow caregivers in the Selangor area. We will provide a heavy tea break at 3:30 PM, allowing for casual conversation and networking. Please wear comfortable clothing and covered shoes for the practical exercises. Parking validation is provided at the front desk.'
    ),

    -- Event 3: Online Webinar - Legal
    (
        'Planning Ahead: Wills, Trusts & Power of Attorney',
        'Legal Aid for Seniors',
        true,
        'Google Meet',
        NULL,
        NOW() + interval '8 days' + interval '19 hours',
        NOW() + interval '8 days' + interval '20 hours 30 minutes',
        'https://meet.google.com/xyz-mnop-qrs',
        'Discussing legal matters and end-of-life planning is never easy, especially when a diagnosis of dementia is involved. However, having these safeguards in place is an act of love that protects both the patient and the caregiver from future bureaucratic hurdles. This webinar aims to demystify the legal jargon and provide a clear, step-by-step roadmap for families in Malaysia.

        Our guest speaker, a lawyer specializing in family and probate law, will cover the critical differences between a Will and a Power of Attorney (PA), and why a specific clause for "mental incapacity" is vital. We will discuss the concept of the "Donatio Mortis Causa" and how to manage bank accounts when a loved one is no longer able to sign documents. The tone will be gentle but informative, acknowledging the emotional weight of these tasks.

        The session includes a downloadable PDF checklist sent via Google Drive to all attendees. We will leave ample time for a Q&A session where you can ask general questions about the process without revealing sensitive personal details. Let us help you navigate the paperwork so you can focus on what matters most: caring for your loved one.'
    ),

    -- Event 4: Physical - Art Therapy (Penang)
    (
        'Sensory Art Therapy for Couples',
        'Penang Dementia Care Home',
        false,
        'Georgetown Community Hall',
        'Penang',
        NOW() + interval '12 days' + interval '10 hours',
        NOW() + interval '12 days' + interval '12 hours',
        'https://docs.google.com/forms/d/e/1FAIpQLSd_art_therapy/viewform',
        'Art has the power to bypass language barriers and connect with emotions deep within. This special Saturday morning session is designed for caregiver-patient pairs to attend together. We invite you to step away from the medical appointments and the strict routines to simply create, play, and be present with one another in a safe, supported environment.

        Facilitated by a certified Art Therapist, we will use non-toxic, sensory-friendly materials like clay, watercolors, and textured fabrics. The goal is not to create a masterpiece, but to engage the senses and spark memories. For those with dementia, art can reduce anxiety and increase social interaction. For the caregiver, it offers a rare moment to interact with your loved one not as a nurse, but as a partner, child, or friend.

        The venue is wheelchair accessible and volunteers will be on hand to assist with bathroom breaks or if anyone becomes overwhelmed. Light refreshments (kuih and coffee) will be served. Please bring an apron or wear clothes that you don''t mind getting a little paint on. Come rediscover the joy of shared activity.'
    ),

    -- Event 5: Online - Medical Info
    (
        'Understanding "Sundowning": Triggers & Management',
        'Geriatric Mental Health Assoc',
        true,
        'Google Meet',
        NULL,
        NOW() + interval '15 days' + interval '20 hours',
        NOW() + interval '15 days' + interval '21 hours 30 minutes',
        'https://meet.google.com/lun-arss-und',
        'As the sun sets, many caregivers witness a distressing change in their loved ones—increased confusion, anxiety, pacing, or even aggression. This phenomenon, known as "Sundowning," can be one of the most exhausting aspects of dementia care, disrupting the evening rest that caregivers so desperately need to recharge for the next day.

        Dr. Sarah Lim will lead this deep-dive webinar into the biological and environmental triggers behind these behaviors. We will move beyond just "coping" to look at prevention strategies: managing light exposure, regulating afternoon naps, and identifying dietary triggers like sugar or caffeine. We will also discuss how your own fatigue levels can inadvertently impact the emotional state of the person you care for.

        This is a practical, strategy-heavy session. We will provide a Google Doc template for a "Sleep & Mood Diary" that you can use to track patterns over two weeks. By understanding the rhythm of the distress, we can often find simple, non-pharmaceutical interventions to bring peace back to your evenings.'
    ),

    -- Event 6: Physical - Social (Johor)
    (
        'Caregiver Coffee Morning: JB Chapter',
        'Johor Family Support',
        false,
        'Old Town White Coffee, Tebrau City',
        'Johor',
        NOW() + interval '3 days' + interval '9 hours',
        NOW() + interval '3 days' + interval '11 hours',
        'https://docs.google.com/forms/d/e/1FAIpQLSe_coffee_jb/viewform',
        'Sometimes the best therapy is just a good cup of coffee and a listener who doesn''t need an explanation. This is an informal, drop-in social gathering for caregivers living in the Johor Bahru area. There is no agenda, no speakers, and no pressure—just a table reserved for us to laugh, vent, and share local resources.

        If you are feeling burnt out or simply haven''t spoken to another adult about something other than medication and schedules this week, please join us. We often share tips on the best local respite care services, where to find affordable medical supplies in JB, and which doctors have the best bedside manner.

        Attendees buy their own drinks/food, but the sense of community is free. Look for the table with the small purple flag in the corner. If you are running late because of care duties, please come anyway—we will be there until 11:00 AM.'
    ),

    -- Event 7: Online - Nutrition
    (
        'Nutrition and Hydration in Late Stage Dementia',
        'Dietitians of Malaysia',
        true,
        'Google Meet',
        NULL,
        NOW() + interval '20 days' + interval '14 hours',
        NOW() + interval '20 days' + interval '15 hours 30 minutes',
        'https://meet.google.com/nut-riti-onx',
        'One of the most worrying aspects of late-stage dementia is the refusal to eat or difficulty swallowing (dysphagia). It brings up primal fears for caregivers about their loved ones starving or dehydrating. This compassionate webinar addresses the physiological changes that occur in the later stages and how to manage them with love and scientific understanding.

        We will cover high-calorie, nutrient-dense food options that are easy to prepare and easier to swallow. We will demonstrate how to thicken fluids correctly and safe feeding postures to prevent aspiration pneumonia. Crucially, we will also have a gentle discussion about the ethical considerations of feeding tubes versus comfort feeding, helping you understand what to ask your medical team.

        Recipes and a "Safe Swallowing Guide" will be shared via Google Drive after the talk. This session is recommended for those caring for loved ones in Stage 6 or 7 of the progression.'
    );

-- 4. Insert Registrations (Linking Users to Events)
-- IDs are 1-7 based on insertion order above.

INSERT INTO event_registrations (event_id, user_uid)
VALUES
    -- Event 1 (Online Support): Popular, attracts many
    (1, 'uid_cg_02'), -- Aishah
    (1, 'uid_cg_03'), -- Rajiv
    (1, 'uid_cg_04'), -- Mei Ling
    (1, 'uid_cg_05'), -- Puan Zaiton

    -- Event 2 (Physical - Selangor): Aishah (local) & Nurse Hafiz (Pro)
    (2, 'uid_cg_02'),
    (2, 'uid_pro_002'),

    -- Event 3 (Online Legal): Rajiv (Needs to manage parents assets) & Uncle Robert
    (3, 'uid_cg_03'),
    (3, 'uid_cg_01'),
    (3, 'uid_cg_04'),

    -- Event 4 (Physical - Penang): Dr Sarah (Pro)
    (4, 'uid_pro_001'),

    -- Event 5 (Online Sundowning): Almost everyone needs this
    (5, 'uid_cg_01'),
    (5, 'uid_cg_02'),
    (5, 'uid_cg_05'),
    (5, 'uid_pro_002'), -- Nurse learning

    -- Event 6 (Physical - Johor): Uncle Robert (Local)
    (6, 'uid_cg_01'),

    -- Event 7 (Online Nutrition): Puan Zaiton & Mei Ling
    (7, 'uid_cg_05'),
    (7, 'uid_cg_04'),
    (7, 'uid_pro_002');

-- 5. Insert Journal Entries (To provide history/context)
INSERT INTO journals (user_uid, title, journal_content, created_at)
VALUES
    ('uid_cg_02', 'A small victory today', 'Mom actually ate her whole lunch without spitting it out. I made that porridge recipe from the forum. It felt good to see her nourished. I need to write down exactly what I did so I can repeat it tomorrow.', NOW() - interval '2 days'),

    ('uid_cg_02', 'Feeling overwhelmed', 'The nights are getting harder. She wakes up at 3am looking for her father, who passed away 20 years ago. It breaks my heart to tell her he is gone, so I have started lying and saying he is at work. I feel guilty about the lying, but it calms her down.', NOW() - interval '5 days'),

    ('uid_cg_03', 'Guilt from far away', 'Called Dad today. He sounds tired. Mom was screaming in the background. I feel so useless being in Perak while they are struggling. I transferred some money for a part-time helper, but I know money is not what they really need. They need me.', NOW() - interval '1 day'),

    ('uid_cg_01', 'She remembered our song', 'We were driving to the clinic and "Getaran Jiwa" came on the radio. For three minutes, the fog lifted. She looked at me, really looked at me, and smiled. She hummed the whole chorus. I held her hand at the traffic light. I will hold onto this memory for the bad days.', NOW() - interval '3 days'),

    ('uid_cg_05', 'Doctor appointment notes', 'Dr. Lim increased the dosage. Said the tremors might get worse before they get better. Need to watch out for dizziness. Must remember to buy the non-slip socks I saw online.', NOW() - interval '12 hours');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE event_registrations, journals, events, users RESTART IDENTITY CASCADE;
-- +goose StatementEnd
