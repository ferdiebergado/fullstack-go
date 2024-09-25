-- name: CreateActivity :one
INSERT INTO
    activities (
        title,
        start_date,
        end_date,
        venue,
        host,
        metadata
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: ListActivities :many
SELECT *
FROM activities
WHERE
    is_deleted = 'N'
ORDER BY start_date DESC;

-- name: FindActivity :one
SELECT * FROM activities WHERE id = $1;

-- name: UpdateActivity :exec
UPDATE activities
SET
    title = $1,
    start_date = $2,
    end_date = $3,
    venue = $4,
    host = $5,
    metadata = $6,
    updated_at = NOW()
WHERE
    id = $7;

-- name: DeleteActivity :exec
UPDATE activities SET is_deleted = 'Y' WHERE id = $1;

-- name: FindActivityByTitle :many
SELECT * FROM activities WHERE title LIKE '%$1%';

-- name: FindActivityByStartDate :many
SELECT * FROM activities WHERE start_date = $1;

-- name: ListAllActivities :many
SELECT * FROM activities ORDER BY start_date DESC;