--Seed files MUST be idempotent by design
INSERT INTO courses (title, hashtags, created_at)
VALUES
    ('Introduction to Biology', array['science', 'hands-on'], '2024-04-21 13:30:00'),
    ('Golang 101', array['programming', 'concurrency', 'code'], '2024-04-21 13:30:00'),
    ('Organic Chemistry', array['science'], '2024-04-21 13:30:00'),
    ('French cuisine', array['cooking', 'international'], '2024-04-21 13:30:00'),
    ('Signals and Systems', array['science', 'EE', 'math', 'control'], '2024-04-21 13:30:00'),
    ('Elixir 101', array['programming', 'concurrency', 'code', 'beam', 'functional'], '2024-04-21 13:30:00')
ON CONFLICT (title) DO NOTHING
RETURNING *;