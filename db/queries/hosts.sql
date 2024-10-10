-- name: CreateHost :one
INSERT INTO hosts (name) VALUES ($1) RETURNING *;

-- name: GetHosts :many
SELECT * FROM hosts ORDER BY name;

-- name: FindHost :one
SELECT * FROM hosts WHERE id = $1;

-- name: FindHostsByName :many
SELECT * FROM hosts WHERE name LIKE '%$1%';

-- name: UpdateHost :exec
UPDATE hosts SET name = $1 WHERE id = $2;

-- name: DeleteHost :exec
UPDATE hosts SET deleted_at = now() WHERE id = $1;