CREATE TABLE IF NOT EXISTS activities (
    id BIGINT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    title VARCHAR(300) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    venue_id INT NOT NULL,
    host_id INT NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE activities ADD CONSTRAINT pk_activities PRIMARY KEY (id);

ALTER TABLE activities
ADD CONSTRAINT chk_end_date_after_start_date CHECK (end_date >= start_date);