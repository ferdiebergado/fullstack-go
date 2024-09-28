#!/usr/bin/env sh

export $(awk -F '=' '/^DB_USER|^DB_PASS|^DB_HOST|^DB_PORT|^DB_NAME/ { gsub(/ /, "", $2); print $1 "=" $2 }' .env)

export DSN="postgres://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

echo $DSN
