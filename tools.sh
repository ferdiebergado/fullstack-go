#!/usr/bin/env sh

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
