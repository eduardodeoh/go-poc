--Seed files MUST be idempotent by design
INSERT INTO courses_users (course_id, user_id)
VALUES
    (1, 3),
    (1, 4),
    (2, 7),
    (2, 1),
    (3, 6),
    (3, 5),
    (4, 5),
    (3, 4),
    (5, 5)
ON CONFLICT (course_id, user_id) DO NOTHING
RETURNING *;