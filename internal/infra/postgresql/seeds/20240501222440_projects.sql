--Seed files MUST be idempotent by design
INSERT INTO projects (name, content, course_id, created_at)
VALUES
    ('Basics of dissection', '<redacted>', 1, '2024-04-21 14:30:00'),
    ('Intro to maps', 'lots of stuff', 2, '2024-04-21 14:30:00'),
    ('Structs', 'User defined structs', 2, '2024-04-21 14:30:00'),
    ('Omelette', 'yummy', 4, '2024-04-21 14:30:00'),
    ('Croissants', 'yummy^2', 4, '2024-04-21 14:30:00'),
    ('Fourier transforms', 'siny and cosy', 5, '2024-04-21 14:30:00'),
    ('Gibbs effect', 'ring-a-ding', 5, '2024-04-21 14:30:00'),
    ('Superposition', 'ax+by', 5, '2024-04-21 14:30:00'),
    ('Convolutions', 'not the ML kind', 5, '2024-04-21 14:30:00')
ON CONFLICT (name) DO NOTHING
RETURNING *;

