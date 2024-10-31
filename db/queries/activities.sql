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

-- name: FindActivity :one
SELECT * FROM activity_details WHERE id = $1;

-- name: FindActiveActivitiesByTitle :many
SELECT * FROM active_activity_details WHERE title LIKE '%$1%';

-- name: FindActivitiesByTitle :many
SELECT * FROM activity_details WHERE title LIKE '%$1%';

-- name: FindActiveActivitiesByStartDate :many
SELECT * FROM active_activity_details WHERE start_date = $1;

-- name: FindActivitiesByStartDate :many
SELECT * FROM activity_details WHERE start_date = $1;

-- name: ListActiveActivities :many
SELECT *, COUNT(*) OVER () AS total_items
FROM active_activity_details
WHERE
    COALESCE($5, '') = ''
    OR title LIKE $5
ORDER BY
    CASE
        WHEN $1 = 'title'
        AND $2 = 'ASC' THEN title
    END ASC,
    CASE
        WHEN $1 = 'title'
        AND $2 = 'DESC' THEN title
    END DESC,
    CASE
        WHEN $1 = 'start_date'
        AND $2 = 'ASC' THEN start_date
    END ASC,
    CASE
        WHEN $1 = 'start_date'
        AND $2 = 'DESC' THEN start_date
    END DESC,
    CASE
        WHEN $1 = 'end_date'
        AND $2 = 'ASC' THEN end_date
    END ASC,
    CASE
        WHEN $1 = 'end_date'
        AND $2 = 'DESC' THEN end_date
    END DESC,
    CASE
        WHEN $1 = 'venue'
        AND $2 = 'ASC' THEN venue
    END ASC,
    CASE
        WHEN $1 = 'venue'
        AND $2 = 'DESC' THEN venue
    END DESC,
    CASE
        WHEN $1 = 'host'
        AND $2 = 'ASC' THEN host
    END ASC,
    CASE
        WHEN $1 = 'host'
        AND $2 = 'DESC' THEN host
    END DESC
LIMIT $3
OFFSET
    $4;

-- name: ListActivities :many
SELECT *, COUNT(*) OVER () AS total_items
FROM activity_details
WHERE
    COALESCE($5, '') = ''
    OR title LIKE $5
ORDER BY
    CASE
        WHEN $1 = 'title'
        AND $2 = 'ASC' THEN title
    END ASC,
    CASE
        WHEN $1 = 'title'
        AND $2 = 'DESC' THEN title
    END DESC,
    CASE
        WHEN $1 = 'start_date'
        AND $2 = 'ASC' THEN start_date
    END ASC,
    CASE
        WHEN $1 = 'start_date'
        AND $2 = 'DESC' THEN start_date
    END DESC,
    CASE
        WHEN $1 = 'end_date'
        AND $2 = 'ASC' THEN end_date
    END ASC,
    CASE
        WHEN $1 = 'end_date'
        AND $2 = 'DESC' THEN end_date
    END DESC,
    CASE
        WHEN $1 = 'venue'
        AND $2 = 'ASC' THEN venue
    END ASC,
    CASE
        WHEN $1 = 'venue'
        AND $2 = 'DESC' THEN venue
    END DESC,
    CASE
        WHEN $1 = 'host'
        AND $2 = 'ASC' THEN host
    END ASC,
    CASE
        WHEN $1 = 'host'
        AND $2 = 'DESC' THEN host
    END DESC
LIMIT $3
OFFSET
    $4;

-- name: CountActiveActivities :one
SELECT COUNT(*) FROM active_activities;

-- name: CountActivities :one
SELECT COUNT(*) FROM activities;