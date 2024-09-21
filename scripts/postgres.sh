#!/usr/bin/env sh

docker run -d --name postgres_db \
    -e POSTGRES_PASSWORD="$DB_PASS" \
    -v "$PWD/postgresql.conf":/etc/postgresql/postgresql.conf \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
	-v "../db/data":/var/lib/postgresql/data \
    postgres:16-bookworm \
    -c 'config_file=/etc/postgresql/postgresql.conf'
