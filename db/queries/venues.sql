-- name: GetVenues :many
SELECT v.*, r.name as region
FROM venues v
    JOIN regions r ON v.region_id = r.region_id
ORDER BY v.name;