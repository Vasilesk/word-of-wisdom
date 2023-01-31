package server

import "time"

type Config struct {
	Port            int
	RequestTimeout  time.Duration
	ShutdownTimeout time.Duration
}
