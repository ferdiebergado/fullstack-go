CREATE TABLE IF NOT EXISTS activities (
    id BIGINT GENERATED BY DEFAULT AS IDENTITY,
    title TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    venue VARCHAR(100),
    host VARCHAR(100),
    metadata JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE activities ADD CONSTRAINT pk_activities PRIMARY KEY (id);