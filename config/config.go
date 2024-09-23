package config

import "time"

const (
	ServerReadTimeout  = 10 * time.Second
	ServerWriteTimeout = 30 * time.Second
	ServerIdleTimeout  = time.Minute
)
