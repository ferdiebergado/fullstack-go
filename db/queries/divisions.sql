-- name: GetDivisions :many
SELECT * FROM divisions ORDER BY name;

-- name: FindDivision :one
SELECT * FROM divisions WHERE id = $1;

-- name: FindDivisionsByRegion :many
SELECT * FROM divisions WHERE region_id = $1 ORDER BY name;

-- name: GetDivisionWithRegion :many
SELECT d.*, r.name as region
FROM divisions d
    JOIN regions r ON r.region_id = d.region_id
ORDER BY d.name;

-- name: GetDivisionOrderedByRegion :many
SELECT * FROM divisions ORDER BY region_id;