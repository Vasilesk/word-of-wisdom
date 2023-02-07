package main

import "time"

type Config struct {
	HTTPClient ConfigClient `yaml:"httpClient"`
}

type ConfigClient struct {
	Timeout time.Duration `yaml:"timeout"`
}
