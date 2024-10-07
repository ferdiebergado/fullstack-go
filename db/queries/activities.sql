-- name: CreateActivity :one
INSERT INTO
    activities (
        title,
        start_date,
        end_date,
        venue_id,
        host_id,
        metadata
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: ListActivities :many
SELECT a.*, v.name as venue, r.name as region
FROM
    active_activities a
    JOIN venues v ON v.id = a.venue_id
    JOIN regions r ON r.region_id = v.region_id
ORDER BY start_date DESC;

-- name: FindActivity :one
SELECT * FROM active_activities WHERE id = $1;

-- name: UpdateActivity :exec
UPDATE activities
SET
    title = $1,
    start_date = $2,
    end_date = $3,
    venue_id = $4,
    host_id = $5,
    metadata = $6,
    updated_at = NOW()
WHERE
    id = $7;

-- name: DeleteActivity :exec
UPDATE active_activities SET deleted_at = NOW() WHERE id = $1;

-- name: RestoreActivity :exec
UPDATE activities SET deleted_at = NULL WHERE id = $1;

-- name: FindActivityByTitle :many
SELECT * FROM active_activities WHERE title LIKE '%$1%';

-- name: FindActivityByStartDate :many
SELECT * FROM active_activities WHERE start_date = $1;

-- name: ListAllActivities :many
SELECT * FROM activities ORDER BY start_date DESC;

-- name: FindActivityAll :one
SELECT * FROM activities WHERE id = $1;