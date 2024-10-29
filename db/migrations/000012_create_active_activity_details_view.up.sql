CREATE VIEW active_activity_details AS
SELECT *
FROM activity_details
WHERE
    deleted_at IS NULL;