CREATE TABLE IF NOT EXISTS activities {
    id SERIAL PRIMARY KEY
    title TEXT NOT NULL
    start_date DATE NOT NULL
    end_date DATE NOT NULL
    venue VARCHAR(100)
    host VARCHAR(100)
    status INT NOT NULL
    metadata JSONB
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    is_deleted BOOLEAN DEFAULT FALSE
}