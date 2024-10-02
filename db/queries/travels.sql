-- name: CreateTravel :one
INSERT INTO
    travels (
        start_date,
        end_date,
        status,
        remarks,
        metadata,
        activity_id
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: ListTravels :many
SELECT * FROM travels ORDER BY start_date DESC;

-- name: FindTravel :one
SELECT * FROM travels WHERE id = $1;

-- name: UpdateTravel :exec
UPDATE travels
SET
    start_date = $1,
    end_date = $2,
    status = $3,
    remarks = $4,
    activity_id = $5,
    metadata = $6,
    updated_at = NOW()
WHERE
    id = $7;

-- name: DeleteTravel :exec
UPDATE travels SET deleted_at = NOW() WHERE id = $1;

-- name: RestoreTravel :exec
UPDATE travels SET deleted_at = NULL WHERE id = $1;

-- name: FindTravelByActivityId :many
SELECT * FROM travels WHERE activity_id = $1;

-- name: FindTravelByActivityTitle :many
SELECT *
FROM travels AS t
    INNER JOIN activities AS a ON t.activity_id = a.id
WHERE
    t.activity_id = $1
ORDER BY t.start_date DESC;

-- name: FindTravelByStartDate :many
SELECT * FROM travels WHERE start_date = $1;