CREATE TABLE IF NOT EXISTS hosts (
    id BIGINT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    name VARCHAR(100) NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE hosts ADD CONSTRAINT pk_hosts PRIMARY KEY (id);