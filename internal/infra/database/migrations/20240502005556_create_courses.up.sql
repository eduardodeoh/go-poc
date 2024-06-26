CREATE TABLE IF NOT EXISTS courses (
    id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    title VARCHAR(255) NOT NULL UNIQUE,
    hashtags VARCHAR(255)[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);