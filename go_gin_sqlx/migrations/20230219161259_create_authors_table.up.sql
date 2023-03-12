BEGIN;

CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_by INT REFERENCES users (id)
);

END;
