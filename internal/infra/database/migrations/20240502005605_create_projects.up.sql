CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name VARCHAR(255) NOT NULL UNIQUE,
    content TEXT NOT NULL,
    course_id INTEGER NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);