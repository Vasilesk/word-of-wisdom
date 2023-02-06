package main

import (
	"time"

	"github.com/vasilesk/word-of-wisdom/pkg/http/server"
)

type Config struct {
	Server     ConfigServer     `yaml:"server"`
	Signer     ConfigSigner     `yaml:"signer"`
	Pow        ConfigPow        `yaml:"pow"`
	PowChecker ConfigPowChecker `yaml:"powChecker"`
}

type ConfigSigner struct {
	Key string `yaml:"key"`
}

type ConfigPow struct {
	Difficulty int `yaml:"difficulty"`
	NonceSize  int `yaml:"nonceSize"`
}

type ConfigPowChecker struct {
	ChallengeValid time.Duration `yaml:"challengeValid"`
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
