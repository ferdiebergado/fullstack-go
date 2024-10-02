-- name: CreatePersonnel :one
INSERT INTO
    personnel (
        lastname,
        firstname,
        mi,
        position_id,
        office_id,
        metadata
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: ListPersonnel :many
SELECT * FROM personnel ORDER BY lastname;

-- name: FindPersonnel :one
SELECT * FROM personnel WHERE id = $1;

-- name: UpdatePersonnel :exec
UPDATE personnel
SET
    lastname = $1,
    firstname = $2,
    mi = $3,
    position_id = $4,
    office_id = $5,
    metadata = $6,
    updated_at = NOW()
WHERE
    id = $7;

-- name: DeletePersonnel :exec
UPDATE personnel SET deleted_at = NOW() WHERE id = $1;

-- name: RestorePersonnel :exec
UPDATE personnel SET deleted_at = NULL WHERE id = $1;

-- name: FindPersonnelByLastname :many
SELECT * FROM personnel WHERE lastname = $1;

-- name: FindPersonnelByFirstname :many
SELECT * FROM personnel WHERE firstname = $1;