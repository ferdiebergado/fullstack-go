#!/usr/bin/env sh

if [ -e $DB_PASS ]; then
    echo "DB_PASS not set"
    exit 1
fi

if [ -e $DB_NAME ]; then
    echo "DB_NAME not set"
    exit 1
fi

container_name=${1:-ptms_db}
postgres_image='postgres:16-alpine3.20'

docker=$(command -v docker)
podman=$(command -v podman)

if [ -n $podman ]; then 
    docker=$podman
fi 

$docker run -d --rm --network host --name $container_name  \
    -e POSTGRES_PASSWORD="$DB_PASS" \
    -v ./postgresql.conf:/etc/postgresql/postgresql.conf:Z \
    $postgres_image \
    -c 'config_file=/etc/postgresql/postgresql.conf'

echo "Waiting for db to start up..."

sleep 5

echo "Creating database..."

$docker exec -it $container_name psql -U postgres -w -c "CREATE DATABASE $DB_NAME;"

$docker logs -f $container_name
