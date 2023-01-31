package main

import (
	"time"

	"github.com/vasilesk/word-of-wisdom/pkg/http/server"
)

type Config struct {
	Server ConfigServer `yaml:"server"`
}

type ConfigServer struct {
	Port            int           `yaml:"port"`
	RequestTimeout  time.Duration `yaml:"requestTimeout"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
}

func (c ConfigServer) ToServerConfig() server.Config {
	return server.Config{
		Port:            c.Port,
		RequestTimeout:  c.RequestTimeout,
		ShutdownTimeout: c.ShutdownTimeout,
	}
}
