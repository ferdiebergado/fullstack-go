CREATE TABLE IF NOT EXISTS offices (
    id INT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name VARCHAR(100),
    short_name VARCHAR(10) NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE offices ADD CONSTRAINT pk_offices PRIMARY KEY (id);