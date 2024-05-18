CREATE TABLE IF NOT EXISTS courses_users (
    course_id INTEGER NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT courses_users_pkey PRIMARY KEY (course_id, user_id)
);