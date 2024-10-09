-- name: GetVenues :many
SELECT v.*, d.name as division, r.name as region
FROM
    venues v
    JOIN divisions d ON d.id = v.division_id
    JOIN regions r ON d.region_id = r.region_id
ORDER BY v.name;

-- name: CreateVenue :one
INSERT INTO venues (name, division_id) VALUES ($1, $2) RETURNING *;