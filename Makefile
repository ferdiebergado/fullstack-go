CONTAINER_NAME := ptms_postgres
CONTAINER_IMAGE := postgres:16-alpine3.20
DSN := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR := db/migrations
CONTAINER := docker
DB_NAME := test

# Capture all additional arguments after the target
ARGS := $(filter-out $@,$(MAKECMDGOALS))

# Load environment variables from .env
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: install postgres migration migrate rollback seed generate setup_test test security critic css.watch js.watch dev

install:
	which migrate || go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	which sqlc || go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	which air || go install github.com/air-verse/air@v1.52.2
	which gosec || go install github.com/securego/gosec/v2/cmd/gosec@latest
	which gocritic || go install github.com/go-critic/go-critic/cmd/gocritic@latest

postgres db:
	$(CONTAINER) run -d --rm --network host --name $(CONTAINER_NAME) -e POSTGRES_PASSWORD="$(DB_PASS)" \
		-v ./db/postgresql.conf:/etc/postgresql/postgresql.conf:Z \
		-v ./db/psqlrc:/root/.psqlrc:Z \
		-v ./db/seeds:/root:Z \
		$(CONTAINER_IMAGE) -c 'config_file=/etc/postgresql/postgresql.conf'

migration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(word 2, $(ARGS))

migrate:
	migrate -database $(DSN) -path $(MIGRATIONS_DIR) up

rollback:
	migrate -database $(DSN) -path $(MIGRATIONS_DIR) down

seed:
	$(CONTAINER) exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -w -c "\copy activities (title,start_date,end_date,venue_id,host_id) FROM '/root/activities.csv' DELIMITER ',' CSV HEADER"
	$(CONTAINER) exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -w -c "\copy venues (name,division_id) FROM '/root/venues.csv' DELIMITER ',' CSV HEADER"
	$(CONTAINER) exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -w -c "\copy hosts (name) FROM '/root/hosts.csv' DELIMITER ',' CSV HEADER"

generate gen:
	export DSN="$(DSN)"; sqlc generate

setup_test:
	$(CONTAINER) exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -w -c "CREATE DATABASE $(DB_NAME);"
	migrate -database $(DSN) -path $(MIGRATIONS_DIR) up

test:
	APP_ENV=test DB_NAME=$(DB_NAME) clear && go test -race ./...

security sec:
	gosec ./...

critic crit:
	gocritic check ./...

css.watch:
	esbuild assets/css/styles.css --bundle --outdir=static/css --watch

js.watch:
	esbuild assets/js/**/*.js --bundle --outdir=static/js --sourcemap --target=es6 --splitting --format=esm --watch

dev:
	air
