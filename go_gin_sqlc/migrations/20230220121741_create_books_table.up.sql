CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    isbn CHAR(13) UNIQUE NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    author_id UUID REFERENCES authors (id) NOT NULL,
    user_id UUID REFERENCES users (id) NOT NULL
);
