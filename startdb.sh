#!/usr/bin/env sh

if [ -e $DB_PASS ]; then
    echo "DB_PASS not set"
    exit 1
fi

container=$(command -v docker)
podman=$(command -v podman)

if [ -n $podman ]; then 
    container=$podman
fi 

$container run -ti --rm -p 5432:5432 --name ptms_postgres  \
    -e POSTGRES_PASSWORD="$DB_PASS" \
    -v ./postgresql.conf:/etc/postgresql/postgresql.conf:Z \
    postgres:16-alpine3.20 \
    -c 'config_file=/etc/postgresql/postgresql.conf'
