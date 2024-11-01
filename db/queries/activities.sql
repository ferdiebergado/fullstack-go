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

-- name: UpdateActivity :exec
UPDATE active_activities
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

-- name: FindActiveActivity :one
SELECT id FROM active_activities WHERE id = $1;

-- name: FindActiveActivityDetails :one
SELECT * FROM active_activity_details WHERE id = $1;

-- name: FindActivityDetail :one
SELECT * FROM activity_details WHERE id = $1;

-- name: FindActivity :one
SELECT * FROM activities WHERE id = $1;

-- name: CountActiveActivities :one
SELECT COUNT(*) FROM active_activities;

-- name: CountActivities :one
SELECT COUNT(*) FROM activities;