CREATE VIEW active_activities AS
SELECT *
FROM activities
WHERE
    is_deleted = FALSE;