#!/usr/bin/env bash

if [ -e $DB_PASS ]; then
    if [ -f .env ]; then
        DB_PASS=$(awk -F '=' '$1 == "DB_PASS" { print $2 }' .env)
    fi
fi

if [ -e $DB_PASS ]; then
    echo "DB_PASS not set"
    exit 1
fi

docker run -ti --rm --network host --name ptms_postgres \
    -e POSTGRES_PASSWORD="$DB_PASS" \
    -v ./postgresql.conf:/etc/postgresql/postgresql.conf:Z \
    postgres:16-alpine3.20 \
    -c 'config_file=/etc/postgresql/postgresql.conf'
