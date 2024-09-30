package config

import (
	"time"

	"github.com/ferdiebergado/fullstack-go/pkg/env"
)

const (
	ServerReadTimeout  = 10 * time.Second
	ServerWriteTimeout = 30 * time.Second
	ServerIdleTimeout  = time.Minute
)

var (
	App = &appOptions{
		Env:  env.Must("APP_ENV"),
		Port: env.Must("APP_PORT"),
	}

	Db = &dbOptions{
		Driver:   "pgx",
		User:     env.Must("DB_USER"),
		Password: env.Must("DB_PASS"),
		Host:     env.Must("DB_HOST"),
		Port:     env.Must("DB_PORT"),
		Name:     env.Must("DB_NAME"),
	}
)
