ALTER TABLE activities DROP COLUMN is_deleted;

ALTER TABLE activities ADD COLUMN deleted_at TIMESTAMPTZ;