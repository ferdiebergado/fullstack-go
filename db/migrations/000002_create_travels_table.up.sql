CREATE TABLE IF NOT EXISTS travels (
    id SERIAL PRIMARY KEY,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status SMALLINT NOT NULL DEFAULT 0,
    remarks TEXT,
    metadata JSONB NOT NULL DEFAULT '{}',
    activity_id INT NOT NULL,
    FOREIGN KEY (activity_id) REFERENCES activities (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);