-- name: FindUserById :one
SELECT
    *
FROM
    "users"
WHERE
    id = $1;

-- name: FindUserByUsername :one
SELECT
    *
FROM
    "users"
WHERE
    "username" = $1;

-- name: AllUsers :many
SELECT
    *
FROM
    "users"
LIMIT 50
OFFSET (50 * ($1::INTEGER - 1))::INTEGER;

-- name: UsersCount :one
SELECT
    COUNT(*) AS total_count,
    (COUNT(*) / 50)::INTEGER AS total_pages
FROM
    "users";

-- name: CreateUser :one
INSERT INTO
    "users" (created_at, updated_at, name, username, password)
VALUES
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $1, $2, $3)
RETURNING id, created_at, updated_at, name, username;

-- name: UpdateUser :one
UPDATE
    "users"
SET
    name = $2,
    username = $3,
    password = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = $1
RETURNING id, created_at, updated_at, name, username;

-- name: DeleteUser :exec
DELETE FROM
    "users"
WHERE
    id = $1;
