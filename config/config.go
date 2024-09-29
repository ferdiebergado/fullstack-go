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
		Env:  env.GetEnv("APP_ENV"),
		Port: env.GetEnv("APP_PORT"),
	}

	Db = &dbOptions{
		User:     env.GetEnv("DB_USER"),
		Password: env.GetEnv("DB_PASS"),
		Host:     env.GetEnv("DB_HOST"),
		Port:     env.GetEnv("DB_PORT"),
		Name:     env.GetEnv("DB_NAME"),
	}
)
