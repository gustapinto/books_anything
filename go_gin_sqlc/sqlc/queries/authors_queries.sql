-- name: FindAuthorByIdAndUser :one
SELECT
    *
FROM
    "authors"
WHERE
    id = $1
    AND user_id = $2;

-- name: AllAuthorsFromUser :many
SELECT
   *
FROM
    "authors"
WHERE
    user_id = $1
LIMIT 50
OFFSET (50 * (sqlc.arg(page)::INTEGER - 1))::INTEGER;

-- name: CreateAuthor :one
INSERT INTO
    "authors" (created_at, updated_at, name, user_id)
VALUES
    (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $1, $2)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE
    "authors"
SET
    name = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = $1
    AND user_id = $2
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM
    "authors"
WHERE
    id = $1
    AND user_id = $2;