-- name: CreatePosition :one
INSERT INTO
    positions (title, metadata)
VALUES ($1, $2)
RETURNING
    *;

-- name: ListPosition :many
SELECT * FROM positions ORDER BY title;

-- name: FindPosition :one
SELECT * FROM positions WHERE id = $1;

-- name: UpdatePosition :exec
UPDATE positions
SET
    title = $1,
    metadata = $2,
    updated_at = NOW()
WHERE
    id = $3;

-- name: DeletePosition :exec
UPDATE positions SET deleted_at = NOW() WHERE id = $1;

-- name: RestorePosition :exec
UPDATE positions SET deleted_at = NULL WHERE id = $1;

-- name: FindPositionByTitle :many
SELECT * FROM positions WHERE title LIKE '%$1%';