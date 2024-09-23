CREATE TABLE IF NOT EXISTS activities {
    id serial PRIMARY KEY
    title TEXT NOT NULL
    start_date VARCHAR(10) NOT NULL
    end_date VARCHAR(10) NOT NULL
    venue VARCHAR(100)
    host VARCHAR(100)
    status INT NOT NULL
    metadata JSONB
    created_at TIMESTAMPTZ DEFAULT NOW()
    updated_at TIMESTAMPTZ DEFAULT NOW()
    is_deleted CHAR DEFAULT 'N'
}