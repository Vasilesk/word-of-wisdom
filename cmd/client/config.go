package main

import "time"

type Config struct {
	HTTPClient ConfigClient `yaml:"httpClient"`
	API        ConfigAPI    `yaml:"api"`
}

type ConfigClient struct {
	Timeout time.Duration `yaml:"timeout"`
}

type ConfigAPI struct {
	Wisdom ConfigWisdom `yaml:"wisdom"`
}

type ConfigWisdom struct {
	BaseURL string `yaml:"baseUrl"`
}
