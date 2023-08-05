CREATE TABLE IF NOT EXISTS authors (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    user_id UUID REFERENCES users (id) NOT NULL
);
