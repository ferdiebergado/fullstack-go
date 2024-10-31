CREATE VIEW activity_details AS
SELECT a.*, v.name venue, r.name region, h.name host
FROM
    activities a
    JOIN venues v ON v.id = a.venue_id
    JOIN divisions d ON d.id = v.division_id
    JOIN regions r ON r.region_id = d.region_id
    JOIN hosts h ON h.id = a.host_id;