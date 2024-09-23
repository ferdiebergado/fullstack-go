-- name: createActivity :one
INSERT INTO activities (title, start_date, end_date, venue, host, status, metadata)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *

-- name: listActivities :many
SELECT * FROM activities ORDER BY start_date DESC

-- name: findActivity :one
SELECT * FROM activities WHERE id = $1

-- name: updateActivity :exec
UPDATE activities 
SET title = $1, start_date = $2, end_date = $3, venue = $4, host = $5, status = $6, metadata = $7, updated_at = NOW()
WHERE id = $8

-- name: deleteActivity :exec
UPDATE activities SET is_deleted = 'Y' WHERE id = $1

-- name: findActivityByTitle :many
SELECT * FROM activities WHERE title LIKE '%$1%'

-- name: findActivityByStartDate :many
SELECT * FROM activities WHERE start_date = $1