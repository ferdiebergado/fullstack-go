#!/usr/bin/env sh

container_name=ptms_db
tool=$(command -v docker)
podman=$(command -v podman)

if [ -n $podman ]; then 
    tool=$podman
fi 

$tool run -d --name $container_name --replace \
    -e POSTGRES_PASSWORD="$DB_PASS" \
    -v ./postgresql.conf:/etc/postgresql/postgresql.conf:Z \
    postgres:16-alpine \
    -c 'config_file=/etc/postgresql/postgresql.conf'

sleep 5

$tool exec -it $container_name psql -U postgres -w -c "CREATE DATABASE $DB_NAME;"