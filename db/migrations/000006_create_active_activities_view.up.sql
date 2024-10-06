CREATE VIEW active_activities AS
SELECT *
FROM activities
WHERE
    deleted_at IS NULL;