CREATE VIEW activity_details AS
SELECT
    a.id,
    a.title,
    a.start_date,
    a.end_date,
    v.name venue,
    r.name region,
    h.name host,
    a.metadata,
    a.created_at,
    a.updated_at,
    a.deleted_at
FROM
    activities a
    JOIN venues v ON v.id = a.venue_id
    JOIN divisions d ON d.id = v.division_id
    JOIN regions r ON r.region_id = d.region_id
    JOIN hosts h ON h.id = a.host_id;