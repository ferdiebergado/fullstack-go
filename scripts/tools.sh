#!/usr/bin/env sh

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/air-verse/air@latest
