-- name: CreateOffice :one
INSERT INTO
    offices (name, metadata)
VALUES ($1, $2)
RETURNING
    *;

-- name: ListOffice :many
SELECT * FROM offices ORDER BY name;

-- name: FindOffice :one
SELECT * FROM offices WHERE id = $1;

-- name: UpdateOffice :exec
UPDATE offices
SET
    name = $1,
    metadata = $2,
    updated_at = NOW()
WHERE
    id = $3;

-- name: DeleteOffice :exec
UPDATE offices SET deleted_at = NOW() WHERE id = $1;

-- name: RestoreOffice :exec
UPDATE offices SET deleted_at = NULL WHERE id = $1;

-- name: FindOfficeByName :many
SELECT * FROM offices WHERE name LIKE '%$1%';